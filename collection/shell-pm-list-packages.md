# shell pm list packages

## Description
Lists all installed packages on the connected Android device.

## Command
```bash
adbjson shell pm list packages
```

## Equivalent ADB Command
```bash
adb shell pm list packages
```

## Sample JSON Output
```json
{
  "packages": [
    {
      "name": "com.android.systemui"
    },
    {
      "name": "com.android.settings"
    }
  ],
  "count": 2
}
```

## Sample YAML Output
```bash
adbjson shell pm list packages --format yaml
```
```yaml
packages:
- name: com.android.systemui
- name: com.android.settings
count: 2
```

## Flags
- `-f, --show-apk-path`: Show APK path
- `-3, --third-party`: List third-party packages only
- `-s, --system`: List system packages only
- `-d, --disabled`: List disabled packages only
- `-e, --enabled`: List enabled packages only
- `-i, --installer`: Show package installer
- `-u, --uninstalled`: Include uninstalled packages
- `--filter string`: Filter by package name
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all packages including system and third-party
- Output is limited to first 5 packages in documentation (full list in actual output)
- Flags correspond to original ADB pm list packages options
