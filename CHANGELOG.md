# Changelog

All notable changes to RacePulse releases are documented here.

## [v0.4.2] — YYYY-MM-DD

RacePulse 0.4.2 adds VR widgets (experimental), maker logos in Relative and Standings, and a launch screen. It also fixes RacePulse opening twice and widgets taking the focus from your game.

### New

- **VR widgets (experimental, free):** see your widgets in your headset: in games running through SteamVR, with any SteamVR headset, and in Direct3D 11 games on other OpenXR runtimes. Games on the old Oculus PC SDK aren't supported; switch them to their OpenXR or SteamVR mode. Turn on **Show widgets in VR** on the new **VR widgets** page; it's off by default, and nothing VR loads until you turn it on. A short guide takes you through choosing widgets, starting VR, and placing them.
- **VR layouts:** VR has its own layouts, kept apart from your desktop layouts, and each VR widget has its own settings. Start from your desktop widgets or pick new ones. VR starts with the layout you selected last.
- **Place widgets from inside the headset:** RacePulse's tab in the SteamVR dashboard shows, hides, moves, resizes, and sets up your VR widgets, and switches layouts. From the keyboard, **Ctrl+Alt+Home** starts or ends editing, **Ctrl+Alt+End** recenters the widgets on your view, and **Ctrl+Alt+Page Down** switches to the next VR layout.
- **Desktop or headset, never both:** while a VR session runs, your desktop widgets hide, and they come back when it ends.
- **Manufacturer logos** in Relative and Standings: each car's maker before the driver's name, in iRacing, Le Mans Ultimate, and F1. Standings shows them by default; in Relative, turn on **Show Manufacturer Logos**.
- **Launch screen:** RacePulse shows a short loading animation while it starts, then the whole window at once.

### Improved

- The Dashboard has its own monitor icon in the side menu, Settings, and its quick-link card.
- The Widget layouts, VR widgets, and Dashboard pages have their own header photos.

### Fixed

- Opening RacePulse while it's already running now brings the open window forward, instead of starting a second copy.
- Widgets no longer take the keyboard focus from your game when they open or reappear.
- **Inputs Telemetry:** the gear and speed no longer get cut off in short windows, and the corners are round again when the steering wheel is hidden.

### Good to know

- VR widgets are experimental. With **Show widgets in VR** off, RacePulse works as before.
- For games that use OpenXR outside SteamVR, start RacePulse before the game. RacePulse registers its VR layer only while it's running, so closing or deleting RacePulse leaves nothing loading into your games.
- The [VR Guide](https://racepulse.racing/vr/) has the recommended settings and setup.

If something looks wrong in your sim or headset, please [open an issue on GitHub](https://github.com/gAlexander77/racepulse-release/issues) with the sim, the widget, and what you saw.

## [v0.4.0] — 2026-10-12

RacePulse 0.4.0 redesigns the main window and most widgets, adds the Standings and Lap Log widgets, brings Relative and Standings to Le Mans Ultimate and F1, and adds widget layouts, a Settings window, and a getting-started Help guide.

**Behavior change:** widgets now hide when no game is sending telemetry. **Only show with telemetry** is on by default, including for existing setups. Locked widgets count down for 5 seconds, then hide, and come back as soon as telemetry returns. They always show while you arrange them. To keep widgets on screen all the time, open the **Overlay** menu in the bottom bar and turn off **Only show with telemetry**, or choose **Keep showing** on the countdown banner.

**Upgrading from 0.3.5:** your widgets, their places, and their settings carry over. They stay in `%USERPROFILE%\.racepulse`. Edit Mode is now **Arrange**, click-through is **Lock for racing**, and each widget's settings are on its own **Customize** page.

### New

- **Standings widget:** the whole field by class in six layouts: Classic Multiclass, Timing Tower, Class Columns, Battle View, Analyst, and Practice and Qualifying. Every column can follow the layout or be shown or hidden. Purple marks a class's fastest lap and green a personal best. Works in iRacing, Le Mans Ultimate, and F1.
- **Lap Log widget:** every lap under a banner of the session best and your best, with its delta, IN and OUT pit tags, and optional sectors, fuel used, temperatures, top speed, and position. A live row shows the lap you're driving. Works in every supported sim.
- **Relative in Le Mans Ultimate and F1:** Relative now works in Le Mans Ultimate and F1 as well as iRacing.
- **Country flags** in Relative and Standings: from profile flair in iRacing, nationality in F1, and nationality on Le Mans Ultimate online servers. Show Country Flags is on by default.
- **Widget layouts:** save the widgets on screen, with their places, sizes, and settings, as named layouts, and switch between them from the bottom bar. The Widget layouts page shows each layout across your monitors, and layouts remember the monitors they were saved on. Choose No layout to use RacePulse without them.
- **Settings window** (the gear at the top of the window, or Ctrl+,): language, the compact sidebar, Only show with telemetry, the selected layout, the connection and each sim's setup, updates, and the data folder in one place, with search.
- **Language:** English (US), English (UK), or Automatic, in Settings › General. English (UK) uses British spelling in the main window and every widget. Automatic decides from your Windows language and time zone, on your computer.
- **Overlay menu** in the bottom bar, with Only show with telemetry, Recenter all widgets, and Turn all off. Recenter and Turn all off now ask first.
- **Help is a getting-started guide:** four steps (connect your sim, turn on widgets, place and lock them, save a layout) that tick off as you go, plus tips and setup for each sim.
- **Telemetry dashboard (experimental):** a car-style dash in its own window for a second screen, with borderless full screen. Open it from Dashboard in the side menu and choose Pop out.
- **Setting hints:** hover a setting name with a dotted underline for a short explanation.
- **Preview backdrops:** on a widget's Customize page, see it over a sunny, stormy, or night track.

### Improved

- **Redesigned main window:** a widget sidebar, an All widgets gallery with live previews and Advanced widgets first, a Customize page for each widget with a large preview, and a bottom bar with the layout picker, the widget count, and Arrange / Lock for racing.
- **Arrange and Lock for racing** replace Edit Mode and Click-Through. Turning a widget on starts Arrange, and you can set widgets up before starting a game. While arranging, each widget has a lock button and a × that turns off just that widget.
- **Relative redesign:** class colors, car numbers, laps-ahead colors, a headset for who's talking on the radio (iRacing), an optional header with strength of field or the session, Hide Parked Cars, Your Row Color, and smooth row animations. Its window fits what it shows.
- **Relative in iRacing:** gaps between classes are measured on your class's lap, so a faster car alongside reads close to 0.0. OUT now marks a car on the lap after it leaves the pits, instead of every car at a race start. From the garage, while spectating, or in a replay, it centers on the car the camera follows.
- **Flat Track Map redesign:** horizontal or vertical, a lane per class, tags with position, number, or initials, sector lines, yellow flag zones, a pit lane row, and warnings when a faster class closes on you.
- **Telemetry Core redesign:** choose each readout from 19 readings, with an optional delta bar, pedal bars, fuel bar, flags, and ABS light.
- **Inputs Telemetry redesign:** a steering wheel at either end that uses the car's own steering lock in iRacing and Le Mans Ultimate, ABS shown on the brake line (iRacing and ACC), and more graph settings. The pedal bars sit closer together, in Inputs too.
- **Basic widgets redesign:** RPM can show rev lights, Speed an optional session top speed and pit limiter state, Lap Times three layouts, and G-Force rings, a trail, and lap peaks. Numbers sit in fixed slots, so they no longer shift as they change.
- **Fuel:** use per lap is now averaged over your recent laps, leaving out pit and refueling laps. Optional fuel to finish, a SAFE, TIGHT, or SHORT status, and fuel to add at your next stop.
- **Widgets a sim can't feed now hide** instead of staying on screen empty, with a notice that says why. They come back in the same place when you switch to a sim that supports them.
- Each widget card shows the sims it works in, and settings that only some sims support are marked.
- Widgets now open as windows of RacePulse itself instead of separate helper programs, and they have their own taskbar button, so the main window is easy to find.
- Delta is now offered only in iRacing, Le Mans Ultimate, and ACC. F1 and Assetto Corsa don't send a live delta, so it used to sit at +0.000 there.
- F1 25's 2026 Season Pack telemetry format is supported.
- **Le Mans Ultimate:** car numbers and team names come from the game, so custom-team cars show their real numbers. Driver names drop LMU's account tag, and class names read cleanly.
- **Le Mans Ultimate:** time left, laps left, and your class position now reach the widgets.
- Clearer connection errors: picking a game that isn't running now says to start it and load into a session.

### Fixed

- **iRacing:** drivers whose names arrive as question marks (characters iRacing's data can't hold) now show as Driver #N with their car number.
- **iRacing:** Delta shows dashes while iRacing marks the delta invalid, instead of 0.000.
- **Le Mans Ultimate:** Delta shows dashes before you have a best lap, instead of a red +0.000.
- Delta shows a plain 0.000 when you're level with your best, instead of a red +0.000.
- Telemetry Core shows a dash instead of P0 when you have no position yet.
- G-Force now shows the force you feel on both axes: out of the corner, up under braking, and down under acceleration.
- **F1:** best laps now show in Lap Times; the Best row used to stay empty.
- **F1:** races were reported as an unknown session, and Sprint Shootouts as races.
- **F1:** the Flat Track Map put every car on the start line. Cars now show where they are on the lap, including you on your out lap.
- **F1:** a car after an empty grid slot could be missing from the field.
- **F1:** the Inputs Telemetry steering wheel turned the wrong way.
- **ACC:** an unset best lap showed as a huge time; it now shows as no time.
- **ACC and Assetto Corsa:** the Inputs Telemetry steering wheel turned the wrong way and only about a third of the way at full lock.
- Retry after a failed update now starts the download again.

### Good to know

If something looks wrong in your sim, please [open an issue on GitHub](https://github.com/gAlexander77/racepulse-release/issues) with the sim, the widget, and what you saw.

- **F1:** if RacePulse doesn't connect, another app may be using the port the game sends to. Pick another of RacePulse's ports (22025, 22024, 22023, or 20777) in the game, or close the other app.
- **Le Mans Ultimate:** country flags need an online server; single-player sessions have none. The Lap Log's session best clears when the driver who set it leaves.
- Standings counts positions gained only if RacePulse was running before the green flag, and counts pit stops and stints from what it has seen.
- A layout's widgets on a monitor that isn't connected stay where they were saved. Use Recenter all widgets to bring them onto your screen.
- The telemetry dashboard is experimental and has few settings for now.

---

## [v0.3.5] — 2026-10-04

### Added

- Smooth opening and closing animations for settings drawers and confirmation dialogs, with reduced-motion support.
- **Star on GitHub** in Home and Help, making it easy to support RacePulse through the public repository.

### Fixed

- More reliable update installation and restart, including downloads staged on a different drive and recovery when replacing the executable fails.
- Safer widget startup, reconnection, and shutdown when connections are slow, interrupted, or closed mid-operation.
- Settings save/reset failures now preserve edits and report the problem instead of appearing to succeed.
- More consistent lap references and sector timing around missing lap counters, session changes, and stationary periods.
- Startup problems now show recovery guidance, and transient event-registration failures retry automatically.

### Improved

- Bounded telemetry queues and lap-analysis buffers reduce unnecessary background work and memory growth during interrupted or stationary sessions.
- The familiar main-page layout and widget designs are preserved, with bundled fonts and icons for offline use.

---

## [v0.3.4] — 2026-07-16

### Fixed

- **Le Mans Ultimate:** RacePulse now works on fresh LMU installs with zero setup. It reads the shared memory interface built into the game, so the separate shared memory plugin DLL is no longer required — just turn on **Settings > Gameplay > Enable Plugins**. Previously, LMU telemetry silently depended on a plugin that other apps happened to install; without it, RacePulse never detected the game. Older LMU versions that predate the built-in interface still fall back to the plugin if installed.

### Improved

- **Le Mans Ultimate:** your car is now identified directly by the game instead of through name-matching heuristics
- In-app Help and the website setup guides for LMU rewritten for the zero-setup flow

---

## [v0.3.3] — 2026-07-16

### Added

- Widgets can now be hidden and shown again without losing them — hiding previously closed the widget for good
- Smarter delta timing: the reference lap is now selected automatically and new personal bests are promoted as you set them

### Fixed

- **iRacing:** tire pressures now show real PSI values (a unit conversion ran the wrong way, giving readings in the thousands)
- **iRacing:** on-track detection works again — it was permanently "off track", which kept dependent widgets dark
- **iRacing:** shift lights now scale to each car's actual redline instead of a fixed RPM
- **iRacing:** checkered and cooldown phases are no longer swapped at the end of a race, and wind speed uses the correct unit
- **iRacing:** the last driver in the session roster is no longer missing from the leaderboard
- **iRacing:** leaderboard gaps to the leader are now accurate
- **ACC:** shift lights work again (a data type mismatch left them permanently dark)
- **ACC:** the clutch bar no longer reads 100% while driving — the sim reports clutch engagement, which is now converted to pedal position
- **ACC & AC:** pausing the game no longer disconnects the session after a few seconds
- **ACC:** the live delta now works from your second lap onward
- **ACC:** all three sector times now populate correctly
- **ACC:** your car now appears at its true position on track maps
- **Assetto Corsa:** flags and pit-lane detection now work (the app was reading ACC's memory layout, which AC doesn't share)
- **Assetto Corsa:** phantom cars no longer appear on track maps
- **Le Mans Ultimate:** tire temperatures now display in Celsius instead of raw Kelvin (~350°)
- **Le Mans Ultimate:** track flags and session types now map correctly — green-flag racing no longer shows as a caution, and races are no longer "unknown" sessions
- **Le Mans Ultimate:** the delta no longer compares against sectors you haven't finished driving
- **Le Mans Ultimate:** torn shared-memory frames are discarded instead of producing corrupted readings, and shared-memory faults no longer stop capture
- **F1 24/25:** F1 24 is now auto-detected (a port bug made it undetectable on a fresh start)
- **F1 24/25:** driver names are no longer garbled in F1 25, and several packet layouts were corrected for both games
- **ACC, AC & LMU:** the rev limiter indicator no longer flashes before the car's redline is known
- Widget layouts can no longer be corrupted by a crash during save — saves are atomic, and a damaged file is set aside and recovered from instead of silently overwritten
- Launching RacePulse twice no longer cross-wires widgets between the two instances
- Widgets automatically reconnect if their connection to the main app drops
- The app can no longer hang forever on exit — shutdown is bounded
- Failed auto-connects now back off progressively instead of retrying every 3 seconds
- Fuel-per-lap estimates are no longer skewed by your first lap out of the pits
- Track map no longer shows a bogus gap to the car ahead while crossing the start/finish line
- Invalid values from the sims no longer cause telemetry frames to silently drop
- The inputs graph no longer floods solid red in widget previews
- Rapidly toggling widgets on the Home page no longer gets them stuck

### Improved

- Live widget previews are far more realistic: coherent lap times, ticking sectors, a true live delta, distance-accurate car movement on the track map, and a proper multiclass field — with no flickering lap times
- Lap time, delta, and gaps on ACC now update smoothly instead of stepping a few times per second
- The Help menu now explains that ACC and AC can only be detected once you've loaded into a session

---

## [v0.3.2] — 2026-07-02

### Added

- Widgets, previews, and pages now show a small error panel instead of going blank if something unexpected goes wrong

### Fixed

- iRacing: projected iRating now reads the correct field and only shows in valid sessions
- Le Mans Ultimate no longer gets stuck in a detect/reconnect loop after the game closes and its shared memory goes stale
- F1: reconnecting now works after capture stops on its own (the UDP port is released properly)
- Fixed hangs that could occur when stopping telemetry capture on iRacing, Le Mans Ultimate, and F1
- Assetto Corsa Competizione and Assetto Corsa no longer hang when the game closes on its own
- Failed auto-connects now retry and surface the error instead of failing silently
- Telemetry no longer bleeds between sessions when switching games

### Improved

- Major performance pass while driving: the Home and dashboard pages stop re-rendering at 60Hz during capture, off-screen widget previews pause, graph canvases only reallocate when resized, and analysis and track-map updates are throttled — noticeably lower CPU and GPU usage
- Home quick links restyled into a single row

---

## [v0.3.1] — 2026-06-02

### Added

- Help menu in the app header with quick-start steps and per-sim setup instructions (F1 UDP port, Le Mans Ultimate plugin, and more)

### Fixed

- Header tooltips no longer get cut off at the top edge of the window — they now open downward

### Improved

- Redesigned the main app header into a single unified control bar — game detection, connection status, and connect/disconnect now live in one cohesive control instead of separate status badges and buttons
- Edit Mode and Click-Through are now clearly labeled toggle switches, so it's obvious what they control and whether they're currently on
- The widget on/off switch is now screen-reader accessible and shows a brief spinner while a widget starts up or closes (and ignores accidental double-clicks)
- Restyled the "reset to default settings" confirmation dialog to match the rest of the app

---

## [v0.3.0] — 2026-05-01

### Fixed

- Mid-session crashes during long Le Mans Ultimate races — several edge cases that could take the app down are now caught and contained
- A single slow widget can no longer freeze every other widget on screen
- Telemetry no longer cuts out for a moment during loading screens and session transitions
- F1 telemetry handles unexpected or partial UDP packets without crashing
- Rare crashes when closing the application
- Pit window warning now correctly fires when your fuel tank reaches empty
- Negative fuel readings from sensor glitches no longer break the fuel display
- Personal best laps are now correctly marked in the lap list and session export
- Lap consistency stat is now accurate (it was slightly too generous before, especially with only a few laps recorded)
- Leaderboard rows no longer flicker when two cars share the same position
- Leaderboard gap to the car ahead is now calculated correctly in all cases
- Track map widgets handle missing track data on loading screens — cars no longer fly off to "infinity"
- Le Mans Ultimate: only your car is now identified as the player car (some sessions previously flagged multiple cars)
- Cars sitting exactly on the start/finish line stay visible on the Flat Track Map widget
- Fuel widget's low-fuel red color now triggers at the right moment when remaining laps include a decimal
- Overlay opacity slider can no longer be broken by an invalid value (which previously made every widget invisible)

### Improved

- The auto-updater now only installs stable releases and skips beta/draft versions
- Better error messages when an update fails to download (you'll see what actually went wrong instead of a confusing "checksum mismatch")

### Compatibility

- Existing widget layouts and per-widget settings carry over from earlier versions — no reconfiguration needed

---

## [v0.2.9-alpha] — 2026-04-08

### Fixed

- Fixed a crash in the Relative widget when iRating data was unavailable
- Relative widget no longer flickers or jumps around during cautions and restarts
- Safety rating now displays correctly for all license levels
- Colored gap times now properly show negative signs for cars ahead of you
- Relative gaps are more accurate during the first lap before anyone sets a time
- iRating change column now only shows during race sessions, not practice or qualifying

### Improved

- Game-specific widgets now show a clear error instead of broken data when used with an unsupported game

---

## [v0.2.8-alpha] — 2026-04-03

### Added

- **Relative widget** — see nearby drivers sorted by proximity on track with live gap times, iRating, safety rating, and license class (iRacing only)
- OUT indicator shows when a driver is on their out lap
- PIT indicator shows when a driver is in the pits
- Projected iRating change display (beta, off by default)
- Colored gap times option — enable green/red coloring with +/- signs in widget settings
- Live widget preview on the home page

---

## [v0.2.7-alpha] — 2026-03-27

### Improved

- Website and support buttons on the home page for quick access

### Fixed

- Widget settings (unit preferences, colors, etc.) now survive application restarts

---

## [v0.2.6-alpha] — 2026-03-22

### Added

- Widget position and size persistence — widgets reopen exactly where you left them after restarting the app
- Recenter button to move all open widgets to the center of the main app's screen
- Reset Size button in the widget customization panel to restore default dimensions
- Optional rev lights bar on the Inputs Telemetry widget

---

## [v0.2.5-alpha] — 2026-03-19

### Added

- Website link in the app footer for quick access to racepulse.racing
- Widget icons for basic widgets
- Toast notifications for connection and widget feedback

### Fixed

- Connection error banner persisting after successfully connecting to a game
- Missing window icon in the Windows taskbar

---

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
