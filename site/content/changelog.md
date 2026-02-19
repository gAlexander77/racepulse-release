---
title: "Changelog"
subtitle: "Release history and notable changes for RacePulse sim racing telemetry overlays."
description: "Track RacePulse release history, feature additions, fixes, and telemetry overlay improvements."
ogImage: "assets/application-preview.gif"
---

## v0.1.2-alpha <span class="changelog-date">2026-02-16</span>

### Added

- Pulse icon branding across the app

---

## v0.1.1-alpha <span class="changelog-date">2026-02-15</span>

### Added

- Hover over any button or control to see what it does
- Trying to enable a widget without a game connected now shows a warning message
- Disabled toggles explain why they're disabled instead of silently ignoring clicks

---

## v0.1.0-alpha <span class="changelog-date">2026-02-14</span>

### Added

- iRacing support via Windows shared memory (auto-detected)
- F1 25 support via UDP telemetry (port 22025, auto-detected)
- F1 24 support via UDP telemetry (port 22024, auto-detected)
- Overlay widget system with independent, always-on-top widget windows
- Core widgets: Speed, Gear, RPM, Lap Time, Delta, Inputs, Fuel, G-Force
- Advanced widgets: Telemetry Core (combined display), Inputs Telemetry (live input graph)
- Per-widget customization with persistent settings
- Configurable frame rate per widget (60, 30, 15 FPS)
- Edit mode for widget positioning, click-through mode for racing
- Session recording to local storage
- Auto game detection (scans for running games every 3 seconds)
- Single executable distribution (~21MB, no installer required)
- Live widget previews on dashboard

### Technical

- Unified telemetry model supporting multiple games through plugin architecture
- F1 provider handles both F1 24 and F1 25 via single codebase (packet format detection)
- Auto port discovery for F1 games (scans common ports in parallel)
- Widget overlay processes communicate via TCP IPC
