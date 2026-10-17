---
title: "VR Guide"
subtitle: "Set up RacePulse widgets in your VR headset with SteamVR, and place them from inside the headset."
description: "How to use RacePulse VR widgets: the recommended SteamVR setup, settings for RacePulse and SteamVR, placing widgets from the SteamVR dashboard, VR layouts, and fixes for common problems."
ogImage: "assets/application-preview.jpg"
---

RacePulse can show your widgets in your VR headset. VR widgets are **experimental**, **free**, and **off by default**, so nothing VR loads until you turn them on. They need RacePulse 0.4.1 or later.

**We recommend running your sim through SteamVR.** RacePulse adds its own tab to the SteamVR dashboard, so you can show, move, resize, and set up widgets with your controller without taking the headset off. It works with any headset SteamVR supports.

## What you need

- RacePulse 0.4.1 or later
- [SteamVR](https://store.steampowered.com/app/250820/SteamVR/), installed through Steam
- A headset SteamVR supports. A Meta Quest connects through Quest Link (cable), Air Link, or Steam Link.
- A sim with a VR mode: iRacing, Assetto Corsa Competizione, Assetto Corsa, or Le Mans Ultimate

## Recommended settings

| Where | Setting | Recommended | Why |
| --- | --- | --- | --- |
| RacePulse › **VR widgets** › **Overview** | Show widgets in VR | **On** | Turns VR widgets on. |
| RacePulse › **VR widgets** › **Headset** | SteamVR overlays | **On** | How your widgets reach the headset through SteamVR. |
| RacePulse › **VR widgets** › **Headset** | Native OpenXR games | **Off** | Only needed for games that run outside SteamVR. With it off, RacePulse loads nothing into your games. |
| SteamVR › **Settings** › **OpenXR** | Current OpenXR runtime | **SteamVR** | Games that use OpenXR then run through SteamVR too, so RacePulse's dashboard tab works in them. |
| Your sim | VR mode | **SteamVR / OpenVR**, or OpenXR with SteamVR as the runtime | Keeps the game in SteamVR. |
| Each VR widget › **Customize** | Frame Rate | **30 FPS** for timing widgets, **60 FPS** for Inputs Telemetry | Relative, Standings, Fuel, and Lap Log don't need more than 30, and a lower rate leaves more of your PC for the game. |

## Set it up

1. In RacePulse, open **VR widgets** in the side menu and turn on **Show widgets in VR**.
2. The **Get set up** guide on the Overview takes you through three steps. For the first, choose **Use my desktop widgets** to bring the widgets you use on your monitor, each with a copy of its settings, or **Choose widgets** to pick others.
3. On the **Headset** tab, check that **SteamVR overlays** is on, and turn **Native OpenXR games** off.
4. In SteamVR's **Settings**, open **OpenXR**, and if SteamVR isn't the current runtime, choose **Set SteamVR as OpenXR Runtime**. If you don't see OpenXR, turn on **Advanced Settings** first.
5. Put on your headset, connect it, and start SteamVR. The Overview's second step shows **Connected through SteamVR** with your headset's name. Your desktop widgets hide while you're in VR, and come back when you leave.
6. Start your sim in its SteamVR mode and get on track. Your widgets show live data in the headset.

## Place widgets from inside the headset

Open the SteamVR dashboard and choose the **RacePulse** tab. Every widget gets a frame while the tab is open, and the selected one is amber. Use the tabs on the left with your controller's laser:

- **Widgets:** show, hide, or remove a widget, open its settings, or add one.
- **Layouts:** switch to another VR layout, duplicate one, or start a new one.
- **Place:** move the selected widget, change its distance, size, turn, and tilt. Click a button for one small step, or hold it to keep moving. **Small/Big steps** changes the step size, and **Recenter on my view** brings every widget in front of you.
- **Settings:** the selected widget's own settings.
- **Headset:** the connection, the frames, Recenter, and **Reset to the seated origin**.

You can also use the keyboard, which works outside the dashboard too:

| Keys | What they do |
| --- | --- |
| **Ctrl+Alt+Home** | Start or end editing |
| **Ctrl+Alt+End** | Recenter the widgets on your view |
| **Ctrl+Alt+Page Down** | Switch to the next VR layout |
| **Tab** (while editing) | Select the next widget |
| **Arrow keys** (while editing) | Move it; **Page Up / Page Down** change its distance |
| **=** and **-** (while editing) | Make it bigger or smaller |
| **Ctrl+arrow keys** (while editing) | Turn and tilt it |
| **Shift** (while editing) | Makes every step five times larger |
| **Enter** or **Escape** | Finish editing |

Recentering moves only your RacePulse widgets, never the game's view.

## VR layouts and settings

- VR has its **own layouts**, kept apart from your desktop layouts. Manage them on the **Layouts** tab of the VR widgets page, or in the headset.
- VR starts with the **layout you selected last**.
- Each VR widget has its **own settings**. Changing a VR widget's settings on the **Widgets** tab, or in the headset, leaves the desktop widget as it is.
- Everything on the VR widgets page works with VR off, so you can set things up before you put the headset on.

## Without SteamVR

RacePulse can also draw its widgets inside Direct3D 11 games running on another OpenXR runtime, such as Meta's own runtime with Quest Link. Turn **Native OpenXR games** on, and set the game's VR mode to OpenXR. This route has no dashboard tab, so you place widgets with the keyboard shortcuts, and a bar under your view shows the selected widget while you edit. We recommend SteamVR whenever your game supports it.

Games in the old Oculus mode aren't supported. Switch them to OpenXR or SteamVR in the game's VR settings.

## If something isn't right

- **No widgets in the headset:** check the Overview. The status should say you're in VR, and the second step should show **Connected through SteamVR**. Widgets follow **Only show with telemetry**: they stay hidden until your sim sends telemetry, unless you're editing them or that setting is off.
- **My desktop widgets disappeared:** that's expected while you're in VR. They come back about 10 seconds after you leave.
- **A shortcut doesn't work:** another app may be using it. The VR widgets page names any shortcut that's taken, and RacePulse picks it up as soon as that app lets it go.
- **Widgets are in the wrong place:** press **Ctrl+Alt+End** to recenter them on your view, or use **Reset to the seated origin** on the Headset tab.

## Turning VR off

Turn off **Show widgets in VR** on the VR widgets page. That also removes RacePulse's VR layer from your games. Your VR layouts stay saved for next time.

Removing RacePulse? Turn off **Show widgets in VR** first, then follow the [Uninstall](/uninstall/) page.

If something looks wrong in your headset, please [open an issue on GitHub](https://github.com/gAlexander77/racepulse-release/issues) with your headset, the sim, and what you saw.
