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

Delete `RacePulse.exe` from where you saved it, and `RacePulse.exe.sha256` if you downloaded it too. After an update there can also be a hidden `RacePulse.exe.old` beside it, which RacePulse normally removes the next time it starts; delete it too if it's there.

## 4. Remove your settings and data (optional)

RacePulse keeps your settings so they're there if you come back. To remove them too, delete these. Paste a path into File Explorer's address bar to open it.

| Folder or file | What's in it |
| --- | --- |
| `%USERPROFILE%\.racepulse` | Your settings, widget layouts, VR layouts, recorded sessions, and RacePulse's random installation ID |
| `%APPDATA%\RacePulse.exe` | The window's browser data. If you renamed `RacePulse.exe`, this folder has the name you gave it. |
| `%LOCALAPPDATA%\RacePulse\VR` | VR widget files, if you used VR widgets |
| `%APPDATA%\RacePulse\overlay_layout.json` | A layout file from older versions |
| `%APPDATA%\racepulse-widget-*.exe` | Widget data from versions before 0.4.0 |

If the `RacePulse` folders in `%LOCALAPPDATA%` and `%APPDATA%` are empty afterwards, you can delete them too.

Settings you changed in your sims for RacePulse, such as F1's UDP telemetry or **Enable Plugins** in Le Mans Ultimate, stay as you set them. Change them back in the game if you like.

## Remove everything with PowerShell

This removes everything in the table above, including your settings, and RacePulse's VR layer if it's still registered. It removes only RacePulse's own files and entry: other apps' files and OpenXR layers stay as they are, and it doesn't need administrator rights. Open PowerShell from the Start menu, paste all of it, and press Enter. If RacePulse is still open, it does nothing and asks you to close it first. If you renamed `RacePulse.exe`, close it yourself before running this, and delete its browser-data folder in `%APPDATA%` by hand.

```powershell
if (Get-Process -Name RacePulse -ErrorAction SilentlyContinue) { Write-Host 'Close RacePulse first, then run this again.' } else {
    $layers = [Microsoft.Win32.Registry]::CurrentUser.OpenSubKey('Software\Khronos\OpenXR\1\ApiLayers\Implicit', $true)
    if ($layers) { $layers.GetValueNames() | Where-Object { $_ -like '*\RacePulse\VR\*' } | ForEach-Object { $layers.DeleteValue($_) }; $layers.Close() }
    $files = "$env:USERPROFILE\.racepulse", "$env:APPDATA\RacePulse.exe", "$env:LOCALAPPDATA\RacePulse\VR", "$env:APPDATA\RacePulse\overlay_layout.json"
    foreach ($p in $files) { if (Test-Path -LiteralPath $p) { Remove-Item -LiteralPath $p -Recurse -Force } }
    foreach ($p in "$env:LOCALAPPDATA\RacePulse", "$env:APPDATA\RacePulse") {
        if ((Test-Path -LiteralPath $p) -and -not (Get-ChildItem -LiteralPath $p -Force)) { Remove-Item -LiteralPath $p }
    }
    Get-ChildItem -LiteralPath $env:APPDATA -Directory -Filter 'racepulse-widget-*.exe' | Remove-Item -Recurse -Force
    if ($files | Where-Object { Test-Path -LiteralPath $_ }) { Write-Host 'Some files were still in use. Wait a few seconds and run this again.' } else { Write-Host 'Done. Now delete RacePulse.exe.' }
}
```

## Already deleted RacePulse with VR widgets on?

RacePulse's VR layer stays registered for your Windows user. It does nothing without RacePulse, but you can remove it, and keep your settings, with just these two lines in PowerShell:

```powershell
$layers = [Microsoft.Win32.Registry]::CurrentUser.OpenSubKey('Software\Khronos\OpenXR\1\ApiLayers\Implicit', $true)
if ($layers) { $layers.GetValueNames() | Where-Object { $_ -like '*\RacePulse\VR\*' } | ForEach-Object { $layers.DeleteValue($_) }; $layers.Close() }
```

Or download RacePulse again, turn off **Show widgets in VR**, and close it.
