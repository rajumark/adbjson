# shell pm disable

## Description
Disable a package on the connected Android device.

## Command
```bash
adbjson shell pm disable <package>
```

## Equivalent ADB Command
```bash
adb shell pm disable <package>
```

## Sample JSON Output
```json
{
  "disable_result": {
    "success": true,
    "message": "Package disabled successfully"
  }
}
```

## Sample YAML Output
```bash
adbjson shell pm disable com.example.app --format yaml
```
```yaml
disable_result:
  success: true
  message: Package disabled successfully
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires package name as argument
- Disables a package, making it inaccessible to the user
