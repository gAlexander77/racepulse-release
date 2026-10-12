---
title: "Free Le Mans Ultimate Overlay"
subtitle: "Lightweight telemetry overlay for Le Mans Ultimate with automatic shared memory detection, 14 customizable widgets, and zero configuration."
description: "Free Le Mans Ultimate telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, and track map widgets for LMU. No plugins to install."
ogImage: "assets/lmu-overlay-preview.jpg"
---

## Free Telemetry Overlay for Le Mans Ultimate

RacePulse is a free telemetry overlay for Le Mans Ultimate. It reads LMU's built-in shared memory interface — the official Studio 397 API that ships with the game. No plugins to install, no web dashboards, no configuration files.

<div style="margin: 2rem 0; overflow: hidden; border: 1px solid #232323;">
  <video playsinline autoplay muted loop preload="metadata" poster="/assets/lmu-overlay-preview.jpg" style="width: 100%; display: block;" aria-label="RacePulse telemetry overlay running in Le Mans Ultimate showing speed, gear, RPM, inputs, and lap time widgets">
    <source src="/assets/lmu-overlay-preview.mp4" type="video/mp4">
    <img src="/assets/lmu-overlay-preview.gif" alt="RacePulse telemetry overlay running in Le Mans Ultimate showing speed, gear, RPM, inputs, and lap time widgets" style="width: 100%; display: block; margin: 0; border: 0;">
  </video>
</div>

Le Mans Ultimate ships with an official shared memory interface that exposes the sim engine's physics, scoring, and vehicle data. RacePulse reads this data at up to 50Hz for physics and 5Hz for scoring, giving you real-time overlays with virtually zero setup.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases) and run it (no installation needed)
2. In LMU, go to **Settings > Gameplay** and turn **Enable Plugins** on
3. Load into a session (practice, qualifying, or race)
4. RacePulse automatically detects the shared memory and begins streaming telemetry

That's the whole setup. RacePulse reads the shared memory interface built into Le Mans Ultimate itself, so there is no plugin DLL to download or install.

> **Troubleshooting:** if RacePulse can't connect to LMU, check that **Settings > Gameplay > Enable Plugins** is on and that you're loaded into a session (not sitting in the main menu). If you're on a very old LMU version without the built-in interface, update the game.

## Available Widgets for Le Mans Ultimate

LMU supports all 14 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed in KPH or MPH
- **Gear** — Large gear indicator for peripheral visibility
- **RPM** — The number, or a row of rev lights
- **Lap Times** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to best lap using sector split comparison
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel, plus optional fuel to finish and fuel to add
- **G-Force** — 2D lateral and longitudinal G-force visualization

### Advanced Widgets

- **Telemetry Core** — Combined display with rev lights, gear, and readouts you choose, such as speed, RPM, position, class position, time left, ERS, and track temperature
- **Inputs Telemetry** — Scrolling input graph with pedal bars, gear, speed, and a steering wheel that uses your car's own steering lock
- **Flat Track Map** — Every car on one line of the lap, with class-colored markers, a lane per class if you like, and your position highlighted
- **Relative** — Nearby drivers sorted by gap on track, with class colors, car numbers, country flags, and pit/out lap indicators. From the pits it follows the best-placed car on track until you head out.
- **Standings** — The field by class in six layouts, with class headers, gaps, laps, and pit stops
- **Lap Log** — Every lap you drive with its delta and sectors, under the session best and your best

> **Note:** On LMU online servers, Relative and Standings show each driver's country flag. LMU sends no nationality in single-player sessions, so those show no flags.

## LMU-Specific Telemetry

RacePulse captures data specific to the LMU engine:

- **ERS Battery** level (for hybrid prototypes)
- **Boost Torque** from hybrid systems
- **Turbo Boost** pressure
- **Competitor gaps** — delta to leader and delta to car ahead from scoring data
- **Multi-class support** — class positions and class-colored track map markers
- **Car numbers, team names, and countries** — read from LMU's own local web interface on your PC while the game is running, so custom-team cars show their real numbers. Countries are only available on online servers.

The Delta widget uses a multi-tier approach for LMU: sector split comparisons after sector boundaries for accuracy, and proportional estimation during the first sector for real-time feedback.

## Dual-Rate Telemetry

LMU's shared memory carries two streams updated at different rates:

- **Telemetry** (~50Hz) — Per-vehicle physics data including speed, G-forces, tire data, fuel, inputs, engine data
- **Scoring** (~5Hz) — Session information, standings, gaps, lap times, sector times, pit status

RacePulse reads both and merges the data into a unified stream. The higher-frequency telemetry provides smooth animations for speed, RPM, and input widgets, while scoring provides accurate timing and position data.

## Customization

Every widget is independently configurable:

- Metric or imperial units
- Toggle individual display elements
- Configurable frame rate (60, 30, or 15 FPS)
- Drag-to-reposition while arranging, then lock for racing
- Persistent settings across sessions

## Why RacePulse for Le Mans Ultimate

- **Free** — No subscription, no paid tier
- **Single executable** — 16MB `.exe`, no installer required
- **One-time setup** — Turn on Enable Plugins in LMU once, then forget about it
- **Endurance-ready** — Fuel estimation, multi-class standings, and a multi-class track map for long races
- **Multi-sim** — Same app works with iRacing, ACC, Assetto Corsa, F1 25, and F1 24

---

[Download RacePulse](/install/) and set up your Le Mans Ultimate overlays. View all widgets on the [features page](/features/) or check the [changelog](/changelog/) for what's new.
