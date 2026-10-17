# RacePulse Features

RacePulse is a desktop telemetry overlay application designed for sim racing. This page matches the [features page](https://racepulse.racing/features/) on the website.

## Supported Games

- **iRacing** — Automatic detection via Windows shared memory. No configuration needed.
- **Assetto Corsa Competizione** — Automatic detection via shared memory once you load into a session. ACC can't be detected from the main menu.
- **Assetto Corsa** — Automatic detection via shared memory once a session is loaded. AC's main menu runs in a separate launcher, so it isn't detected there.
- **Le Mans Ultimate** — Automatic detection via the game's built-in shared memory (enable Plugins in LMU settings once).
- **F1 25** — UDP telemetry. Turn it on in the game's Telemetry Settings and send it to one of RacePulse's ports: 22025, 22024, 22023, or 20777.
- **F1 24** — UDP telemetry, set up the same way as F1 25.

RacePulse auto-detects the running game and begins capturing telemetry. Multiple game support means you can switch between sims without reconfiguring, and the connection pill at the top of the window can also pick a sim by hand. Most widgets work in every supported game. A few need data that only some sims send; see [Widget Support by Game](#widget-support-by-game).

## Telemetry Capture

- Live telemetry ingestion at up to 60Hz
- Normalized telemetry model for consistent widget behavior across all supported games
- Real-time update pipeline for responsive overlays
- Automatic game detection — no manual game selection required

## Core Widgets

Standalone, always-on-top widget overlays that float over your game. The app lists these eight under Basic:

- **Speed** — Live speed in KPH or MPH, as a number or a small pill. Optional top speed for the session and a pit limiter state.
- **Gear** — Large, bold indicator visible in peripheral vision, with its own colors for neutral and reverse
- **RPM** — The number (7840 or 7.8k), or a row of 8 to 15 rev lights with the number under them. The pit limiter can flash every light together.
- **Lap Times** — Current, last, and best lap, in rows, with a large current lap, or on one line. Optional last-vs-best and a 5-second hold of the lap you just finished.
- **Delta** — Real-time delta to best lap (green = faster, red = slower). iRacing, Le Mans Ultimate, and ACC.
- **Inputs** — Throttle, brake, and clutch as vertical bars, with percentages
- **Fuel** — Fuel left and laps left, with fuel use averaged over your recent laps (pit and refueling laps are left out). Optional fuel per lap, a fuel bar, fuel to finish, a SAFE, TIGHT, or SHORT status, and fuel to add at your next stop.
- **G-Force** — 2D lateral and longitudinal G-force visualization, with rings at 1 g and 2 g, a fading trail, and optional peaks for the lap

## Advanced Widgets

- **Telemetry Core** — All-in-one display combining animated rev lights, a large gear, two readouts beside it, and an info row of four slots. Each readout and slot is your choice of 19 readings, such as speed, RPM, position, lap, time left, delta, fuel, temperatures, brake bias, and ERS. Supports metric and imperial units. Optional delta bar, pedal bars, fuel bar, flags, and ABS light.
- **Inputs Telemetry** — Live scrolling graph of throttle and brake inputs (and optionally clutch) with color-coded traces, combined with real-time pedal bars, gear, speed, and a steering wheel that turns with yours. The brake line can change color while ABS is working (iRacing and ACC).
- **Flat Track Map** — Every car on one line of the lap, across the screen or down its edge, in one lane or a lane per class. Markers can be dots or tags with each car's position, class position, number, or initials. Optional sector lines, yellow flag zones, a pit lane row, lapped-car marks, and a warning when a faster class closes on you. Supports iRacing, LMU, and F1.
- **Relative** — Nearby drivers sorted by track proximity with live gap times, class colors, car numbers, country flags, and pit/out lap indicators. In iRacing it also shows iRating, safety rating, license class, and who's talking on the radio; projected iRating change is in beta. An optional header shows strength of field or the session. Supports iRacing, LMU, and F1.
- **Standings** — The whole field by class in six layouts: Classic Multiclass, Timing Tower, Class Columns, Battle View, Analyst, and Practice and Qualifying. Each column can follow the layout or be shown or hidden. Purple marks a class's fastest lap and green a personal best. Fill Window lists more drivers when the window has room. Supports iRacing, LMU, and F1.
- **Lap Log** — Every lap under a banner of the session best and your best, with its delta to your best, the session best, the previous lap, or the average of your last five. IN and OUT pit tags, and optional sectors, fuel used, track and air temperature, top speed, and position. A live row shows the lap you're driving. Works in every supported game.

Country flags come from each driver's profile flair in iRacing, their nationality in F1, and their nationality on Le Mans Ultimate online servers. LMU single-player sessions have no flags.

## Widget Support by Game

| Widget | iRacing | Le Mans Ultimate | F1 25 / F1 24 | ACC | Assetto Corsa |
| --- | --- | --- | --- | --- | --- |
| Speed, Gear, RPM, Lap Times, Inputs, Fuel, G-Force | Yes | Yes | Yes | Yes | Yes |
| Delta | Yes | Yes | No | Yes | No |
| Telemetry Core, Inputs Telemetry, Lap Log | Yes | Yes | Yes | Yes | Yes |
| Flat Track Map, Relative, Standings | Yes | Yes | Yes | No | No |

F1 and Assetto Corsa don't send a live delta, so Delta isn't offered there. Some settings also depend on the game, such as iRating and license in Relative (iRacing only). Each widget card shows the sims it works in, and settings that only some sims support are marked with those sims.

If you connect a game that can't feed an open widget, such as Relative in ACC, the widget hides instead of closing, and a notice says why. It comes back in the same place when you switch to a game that supports it.

## Widget Customization

- Each widget has a Customize page with a large live preview and its settings in Display, Format, and Window tabs
- Per-widget settings (units, colors, layouts, visibility toggles, display options)
- Hover a setting name with a dotted underline for a short explanation of what it does
- Preview a widget over a sunny, stormy, or night track backdrop to judge how it will look in game
- Configurable frame rate (60, 30, or 15 FPS) and background opacity per widget
- Conditional settings that gray out when their parent setting is disabled
- **Save settings** keeps your changes, **Reset defaults** restores the widget's defaults, and **Reset size** restores its default size
- Persistent settings that survive app restarts

## Widget Layouts

Save the widgets on screen as a named layout: which widgets are on, where they sit, their sizes, and their settings. Pick a layout from the layout picker in the bottom bar. Choose No layout to use RacePulse without layouts.

- The Widget layouts page shows each layout's desktop across your monitors, with a table of its widgets
- Select, deselect, rename, duplicate, or delete a layout from its details, or right-click it in the list
- While a layout is selected, moving, resizing, and saving a widget's settings save to it. Turning a widget on or off asks first.
- Layouts remember the monitors they were saved on. If one isn't connected, the page shows it and selecting the layout warns you; **Recenter all widgets** brings them onto your screen.

## Arrange and Lock for Racing

- Turning a widget on starts **Arrange**, so you can drag it where you want it on your monitor. You don't need a game running.
- While arranging, every widget has an amber frame, a lock button that locks all widgets, and a × that turns off just that widget
- **Lock for racing** in the bottom bar makes widgets click-through, so your mouse goes to the game
- **Only show with telemetry** (on by default) hides locked widgets while no game is sending telemetry. Each widget counts down for 5 seconds first, and they come back as soon as telemetry returns. Widgets always show while you arrange. Turn it off in the bottom bar's **Overlay** menu to keep widgets on screen.
- The Overlay menu also has **Recenter all widgets**, which brings every widget to the main window's screen, and **Turn all off**. Both ask first.
- Widgets you turned on reopen when you start RacePulse
- Widget windows have their own taskbar button, so the main window is easy to find

## Settings, Language, and Help

- **Settings** (the gear at the top of the window, or Ctrl+,) gathers RacePulse's settings in one place: language, the compact sidebar, Only show with telemetry, Recenter, the selected layout, the connection and each sim's setup, the dashboard, updates, and the data folder. Search filters every page, and changes apply at once.
- **Language** (Settings › General) is English (US), English (UK), or Automatic. English (UK) uses British spelling, such as colour, tyre, and licence, in the main window and every widget. Automatic decides from your Windows language and time zone, on your computer; nothing is sent anywhere.
- **Help** (the ? at the top of the window) walks you through getting started in four steps: connect your sim, turn on widgets, place and lock them, and save a layout. Each step ticks off as you do it. It also has tips and the setup for each sim.

## Telemetry Dashboard (Experimental)

A car-style dash for a second screen, in a normal window of its own. Open **Dashboard** in the side menu to see what it shows, preview it, and change its settings, then choose **Pop out**. The window can go borderless full screen on any monitor: F11 toggles full screen and Esc leaves it. The dash shows rev lights, gear, speed, delta, lap times, position and lap, fuel, temperatures, and inputs, and each can be turned off. It is experimental, with few settings for now.

## Session Recording (Coming Soon)

- Telemetry sessions will be recorded to local storage automatically
- Captures session metadata (game, track, car, duration)
- Records lap times and validity
- Stores full telemetry frame history

## Quality of Life

- Single executable — no installer, no dependencies
- Compact distribution (~19MB)
- Widgets open as windows of RacePulse itself, with no separate helper programs
- Widget positions, sizes, and settings persist across sessions
- Set widgets up before you start a game, so they're ready when you load in
- RacePulse checks for updates when it starts, and Settings › Updates can check any time
