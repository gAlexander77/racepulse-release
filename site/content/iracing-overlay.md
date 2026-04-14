---
title: "Free iRacing Overlay"
subtitle: "Lightweight telemetry overlay for iRacing with automatic shared memory detection, 12 customizable widgets, and zero configuration."
description: "Free iRacing telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, relative, and track map widgets. No plugins, no configuration — just download and race."
ogImage: "assets/iracing-overlay-preview.gif"
---

## Free Telemetry Overlay for iRacing

RacePulse is a free, standalone telemetry overlay built for iRacing. It reads telemetry directly from iRacing's shared memory interface — no plugins, no configuration files, and no third-party dependencies. Download the single `.exe`, launch it alongside iRacing, and your overlays appear on screen automatically.

<div style="margin: 2rem 0; border-radius: 0.75rem; overflow: hidden; border: 1px solid #2A2D3A;">
  <img src="/assets/iracing-overlay-preview.gif" alt="RacePulse telemetry overlay running in iRacing showing speed, gear, RPM, inputs, and lap time widgets" style="width: 100%; display: block;">
</div>

Unlike SimHub-based overlays that require plugin installation and dashboard configuration, RacePulse works out of the box. There is nothing to configure inside iRacing — shared memory is always available when the sim is running.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch iRacing and enter a session
4. RacePulse automatically detects iRacing and begins streaming telemetry

That's it. No shared memory plugins to install, no ports to configure, no config files to edit. iRacing exposes telemetry via Windows shared memory by default, and RacePulse reads it directly.

> **Tip:** Run iRacing in **Borderless Windowed** mode for the best overlay experience. This allows widgets to float on top of the sim without alt-tab issues.

## Available Widgets for iRacing

iRacing has full support for all 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in km/h or mph
- **Gear** — Large, bold gear indicator visible in peripheral vision
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to your best lap (green = faster, red = slower)
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — All-in-one display with animated rev lights, gear, RPM, speed, race position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs with pedal bars, optional steering visualization, gear, and speed readout
- **Flat Track Map** — Horizontal bar showing all cars on track with class-colored markers and your position highlighted
- **Relative** — Nearby drivers sorted by track proximity with live gap times, iRating, safety rating, license class, and pit/out lap indicators

> **Note:** The Relative widget is exclusive to iRacing. It uses iRacing's detailed competitor data to show iRating, safety rating, and license class alongside gap times.

## Customization

Every widget is independently customizable:

- Choose between metric and imperial units
- Toggle individual display elements on or off
- Set per-widget frame rates (60, 30, or 15 FPS)
- Drag widgets to any position in Edit Mode
- All settings persist across sessions automatically

<div style="margin: 2rem 0; border-radius: 0.75rem; overflow: hidden; border: 1px solid #2A2D3A;">
  <img src="/assets/widget-customization-preview.gif" alt="RacePulse widget customization panel showing per-widget settings for iRacing overlays" style="width: 100%; display: block;">
</div>

## How It Works

RacePulse connects to iRacing's shared memory interface (`irsdk`) using direct Windows API calls. This is the same interface used by tools like iSpeed, Kapps, and Trading Paints. The connection is read-only — RacePulse never writes to shared memory or modifies any iRacing state.

Telemetry is captured at up to 60Hz and normalized into a unified format. Each widget runs as an independent, transparent, always-on-top window that floats over your sim. Widgets are click-through during racing so they never interfere with mouse input.

## Why RacePulse for iRacing

- **Free and open source** — No subscription, no trial period, no feature gates
- **Single executable** — One 21MB `.exe` file, no installer, no dependencies
- **Zero configuration** — No plugins to install, no ports to open, no JSON files to edit
- **Lightweight** — Minimal CPU and memory footprint designed for racing
- **Multi-sim support** — The same app also works with ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24 if you race across multiple sims

---

Ready to try it? [Download RacePulse](/install/) and have your overlays running in under a minute. Check out the full [feature list](/features/) or see what's new in the [changelog](/changelog/).
