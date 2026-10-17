---
title: "Free F1 25 Overlay"
subtitle: "Lightweight telemetry overlay for F1 25 with automatic UDP detection, 13 customizable widgets, and simple one-time setup."
description: "Free F1 25 telemetry overlay with real-time speed, RPM, inputs, lap times, relative, fuel, G-force, and track map widgets. Enable UDP telemetry and go."
ogImage: "assets/application-preview.jpg"
---

## Free Telemetry Overlay for F1 25

RacePulse is a free telemetry overlay for EA Sports F1 25. It listens for F1 25's UDP telemetry broadcast to display real-time speed, RPM, inputs, lap times, standings, fuel, and more as transparent overlay widgets floating over your game.

Setup takes less than a minute: enable UDP telemetry in F1 25's settings, launch RacePulse, and your overlays appear automatically. No additional software, no web dashboards, no JSON configuration.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. In F1 25, go to **Settings > Telemetry Settings**
4. Set **UDP Telemetry** to **On**
5. Set **UDP Port** to one of the ports RacePulse listens on: **22025**, **22024**, **22023**, or **20777**
6. Set **UDP Format** to **2025**. With the 2026 Season Pack, its 2026 format works too.
7. Start a session — RacePulse auto-detects F1 25 and begins displaying telemetry

> **First time only:** You need to enable UDP telemetry in F1 25 once. After that, the setting persists across game restarts. RacePulse listens on all four ports at once, so any of them works.

> **If it doesn't connect:** another app may be using that port. Pick another of RacePulse's ports in the game, or close the other app.

If Windows Firewall prompts you to allow RacePulse network access, click Allow. RacePulse only listens for UDP packets — it does not upload your racing data anywhere.

## Available Widgets for F1 25

F1 25 supports 13 of the 14 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in KPH or MPH
- **Gear** — Large gear indicator visible in peripheral vision
- **RPM** — The number, or a row of rev lights in F1's green, red, and blue if you like
- **Lap Times** — Current lap, last lap, and personal best, with an optional strike-through when F1 invalidates your lap
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force visualization

### Advanced Widgets

- **Telemetry Core** — All-in-one display with rev lights, gear, and readouts you choose, such as speed, RPM, position, time left, elapsed time, ERS, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, a steering wheel, gear, and speed
- **Flat Track Map** — Every car on one line of the lap, with your position highlighted, sector lines, and yellow flag zones
- **Relative** — Nearby drivers sorted by gap on track, with car numbers, country flags, and pit/out lap indicators
- **Standings** — The whole field in six layouts, with gaps, best and last laps, pit stops, positions gained, and each car's tire (beta)
- **Lap Log** — Every lap you drive with its delta, sectors, and lap validity, under the session best and your best

### Not Available

- **Delta** — F1 25 does not send a live delta to your best lap, so the Delta widget isn't offered. The Lap Log still compares each finished lap with your best.

> **Note:** Online drivers who hide their names in F1 25 show as their team in Relative and Standings. Country flags show for drivers whose nationality the game sends.

## F1 25-Specific Data

RacePulse captures F1-specific telemetry beyond standard fields:

- **ERS** — Energy Recovery System deployment and harvest levels
- **DRS** — DRS availability and activation status
- **Fuel Mix** — Current fuel mixture setting
- **Tire Compound** — Active tire compound (soft, medium, hard, intermediate, wet)
- **Tire Wear** — Per-wheel wear percentage
- **Tire Temperature** — Inner, outer, and surface temperatures per wheel

F1 25 broadcasts multiple UDP packet types (motion, session, lap data, car telemetry, car status, and more). RacePulse aggregates all relevant packets and merges them into a unified 60Hz stream for smooth, responsive overlays.

## How UDP Telemetry Works

F1 25 uses the same UDP telemetry system as previous Codemasters/EA F1 titles. When enabled, the game sends telemetry packets to the IP address and port set in its Telemetry Settings. RacePulse listens on ports **22025**, **22024**, **22023**, and **20777**, and connects on whichever one the game sends to.

RacePulse names the game by the UDP format it receives, so F1 25 set to the 2024 format shows as F1 24. Formats older than 2024 aren't read.

The telemetry broadcast is one-way — the game sends data, and RacePulse receives it. No racing telemetry is sent back to the game or uploaded to RacePulse servers.

Each packet type arrives at different rates, and RacePulse normalizes them into a consistent stream. This means your overlays update smoothly even though different data fields arrive at different frequencies.

## Customization

All widgets are independently customizable:

- Choose KPH or MPH for speed display
- Toggle individual UI elements on or off
- Set per-widget frame rates (60, 30, or 15 FPS)
- Drag widgets to any position on screen while arranging, then lock them for racing
- Settings persist automatically

## Why RacePulse for F1 25

- **Free** — No subscription, no trial, no locked features
- **Single executable** — 19MB `.exe`, no installer or dependencies
- **One-time setup** — Enable UDP telemetry once in F1 25, done forever
- **Lightweight** — Minimal resource footprint alongside the game
- **Multi-sim** — Same app also works with iRacing, ACC, Assetto Corsa, Le Mans Ultimate, and F1 24

---

[Download RacePulse](/install/) and set up your F1 25 telemetry overlays in under a minute. See all widgets on the [features page](/features/) or check the [changelog](/changelog/) for the latest updates.
