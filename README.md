# adbjson

A cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands and outputs structured JSON.

## Features

- Execute ADB commands and get JSON output
- Bundled ADB support (no system installation required)
- Pretty-printed or compact JSON output
- Extensible architecture for adding more commands

## Installation

### Option 1: Build from Source

#### Prerequisites

1. Install Go (https://golang.org/dl/)

#### Step-by-Step Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/rajumark/adbjson.git
   cd adbjson
   ```

2. Download platform-tools for your platform:
   ```bash
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
   ```

3. Build the project:
   ```bash
   go mod tidy
   go build -o adbjson
   ```

4. (Optional) Move to PATH:
   ```bash
   # macOS/Linux
   sudo mv adbjson /usr/local/bin/

   # Windows
   # Add adbjson.exe to your PATH
   ```

### Option 2: Download Pre-built Binary

Download pre-built binaries from the [Releases](https://github.com/rajumark/adbjson/releases) page. Each release includes:
- `adbjson-ubuntu-latest.tar.gz` - Linux binary with platform-tools
- `adbjson-macos-latest.tar.gz` - macOS binary with platform-tools
- `adbjson-windows-latest.zip` - Windows binary with platform-tools

Extract and run:
```bash
# macOS/Linux
tar -xzf adbjson-macos-latest.tar.gz
cd adbjson-macos-latest
./adbjson devices

# Windows
unzip adbjson-windows-latest.zip
cd adbjson-windows-latest
adbjson.exe devices
```

## Usage

### List Devices

```bash
./adbjson devices
```

Output:
```json
{
  "devices": [
    {
      "id": "emulator-5554",
      "status": "device"
    }
  ]
}
```

### Get ADB Version

```bash
./adbjson adb-version
```

Output:
```json
{
  "version": "1.0.41"
}
```

### Flags

- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output (overrides --pretty)
- `--version` - Show version

### Examples

```bash
# Pretty print (default)
./adbjson devices

# Compact output
./adbjson devices --compact

# Show version
./adbjson --version
```

## Project Structure

```
adbjson/
 ├── main.go
 ├── cmd/
 │    ├── root.go
 │    └── devices.go
 ├── internal/
 │    ├── adb/
 │    │    └── executor.go
 │    ├── parser/
 │    │    └── devices_parser.go
 │    ├── model/
 │    │    └── device.go
 │    └── platform/
 │         └── platform.go
 ├── platform-tools/
 │    ├── platform-tools-darwin/
 │    ├── platform-tools-linux/
 │    └── platform-tools-windows/
 ├── go.mod
 └── go.sum
```

## Supported Commands

- `devices` - List connected ADB devices
- `adb-version` - Get ADB version information

## Future Extensibility

The architecture is designed to easily add more commands:
- `packages` - List installed packages
- `battery` - Get battery status
- `logcat` - Stream logcat output

## License

MIT
