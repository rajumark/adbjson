# shell uname

## Description
Gets system information from the connected Android device.

## Command
```bash
adbjson shell uname
```

## Equivalent ADB Command
```bash
adb shell uname -a
```

## Sample JSON Output
```json
{
  "system_info": {
    "kernel_name": "Linux",
    "node_name": "localhost",
    "kernel_release": "4.14.0-android-g5f8c9a9",
    "kernel_version": "#1",
    "machine": "aarch64",
    "processor": "unknown",
    "hardware": "unknown",
    "os": "unknown"
  }
}
```

## Sample YAML Output
```bash
adbjson shell uname --format yaml
```
```yaml
system_info:
  kernel_name: Linux
  node_name: localhost
  kernel_release: 4.14.0-android-g5f8c9a9
  kernel_version: "#1"
  machine: aarch64
  processor: unknown
  hardware: unknown
  os: unknown
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Returns kernel, node, and machine information
