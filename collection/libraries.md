# libraries

## Description
Lists all device libraries on the connected Android device.

## Command
```bash
adbjson libraries
```

## Equivalent ADB Command
```bash
adb shell pm list libraries
```

## Sample JSON Output
```json
{
  "libraries": [
    {
      "name": "android.hardware.camera"
    },
    {
      "name": "android.hardware.wifi"
    }
  ],
  "count": 2
}
```

## Sample YAML Output
```bash
adbjson libraries --format yaml
```
```yaml
libraries:
- name: android.hardware.camera
- name: android.hardware.wifi
count: 2
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all shared libraries
- Output is limited to first 5 libraries in documentation (full list in actual output)
