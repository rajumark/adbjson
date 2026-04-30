# shell pm list features

## Description
Lists all device features on the connected Android device.

## Command
```bash
adbjson shell pm list features
```

## Equivalent ADB Command
```bash
adb shell pm list features
```

## Sample JSON Output
```json
{
  "features": [
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
adbjson shell pm list features --format yaml
```
```yaml
features:
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
- Lists all hardware and software features
- Output is limited to first 5 features in documentation (full list in actual output)
