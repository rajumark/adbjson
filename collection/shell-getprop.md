# shell getprop

## Description
Gets all system properties from the connected Android device.

## Command
```bash
adbjson shell getprop
```

## Equivalent ADB Command
```bash
adb shell getprop
```

## Sample JSON Output
```json
{
  "properties": [
    {
      "key": "ro.build.id",
      "value": "OPM1.171019.021"
    },
    {
      "key": "ro.build.display.id",
      "value": "OPM1.171019.021"
    }
  ],
  "count": 2
}
```

## Sample YAML Output
```bash
adbjson shell getprop --format yaml
```
```yaml
properties:
- key: ro.build.id
  value: OPM1.171019.021
- key: ro.build.display.id
  value: OPM1.171019.021
count: 2
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all system properties
- Output is limited to first 5 properties in documentation (full list in actual output)
