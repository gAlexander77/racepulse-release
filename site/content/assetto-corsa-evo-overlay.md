---
title: "Free Assetto Corsa EVO Overlay"
subtitle: "Lightweight telemetry overlay for Assetto Corsa EVO with automatic shared memory detection, customizable widgets, and zero configuration."
description: "Free Assetto Corsa EVO telemetry overlay with real-time speed, RPM, inputs, lap times, delta, fuel, G-force, lap log, and LED flag panel widgets. No plugins, no mods — just download and drive."
ogImage: "assets/assetto-corsa-evo-overlay-preview.jpg"
videoUploadDate: "2026-11-25"
---

## Free Telemetry Overlay for Assetto Corsa EVO

RacePulse is a free telemetry overlay for Assetto Corsa EVO. It reads the shared memory EVO publishes for external tools, and shows your speed, RPM, inputs, lap times, delta, fuel, and more as transparent widgets over the game.

<div style="margin: 2rem 0; overflow: hidden; border: 1px solid #232323;">
  <video playsinline autoplay muted loop preload="metadata" poster="/assets/assetto-corsa-evo-overlay-preview.jpg" style="width: 100%; display: block;" aria-label="The RacePulse Inputs Telemetry widget showing throttle, brake, gear, speed, and steering as an overlay in Assetto Corsa EVO">
    <source src="/assets/assetto-corsa-evo-overlay-preview.mp4" type="video/mp4">
    <img src="/assets/assetto-corsa-evo-overlay-preview.gif" alt="The RacePulse Inputs Telemetry widget showing throttle, brake, gear, speed, and steering as an overlay in Assetto Corsa EVO" style="width: 100%; display: block; margin: 0; border: 0;">
  </video>
</div>

No mods, no in-game apps, no dashboards to import. EVO's shared memory is built into the game, so there is nothing to install or turn on.

## Setup

1. Download `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases)
2. Run `RacePulse.exe` (no installation needed)
3. Launch Assetto Corsa EVO and load into a session
4. RacePulse automatically detects EVO and begins streaming telemetry

EVO publishes its telemetry from game version 0.6 onwards, and only once a session is loaded, so RacePulse finds the game when you're on track, not in the main menu.

**Last tested:** RacePulse's Assetto Corsa EVO support was last tested and confirmed with EVO **v0.9.1**.

> **Tip:** Use **Borderless Windowed** mode in EVO for the best overlay experience. Widgets float on top of the game without interrupting fullscreen rendering.

## Available Widgets for Assetto Corsa EVO

Assetto Corsa EVO supports 12 of the 15 RacePulse widgets:

### Core Widgets

- **Speed** — Live speed readout in KPH or MPH
- **Gear** — Large, bold gear indicator
- **RPM** — The number, or a row of rev lights
- **Lap Times** — Current lap, last lap, and personal best, with an optional strike-through when EVO invalidates your lap
- **Delta** — Live delta to your best lap, from EVO's own calculation
- **Inputs** — Throttle, brake, and clutch as vertical bars
- **Fuel** — Fuel remaining with estimated laps of fuel
- **G-Force** — 2D lateral and longitudinal G-force dot

### Advanced Widgets

- **Telemetry Core** — All-in-one display with rev lights, gear, and readouts you choose, such as speed, RPM, position, brake bias, TC and ABS settings, water temperature, and laps of fuel. EVO doesn't send elapsed session time, so that reading shows a dash.
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs with pedal bars, gear, speed, and a steering wheel that turns as far as your car's wheel does. ABS can color the brake line.
- **Lap Log** — Every lap you drive with its delta and lap validity, under your best lap. EVO doesn't send sector times, so Lap Log has no sector columns in EVO.
- **Flag Panel** — An LED flag box that lights with EVO's flag: green, yellow, blue, white, red, black, the meatball, debris (red and yellow stripes), and checkered. Run a second one with Flag Panel 2.

### Not Available

- **Flat Track Map**, **Relative**, and **Standings** — EVO's shared memory doesn't include other drivers' names, laps, or times, so these aren't offered. If one is open when you connect EVO, it hides with a notice and comes back when you switch to a supported sim.

## EVO-Specific Data

RacePulse reads EVO-specific telemetry beyond the basics:

- **Brake Bias** — front share as a percentage
- **Traction Control** and **ABS** settings, and when ABS is working
- **Fuel Per Lap** and estimated laps of fuel
- **Steering angle** in degrees, so the Inputs Telemetry wheel turns as far as the car's
- **Flags** — green, yellow, blue, white, red, black, meatball, checkered, and debris

## Customization

Every widget is independently customizable:

- Choose between KPH and MPH
- Toggle display elements on or off
- Configurable frame rate (60, 30, or 15 FPS)
- Reposition widgets while arranging, then lock them for racing
- All settings persist across sessions

## How It Works

Assetto Corsa EVO exposes telemetry through three Windows shared memory buffers of its own, separate from those of Assetto Corsa and ACC:

- **Physics** — speed, inputs, G-forces, tire data, fuel, temperatures
- **Graphics** — lap times, delta, position, flags, session state, and electronics settings
- **Static** — track, session type, and track length

RacePulse only reads the buffers the game creates; it never creates or changes them. Each widget is a separate transparent, always-on-top, click-through window. There is no performance impact on the game beyond standard shared memory reads.

## Why RacePulse for Assetto Corsa EVO

- **Free** — No subscription, no paid features
- **Single executable** — One 20MB file, no installer needed
- **No mods** — Nothing to install inside EVO, no in-game apps, no dashboards to import
- **Lightweight** — Minimal resource usage, won't impact your FPS
- **Multi-sim** — Same overlay app works across iRacing, ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24

---

[Download RacePulse](/install/) and have your Assetto Corsa EVO overlays running in under a minute. See all available widgets on the [features page](/features/) or check the [changelog](/changelog/) for updates.
