# shell pm list permissions

## Description
Lists all permissions on the connected Android device.

## Command
```bash
adbjson shell pm list permissions
```

## Equivalent ADB Command
```bash
adb shell pm list permissions
```

## Sample JSON Output
```json
{
  "permissions": [
    {
      "name": "android.permission.INTERNET"
    },
    {
      "name": "android.permission.CAMERA"
    }
  ],
  "count": 2
}
```

## Sample YAML Output
```bash
adbjson shell pm list permissions --format yaml
```
```yaml
permissions:
- name: android.permission.INTERNET
- name: android.permission.CAMERA
count: 2
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all system and app permissions
- Output is limited to first 5 permissions in documentation (full list in actual output)
