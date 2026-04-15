---
title: "Free Le Mans Ultimate Overlay"
subtitle: "Lightweight telemetry overlay for Le Mans Ultimate with automatic shared memory detection, 12 customizable widgets, and zero configuration."
description: "Free Le Mans Ultimate telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, and track map widgets for LMU. No plugins to install."
ogImage: "assets/lmu-overlay-preview.gif"
---

## Free Telemetry Overlay for Le Mans Ultimate

RacePulse is a free telemetry overlay for Le Mans Ultimate. It reads telemetry directly from LMU's rFactor 2 shared memory plugin — the same interface used by the sim engine to expose physics, scoring, and vehicle data. No additional plugins, no web dashboards, no configuration files.

<div style="margin: 2rem 0; border-radius: 0.75rem; overflow: hidden; border: 1px solid #2A2D3A;">
  <img src="/assets/lmu-overlay-preview.gif" alt="RacePulse telemetry overlay running in Le Mans Ultimate showing speed, gear, RPM, inputs, and lap time widgets" style="width: 100%; display: block;">
</div>

Le Mans Ultimate is built on the rFactor 2 engine, which exposes telemetry through a shared memory plugin. RacePulse reads this data at up to 50Hz for physics and 5Hz for scoring, giving you real-time overlays with virtually zero setup.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. In LMU, go to **Settings > Gameplay** and ensure **Enable Plugins** is turned on
4. Launch a session in LMU
5. RacePulse automatically detects the shared memory and begins streaming telemetry

> **Important:** The rFactor 2 shared memory plugin must be enabled in LMU. Go to **Settings > Gameplay > Enable Plugins** and make sure it is turned on. This is the only configuration step required.

## Available Widgets for Le Mans Ultimate

LMU supports 11 of the 12 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in km/h or mph
- **Gear** — Large gear indicator for peripheral visibility
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to best lap using sector split comparison
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel
- **G-Force** — 2D lateral and longitudinal G-force visualization

### Advanced Widgets

- **Telemetry Core** — Combined display with rev lights, gear, RPM, speed, position, track temperature, elapsed time, and lap count
- **Inputs Telemetry** — Scrolling input graph with pedal bars, optional steering wheel, gear, and speed
- **Flat Track Map** — Horizontal bar showing all cars on track with class-colored markers and your position highlighted

### Not Available

- **Relative** — Not available for LMU (iRacing exclusive, requires iRating and license class data)

## LMU-Specific Telemetry

RacePulse captures data specific to the rFactor 2 / LMU engine:

- **ERS Battery** level (for hybrid prototypes)
- **Boost Torque** from hybrid systems
- **Turbo Boost** pressure
- **Competitor gaps** — delta to leader and delta to car ahead from scoring data
- **Multi-class support** — class positions and class-colored track map markers

The Delta widget uses a multi-tier approach for LMU: sector split comparisons after sector boundaries for accuracy, and proportional estimation during the first sector for real-time feedback.

## Dual-Buffer Telemetry

LMU's rFactor 2 engine exposes two separate shared memory buffers:

- **Telemetry buffer** (~50Hz) — Per-vehicle physics data including speed, G-forces, tyre data, fuel, inputs, engine data
- **Scoring buffer** (~5Hz) — Session information, standings, gaps, lap times, sector times, pit status

RacePulse reads both buffers and merges the data into a unified stream. The higher-frequency telemetry buffer provides smooth animations for speed, RPM, and input widgets, while the scoring buffer provides accurate timing and position data.

## Customization

Every widget is independently configurable:

- Metric or imperial units
- Toggle individual display elements
- Configurable frame rate (60, 30, or 15 FPS)
- Drag-to-reposition in Edit Mode
- Persistent settings across sessions

## Why RacePulse for Le Mans Ultimate

- **Free** — No subscription, no paid tier
- **Single executable** — 21MB `.exe`, no installer required
- **Minimal setup** — Just enable the LMU plugin once, then forget about it
- **Endurance-ready** — Fuel estimation and multi-class track map for long races
- **Multi-sim** — Same app works with iRacing, ACC, Assetto Corsa, F1 25, and F1 24

---

[Download RacePulse](/install/) and set up your Le Mans Ultimate overlays in under a minute. View all widgets on the [features page](/features/) or check the [changelog](/changelog/) for what's new.
