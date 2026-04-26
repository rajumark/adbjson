# adbjson

A cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands and outputs structured JSON.

## Features

- Execute ADB commands and get JSON output
- Bundled ADB support (no system installation required)
- Pretty-printed or compact JSON output
- Extensible architecture for adding more commands

## Installation

### Prerequisites

1. Install Go (https://golang.org/dl/)
2. Download platform-tools for your platform from https://developer.android.com/tools/releases/platform-tools

### Setup

1. Clone or download this project
2. Download the platform-tools zip files:
   - macOS: `platform-tools-latest-darwin.zip`
   - Linux: `platform-tools-latest-linux.zip`
   - Windows: `platform-tools-latest-windows.zip`

3. Extract each zip to the `platform-tools` directory:
   ```
   adbjson/
    ├── platform-tools/
    │    ├── platform-tools-darwin/    # Extract macOS zip here
    │    ├── platform-tools-linux/     # Extract Linux zip here
    │    └── platform-tools-windows/   # Extract Windows zip here
   ```

4. Build the project:
   ```bash
   go mod tidy
   go build -o adbjson
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
