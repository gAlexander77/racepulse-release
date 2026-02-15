# RacePulse

Real-time telemetry overlays for sim racing. Lightweight, modular widgets that float over your game — speed, inputs, lap times, fuel, and more. Auto-detects iRacing, F1 25, and F1 24. Single .exe, no installer, no config.

**[racepulse.racing](https://racepulse.racing/)** — See all features, previews, and download the latest release.

## Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![Wails](https://img.shields.io/badge/Wails_v2-DF0000?style=for-the-badge&logo=webassembly&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-0078D4?style=for-the-badge&logo=windows&logoColor=white)

## Preview

**Main Application**

![RacePulse Preview](site/static/assets/application-preview.gif)

**Widget Customization**

![Widget customization preview](site/static/assets/widget-customization-preview.gif)

**Overlay Widget**

![iRacing inputs telemetry overlay at Sebring](site/static/assets/iracing-input-telemetry-overlay-sebring.gif)

## What this repository is for

This repo is the public release and documentation home for RacePulse.

- Download signed Windows executables
- Read feature documentation and setup guides
- Track release notes and known issues

Source code is not published in this repository.

## Download

Get the latest build from [Releases](../../releases).

Each release includes:

- `racepulse.exe` (main application)
- Optional release notes and checksums

## Supported Games

| Game | Transport | Setup |
|------|-----------|-------|
| **iRacing** | Shared memory | No configuration needed — detected automatically when iRacing is running |
| **F1 25** | UDP | Enable UDP telemetry in game settings (default port 22025) |
| **F1 24** | UDP | Enable UDP telemetry in game settings (default port 22024) |

RacePulse automatically detects which game is running and begins capturing telemetry.

## Core Features

- Real-time racing telemetry overlays for iRacing, F1 25, and F1 24
- Modular widget system (speed, gear, RPM, delta, inputs, fuel, G-force, and more)
- Per-widget customization with persistent settings
- Drag-and-drop widget positioning with click-through overlay mode
- Session recording to local storage
- Single executable distribution — no installer required

## System Requirements

- Windows 10/11 (64-bit)
- Graphics driver with desktop composition enabled
- For F1 games: UDP telemetry enabled in game settings

## Documentation

- See `docs/features.md` for an end-user feature overview
- See `docs/install.md` for installation and first-run setup
- See `CHANGELOG.md` for release history

## License

This repository is distributed under a proprietary binary-use license.
See `LICENSE` and `EULA.md`.
