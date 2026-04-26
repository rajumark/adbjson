# connect

Connect to a device via TCP/IP.

## Command
```bash
adbjson connect <host:port>
```

## Description
Executes `adb connect <host:port>` and outputs the result as structured JSON.

## Examples

### Connect to device
```bash
./adbjson connect 192.168.1.100:5555
```

**Output:**
```json
{
  "connected": true,
  "target": "192.168.1.100:5555",
  "message": "connected to 192.168.1.100:5555"
}
```

### Failed connection
```bash
./adbjson connect 192.168.1.100:5555
```

**Output:**
```json
{
  "connected": false,
  "target": "192.168.1.100:5555",
  "message": "failed to connect to '192.168.1.100:5555': Connection refused"
}
```

## Response Fields

- **connected** (boolean): Whether the connection was successful
- **target** (string): The host:port that was attempted to connect to
- **message** (string): The raw message from ADB command

## Notes

- The device must be listening on the specified port for connection to succeed
- Use `adb tcpip 5555` on the device first to enable TCP/IP debugging
- Common connection failures include:
  - Connection refused: Device not listening on port
  - Network unreachable: Device not accessible on network
  - Timeout: Network issues or firewall blocking
