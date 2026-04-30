# shell settings list

List Android settings from system namespace.

## Command
```bash
adbjson shell settings list <namespace>
```

## Description
Executes `adb shell settings list <namespace>` and outputs the result as structured JSON.

## Examples

### List system settings
```bash
./adbjson shell settings list system
```

**Output:**
```json
{
  "namespace": "system",
  "settings": [
    {
      "key": "accelerometer_rotation",
      "value": "0"
    },
    {
      "key": "screen_brightness",
      "value": "52"
    },
    {
      "key": "font_scale",
      "value": "1.3"
    }
  ],
  "count": 83
}
```

### List secure settings
```bash
./adbjson shell settings list secure
```

**Output:**
```json
{
  "namespace": "secure",
  "settings": [
    {
      "key": "android_id",
      "value": "1234567890abcdef"
    },
    {
      "key": "wifi_on",
      "value": "1"
    }
  ],
  "count": 45
}
```

### List global settings
```bash
./adbjson shell settings list global
```

**Output:**
```json
{
  "namespace": "global",
  "settings": [
    {
      "key": "adb_enabled",
      "value": "1"
    },
    {
      "key": "airplane_mode_on",
      "value": "0"
    }
  ],
  "count": 67
}
```

### Invalid namespace
```bash
./adbjson shell settings list invalid
```

**Output:**
```
Error: namespace must be one of: system, secure, global
```

## Response Fields

- **namespace** (string): The settings namespace (system, secure, or global)
- **settings** (array): List of key-value setting pairs
- **count** (number): Total number of settings

## Namespaces

- **system**: User-configurable system settings (display, sound, etc.)
- **secure**: Security-sensitive settings (device identifiers, etc.)
- **global**: Global system-wide settings (airplane mode, etc.)

## Common Settings

### System Settings
- `screen_brightness` - Screen brightness level
- `font_scale` - Font scaling factor
- `screen_off_timeout` - Screen timeout in milliseconds
- `volume_music` - Music volume level
- `accelerometer_rotation` - Auto-rotate screen

### Secure Settings
- `android_id` - Unique device identifier
- `wifi_on` - WiFi enabled status
- `bluetooth_on` - Bluetooth enabled status

### Global Settings
- `adb_enabled` - ADB debugging enabled
- `airplane_mode_on` - Airplane mode status
- `hidden_api_policy_pre_p_apps` - Non-SDK API policy

## Notes

- Settings are returned as key=value pairs
- Empty string values indicate settings with no value
- Some settings may be read-only depending on permissions
- Use `adbjson shell settings get <namespace> <key>` to get specific setting
- Use `adbjson shell settings put <namespace> <key> <value>` to modify settings
