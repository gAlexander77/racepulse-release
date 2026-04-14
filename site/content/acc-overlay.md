---
title: "Free ACC Overlay"
subtitle: "Lightweight telemetry overlay for Assetto Corsa Competizione with automatic shared memory detection, 12 customizable widgets, and zero configuration."
description: "Free ACC telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, and track map widgets for Assetto Corsa Competizione. No plugins needed."
ogImage: "assets/application-preview.gif"
---

## Free Telemetry Overlay for Assetto Corsa Competizione

RacePulse is a free, standalone telemetry overlay for Assetto Corsa Competizione. It reads telemetry directly from ACC's shared memory buffers — the same interface the game uses to expose physics, graphics, and session data. No plugins, no SimHub dashboards, no configuration. Download and race.

Unlike overlay tools that require importing dashboard layouts or installing additional software, RacePulse works immediately. ACC exposes three shared memory buffers (Physics, Graphics, and Static) whenever the sim is running, and RacePulse reads them directly at up to 60Hz.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch ACC and enter a session
4. RacePulse automatically detects ACC and begins streaming telemetry

No plugins to enable, no ports to configure. ACC's shared memory is always available when the game is running.

> **Tip:** For best results, run ACC in **Borderless Windowed** mode. This allows RacePulse widgets to float over the game without alt-tab interruptions.

## Available Widgets for ACC

ACC supports 11 of the 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in km/h or mph
- **Gear** — Large gear indicator for peripheral visibility
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to best lap using ACC's built-in delta calculation
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel
- **G-Force** — 2D lateral and longitudinal G-force visualization

### Advanced Widgets

- **Telemetry Core** — Combined display with rev lights, gear, RPM, speed, position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, optional steering wheel, gear, and speed

### Partial Support

- **Flat Track Map** — Not currently available for ACC (requires per-car coordinate data that ACC's shared memory does not expose in a compatible format)
- **Relative** — Not available for ACC (iRacing exclusive, requires iRating and license data)

## ACC-Specific Telemetry

RacePulse captures ACC-specific data beyond standard telemetry:

- **Traction Control** level and activation
- **ABS** level and activation
- **Engine Map** setting
- **Brake Bias** front/rear percentage
- **Fuel Per Lap** estimation
- **Rain Intensity** and forecast
- **Track Grip** status (Green, Fast, Optimum, etc.)
- **Weather data** — ambient temperature, track temperature, wind speed, rain level

The Delta widget uses ACC's native `iDeltaLapTime` field, giving you accurate delta-to-best directly from the sim's own calculation.

## Customization

Every widget can be individually customized:

- Metric or imperial units per widget
- Toggle display elements on or off
- Configurable frame rate (60, 30, or 15 FPS)
- Drag-to-reposition in Edit Mode
- Settings persist automatically between sessions

## How It Works

ACC exposes telemetry through three Windows shared memory buffers:

- **Physics** (~60Hz) — speed, G-forces, tyre data, fuel, inputs, temperatures
- **Graphics** (~5Hz) — session info, positions, gaps, delta, flags, weather
- **Static** (once on connect) — car model, track name, max RPM, pit window

RacePulse reads these buffers directly using Windows API calls. The connection is read-only and adds minimal overhead. Each widget runs as a separate transparent window that floats over ACC, passing through mouse clicks so they never interfere with the game.

## Why RacePulse for ACC

- **Free** — No subscription, no premium tier, no feature locks
- **Single executable** — 21MB `.exe`, no installer, no runtime dependencies
- **Zero configuration** — No SimHub import, no plugin setup, no dashboard files
- **Lightweight** — Designed to run alongside demanding sims without impacting FPS
- **Multi-sim support** — The same app works with iRacing, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24

---

Get started in under a minute. [Download RacePulse](/install/), launch ACC, and your overlays appear automatically. See the full [feature list](/features/) or check the [changelog](/changelog/) for recent updates.
