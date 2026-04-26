# adbjson

A cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands and outputs structured JSON.

## Quick Start

### Installation

#### Option 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/rajumark/adbjson.git
cd adbjson

# Download platform-tools for your platform
# macOS
curl -L -o platform-tools-darwin.zip https://dl.google.com/android/repository/platform-tools-latest-darwin.zip
unzip -q platform-tools-darwin.zip -d platform-tools
mv platform-tools/platform-tools platform-tools/platform-tools-darwin
rm platform-tools-darwin.zip

# Linux
curl -L -o platform-tools-linux.zip https://dl.google.com/android/repository/platform-tools-latest-linux.zip
unzip -q platform-tools-linux.zip -d platform-tools
mv platform-tools/platform-tools platform-tools/platform-tools-linux
rm platform-tools-linux.zip

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://dl.google.com/android/repository/platform-tools-latest-windows.zip" -OutFile "platform-tools-windows.zip"
Expand-Archive -Path "platform-tools-windows.zip" -DestinationPath "platform-tools"
Move-Item "platform-tools\platform-tools" "platform-tools\platform-tools-windows"
Remove-Item "platform-tools-windows.zip"

# Build
go mod tidy
go build -o adbjson
```

#### Option 2: Download Pre-built Binary

Download from [Releases](https://github.com/rajumark/adbjson/releases) and extract.

### Usage

```bash
# List connected devices
./adbjson devices

# Get ADB version
./adbjson adb-version

# Compact JSON output
./adbjson devices --compact

# Show CLI version
./adbjson --version
```

## Features

- Execute ADB commands and get JSON output
- Bundled ADB support (no system installation required)
- Pretty-printed or compact JSON output
- Cross-platform (macOS, Linux, Windows)

## Documentation

For detailed project documentation, see [PROJECT.md](PROJECT.md).
