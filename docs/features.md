# RacePulse Features

RacePulse is a desktop telemetry overlay application designed for sim racing.

## Supported Games

- **iRacing** — Automatic detection via Windows shared memory. No configuration needed.
- **F1 25** — UDP telemetry on port 22025 (enable in game settings under Telemetry).
- **F1 24** — UDP telemetry on port 22024 (enable in game settings under Telemetry).

RacePulse auto-detects the running game and begins capturing telemetry. Multiple game support means you can switch between sims without reconfiguring.

## Telemetry Capture

- Live telemetry ingestion at up to 60Hz
- Normalized telemetry model for consistent behavior across all supported games
- Real-time update pipeline for responsive overlays
- Automatic game detection — no manual game selection required

## Overlay Widgets

Standalone, always-on-top widget overlays that float over your game:

**Core widgets:**
- Speed display (km/h or mph)
- Gear indicator
- RPM bar with shift light integration
- Lap timing (current, last, best)
- Delta time (to best lap, to leader, to car ahead)
- Input bars (throttle, brake, clutch)
- Fuel remaining

**Advanced widgets:**
- Telemetry Core — combines rev lights, gear, speed, position, and session info
- Inputs Telemetry — live scrolling input graph with pedal bars
- G-force display

## Widget Customization

- Per-widget settings (units, labels, colors, visibility toggles)
- Persistent defaults that survive app restarts
- Quick reset to default settings with confirmation
- Live preview of widgets on the dashboard before enabling overlays

## Analysis

- Lap timing and delta comparisons
- Sector time tracking
- Session-level insights and telemetry history
- Fuel strategy data (remaining laps, consumption)

## F1-Specific Data

When connected to F1 24 or F1 25, additional data is available:
- ERS energy store and deployment mode
- DRS status (allowed/active)
- Tyre compound (soft/medium/hard/inter/wet) and tyre age
- Fuel mix mode
- Safety car status
- Full competitor grid with positions, gaps, and pit status

## Quality of Life

- Single executable — no installer, no dependencies
- Compact distribution (~21MB)
- Widget positions and settings persist across sessions
- Edit mode for repositioning widgets, click-through mode for racing
