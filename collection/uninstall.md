# uninstall

## Description
Uninstall a package from the connected Android device.

## Command
```bash
adbjson uninstall <package>
```

## Equivalent ADB Command
```bash
adb uninstall <package>
```

## Sample JSON Output
```json
{
  "uninstall_result": {
    "success": true,
    "message": "Uninstallation successful"
  }
}
```

## Sample YAML Output
```bash
adbjson uninstall com.example.app --format yaml
```
```yaml
uninstall_result:
  success: true
  message: Uninstallation successful
```

## Flags
- `-k, --keep-data`: Uninstall but keep data
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires package name as argument
- Flags correspond to original ADB uninstall options
