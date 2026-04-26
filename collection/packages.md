# packages

## Description
Lists all installed packages on the connected Android device.

## Command
```bash
adbjson packages
```

## Equivalent ADB Command
```bash
adb shell pm list packages
```

## Sample JSON Output
```json
{
  "packages": [
    {
      "name": "com.android.systemui"
    },
    {
      "name": "com.android.settings"
    }
  ],
  "count": 2
}
```

## Sample YAML Output
```bash
adbjson packages --format yaml
```
```yaml
packages:
- name: com.android.systemui
- name: com.android.settings
count: 2
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all packages including system and third-party
- Output is limited to first 5 packages in documentation (full list in actual output)
