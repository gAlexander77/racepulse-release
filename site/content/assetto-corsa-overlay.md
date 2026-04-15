---
title: "Free Assetto Corsa Overlay"
subtitle: "Lightweight telemetry overlay for Assetto Corsa with automatic shared memory detection, customizable widgets, and zero configuration."
description: "Free Assetto Corsa telemetry overlay with real-time speed, RPM, inputs, lap times, fuel, and G-force widgets. No plugins, no Content Manager setup — just download and drive."
ogImage: "assets/application-preview.gif"
---

## Free Telemetry Overlay for Assetto Corsa

RacePulse is a free telemetry overlay for Assetto Corsa. It connects directly to AC's shared memory buffers to display real-time speed, RPM, inputs, lap times, fuel, and more as transparent overlay widgets floating over the game.

No Content Manager plugins, no SimHub dashboards, no Python apps. RacePulse reads AC's shared memory natively — the same interface that AC exposes for any external tool. Download, launch, and drive.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch Assetto Corsa and enter a session
4. RacePulse automatically detects AC and begins streaming telemetry

No setup inside Assetto Corsa or Content Manager is required. Shared memory is always available when AC is running.

> **Tip:** Use **Borderless Windowed** mode in AC for the best overlay experience. Widgets float on top of the game without interrupting fullscreen rendering.

## Available Widgets for Assetto Corsa

Assetto Corsa supports 10 of the 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed readout in km/h or mph
- **Gear** — Large, bold gear indicator
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps
- **G-Force** — 2D lateral and longitudinal G-force dot
- **Delta** — Delta time display (note: AC does not expose a native delta calculation via shared memory, so accuracy depends on RacePulse's internal best-lap comparison)

### Advanced Widgets

- **Telemetry Core** — All-in-one display with rev lights, gear, RPM, speed, position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs with pedal bars, optional steering wheel, gear, and speed

### Not Available

- **Flat Track Map** — Not supported for AC (requires competitor coordinate data in a compatible format)
- **Relative** — iRacing exclusive

## AC-Specific Data

RacePulse captures Assetto Corsa-specific telemetry fields:

- **Brake Bias** — front/rear distribution percentage
- **Turbo Boost** — turbo pressure for forced-induction cars

AC uses three shared memory buffers (Physics, Graphics, Static) with the same memory map names as ACC but with different struct layouts. RacePulse automatically distinguishes between AC and ACC by checking which game process is running.

## Customization

Every widget is independently customizable:

- Choose between km/h and mph
- Toggle display elements on or off
- Configurable frame rate (60, 30, or 15 FPS)
- Reposition widgets by entering Edit Mode
- All settings persist across sessions

## How It Works

Assetto Corsa exposes telemetry through three Windows shared memory buffers:

- **Physics** (~60Hz) — speed, inputs, G-forces, tyre data, fuel, temperatures
- **Graphics** (~5Hz) — session info, positions, flags, lap info
- **Static** (once on connect) — car name, track name, max RPM

RacePulse reads these buffers directly. Each widget is a separate transparent, always-on-top, click-through window. There is no performance impact on the game beyond standard shared memory reads.

## Why RacePulse for Assetto Corsa

- **Free** — No subscription, no paid features
- **Single executable** — One 21MB file, no installer needed
- **No plugin setup** — No Content Manager apps, no Python scripts, no SimHub
- **Lightweight** — Minimal resource usage, won't impact your FPS
- **Multi-sim** — Same overlay app works across iRacing, ACC, Le Mans Ultimate, F1 25, and F1 24

---

[Download RacePulse](/install/) and have your Assetto Corsa overlays running in under a minute. See all available widgets on the [features page](/features/) or check the [changelog](/changelog/) for updates.
