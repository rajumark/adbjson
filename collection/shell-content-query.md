# shell content query

Query content provider data with structured JSON output.

## Command
```bash
adbjson shell content query --uri <uri> [flags]
```

## Description
Executes `adb shell content query` and outputs the result as structured JSON. Queries Android content providers for system and application data with filtering and selection options.

## Examples

### Query system settings
```bash
./adbjson shell content query --uri content://settings/system/
```

**Output:**
```json
{
  "uri": "content://settings/system/",
  "count": 82,
  "rows": [
    {
      "id": "43",
      "name": "tap_app_quick_single",
      "value": "0",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "1",
      "name": "volume_ring",
      "value": "5",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "6151",
      "name": "screen_brightness",
      "value": "52",
      "is_preserved_in_restore": "true"
    }
  ],
  "fields": ["_id", "name", "value", "is_preserved_in_restore"]
}
```

### Query specific setting with WHERE clause
```bash
./adbjson shell content query --uri content://settings/system/ --where "name='screen_brightness'"
```

**Output:**
```json
{
  "uri": "content://settings/system/",
  "count": 1,
  "rows": [
    {
      "id": "6151",
      "name": "screen_brightness",
      "value": "52",
      "is_preserved_in_restore": "true"
    }
  ],
  "fields": ["_id", "name", "value", "is_preserved_in_restore"]
}
```

### Query with specific columns
```bash
./adbjson shell content query --uri content://settings/system/ --selection "name,value"
```

**Output:**
```json
{
  "uri": "content://settings/system/",
  "count": 82,
  "rows": [
    {
      "id": "43",
      "name": "tap_app_quick_single",
      "value": "0",
      "is_preserved_in_restore": "false"
    }
  ],
  "fields": ["_id", "name", "value", "is_preserved_in_restore"]
}
```

### Query secure settings
```bash
./adbjson shell content query --uri content://settings/secure/
```

**Output:**
```json
{
  "uri": "content://settings/secure/",
  "count": 149,
  "rows": [
    {
      "id": "1",
      "name": "android_id",
      "value": "9774d56d682e549c",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "43",
      "name": "skip_gesture",
      "value": "0",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "130",
      "name": "bluetooth_address",
      "value": "24:D5:3B:CA:92:4B",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "29",
      "name": "multi_press_timeout",
      "value": "300",
      "is_preserved_in_restore": "false"
    },
    {
      "id": "99",
      "name": "clock_seconds",
      "value": "NULL",
      "is_preserved_in_restore": "false"
    }
  ],
  "fields": ["_id", "name", "value", "is_preserved_in_restore"]
}
```

## Response Fields

- **uri** (string): The content provider URI that was queried
- **count** (number): Total number of rows returned
- **rows** (array): Array of content provider rows
- **fields** (array): Array of field names in each row

### Row Fields

- **id** (string): Row ID from the content provider
- **name** (string): Setting name or key
- **value** (string): Setting value
- **is_preserved_in_restore** (string): Whether setting is preserved during device restore

## Flags

- `--uri`: Content provider URI (required)
- `--where`: WHERE clause for filtering results
- `--bind`: BIND clause for parameter binding
- `--selection`: Selection columns to return
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Common Content Provider URIs

### System Settings
- `content://settings/system/` - System-wide settings
- `content://settings/secure/` - Secure settings (device-specific)
- `content://settings/global/` - Global settings

### Media Storage
- `content://media/external/images/` - External images
- `content://media/external/video/` - External videos
- `content://media/external/audio/` - External audio files

### Applications
- `content://applications/` - Installed applications
- `content://packages/` - Package information

## Use Cases

- **System Configuration**: Read system settings and preferences
- **Device Analysis**: Analyze device configuration and state
- **Application Data**: Query application-specific data
- **Debugging**: Debug content provider issues
- **Data Migration**: Export settings for backup/restore
- **Monitoring**: Monitor system configuration changes

## Compatibility

- **Android Versions**: All Android versions supporting content providers
- **Device Requirements**: Any Android device
- **Permissions**: May require special permissions for some content providers
- **Root Access**: Some content providers may require root access

## Notes

- Content provider access is subject to Android permissions
- Some URIs may be restricted or require special permissions
- Large result sets may be truncated for performance
- WHERE clause syntax follows standard SQL format
- BIND clause allows safe parameter substitution
- Selection can limit returned columns for better performance

## Privacy Considerations

- **Safe URIs**: System settings, configuration data
- **Sensitive URIs**: Contacts, SMS, call logs (requires explicit permission)
- **App Data**: Application-specific data may contain user information
- **Media Files**: Media content may contain personal files

## Related Commands

- `adbjson shell getprop` - Get system properties
- `adbjson shell settings list` - List Android settings
- `adbjson shell dumpsys` - Get system service information
- `adbjson shell content update` - Update content provider data
- `adbjson shell content insert` - Insert content provider data

## Error Handling

- **Permission Denied**: Check if you have access to the content provider
- **URI Not Found**: Verify the content provider URI is correct
- **Invalid WHERE Clause**: Check SQL syntax in WHERE clause
- **Empty Results**: Query may return no rows if no matching data
