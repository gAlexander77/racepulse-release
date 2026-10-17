---
title: "Free F1 24 Overlay"
subtitle: "Lightweight telemetry overlay for F1 24 with automatic UDP detection, 13 customizable widgets, and simple one-time setup."
description: "Free F1 24 telemetry overlay with real-time speed, RPM, inputs, lap times, relative, fuel, G-force, and track map widgets. Enable UDP telemetry and go."
ogImage: "assets/application-preview.jpg"
---

## Free Telemetry Overlay for F1 24

RacePulse is a free telemetry overlay for EA Sports F1 24. It receives F1 24's UDP telemetry broadcast and displays real-time speed, RPM, inputs, lap times, standings, fuel, and more as transparent overlay widgets over your game.

One-time setup: enable UDP telemetry in F1 24's settings, launch RacePulse, and overlays appear automatically. No extra software, no web apps, no configuration files.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. In F1 24, go to **Settings > Telemetry Settings**
4. Set **UDP Telemetry** to **On**
5. Set **UDP Port** to one of the ports RacePulse listens on: **22025**, **22024**, **22023**, or **20777**
6. Set **UDP Format** to **2024**
7. Start a session — RacePulse auto-detects F1 24 and begins displaying telemetry

> **Note:** You only need to enable UDP telemetry once. The setting persists in F1 24 across restarts. RacePulse listens on all four ports at once, so any of them works.

> **If it doesn't connect:** another app may be using that port. Pick another of RacePulse's ports in the game, or close the other app.

If Windows Firewall asks to allow RacePulse, click Allow. RacePulse only listens for UDP packets and never sends your racing data anywhere.

## Available Widgets for F1 24

F1 24 supports 13 of the 14 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in KPH or MPH
- **Gear** — Large gear indicator for peripheral visibility
- **RPM** — The number, or a row of rev lights in F1's green, red, and blue if you like
- **Lap Times** — Current lap, last lap, and personal best, with an optional strike-through when F1 invalidates your lap
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — Combined display with rev lights, gear, and readouts you choose, such as speed, RPM, position, time left, elapsed time, ERS, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, a steering wheel, gear, and speed
- **Flat Track Map** — Every car on one line of the lap, with your position highlighted, sector lines, and yellow flag zones
- **Relative** — Nearby drivers sorted by gap on track, with car numbers, country flags, and pit/out lap indicators
- **Standings** — The whole field in six layouts, with gaps, best and last laps, pit stops, positions gained, and each car's tire (beta)
- **Lap Log** — Every lap you drive with its delta, sectors, and lap validity, under the session best and your best

### Not Available

- **Delta** — F1 24 does not send a live delta to your best lap, so the Delta widget isn't offered. The Lap Log still compares each finished lap with your best.

## F1 24-Specific Data

RacePulse captures F1-specific telemetry beyond standard fields:

- **ERS** — Energy Recovery System deployment and harvest
- **DRS** — Availability and activation status
- **Fuel Mix** — Current fuel mixture setting
- **Tire Compound** — Active compound (soft, medium, hard, intermediate, wet)
- **Tire Wear** — Per-wheel wear percentage
- **Tire Temperature** — Inner, outer, and surface temperatures

RacePulse handles both F1 24 and F1 25 UDP formats automatically. The game's packet header identifies the format version (2024 vs 2025), and RacePulse parses each accordingly. If you play both games, you don't need to change anything — RacePulse names the game from the packet format it receives.

## How UDP Telemetry Works

F1 24 sends telemetry as UDP packets to the IP address and port set in its Telemetry Settings. RacePulse listens on ports **22025**, **22024**, **22023**, and **20777**, and connects on whichever one the game sends to. The connection is receive-only with no data sent back.

Multiple packet types are broadcast at different rates (motion, session, lap data, car telemetry, car status, etc.). RacePulse aggregates them into a unified 60Hz telemetry stream for smooth widget updates.

## Customization

Every widget is independently configurable:

- Metric or imperial units
- Toggle display elements on or off
- Per-widget frame rate (60, 30, or 15 FPS)
- Reposition by dragging while arranging, then lock for racing
- Persistent settings across sessions

## Still Playing F1 24?

RacePulse supports both F1 24 and F1 25 simultaneously. If you have both games installed, RacePulse will connect to whichever one is actively broadcasting telemetry. There is no manual game selection — detection is fully automatic.

## Why RacePulse for F1 24

- **Free** — No cost, no subscription, no feature limits
- **Single executable** — 19MB `.exe`, no installer needed
- **Simple setup** — Enable UDP telemetry once in F1 24 and forget about it
- **Lightweight** — Won't impact your game performance
- **Multi-sim** — Same app works with iRacing, ACC, Assetto Corsa, Le Mans Ultimate, and F1 25

---

[Download RacePulse](/install/) and have your F1 24 overlays running in under a minute. See all available widgets on the [features page](/features/) or check the [changelog](/changelog/) for recent updates.
