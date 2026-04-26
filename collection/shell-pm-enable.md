# shell pm enable

## Description
Enable a package on the connected Android device.

## Command
```bash
adbjson shell pm enable <package>
```

## Equivalent ADB Command
```bash
adb shell pm enable <package>
```

## Sample JSON Output
```json
{
  "enable_result": {
    "success": true,
    "message": "Package enabled successfully"
  }
}
```

## Sample YAML Output
```bash
adbjson shell pm enable com.example.app --format yaml
```
```yaml
enable_result:
  success: true
  message: Package enabled successfully
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires package name as argument
- Enables a previously disabled package
