# root

Restart adbd as root.

## Command
```bash
adbjson root
```

## Description
Executes `adb root` and outputs the result as structured JSON.

## Examples

### Root successful (on rooted device)
```bash
./adbjson root
```

**Output:**
```json
{
  "success": true,
  "message": "restarting adbd as root"
}
```

### Root failed (production build)
```bash
./adbjson root
```

**Output:**
```json
{
  "success": false,
  "message": "adbd cannot run as root in production builds"
}
```

### Root failed (other error)
```bash
./adbjson root
```

**Output:**
```json
{
  "success": false,
  "message": "failed to restart adbd as root"
}
```

## Response Fields

- **success** (boolean): Whether the root operation was successful
- **message** (string): The raw message from ADB command

## Notes

- This command requires root access on the device
- Most production Android builds prevent adbd from running as root
- On rooted devices, this will restart the ADB daemon with root privileges
- After successful root, you can execute commands that require root access
- Use `adbjson unroot` to restore normal (non-root) ADB daemon
- Common failure reasons:
  - Device not rooted
  - Production build restrictions
  - Insufficient permissions
