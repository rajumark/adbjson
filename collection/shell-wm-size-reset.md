# shell wm size reset

Reset screen resolution to default.

## Command
```bash
adbjson shell wm size reset
```

## Description
Executes `adb shell wm size reset` and outputs the result as structured JSON.

## Examples

### Reset screen resolution
```bash
./adbjson shell wm size reset
```

**Output:**
```json
{
  "success": true,
  "message": ""
}
```

### Failed reset (permission denied)
```bash
./adbjson shell wm size reset
```

**Output:**
```json
{
  "success": false,
  "message": "Permission denied"
}
```

## Response Fields

- **success** (boolean): Whether the screen resolution reset was successful
- **message** (string): Error message if any (empty on success)

## Notes

- This command resets the screen resolution to the device's default/native resolution
- Reset commands typically have no output on success
- May require certain permissions on some devices
- Useful after testing custom resolutions
- The change takes effect immediately
- Some devices may restart the display system

## Related Commands

- `adbjson shell wm size` - Get current screen size
- `adbjson shell wm size <WxH>` - Set custom screen resolution
- `adbjson shell wm density reset` - Reset screen density
- `adbjson shell wm density` - Get current screen density
