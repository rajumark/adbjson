# shell pm path

## Description
Gets the APK path of a specific package on the connected Android device.

## Command
```bash
adbjson shell pm path <package>
```

## Equivalent ADB Command
```bash
adb shell pm path <package>
```

## Sample JSON Output
```json
{
  "package_path": {
    "package": "package",
    "path": "/data/app/com.example/base.apk"
  }
}
```

## Sample YAML Output
```bash
adbjson shell pm path android --format yaml
```
```yaml
package_path:
  package: package
  path: /data/app/com.example/base.apk
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Requires package name as argument
- Returns the APK path for the specified package
