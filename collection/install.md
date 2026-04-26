# install

## Description
Install an APK file on the connected Android device.

## Command
```bash
adbjson install <apk>
```

## Equivalent ADB Command
```bash
adb install <apk>
```

## Sample JSON Output
```json
{
  "install_result": {
    "success": true,
    "message": "Installation successful"
  }
}
```

## Sample YAML Output
```bash
adbjson install app.apk --format yaml
```
```yaml
install_result:
  success: true
  message: Installation successful
```

## Flags
- `-r, --reinstall`: Reinstall package
- `-l, --protect`: Protect installation directory
- `-t, --test`: Install test-only apps
- `-s, --sdcard`: Install to sdcard
- `-d, --downgrade`: Allow downgrade
- `-g, --grant`: Grant all runtime permissions
- `--abi string`: Force specific ABI
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires APK file path as argument
- Flags correspond to original ADB install options
