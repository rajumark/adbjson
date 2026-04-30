# shell wm density reset

Reset screen density to default.

## Command
```bash
adbjson shell wm density reset
```

## Description
Executes `adb shell wm density reset` and outputs the result as structured JSON.

## Examples

### Reset screen density
```bash
./adbjson shell wm density reset
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
./adbjson shell wm density reset
```

**Output:**
```json
{
  "success": false,
  "message": "Permission denied"
}
```

## Response Fields

- **success** (boolean): Whether the screen density reset was successful
- **message** (string): Error message if any (empty on success)

## Notes

- This command resets the screen density to the device's default/native density
- Reset commands typically have no output on success
- May require certain permissions on some devices
- Useful after testing custom density values
- The change takes effect immediately
- Affects UI scaling and app display size
- Some devices may restart the display system

## Related Commands

- `adbjson shell wm density` - Get current screen density
- `adbjson shell wm density <dpi>` - Set custom screen density
- `adbjson shell wm size reset` - Reset screen resolution
- `adbjson shell wm size` - Get current screen size
