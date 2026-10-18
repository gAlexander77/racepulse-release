---
title: "FAQ"
subtitle: "Frequently asked questions about RacePulse — setup, supported games, widgets, troubleshooting, and more."
description: "Answers to common questions about RacePulse: free sim racing telemetry overlay for iRacing, ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24."
ogImage: "assets/application-preview.jpg"
---

<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "FAQPage",
  "mainEntity": [
    {
      "@type": "Question",
      "name": "Is RacePulse free?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Yes. RacePulse is completely free with no subscription, no trial period, and no locked features. Every widget and every supported game is available at no cost."
      }
    },
    {
      "@type": "Question",
      "name": "What sim racing games does RacePulse support?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse supports six games: iRacing, Assetto Corsa Competizione (ACC), Assetto Corsa, Le Mans Ultimate (LMU), F1 25, and F1 24. iRacing, ACC, AC, and LMU use Windows shared memory (LMU needs the Enable Plugins toggle turned on once). F1 25 and F1 24 use UDP telemetry (one-time setting in the game)."
      }
    },
    {
      "@type": "Question",
      "name": "Do I need to install plugins or configure anything?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "No plugins to install for any game. For iRacing, ACC, and Assetto Corsa, RacePulse detects the game automatically via shared memory; ACC and Assetto Corsa are detected once you load into a session. For Le Mans Ultimate, turn on Enable Plugins in Settings > Gameplay (one-time). RacePulse reads the shared memory interface built into the game. For F1 25 and F1 24, enable UDP Telemetry in the game's settings and set the UDP port to one RacePulse listens on: 22025, 22024, 22023, or 20777."
      }
    },
    {
      "@type": "Question",
      "name": "Does RacePulse work in VR?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Yes. From version 0.4.1, RacePulse can show its widgets inside your VR headset in iRacing, Assetto Corsa Competizione, Assetto Corsa, and Le Mans Ultimate. VR widgets are free, experimental, and off by default. We recommend running your sim through SteamVR, which lets you place widgets from RacePulse's tab in the SteamVR dashboard. The VR Guide on racepulse.racing has the recommended settings."
      }
    },
    {
      "@type": "Question",
      "name": "How many widgets does RacePulse have?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse has 14 overlay widgets: Speed, Gear, RPM, Lap Times, Delta, Inputs, Fuel, G-Force, Telemetry Core (all-in-one display), Inputs Telemetry (scrolling input graph), Flat Track Map (car positions on track), Relative (nearby drivers with gap times), Standings (the whole field by class), and Lap Log (every lap with its delta)."
      }
    },
    {
      "@type": "Question",
      "name": "Does RacePulse require an installer?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "No. RacePulse is a single portable executable (about 20MB). Download RacePulse.exe, run it, and you're ready. No installation, no dependencies, no runtime to install."
      }
    },
    {
      "@type": "Question",
      "name": "What operating system does RacePulse run on?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse runs on Windows 10 and Windows 11 (64-bit). iRacing, ACC, Assetto Corsa, and Le Mans Ultimate support is Windows-only because those games use Windows shared memory. F1 25 and F1 24 use UDP which is cross-platform in principle, but the overlay windows currently require Windows."
      }
    },
    {
      "@type": "Question",
      "name": "Will RacePulse affect my game performance?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse is designed to be lightweight. It reads telemetry from shared memory or UDP — both are passive, read-only operations that add no load to the game itself. Each widget runs as a small transparent window. Most users see no measurable FPS impact."
      }
    },
    {
      "@type": "Question",
      "name": "Can I use RacePulse while streaming?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Yes. RacePulse widgets are standard windows, so streaming and recording software can capture them. A display capture picks them up automatically. A game capture may not show them; add each widget as its own window capture, or use a display capture instead."
      }
    },
    {
      "@type": "Question",
      "name": "How do I update RacePulse?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse checks for new versions when it starts and shows an update banner when one is out. Choose Update now and it downloads, checks, and restarts into the new version. You can also check from Settings, or download the latest RacePulse.exe from the releases page and replace your existing file. Widget positions, sizes, layouts, and settings are stored separately and will be preserved across updates."
      }
    },
    {
      "@type": "Question",
      "name": "Why do my widgets disappear when I'm not driving?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Only show with telemetry is on by default. It hides locked widgets while no game is sending telemetry, after a 5-second countdown on each widget, and shows them again as soon as telemetry returns. Widgets always show while you arrange them. To keep them on screen all the time, turn off Only show with telemetry in the Overlay menu in RacePulse's bottom bar."
      }
    },
    {
      "@type": "Question",
      "name": "Where does RacePulse keep my settings?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "In the .racepulse folder in your Windows user profile. It holds your widgets' places, sizes, and settings, your widget layouts, your VR layouts, and the dashboards' settings. Settings > Advanced in RacePulse opens the folder. Updating or replacing RacePulse.exe doesn't touch it."
      }
    },
    {
      "@type": "Question",
      "name": "How do I uninstall RacePulse?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse has no installer, so there's no uninstaller either. If you used VR widgets, turn off Show widgets in VR first. Then close RacePulse and delete RacePulse.exe. Your settings stay in the .racepulse folder in your user profile until you delete it too. The Uninstall Guide on racepulse.racing lists every folder RacePulse uses and has a PowerShell command that removes them."
      }
    }
  ]
}
</script>

## General

### Is RacePulse free?

Yes. RacePulse is completely free — no subscription, no trial period, no premium tier, and no locked features. Every widget and every supported game is available at no cost.

### Does RacePulse require an installer?

No. RacePulse is a single portable executable, about 20MB. Download `RacePulse.exe`, run it, and you're ready. No installation wizard and no runtimes to install. To remove it, see the [Uninstall Guide](/uninstall/).

### What operating system does RacePulse run on?

Windows 10 and Windows 11 (64-bit). iRacing, ACC, Assetto Corsa, and Le Mans Ultimate support is Windows-only because those games use Windows shared memory for telemetry. F1 25 and F1 24 use UDP which is cross-platform in principle, but the overlay windows currently require Windows.

### How do I update RacePulse?

RacePulse checks for new versions when it starts. When one is out, an update banner appears under the title bar: choose **Update now**, and RacePulse downloads the new version, checks it, and restarts into it. You can also check from **Settings › Updates**.

To update by hand, download the latest `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases) and replace your existing file. Widget positions, sizes, layouts, and all customized settings are stored separately in `%USERPROFILE%\.racepulse` and will be preserved across updates.

### Where does RacePulse keep my settings?

In `%USERPROFILE%\.racepulse`. It holds your widgets' places, sizes, and settings, your widget layouts, your VR layouts, and the dashboards' settings. **Settings › Advanced** in RacePulse opens the folder. Updating or replacing `RacePulse.exe` doesn't touch it. Close RacePulse before backing it up or editing its files.

### How do I uninstall RacePulse?

RacePulse has no installer, so there's no uninstaller either. If you used VR widgets, turn off **Show widgets in VR** first. Then close RacePulse and delete `RacePulse.exe`. Your settings stay in `%USERPROFILE%\.racepulse` until you delete that folder too. The [Uninstall Guide](/uninstall/) lists every folder RacePulse uses and has a PowerShell command that removes them.

---

## Supported Games

### What sim racing games does RacePulse support?

RacePulse supports six games:

- [**iRacing**](/iracing-overlay/) — Windows shared memory, no setup needed
- [**Assetto Corsa Competizione**](/acc-overlay/) — Windows shared memory, no setup needed; load into a session to be detected
- [**Assetto Corsa**](/assetto-corsa-overlay/) — Windows shared memory, no setup needed; load into a session to be detected
- [**Le Mans Ultimate**](/lmu-overlay/) — built-in shared memory, enable Plugins in LMU settings once
- [**F1 25**](/f1-25-overlay/) — UDP telemetry, enable in game settings once and pick one of RacePulse's ports
- [**F1 24**](/f1-24-overlay/) — UDP telemetry, enable in game settings once and pick one of RacePulse's ports

### Do I need to install plugins or configure anything?

For **iRacing, ACC, and Assetto Corsa** — no. RacePulse detects these games automatically via shared memory with zero configuration. ACC and Assetto Corsa can't be detected from their main menus, so load into a session first.

For **Le Mans Ultimate** — one-time step: go to **Settings > Gameplay** and turn on **Enable Plugins**. RacePulse reads the shared memory interface built into the game, so there is no plugin DLL to install.

For **F1 25 and F1 24** — one-time step: go to **Settings > Telemetry Settings**, set **UDP Telemetry** to **On**, and set the **UDP Port** to one RacePulse listens on: 22025, 22024, 22023, or 20777. Use UDP format 2025 for F1 25 and 2024 for F1 24. If RacePulse doesn't connect, another app may be using that port; pick another of these ports in the game. The [install guide](/install/) has the full steps.

### Can I switch between games without reconfiguring?

Yes. RacePulse auto-detects which supported game is currently running. Close one sim, launch another, and RacePulse connects to the new game automatically. Your widget layout and settings carry over between games. If the new game can't feed a widget, such as Relative in ACC, that widget hides with a notice and comes back when you switch to a game that supports it.

### Does RacePulse work in VR?

Yes. From version 0.4.1, RacePulse can show its widgets inside your VR headset in iRacing, Assetto Corsa Competizione, Assetto Corsa, and Le Mans Ultimate. VR widgets are free, experimental, and off by default: turn on **Show widgets in VR** on the **VR widgets** page. We recommend running your sim through SteamVR, which lets you place widgets from RacePulse's tab in the SteamVR dashboard. The [VR Guide](/vr/) has the recommended settings.

---

## Widgets

### How many widgets does RacePulse have?

RacePulse has **14 overlay widgets**:

- **Core (8):** Speed, Gear, RPM, Lap Times, Delta, Inputs, Fuel, G-Force (listed as Basic in the app)
- **Advanced (6):** Telemetry Core (all-in-one display), Inputs Telemetry (scrolling input graph), Flat Track Map (car positions), Relative (nearby drivers with gaps), Standings (the whole field by class, in six layouts), Lap Log (every lap with its delta)

See the full [features page](/features/) for details on each widget.

### Do all widgets work with every game?

Most widgets work across all six supported games. The exceptions:

- **Relative, Standings, and Flat Track Map** — iRacing, Le Mans Ultimate, and F1 (they need data about every car that ACC and Assetto Corsa don't send)
- **Delta** — iRacing, Le Mans Ultimate, and ACC (F1 and Assetto Corsa don't send a live delta)

Some settings depend on the game too. iRating, license, and who's talking in Relative and Standings are iRacing only, and country flags show in iRacing, F1, and Le Mans Ultimate online servers. The app marks each widget and setting with the sims it works in.

### Can I customize widget appearance and behavior?

Yes. Every widget has its own Customize page with a live preview:

- **Units** — KPH or MPH for speed, Celsius or Fahrenheit for temperatures, liters or US gallons for fuel
- **Visibility toggles** — show or hide individual display elements
- **Layouts and colors** — for example, RPM as a number or rev lights, and Standings in six layouts
- **Frame rate** — 60, 30, or 15 FPS per widget
- **Position and size** — drag to reposition while arranging, resize as needed

All settings persist across sessions automatically. **Reset defaults** on a widget's Customize page restores its default settings.

### How do I reposition widgets on screen?

Turning a widget on starts **Arrange**: widgets become draggable, so you can move them anywhere on your screen. When you're done, choose **Lock for racing** in the bottom bar. Locked widgets are click-through, so they don't interfere with your game. Switch back to **Arrange** any time to move them again.

To keep different setups, save them as widget layouts from the layout picker in the bottom bar.

### Why do my widgets disappear when I'm not driving?

**Only show with telemetry** is on by default. It hides locked widgets while no game is sending telemetry, for example in menus or between sessions. Each widget counts down for 5 seconds first, and they all come back as soon as telemetry returns. Widgets always show while you arrange them.

To keep widgets on screen all the time, open the **Overlay** menu in RacePulse's bottom bar and turn off **Only show with telemetry**. You can also choose **Keep showing** on the countdown banner in the main window.

---

## Performance & Streaming

### Will RacePulse affect my game performance?

RacePulse is designed to be lightweight. It reads telemetry passively — shared memory reads and UDP packet reception add no processing load to the game itself. Each widget runs as a small transparent window. Most users see no measurable FPS impact.

### Can I use RacePulse while streaming or recording?

Yes. RacePulse widgets are standard windows that streaming and recording software can capture.

- **Display capture** picks up widgets automatically
- **Game capture** may not show widgets. Add each widget as its own **window capture** source, or switch to display capture.
- Widgets have transparent backgrounds, so they overlay cleanly in both the game and the stream

### Privacy & Data Collection

RacePulse processes racing telemetry locally on your device. It does not upload racing telemetry, lap times, racing inputs, session recordings, driver names, iRacing IDs, or detailed driving data to RacePulse servers.

RacePulse does collect limited app usage analytics to understand whether the app is being used, which features and supported games need the most development attention, and whether updates and connections work. This may include app startup events, app version, operating system and Windows version, supported game detected, and a randomly generated installation ID. The installation ID is not based on hardware information and is used only to measure app usage. From version 0.4.1, RacePulse also sends a short summary of each time you use it the next time it starts, such as how long it was open, which widgets and features were used, and whether it closed normally. The [Privacy Policy](/privacy/) lists exactly what it includes.

RacePulse may keep limited server logs for security, abuse prevention, rate limiting, and blocking malicious traffic. Website analytics may also be used to understand page views, downloads, referrers, browser/device information, and general product interest.

RacePulse does not sell personal data and does not use analytics to collect racing telemetry.

---

## Troubleshooting

### RacePulse doesn't detect my game

- Make sure you are **in a session** (on track or in replay), not in the main menu — shared memory is only populated while on track
- For **ACC and Assetto Corsa**: they can't be detected from the main menu, so load into a session first
- For **Le Mans Ultimate**: verify that **Settings > Gameplay > Enable Plugins** is turned on
- For **F1 25/F1 24**: verify that **UDP Telemetry** is enabled and the **UDP Port** is 22025, 22024, 22023, or 20777. If it still doesn't connect, another app may be using that port; pick another of these ports in the game, or close the other app.
- Check the connection pill at the top of RacePulse. It reads **LIVE** when telemetry arrives. If it reads **Not connected**, click it to auto-detect, or open its menu to pick your sim.

### Overlays don't appear on screen

- Make sure at least one widget is **on** in All widgets
- With **Only show with telemetry** on, locked widgets stay hidden until a game sends telemetry. Switch to **Arrange** to see them, or turn the setting off in the bottom bar's **Overlay** menu.
- If a widget card says **Hidden in** your sim, that game doesn't send the data the widget needs
- Run your game in **Borderless Windowed** mode for best overlay compatibility
- If using multiple monitors, **Overlay › Recenter all widgets** brings every widget to the main window's screen

### Windows SmartScreen blocks RacePulse

On first launch, Windows may show a SmartScreen warning because the app is not code-signed. Click **"More info"** then **"Run anyway"**. This only happens once.

### Widgets show stale data after closing a game

RacePulse disconnects automatically shortly after the game stops sending telemetry. You can also choose **Disconnect** from the connection pill's menu to clear stale data immediately.

---

Still have questions? Open an issue on [GitHub](https://github.com/gAlexander77/racepulse-release/issues) or check the [installation guide](/install/) for step-by-step setup instructions.
