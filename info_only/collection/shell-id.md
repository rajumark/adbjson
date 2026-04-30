# shell id

## Description
Gets user ID information from the connected Android device.

## Command
```bash
adbjson shell id
```

## Equivalent ADB Command
```bash
adb shell id
```

## Sample JSON Output
```json
{
  "user_info": {
    "user_id": "0(root)",
    "group_id": "0(root)"
  }
}
```

## Sample YAML Output
```bash
adbjson shell id --format yaml
```
```yaml
user_info:
  user_id: 0(root)
  group_id: 0(root)
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Returns user and group IDs
