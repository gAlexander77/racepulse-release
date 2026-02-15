package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type lineType struct {
	Type string `json:"type"`
}

type packetLine struct {
	Type       string `json:"type"`
	DeltaUS    int64  `json:"dt_us"`
	Size       int    `json:"size"`
	DataBase64 string `json:"data_b64"`
}

func main() {
	var inPath string
	var host string
	var port int
	var speed float64
	var loop bool

	flag.StringVar(&inPath, "in", "", "Input recording .jsonl file")
	flag.StringVar(&host, "host", "127.0.0.1", "Destination host")
	flag.IntVar(&port, "port", 22025, "Destination UDP port")
	flag.Float64Var(&speed, "speed", 1.0, "Playback speed multiplier")
	flag.BoolVar(&loop, "loop", false, "Loop playback forever")
	flag.Parse()

	if strings.TrimSpace(inPath) == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run ./tools/f1playback -in <file.jsonl> [-host 127.0.0.1] [-port 22025] [-speed 1.0] [--loop]")
		os.Exit(2)
	}
	if speed <= 0 {
		fmt.Fprintln(os.Stderr, "speed must be > 0")
		os.Exit(2)
	}

	target := fmt.Sprintf("%s:%d", host, port)
	addr, err := net.ResolveUDPAddr("udp4", target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve destination %s: %v\n", target, err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open UDP socket to %s: %v\n", target, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Playback %s -> %s (%.2fx)\n", inPath, target, speed)
	if loop {
		fmt.Println("Loop mode enabled")
	}

	run := 0
	for {
		run++
		count, totalBytes, err := playOnce(conn, inPath, speed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "playback failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Run %d complete: %d packets (%d bytes)\n", run, count, totalBytes)

		if !loop {
			break
		}
	}
}

func playOnce(conn *net.UDPConn, path string, speed float64) (int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	packets := 0
	totalBytes := 0

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}

		var lt lineType
		if err := json.Unmarshal([]byte(line), &lt); err != nil {
			return packets, totalBytes, err
		}
		if lt.Type != "packet" {
			continue
		}

		var p packetLine
		if err := json.Unmarshal([]byte(line), &p); err != nil {
			return packets, totalBytes, err
		}

		if p.DeltaUS > 0 {
			time.Sleep(time.Duration(float64(p.DeltaUS)/speed) * time.Microsecond)
		}

		data, err := base64.StdEncoding.DecodeString(p.DataBase64)
		if err != nil {
			return packets, totalBytes, err
		}
		if p.Size > 0 && len(data) != p.Size {
			return packets, totalBytes, fmt.Errorf("packet size mismatch: expected %d got %d", p.Size, len(data))
		}

		n, err := conn.Write(data)
		if err != nil {
			return packets, totalBytes, err
		}

		packets++
		totalBytes += n
	}

	if err := s.Err(); err != nil {
		return packets, totalBytes, err
	}

	return packets, totalBytes, nil
}
