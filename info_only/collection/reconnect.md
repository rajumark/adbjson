# reconnect

Reconnect to the device.

## Command
```bash
adbjson reconnect
```

## Description
Executes `adb reconnect` and outputs the result as structured JSON.

## Examples

### Reconnect device
```bash
./adbjson reconnect
```

**Output (successful):**
```json
{
  "success": true,
  "message": "reconnecting ZD222XW5RL [device]"
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
- **message** (string): The raw message from ADB command

## Notes

- This command attempts to reconnect to the currently connected device
- Works even when the device connection is unstable
- Useful for fixing connection issues without unplugging the device
- No arguments required - reconnects to the last known device
- May take a few seconds to complete

## Related Commands

- `adbjson devices` - List connected devices
- `adbjson connect <host:port>` - Connect to device via TCP/IP
- `adbjson disconnect <host:port>` - Disconnect from device
