---
title: "Uninstall"
subtitle: "Remove RacePulse from your PC, with or without your settings."
description: "How to uninstall RacePulse: turn off VR widgets, delete RacePulse.exe, and remove its settings, browser data, and VR files, by hand or with one PowerShell command."
ogImage: "assets/application-preview.jpg"
---

RacePulse is a single portable file with no installer, so there's no uninstaller either. Removing it takes a minute.

## 1. Turn off VR widgets

If you used VR widgets, open **VR widgets** in RacePulse and turn off **Show widgets in VR**. That removes RacePulse's VR layer from your games. If you never turned VR widgets on, skip this step.

## 2. Close RacePulse

Close RacePulse's main window. Its widgets close with it.

## 3. Delete RacePulse.exe

Delete `RacePulse.exe` from where you saved it, and `RacePulse.exe.sha256` if you downloaded it too.

## 4. Remove your settings and data (optional)

RacePulse keeps your settings so they're there if you come back. To remove them too, delete these folders. Paste a path into File Explorer's address bar to open it.

| Folder | What's in it |
| --- | --- |
| `%USERPROFILE%\.racepulse` | Your settings, widget layouts, VR layouts, recorded sessions, and RacePulse's random installation ID |
| `%APPDATA%\RacePulse.exe` | The window's browser data. If you renamed `RacePulse.exe`, this folder has the name you gave it. |
| `%LOCALAPPDATA%\RacePulse` | VR widget files, if you used VR widgets |
| `%APPDATA%\RacePulse` | A layout file from older versions |
| `%APPDATA%\racepulse-widget-*.exe` | Widget data from versions before 0.4.0 |

Settings you changed in your sims for RacePulse, such as F1's UDP telemetry or **Enable Plugins** in Le Mans Ultimate, stay as you set them. Change them back in the game if you like.

## Remove everything with PowerShell

Close RacePulse first. This removes every folder above, including your settings, and RacePulse's VR layer if it's still registered. It doesn't touch other apps' files or OpenXR layers.

```powershell
$k = 'HKCU:\Software\Khronos\OpenXR\1\ApiLayers\Implicit'
if (Test-Path $k) { (Get-Item $k).Property | Where-Object { $_ -like '*\RacePulse\VR\*' } | ForEach-Object { Remove-ItemProperty -Path $k -Name $_ } }
Remove-Item -Recurse -Force "$env:USERPROFILE\.racepulse", "$env:APPDATA\RacePulse.exe", "$env:LOCALAPPDATA\RacePulse", "$env:APPDATA\RacePulse" -ErrorAction SilentlyContinue
Get-ChildItem $env:APPDATA -Directory -Filter 'racepulse-widget-*.exe' | Remove-Item -Recurse -Force
```

Then delete `RacePulse.exe`.

## Already deleted RacePulse with VR widgets on?

RacePulse's VR layer stays registered for your Windows user. It does nothing without RacePulse, but you can remove it, and keep your settings, with just these two lines in PowerShell:

```powershell
$k = 'HKCU:\Software\Khronos\OpenXR\1\ApiLayers\Implicit'
if (Test-Path $k) { (Get-Item $k).Property | Where-Object { $_ -like '*\RacePulse\VR\*' } | ForEach-Object { Remove-ItemProperty -Path $k -Name $_ } }
```

Or download RacePulse again, turn off **Show widgets in VR**, and close it.
