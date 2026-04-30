# reconnect offline

Reconnect to an offline device.

## Command
```bash
adbjson reconnect offline
```

## Description
Executes `adb reconnect offline` and outputs the result as structured JSON.

## Examples

### Reconnect offline device
```bash
./adbjson reconnect offline
```

**Output (successful):**
```json
{
  "success": true,
  "message": ""
}
```

**Output (no device):**
```json
{
  "success": false,
  "message": "no devices/emulators found"
}
```

## Response Fields

- **success** (boolean): Whether the reconnection was successful
- **message** (string): The raw message from ADB command (empty on success)

## Notes

- This command attempts to reconnect to an offline device
- Returns no output on successful reconnection (empty message)
- Useful for recovering devices that have gone offline
- Works even when the device connection is unstable
- No arguments required - reconnects to the last known device
- May take a few seconds to complete

## Related Commands

- `adbjson reconnect` - Reconnect device (general)
- `adbjson reconnect device` - Reconnect specific device
- `adbjson devices` - List connected devices
- `adbjson connect <host:port>` - Connect to device via TCP/IP
