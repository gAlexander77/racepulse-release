---
title: "Free F1 25 Overlay"
subtitle: "Lightweight telemetry overlay for F1 25 with automatic UDP detection, 11 customizable widgets, and simple one-time setup."
description: "Free F1 25 telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, and track map widgets. Enable UDP telemetry and go."
ogImage: "assets/application-preview.gif"
---

## Free Telemetry Overlay for F1 25

RacePulse is a free telemetry overlay for EA Sports F1 25. It listens for F1 25's UDP telemetry broadcast to display real-time speed, RPM, inputs, lap times, delta, fuel, and more as transparent overlay widgets floating over your game.

Setup takes less than a minute: enable UDP telemetry in F1 25's settings, launch RacePulse, and your overlays appear automatically. No additional software, no web dashboards, no JSON configuration.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. In F1 25, go to **Settings > Telemetry Settings**
4. Set **UDP Telemetry** to **On**
5. Ensure **UDP Port** is set to **22025** (this is the default)
6. Start a session — RacePulse auto-detects F1 25 and begins displaying telemetry

> **First time only:** You need to enable UDP telemetry in F1 25 once. After that, the setting persists across game restarts. RacePulse scans for F1 25's telemetry broadcast on port 22025 automatically.

If Windows Firewall prompts you to allow RacePulse network access, click Allow. RacePulse only listens for local UDP packets — it does not send data anywhere.

## Available Widgets for F1 25

F1 25 supports 11 of the 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in km/h or mph
- **Gear** — Large gear indicator visible in peripheral vision
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to your best lap
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force visualization

### Advanced Widgets

- **Telemetry Core** — All-in-one display with rev lights, gear, RPM, speed, position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, optional steering wheel, gear, and speed
- **Flat Track Map** — Horizontal bar showing all 20 cars on track with your position highlighted

### Not Available

- **Relative** — Not available for F1 25 (iRacing exclusive)

## F1 25-Specific Data

RacePulse captures F1-specific telemetry beyond standard fields:

- **ERS** — Energy Recovery System deployment and harvest levels
- **DRS** — DRS availability and activation status
- **Fuel Mix** — Current fuel mixture setting
- **Tyre Compound** — Active tyre compound (soft, medium, hard, intermediate, wet)
- **Tyre Wear** — Per-wheel wear percentage
- **Tyre Temperature** — Inner, outer, and surface temperatures per wheel

F1 25 broadcasts multiple UDP packet types (motion, session, lap data, car telemetry, car status, and more). RacePulse aggregates all relevant packets and merges them into a unified 60Hz stream for smooth, responsive overlays.

## How UDP Telemetry Works

F1 25 uses the same UDP telemetry system as previous Codemasters/EA F1 titles. When enabled, the game broadcasts telemetry packets to `localhost` on port **22025** (the port number matches the game year).

RacePulse listens for these packets locally. The telemetry broadcast is one-way — the game sends data, and RacePulse receives it. No data is sent back to the game or to any external server.

Each packet type arrives at different rates, and RacePulse normalizes them into a consistent stream. This means your overlays update smoothly even though different data fields arrive at different frequencies.

## Customization

All widgets are independently customizable:

- Choose km/h or mph for speed display
- Toggle individual UI elements on or off
- Set per-widget frame rates (60, 30, or 15 FPS)
- Drag widgets to any position on screen in Edit Mode
- Settings persist automatically

## Why RacePulse for F1 25

- **Free** — No subscription, no trial, no locked features
- **Single executable** — 21MB `.exe`, no installer or dependencies
- **One-time setup** — Enable UDP telemetry once in F1 25, done forever
- **Lightweight** — Minimal resource footprint alongside the game
- **Multi-sim** — Same app also works with iRacing, ACC, Assetto Corsa, Le Mans Ultimate, and F1 24

---

[Download RacePulse](/install/) and set up your F1 25 telemetry overlays in under a minute. See all widgets on the [features page](/features/) or check the [changelog](/changelog/) for the latest updates.
