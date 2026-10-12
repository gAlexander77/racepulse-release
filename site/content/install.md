---
title: "Installation"
subtitle: "Download, launch, and configure RacePulse in under a minute. Setup guides for iRacing, ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24."
description: "Install RacePulse in under a minute with step-by-step setup for iRacing, ACC, Assetto Corsa, and Le Mans Ultimate shared memory and F1 25/F1 24 UDP telemetry."
ogImage: "assets/application-preview.jpg"
---

## 1. Download

Download the latest `RacePulse.exe` from the [Releases](https://github.com/gAlexander77/racepulse-release/releases) page.

Download the matching `RacePulse.exe.sha256` asset and verify the executable SHA-256 before running. Compare it with `Get-FileHash -LiteralPath .\RacePulse.exe -Algorithm SHA256` in PowerShell; the first field in the sidecar is the expected hash.

## 2. Launch

Run `RacePulse.exe`. No installation required.

> **Note:** On first launch, Windows may show a SmartScreen prompt since the app is unsigned. Click "More info" then "Run anyway".

## 3. Game Setup

### iRacing

No configuration needed. RacePulse automatically detects iRacing via shared memory when the sim is running. Load into a session to see live telemetry.

### Assetto Corsa Competizione

No configuration needed. RacePulse reads ACC's shared memory, but ACC can't be detected from the main menu. Load into a session (practice, qualifying, or race) first, and RacePulse detects it.

### Assetto Corsa

No configuration needed. Assetto Corsa is detected once a session is loaded. Its main menu runs in a separate launcher, so RacePulse can't see the game there.

### Le Mans Ultimate

One-time step: in LMU, go to **Settings > Gameplay** and turn **Enable Plugins** on. RacePulse reads the shared memory interface built into the game, with no plugin DLL to install. Once a session is loaded, RacePulse detects LMU automatically.

While LMU is running, RacePulse also reads car numbers, team names, and (on online servers) driver countries from the game's own local web interface on your PC. Nothing to set up; it feeds Relative and Standings.

### F1 25

1. In F1 25, go to **Settings > Telemetry Settings**
2. Set **UDP Telemetry** to **On**
3. Set **UDP Port** to one of the ports RacePulse listens on: **22025**, **22024**, **22023**, or **20777**
4. Set **UDP Format** to **2025**. With the 2026 Season Pack, its 2026 format works too.
5. Start a session. RacePulse auto-detects the game when telemetry is broadcasting.

### F1 24

1. In F1 24, go to **Settings > Telemetry Settings**
2. Set **UDP Telemetry** to **On**
3. Set **UDP Port** to one of the ports RacePulse listens on: **22025**, **22024**, **22023**, or **20777**
4. Set **UDP Format** to **2024**
5. Start a session. RacePulse auto-detects the game when telemetry is broadcasting.

> **If F1 doesn't connect:** another app may be using that port. Pick another of RacePulse's ports in the game, or close the other app. RacePulse names the game by the UDP format it receives, so F1 25 sending the 2024 format shows as F1 24. Older formats aren't read.

## 4. Configure Widgets

- In **All widgets**, flip the switch on any widget card. You don't need a game running.
- Turning a widget on starts **Arrange**: drag your widgets where you want them on your monitor
- Choose **Lock for racing** in the bottom bar when you're done. Locked widgets let your clicks through to the game.
- Choose **Customize** on a card to change units, colors, and what the widget shows, with a live preview. **Save settings** keeps your changes.
- Optional: save what's on screen as a widget layout, from the layout picker in the bottom bar or the **Widget layouts** page
- Settings persist automatically across sessions

**Help** (the ? at the top of the window) walks through these steps and ticks each one off as you go.

> **Note:** With **Only show with telemetry** on (the default), locked widgets hide while no game is sending telemetry and come back when it returns. To keep them on screen all the time, turn it off in the bottom bar's **Overlay** menu.

## 5. Updating

RacePulse checks for new versions when it starts. When one is out, an update banner appears under the title bar: choose **Update now**, and RacePulse downloads the new version, checks it, and restarts into it. You can also check from **Settings › Updates**.

To update by hand, close RacePulse, download the newer release, and replace the previous `RacePulse.exe`.

Your widget settings, positions, and layouts are stored separately in `%USERPROFILE%\.racepulse` and will be preserved. **Settings › Advanced** opens that folder. Close RacePulse before backing it up or editing its files.

**Upgrading from 0.3.5:** your widgets and their settings carry over. A few things work differently:

- Edit Mode is now **Arrange**, and click-through is **Lock for racing**, both in the bottom bar
- Each widget's settings are on its own **Customize** page instead of the Config drawer
- Widgets now hide while no game is sending telemetry. Turn off **Only show with telemetry** in the **Overlay** menu if you want them on screen all the time.
- **Recenter all widgets** and **Turn all off** are in the **Overlay** menu, and both ask first

---

## Troubleshooting

### Overlays do not appear

- Ensure your game is running and a session is active (on track, not in menus)
- Check the connection pill at the top of RacePulse: it reads **LIVE** when telemetry arrives
- With **Only show with telemetry** on, locked widgets stay hidden until a game sends telemetry. Switch to **Arrange** to see them, or turn the setting off in the **Overlay** menu.
- If a widget card says **Hidden in** your sim, that game doesn't send the data the widget needs
- Lost a widget on another monitor? **Overlay › Recenter all widgets** brings every widget to the main window's screen

### No telemetry data (F1 games)

- Confirm UDP Telemetry is enabled in the F1 game settings
- Check that the UDP port is one RacePulse listens on: 22025, 22024, 22023, or 20777
- Check that UDP Format is 2025 for F1 25 (or 2026 with the 2026 Season Pack) and 2024 for F1 24
- If it still doesn't connect, another app may be using that port. Pick another of these ports in the game, or close the other app.
- Allow RacePulse through Windows Firewall if prompted

### No telemetry data (iRacing / ACC / Assetto Corsa / Le Mans Ultimate)

- Ensure the sim is running and you are in a session (not the main menu)
- Shared memory is only available while on track or in replay
- **ACC and Assetto Corsa:** these are only detected once a session is loaded
- **LMU only:** verify **Settings > Gameplay > Enable Plugins** is turned on

### App won't start or crashes

- Ensure you are on Windows 10 or 11 (64-bit)
- Try running as Administrator
- Check that desktop composition is enabled (default on Windows 10/11)

### Widgets show stale data after closing a game

- RacePulse disconnects automatically shortly after the game stops sending data
- You can also choose **Disconnect** from the connection pill's menu
