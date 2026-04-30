# tcpip

Allow device to listen on TCP/IP port.

## Command
```bash
adbjson tcpip <port>
```

## Description
Executes `adb tcpip <port>` and outputs the result as structured JSON.

## Examples

### Enable TCP/IP debugging on port 5555
```bash
./adbjson tcpip 5555
```

**Output:**
```json
{
  "success": true,
  "port": "5555",
  "message": "restarting in TCP mode port: 5555"
}
```

### Enable TCP/IP debugging on custom port
```bash
./adbjson tcpip 8888
```

**Output:**
```json
{
  "success": true,
  "port": "8888",
  "message": "restarting in TCP mode port: 8888"
}
```

### No device connected
```bash
./adbjson tcpip 5555
```

**Output:**
```json
{
  "success": false,
  "port": "5555",
  "message": "error: no device"
}
```

## Response Fields

- **success** (boolean): Whether the TCP/IP mode was enabled successfully
- **port** (string): The port that was configured
- **message** (string): The raw message from ADB command

## Notes

- This command restarts the ADB daemon on the device in TCP/IP mode
- After running this command, you can connect to the device via TCP/IP using `adbjson connect <ip>:<port>`
- The device must be connected via USB when running this command
- Common port numbers:
  - 5555: Default ADB TCP/IP port
  - 5556: Alternative port (if 5555 is in use)
- The device will restart its ADB daemon, temporarily disconnecting from USB
- Use `adbjson connect <device_ip>:<port>` to connect via TCP/IP after enabling

## Related Commands

- `adbjson connect <host:port>` - Connect to device via TCP/IP
- `adbjson disconnect <host:port>` - Disconnect from device via TCP/IP
- `adbjson devices` - List connected devices
