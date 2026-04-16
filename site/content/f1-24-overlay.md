---
title: "Free F1 24 Overlay"
subtitle: "Lightweight telemetry overlay for F1 24 with automatic UDP detection, 11 customizable widgets, and simple one-time setup."
description: "Free F1 24 telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, and track map widgets. Enable UDP telemetry and go."
ogImage: "assets/application-preview.gif"
---

## Free Telemetry Overlay for F1 24

RacePulse is a free telemetry overlay for EA Sports F1 24. It receives F1 24's UDP telemetry broadcast and displays real-time speed, RPM, inputs, lap times, delta, fuel, and more as transparent overlay widgets over your game.

One-time setup: enable UDP telemetry in F1 24's settings, launch RacePulse, and overlays appear automatically. No extra software, no web apps, no configuration files.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. In F1 24, go to **Settings > Telemetry Settings**
4. Set **UDP Telemetry** to **On**
5. Ensure **UDP Port** is set to **22024** (this is the default)
6. Start a session — RacePulse auto-detects F1 24 and begins displaying telemetry

> **Note:** You only need to enable UDP telemetry once. The setting persists in F1 24 across restarts. RacePulse scans for the broadcast on port 22024 automatically.

If Windows Firewall asks to allow RacePulse, click Allow. RacePulse only listens for local UDP packets and never sends data externally.

## Available Widgets for F1 24

F1 24 supports 11 of the 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in km/h or mph
- **Gear** — Large gear indicator for peripheral visibility
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to your best lap
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — Combined display with rev lights, gear, RPM, speed, position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, optional steering wheel, gear, and speed
- **Flat Track Map** — Horizontal bar showing all 20 cars on track with your position highlighted

### Not Available

- **Relative** — Not available for F1 24 (iRacing exclusive)

## F1 24-Specific Data

RacePulse captures F1-specific telemetry beyond standard fields:

- **ERS** — Energy Recovery System deployment and harvest
- **DRS** — Availability and activation status
- **Fuel Mix** — Current fuel mixture setting
- **Tyre Compound** — Active compound (soft, medium, hard, intermediate, wet)
- **Tyre Wear** — Per-wheel wear percentage
- **Tyre Temperature** — Inner, outer, and surface temperatures

RacePulse handles both F1 24 and F1 25 UDP formats automatically. The game's packet header identifies the format version (2024 vs 2025), and RacePulse parses each accordingly. If you play both games, you don't need to change anything — RacePulse detects which game is broadcasting based on the port and packet format.

## How UDP Telemetry Works

F1 24 broadcasts telemetry as UDP packets to `localhost` on port **22024**. The port number follows the game year (F1 24 = 22024, F1 25 = 22025). RacePulse listens for these packets locally — the connection is receive-only with no data sent back.

Multiple packet types are broadcast at different rates (motion, session, lap data, car telemetry, car status, etc.). RacePulse aggregates them into a unified 60Hz telemetry stream for smooth widget updates.

## Customization

Every widget is independently configurable:

- Metric or imperial units
- Toggle display elements on or off
- Per-widget frame rate (60, 30, or 15 FPS)
- Reposition by dragging in Edit Mode
- Persistent settings across sessions

## Still Playing F1 24?

RacePulse supports both F1 24 and F1 25 simultaneously. If you have both games installed, RacePulse will connect to whichever one is actively broadcasting telemetry. There is no manual game selection — detection is fully automatic.

## Why RacePulse for F1 24

- **Free** — No cost, no subscription, no feature limits
- **Single executable** — 21MB `.exe`, no installer needed
- **Simple setup** — Enable UDP telemetry once in F1 24 and forget about it
- **Lightweight** — Won't impact your game performance
- **Multi-sim** — Same app works with iRacing, ACC, Assetto Corsa, Le Mans Ultimate, and F1 25

---

[Download RacePulse](/install/) and have your F1 24 overlays running in under a minute. See all available widgets on the [features page](/features/) or check the [changelog](/changelog/) for recent updates.
