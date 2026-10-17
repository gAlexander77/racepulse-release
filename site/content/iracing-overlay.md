---
title: "Free iRacing Overlay"
subtitle: "Lightweight telemetry overlay for iRacing with automatic shared memory detection, 14 customizable widgets, and zero configuration."
description: "Free iRacing telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, relative, and track map widgets. No plugins, no configuration — just download and race."
ogImage: "assets/iracing-overlay-preview.jpg"
---

## Free Telemetry Overlay for iRacing

RacePulse is a free, standalone telemetry overlay built for iRacing. It reads telemetry directly from iRacing's shared memory interface — no plugins, no configuration files, and no third-party dependencies. Download the single `.exe`, launch it alongside iRacing, and your overlays appear on screen automatically.

<div style="margin: 2rem 0; overflow: hidden; border: 1px solid #232323;">
  <video playsinline autoplay muted loop preload="metadata" poster="/assets/iracing-overlay-preview.jpg" style="width: 100%; display: block;" aria-label="RacePulse telemetry overlay running in iRacing showing speed, gear, RPM, inputs, and lap time widgets">
    <source src="/assets/iracing-overlay-preview.mp4" type="video/mp4">
    <img src="/assets/iracing-overlay-preview.gif" alt="RacePulse telemetry overlay running in iRacing showing speed, gear, RPM, inputs, and lap time widgets" style="width: 100%; display: block; margin: 0; border: 0;">
  </video>
</div>

Unlike overlay setups that require plugin installation and dashboard configuration, RacePulse works out of the box. There is nothing to configure inside iRacing — shared memory is always available when the sim is running.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch iRacing and enter a session
4. RacePulse automatically detects iRacing and begins streaming telemetry

That's it. No shared memory plugins to install, no ports to configure, no config files to edit. iRacing exposes telemetry via Windows shared memory by default, and RacePulse reads it directly.

> **Tip:** Run iRacing in **Borderless Windowed** mode for the best overlay experience. This allows widgets to float on top of the sim without alt-tab issues.

## Available Widgets for iRacing

iRacing has full support for all 14 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in KPH or MPH, with optional session top speed and pit limiter state
- **Gear** — Large, bold gear indicator visible in peripheral vision
- **RPM** — The number, or a row of rev lights that scale to each car's redline
- **Lap Times** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to your best lap (green = faster, red = slower), with dashes while iRacing marks the delta invalid
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel, plus optional fuel to finish and fuel to add
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — All-in-one display with animated rev lights, gear, and readouts you choose, such as speed, RPM, race position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs with pedal bars, gear, speed, and a steering wheel that uses your car's own steering lock. ABS can color the brake line.
- **Flat Track Map** — Every car on one line of the lap, with class-colored markers, your position highlighted, and optional faster-class warnings and who's talking on the radio
- **Relative** — Nearby drivers sorted by track proximity with live gap times, iRating, safety rating, license class, country flags, and pit/out lap indicators. Gaps between classes are measured on your class's lap, and a headset marks the driver talking on the radio.
- **Standings** — The whole field by class in six layouts, with strength of field per class, gaps, best and last laps, pit stops, stints, and positions gained
- **Lap Log** — Every lap you drive with its delta, under the session best and your best

> **Note:** iRating, safety rating, license class, and who's talking are iRacing-only parts of Relative and Standings. Country flags come from each driver's profile flair. Projected iRating change is in beta.

## Customization

Every widget is independently customizable:

- Choose between metric and imperial units
- Toggle individual display elements on or off
- Set per-widget frame rates (60, 30, or 15 FPS)
- Drag widgets to any position while arranging, then lock them for racing
- Save widget layouts for different cars or series
- All settings persist across sessions automatically

<div style="margin: 2rem 0; overflow: hidden; border: 1px solid #232323;">
  <video playsinline autoplay muted loop preload="metadata" poster="/assets/widget-customization-preview.jpg?v=2026-10-12" width="1600" height="1000" style="width: 100%; height: auto; display: block;" aria-label="RacePulse widget customization: changing Inputs Telemetry settings updates its live preview">
    <source src="/assets/widget-customization-preview.mp4?v=2026-10-12" type="video/mp4">
    <img src="/assets/widget-customization-preview.gif?v=2026-10-12" alt="RacePulse widget customization: changing Inputs Telemetry settings updates its live preview" width="1000" height="626" loading="lazy" style="width: 100%; height: auto; display: block; margin: 0; border: 0;">
  </video>
</div>

## How It Works

RacePulse connects to iRacing's shared memory interface (`irsdk`) using direct Windows API calls. This is the same interface other iRacing tools read. The connection is read-only — RacePulse never writes to shared memory or modifies any iRacing state.

Telemetry is captured at up to 60Hz and normalized into a unified format. Each widget runs as an independent, transparent, always-on-top window that floats over your sim. Widgets are click-through during racing so they never interfere with mouse input.

## Why RacePulse for iRacing

- **Free** — No subscription, no trial period, no feature gates
- **Single executable** — One 19MB `.exe` file, no installer, no dependencies
- **Zero configuration** — No plugins to install, no ports to open, no JSON files to edit
- **Lightweight** — Minimal CPU and memory footprint designed for racing
- **Multi-sim support** — The same app also works with ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24 if you race across multiple sims

---

Ready to try it? [Download RacePulse](/install/) and have your overlays running in under a minute. Check out the full [feature list](/features/) or see what's new in the [changelog](/changelog/).
