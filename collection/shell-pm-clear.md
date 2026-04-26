# shell pm clear

## Description
Clear all data associated with a package on the connected Android device.

## Command
```bash
adbjson shell pm clear <package>
```

## Equivalent ADB Command
```bash
adb shell pm clear <package>
```

## Sample JSON Output
```json
{
  "clear_result": {
    "success": true,
    "message": "Package data cleared successfully"
  }
}
```

## Sample YAML Output
```bash
adbjson shell pm clear com.example.app --format yaml
```
```yaml
clear_result:
  success: true
  message: Package data cleared successfully
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires package name as argument
- Clears all application data including cache
