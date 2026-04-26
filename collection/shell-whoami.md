# shell whoami

## Description
Gets the current user from the connected Android device.

## Command
```bash
adbjson shell whoami
```

## Equivalent ADB Command
```bash
adb shell whoami
```

## Sample JSON Output
```json
{
  "current_user": {
    "username": "root"
  }
}
```

## Sample YAML Output
```bash
adbjson shell whoami --format yaml
```
```yaml
current_user:
  username: root
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Returns current username
