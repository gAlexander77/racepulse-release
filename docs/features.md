# RacePulse Features

RacePulse is a desktop telemetry overlay application designed for sim racing.

## Supported Games

- **iRacing** — Automatic detection via Windows shared memory. No configuration needed.
- **Le Mans Ultimate** — Automatic detection via rFactor 2 shared memory plugin. No configuration needed.
- **F1 25** — UDP telemetry on port 22025 (enable in game settings under Telemetry).
- **F1 24** — UDP telemetry on port 22024 (enable in game settings under Telemetry).

RacePulse auto-detects the running game and begins capturing telemetry. Multiple game support means you can switch between sims without reconfiguring. All widgets work across all supported games through a unified telemetry model.

## Telemetry Capture

- Live telemetry ingestion at up to 60Hz
- Normalized telemetry model for consistent widget behavior across all supported games
- Real-time update pipeline for responsive overlays
- Automatic game detection — no manual game selection required

## Overlay Widgets

Standalone, always-on-top widget overlays that float over your game:

**Core widgets:**
- Speed display
- Gear indicator
- RPM bar with redline indicator
- Lap timing (current, last, best)
- Delta time to best lap
- Input bars (throttle, brake, clutch)
- Fuel remaining with estimated laps
- G-force visualization

**Advanced widgets:**
- Telemetry Core — combines animated rev lights, gear, RPM, speed, position, track temp, elapsed time, and lap count with metric/imperial support
- Inputs Telemetry — live scrolling input graph with pedal bars, optional steering wheel visualization, gear, and speed

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
