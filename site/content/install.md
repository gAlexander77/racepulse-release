---
title: "Installation"
subtitle: "Get RacePulse running in under a minute. No installer required."
---

## 1. Download

Download the latest `racepulse.exe` from the [Releases](https://github.com/gAlexander77/racepulse-release/releases) page.

If a checksum file is included in the release, verify file integrity before running.

## 2. Launch

Run `racepulse.exe`. No installation required.

> **Note:** On first launch, Windows may show a SmartScreen prompt since the app is unsigned. Click "More info" then "Run anyway".

## 3. Game Setup

### iRacing

No configuration needed. RacePulse automatically detects iRacing via shared memory when the sim is running.

### F1 25

1. In F1 25, go to **Settings > Telemetry Settings**
2. Set **UDP Telemetry** to **On**
3. Ensure **UDP Port** is set to **22025** (default)
4. RacePulse will auto-detect the game when telemetry is broadcasting

### F1 24

1. In F1 24, go to **Settings > Telemetry Settings**
2. Set **UDP Telemetry** to **On**
3. Ensure **UDP Port** is set to **22024** (default)
4. RacePulse will auto-detect the game when telemetry is broadcasting

## 4. Configure Widgets

- Enable desired widgets from the dashboard
- Click "Edit Mode" to reposition overlays on your display
- Customize widget settings (units, colors, visibility) via the Customize panel
- Settings persist automatically across sessions

## 5. Updating

Download a newer release and replace the previous `racepulse.exe`. Your widget settings and positions are stored separately and will be preserved.

---

## Troubleshooting

### Overlays do not appear

- Ensure your game is running and a session is active (on track, not in menus)
- Check that RacePulse shows "Connected" in the dashboard
- Try toggling Edit Mode on and off

### No telemetry data (F1 games)

- Confirm UDP Telemetry is enabled in the F1 game settings
- Check the UDP port matches the game year (F1 25 = 22025, F1 24 = 22024)
- Allow RacePulse through Windows Firewall if prompted

### No telemetry data (iRacing)

- Ensure iRacing is running and you are in a session (not the main menu)
- iRacing shared memory is only available while on track or in replay

### App won't start or crashes

- Ensure you are on Windows 10 or 11 (64-bit)
- Try running as Administrator
- Check that desktop composition is enabled (default on Windows 10/11)

### Widgets show stale data after closing a game

- RacePulse will automatically disconnect after ~10 seconds of no data
- You can also click Disconnect manually from the dashboard
