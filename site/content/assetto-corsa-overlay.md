---
title: "Free Assetto Corsa Overlay"
subtitle: "Lightweight telemetry overlay for Assetto Corsa with automatic shared memory detection, customizable widgets, and zero configuration."
description: "Free Assetto Corsa telemetry overlay with real-time speed, RPM, inputs, lap times, fuel, and G-force widgets. No plugins, no launcher setup — just download and drive."
ogImage: "assets/application-preview.jpg"
---

## Free Telemetry Overlay for Assetto Corsa

RacePulse is a free telemetry overlay for Assetto Corsa. It connects directly to AC's shared memory buffers to display real-time speed, RPM, inputs, lap times, fuel, and more as transparent overlay widgets floating over the game.

No launcher plugins, no dashboards to import, no in-game apps. RacePulse reads AC's shared memory natively — the same interface that AC exposes for any external tool. Download, launch, and drive.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch Assetto Corsa and enter a session
4. RacePulse automatically detects AC and begins streaming telemetry

No setup inside Assetto Corsa or its launcher is required. AC is detected once a session is loaded; its main menu runs in a separate launcher, so RacePulse can't see the game there.

> **Tip:** Use **Borderless Windowed** mode in AC for the best overlay experience. Widgets float on top of the game without interrupting fullscreen rendering.

## Available Widgets for Assetto Corsa

Assetto Corsa supports 10 of the 14 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed readout in KPH or MPH
- **Gear** — Large, bold gear indicator
- **RPM** — The number, or a row of rev lights
- **Lap Times** — Current lap, last lap, and personal best
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — All-in-one display with rev lights, gear, and readouts you choose, such as speed, RPM, position, track temperature, and lap count. AC doesn't send elapsed session time, so that reading shows a dash.
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs with pedal bars, a steering wheel, gear, and speed
- **Lap Log** — Every lap you drive with its delta and sectors, under your best lap, with the track's readable name

### Not Available

- **Delta** — AC does not send a live delta to your best lap, so the Delta widget isn't offered. The Lap Log still compares each finished lap with your best.
- **Flat Track Map** — Not supported for AC (requires competitor data in a compatible format)
- **Relative** and **Standings** — Not supported for AC, for the same reason. If one is open when you connect AC, it hides with a notice and comes back when you switch to a supported sim.

## AC-Specific Data

RacePulse captures Assetto Corsa-specific telemetry fields:

- **Brake Bias** — front/rear distribution percentage
- **Turbo Boost** — turbo pressure for forced-induction cars

AC uses three shared memory buffers (Physics, Graphics, Static) with the same memory map names as ACC but with different struct layouts. RacePulse automatically distinguishes between AC and ACC by checking which game process is running.

## Customization

Every widget is independently customizable:

- Choose between KPH and MPH
- Toggle display elements on or off
- Configurable frame rate (60, 30, or 15 FPS)
- Reposition widgets while arranging, then lock them for racing
- All settings persist across sessions

## How It Works

Assetto Corsa exposes telemetry through three Windows shared memory buffers:

- **Physics** (~60Hz) — speed, inputs, G-forces, tire data, fuel, temperatures
- **Graphics** (~5Hz) — session info, positions, flags, lap info
- **Static** (once on connect) — car name, track name, max RPM

RacePulse reads these buffers directly. Each widget is a separate transparent, always-on-top, click-through window. There is no performance impact on the game beyond standard shared memory reads.

## Why RacePulse for Assetto Corsa

- **Free** — No subscription, no paid features
- **Single executable** — One 20MB file, no installer needed
- **VR overlays** — Your widgets in your headset through SteamVR (experimental). See the [VR Guide](/vr/)
- **No plugin setup** — No launcher apps, no scripts, no dashboards to import
- **Lightweight** — Minimal resource usage, won't impact your FPS
- **Multi-sim** — Same overlay app works across iRacing, ACC, Le Mans Ultimate, F1 25, and F1 24

---

[Download RacePulse](/install/) and have your Assetto Corsa overlays running in under a minute. See all available widgets on the [features page](/features/) or check the [changelog](/changelog/) for updates.
