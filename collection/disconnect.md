# disconnect

Disconnect from a device via TCP/IP.

## Command
```bash
adbjson disconnect <host:port>
```

## Description
Executes `adb disconnect <host:port>` and outputs the result as structured JSON.

## Examples

### Disconnect from connected device
```bash
./adbjson disconnect 192.168.1.100:5555
```

**Output:**
```json
{
  "disconnected": true,
  "target": "192.168.1.100:5555",
  "message": "disconnected 192.168.1.100:5555"
}
```

### Disconnect from non-existent device
```bash
./adbjson disconnect 192.168.1.100:5555
```

**Output:**
```json
{
  "disconnected": false,
  "target": "192.168.1.100:5555",
  "message": "error: no such device '192.168.1.100:5555'"
}
```

### Disconnect from all devices
```bash
./adbjson disconnect
```

**Output:**
```json
{
  "disconnected": true,
  "target": "all",
  "message": "disconnected everything"
}
```

## Response Fields

- **disconnected** (boolean): Whether the disconnection was successful
- **target** (string): The host:port that was disconnected (or "all" for all devices)
- **message** (string): The raw message from ADB command

## Notes

- The command will succeed even if the device was not connected
- Use `adbjson devices` to check current device connections
- Common scenarios:
  - Device not connected: Returns error message but command succeeds
  - Device connected: Successfully disconnects the device
  - Multiple devices: Can disconnect specific device or all devices
