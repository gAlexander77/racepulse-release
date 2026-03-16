# Changelog

All notable changes to RacePulse releases are documented here.

## [v0.2.4-alpha] — 2026-03-16

### Added

- Game source selector replaced with visual icon grid showing colored badges for each supported game

### Fixed

- Edit mode banner no longer appears on startup when no widgets are open

### Changed

- Executable renamed from `racepulse.exe` to `RacePulse.exe`

---

## [v0.2.3-alpha] — 2026-03-14

### Added

- Edit mode indicator on each widget overlay — yellow badge shows "Edit Mode Enabled", changes to "Disable Edit Mode" on hover to disable edit mode globally
- Flat Track Map live preview in widget customization panel

---

## [v0.2.2-alpha] — 2026-03-14

### Added

- Flat Track Map widget — horizontal bar showing all cars on track with class-colored dots, player highlight, and position label (supports iRacing, LMU, and F1)
- iRacing competitor data support (car positions, classes, pit status via CarIdx arrays and session info YAML)
- Live preview for Flat Track Map widget with multi-class mock competitors
- Per-game widget support filtering — widgets can declare which games they support, with "Unsupported" badge and tooltip

### Fixed

- Fixed Flat Track Map preview glitch on initial load (state mismatch between simulation and initial telemetry)
- Fixed track map dot bounce when cars complete a lap (CSS calc and section transition timing)

---

## [v0.2.1-alpha] — 2026-03-09

### Added

- Background opacity setting for all widgets
- Graph background opacity setting for Inputs Telemetry widget
- Reset-to-default button with tooltip on number slider settings

### Improved

- Widgets now start in edit mode by default so they're immediately movable
- Revamped all basic widgets with new settings and polished styling
- Simplified RPM widget display
- Streamlined Fuel widget to focus on laps remaining estimation

---

## [v0.2.0-alpha] — 2026-03-08

### Fixed

- Fixed false iRacing detection on app startup when iRacing is not running
- Fixed LMU not reconnecting when joining a session after app is already open
- Fixed LMU not auto-disconnecting when the game is closed
- Fixed LMU input telemetry showing ECU-filtered values instead of actual pedal inputs

### Improved

- LMU throttle, brake, and clutch now show raw pedal inputs (no more phantom spikes from auto-blip, TC, or shift cuts)
- Removed unnecessary debug log file (widget-embed.log) from production builds

---

## [v0.1.9-alpha] — 2026-03-06

### Fixed

- Fixed telemetry provider not cleaning up properly on game disconnect

---

## [v0.1.8-alpha] — 2026-03-06

### Added

- App now prevents multiple instances from running at the same time

### Improved

- Telemetry Core widget no longer shifts when values change
- Improved update process

---

## [v0.1.7-alpha] — 2026-03-04

### Improved

- Redesigned home page layout — cleaner Telemetry Dashboard button and removed cluttered quick actions section

---

## [v0.1.6-alpha] — 2026-03-03

### Added

- Assetto Corsa Competizione support via shared memory (auto-detected, no configuration needed)
- Assetto Corsa support via shared memory (auto-detected, no configuration needed)

### Fixed

- Auto-update now correctly detects new pre-release versions

---

## [v0.1.5-alpha] — 2026-03-01

### Added

- RacePulse now checks for updates on startup and notifies you when a new version is available
- One-click update — download and install without leaving the app
- Updates keep your exe in the same location, no reinstall needed

---

## [v0.1.4-alpha] — 2026-02-27

### Added

- Le Mans Ultimate support via rFactor 2 shared memory (auto-detected, no configuration needed)
- Rounded corners setting for the Inputs Telemetry graph
- New app icon

### Fixed

- More accurate delta-to-best using sector split comparisons instead of estimated lap time
- Fuel now displays in consistent units across all games

---

## [v0.1.3-alpha] — 2026-02-24

### Fixed

- Fixed a crash that could happen when stopping iRacing telemetry capture
- Fixed iRacing telemetry reader freezing under heavy load
- Fixed widget processes not closing properly on app shutdown
- Fixed game providers ignoring app shutdown, causing lingering background activity

---

## [v0.1.2-alpha] — 2026-02-16

### Added

- Pulse icon branding across the app

---

## [v0.1.1-alpha] — 2026-02-15

### Added

- Hover over any button or control to see what it does
- Trying to enable a widget without a game connected now shows a warning message
- Disabled toggles explain why they're disabled instead of silently ignoring clicks

## [v0.1.0-alpha] — 2026-02-14

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
