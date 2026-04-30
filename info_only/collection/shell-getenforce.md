# shell getenforce

## Description
Gets the SELinux status from the connected Android device.

## Command
```bash
adbjson shell getenforce
```

## Equivalent ADB Command
```bash
adb shell getenforce
```

## Sample JSON Output
```json
{
  "selinux_status": {
    "status": "Enforcing"
  }
}
```

## Sample YAML Output
```bash
adbjson shell getenforce --format yaml
```
```yaml
selinux_status:
  status: Enforcing
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Returns SELinux status: Enforcing, Permissive, or Disabled
