---
title: "FAQ"
subtitle: "Frequently asked questions about RacePulse — setup, supported games, widgets, troubleshooting, and more."
description: "Answers to common questions about RacePulse: free sim racing telemetry overlay for iRacing, ACC, Assetto Corsa, Le Mans Ultimate, F1 25, and F1 24."
ogImage: "assets/application-preview.gif"
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
        "text": "RacePulse supports six games: iRacing, Assetto Corsa Competizione (ACC), Assetto Corsa, Le Mans Ultimate (LMU), F1 25, and F1 24. iRacing, ACC, AC, and LMU use Windows shared memory (no setup needed). F1 25 and F1 24 use UDP telemetry (one-time toggle in game settings)."
      }
    },
    {
      "@type": "Question",
      "name": "Do I need to install plugins or configure anything?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "For iRacing, ACC, and Assetto Corsa — no. RacePulse detects these games automatically via shared memory. For Le Mans Ultimate, you need to enable the Plugins option in Settings > Gameplay (one-time). For F1 25 and F1 24, you need to enable UDP Telemetry in the game's settings (one-time)."
      }
    },
    {
      "@type": "Question",
      "name": "Does RacePulse work in VR?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse overlays are standard Windows windows, so they appear on your desktop monitor rather than inside a VR headset. If your sim supports a desktop mirror or spectator view, the overlays will be visible there. For in-headset overlays, you would need a VR overlay tool like OVR Toolkit to bring the RacePulse widgets into your VR environment."
      }
    },
    {
      "@type": "Question",
      "name": "How many widgets does RacePulse have?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "RacePulse has 12 overlay widgets: Speed, Gear, RPM, Lap Time, Delta, Inputs, Fuel, G-Force, Telemetry Core (all-in-one display), Inputs Telemetry (scrolling input graph), Flat Track Map (car positions on track), and Relative (nearby drivers with gap times — iRacing only)."
      }
    },
    {
      "@type": "Question",
      "name": "Does RacePulse require an installer?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "No. RacePulse is a single portable executable (about 21MB). Download RacePulse.exe, run it, and you're ready. No installation, no dependencies, no runtime to install."
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
        "text": "Yes. RacePulse widgets are standard windows that appear in OBS window and game captures. If you use Game Capture in OBS, you may need to add each widget as a separate Window Capture source, or switch to Display Capture to pick them up automatically."
      }
    },
    {
      "@type": "Question",
      "name": "How do I update RacePulse?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Download the latest RacePulse.exe from the releases page and replace your existing file. Widget positions, sizes, and settings are stored separately and will be preserved across updates."
      }
    }
  ]
}
</script>

## General

### Is RacePulse free?

Yes. RacePulse is completely free — no subscription, no trial period, no premium tier, and no locked features. Every widget and every supported game is available at no cost.

### Does RacePulse require an installer?

No. RacePulse is a single portable executable, about 21MB. Download `RacePulse.exe`, run it, and you're ready. No installation wizard, no runtime dependencies, no `.NET` or `Java` to install.

### What operating system does RacePulse run on?

Windows 10 and Windows 11 (64-bit). iRacing, ACC, Assetto Corsa, and Le Mans Ultimate support is Windows-only because those games use Windows shared memory for telemetry. F1 25 and F1 24 use UDP which is cross-platform in principle, but the overlay windows currently require Windows.

### How do I update RacePulse?

Download the latest `RacePulse.exe` from the [Releases page](https://github.com/gAlexander77/racepulse-release/releases) and replace your existing file. Widget positions, sizes, and all customized settings are stored separately in `~/.racepulse/` and will be preserved across updates.

---

## Supported Games

### What sim racing games does RacePulse support?

RacePulse supports six games:

- [**iRacing**](/iracing-overlay/) — Windows shared memory, no setup needed
- [**Assetto Corsa Competizione**](/acc-overlay/) — Windows shared memory, no setup needed
- [**Assetto Corsa**](/assetto-corsa-overlay/) — Windows shared memory, no setup needed
- [**Le Mans Ultimate**](/lmu-overlay/) — rFactor 2 shared memory plugin, enable Plugins in LMU settings once
- [**F1 25**](/f1-25-overlay/) — UDP telemetry on port 22025, enable in game settings once
- [**F1 24**](/f1-24-overlay/) — UDP telemetry on port 22024, enable in game settings once

### Do I need to install plugins or configure anything?

For **iRacing, ACC, and Assetto Corsa** — no. RacePulse detects these games automatically via shared memory with zero configuration.

For **Le Mans Ultimate** — one-time step: go to **Settings > Gameplay** and enable **Plugins**. This activates the rFactor 2 shared memory plugin that RacePulse reads from.

For **F1 25 and F1 24** — one-time step: go to **Settings > Telemetry Settings** and set **UDP Telemetry** to **On**. The default UDP port (22025 for F1 25, 22024 for F1 24) works automatically.

### Can I switch between games without reconfiguring?

Yes. RacePulse auto-detects which supported game is currently running. Close one sim, launch another, and RacePulse connects to the new game automatically. Your widget layout and settings carry over between games.

### Does RacePulse work with iRacing in VR?

RacePulse overlays are standard Windows windows that appear on your desktop monitor, not inside the VR headset. If you use a desktop mirror or spectator view, overlays will be visible there. To bring RacePulse widgets into your VR environment, you can use a VR overlay tool like **OVR Toolkit** or **Desktop+** to project the windows into headset space.

---

## Widgets

### How many widgets does RacePulse have?

RacePulse has **12 overlay widgets**:

- **Core (8):** Speed, Gear, RPM, Lap Time, Delta, Inputs, Fuel, G-Force
- **Advanced (4):** Telemetry Core (all-in-one display), Inputs Telemetry (scrolling input graph), Flat Track Map (car positions), Relative (nearby drivers with gaps)

See the full [features page](/features/) for details on each widget.

### Do all widgets work with every game?

Most widgets work across all six supported games. Two exceptions:

- **Relative** — iRacing only (requires iRating, safety rating, and license class data)
- **Flat Track Map** — iRacing, Le Mans Ultimate, and F1 only (requires per-car position data)

All other widgets work identically across iRacing, ACC, Assetto Corsa, LMU, F1 25, and F1 24.

### Can I customize widget appearance and behavior?

Yes. Every widget has per-instance settings:

- **Units** — km/h or mph for speed, Celsius or Fahrenheit for temperatures
- **Visibility toggles** — show or hide individual display elements
- **Frame rate** — 60, 30, or 15 FPS per widget
- **Position and size** — drag to reposition in Edit Mode, resize as needed

All settings persist across sessions automatically. You can reset any widget to defaults through the customization panel.

### How do I reposition widgets on screen?

Click **Edit Mode** in the RacePulse dashboard. Widgets become draggable — move them anywhere on your screen. Click Edit Mode again to lock them in place with click-through enabled so they don't interfere with your game.

---

## Performance & Streaming

### Will RacePulse affect my game performance?

RacePulse is designed to be lightweight. It reads telemetry passively — shared memory reads and UDP packet reception add no processing load to the game itself. Each widget runs as a small transparent window. Most users see no measurable FPS impact.

### Can I use RacePulse while streaming or recording?

Yes. RacePulse widgets are standard windows that appear in OBS and other capture tools.

- **Display Capture** in OBS picks up widgets automatically
- **Game Capture** may not show widgets — add each widget as a separate **Window Capture** source, or switch to Display Capture
- Widgets have transparent backgrounds, so they overlay cleanly in both the game and the stream

### Does RacePulse send data to the internet?

RacePulse reads telemetry locally from your machine (shared memory or localhost UDP). It does not send your telemetry data, lap times, or any racing data to external servers. The only network activity is checking for updates and anonymous usage analytics (page views), both of which can be blocked by a firewall if preferred.

---

## Troubleshooting

### RacePulse doesn't detect my game

- Make sure you are **in a session** (on track or in replay), not in the main menu — shared memory is only populated while on track
- For **Le Mans Ultimate**: verify that **Settings > Gameplay > Enable Plugins** is turned on
- For **F1 25/F1 24**: verify that **UDP Telemetry** is enabled in the game's Telemetry Settings
- Check that the RacePulse dashboard shows "Searching" — if it shows "Disconnected," click Connect

### Overlays don't appear on screen

- Make sure at least one widget is **enabled** (toggled on) in the dashboard
- Try toggling **Edit Mode** on and off
- Run your game in **Borderless Windowed** mode for best overlay compatibility
- If using multiple monitors, check that widgets aren't positioned on a different display

### Windows SmartScreen blocks RacePulse

On first launch, Windows may show a SmartScreen warning because the app is not code-signed. Click **"More info"** then **"Run anyway"**. This only happens once.

### Widgets show stale data after closing a game

RacePulse automatically disconnects after approximately 10 seconds of no telemetry data. You can also click **Disconnect** manually in the dashboard to clear stale data immediately.

---

Still have questions? Open an issue on [GitHub](https://github.com/gAlexander77/racepulse-release/issues) or check the [installation guide](/install/) for step-by-step setup instructions.
