---
title: "Features"
subtitle: "10 overlay widgets, per-widget customization, multi-game support, and session recording (coming soon) for iRacing, F1 25, and F1 24."
description: "Explore RacePulse features: 10 overlay widgets, per-widget customization, and unified telemetry support for iRacing, F1 25, and F1 24."
ogImage: "assets/widget-customization-preview.gif"
---

## Supported Games

- **iRacing** — Automatic detection via Windows shared memory. No configuration needed.
- **F1 25** — UDP telemetry on port 22025 (enable in game settings under Telemetry).
- **F1 24** — UDP telemetry on port 22024 (enable in game settings under Telemetry).

RacePulse auto-detects the running game and begins capturing telemetry. Multiple game support means you can switch between sims without reconfiguring. All widgets work across all supported games through a unified telemetry model.

## Telemetry Capture

- Live telemetry ingestion at up to 60Hz
- Normalized telemetry model for consistent widget behavior across all supported games
- Real-time update pipeline for responsive overlays
- Automatic game detection — no manual game selection required

## Core Widgets

Standalone, always-on-top widget overlays that float over your game:

- **Speed** — Live speed readout
- **Gear** — Large, bold indicator visible in peripheral vision
- **RPM** — Horizontal bar with redline indicator
- **Lap Time** — Current lap, last lap, and personal best
- **Delta** — Real-time delta to best lap (green = faster, red = slower)
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel
- **G-Force** — 2D lateral and longitudinal G-force visualization

## Advanced Widgets

- **Telemetry Core** — All-in-one display combining animated rev lights, gear indicator, RPM, speed, race position, track temperature, elapsed time, and lap count. Supports metric and imperial units. Includes pit limiter flash indicator.
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs over time with color-coded traces, combined with real-time pedal bars. Optional steering wheel visualization, gear, and speed readout.

## Widget Customization

- Per-widget settings (units, visibility toggles, display options)
- Configurable frame rate (60, 30, or 15 FPS) per widget
- Advanced widgets have detailed toggle controls for each display element
- Conditional settings that grey out when their parent setting is disabled
- Persistent settings that survive app restarts
- Live preview of advanced widgets on the dashboard before enabling overlays

## Session Recording (Coming Soon)

- Telemetry sessions will be recorded to local storage automatically
- Captures session metadata (game, track, car, duration)
- Records lap times and validity
- Stores full telemetry frame history

## Quality of Life

- Single executable — no installer, no dependencies
- Compact distribution (~21MB)
- Widget positions and settings persist across sessions
- Edit mode for repositioning widgets, click-through mode for racing
