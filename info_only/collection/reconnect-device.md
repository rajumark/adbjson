# reconnect device

Reconnect to a specific device.

## Command
```bash
adbjson reconnect device
```

## Description
Executes `adb reconnect device` and outputs the result as structured JSON.

## Examples

### Reconnect device
```bash
./adbjson reconnect device
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

- This command attempts to reconnect to the currently connected device
- Returns no output on successful reconnection (empty message)
- Works even when the device connection is unstable
- Useful for fixing connection issues without unplugging the device
- No arguments required - reconnects to the last known device
- May take a few seconds to complete

## Related Commands

- `adbjson reconnect` - Reconnect device (general)
- `adbjson reconnect offline` - Reconnect offline device
- `adbjson devices` - List connected devices
- `adbjson connect <host:port>` - Connect to device via TCP/IP
