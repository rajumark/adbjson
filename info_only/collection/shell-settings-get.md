# shell settings get

Get Android system settings values.

## Command
```bash
adbjson shell settings get [--user USER_ID] namespace key
```

## Description
Executes `adb shell settings get` and outputs the result as structured JSON. Retrieves system settings values from different namespaces.

## Sample Output
```json
{
  "namespace": "global",
  "key": "adb_enabled",
  "value": "1",
  "user_id": 0,
  "type": "integer"
}
```

## Parameters
- `namespace` - Settings namespace (system, secure, global)
- `key` - Setting key name
- `--user USER_ID` - User ID (default: current user)

## Namespaces
- `system` - System-wide settings
- `secure` - Secure settings (requires special permissions)
- `global` - Global settings across all users

## Examples
```bash
# Get ADB enabled status
adbjson shell settings get global adb_enabled

# Get airplane mode status
adbjson shell settings get global airplane_mode_on

# Get specific user setting
adbjson shell settings get --user 0 system time_12_24
```

## Common Setting Keys
- `adb_enabled` - ADB over USB enabled
- `adb_wifi_enabled` - ADB over Wi-Fi enabled
- `airplane_mode_on` - Airplane mode status
- `development_settings_enabled` - Developer options enabled
- `time_12_24` - Time format (12/24 hour)
- `auto_time` - Automatic time setting
- `auto_timezone` - Automatic timezone setting

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
