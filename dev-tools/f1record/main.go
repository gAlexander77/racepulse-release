package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type metaLine struct {
	Type      string `json:"type"`
	Format    string `json:"format"`
	Game      string `json:"game"`
	CreatedAt string `json:"createdAt"`
	Port      int    `json:"port"`
}

type packetLine struct {
	Type       string `json:"type"`
	DeltaUS    int64  `json:"dt_us"`
	Size       int    `json:"size"`
	DataBase64 string `json:"data_b64"`
}

func main() {
	var port int
	var outPath string
	var duration time.Duration

	flag.IntVar(&port, "port", 22025, "UDP port to record")
	flag.StringVar(&outPath, "out", "", "Output .jsonl file path")
	flag.DurationVar(&duration, "duration", 0, "Optional duration, e.g. 30s or 2m")
	flag.Parse()

	if outPath == "" {
		outPath = fmt.Sprintf("f1-telemetry-%s.jsonl", time.Now().Format("20060102-150405"))
	}

	cleanOut := filepath.Clean(outPath)
	if dir := filepath.Dir(cleanOut); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "failed to create output directory: %v\n", err)
			os.Exit(1)
		}
	}

	f, err := os.Create(cleanOut)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	enc := json.NewEncoder(w)
	if err := enc.Encode(metaLine{
		Type:      "meta",
		Format:    "racepulse-f1-udp-v1",
		Game:      "f1",
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		Port:      port,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write metadata: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen on UDP %d: %v\n", port, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Recording F1 telemetry on :%d -> %s\n", port, cleanOut)
	if duration > 0 {
		fmt.Printf("Auto-stop in %s\n", duration)
	}
	fmt.Println("Press Ctrl+C to stop")

	var stopAt time.Time
	if duration > 0 {
		stopAt = time.Now().Add(duration)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	buf := make([]byte, 4096)
	start := time.Now()
	last := start
	packets := 0
	totalBytes := 0

	for {
		if !stopAt.IsZero() && time.Now().After(stopAt) {
			break
		}

		_ = conn.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				select {
				case <-sigCh:
					goto done
				default:
					continue
				}
			}
			fmt.Fprintf(os.Stderr, "UDP read error: %v\n", err)
			continue
		}

		now := time.Now()
		delta := now.Sub(last).Microseconds()
		if packets == 0 {
			delta = now.Sub(start).Microseconds()
		}
		last = now

		line := packetLine{
			Type:       "packet",
			DeltaUS:    delta,
			Size:       n,
			DataBase64: base64.StdEncoding.EncodeToString(buf[:n]),
		}

		if err := enc.Encode(line); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write packet: %v\n", err)
			os.Exit(1)
		}

		packets++
		totalBytes += n

		select {
		case <-sigCh:
			goto done
		default:
		}
	}

done:
	if err := w.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "flush failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Recorded %d packets (%d bytes)\n", packets, totalBytes)
}
