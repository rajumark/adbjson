# shell setenforce

Set SELinux enforcing mode.

## Command
```bash
adbjson shell setenforce <0|1>
```

## Description
Executes `adb shell setenforce <0|1>` and outputs the result as structured JSON.

## Examples

### Set SELinux to permissive mode
```bash
./adbjson shell setenforce 0
```

**Output (success on rooted device):**
```json
{
  "success": true,
  "mode": "0",
  "message": "SELinux enforcing mode set successfully"
}
```

**Output (permission denied):**
```json
{
  "success": false,
  "mode": "0",
  "message": "Permission denied: requires root access"
}
```

### Set SELinux to enforcing mode
```bash
./adbjson shell setenforce 1
```

**Output (success on rooted device):**
```json
{
  "success": true,
  "mode": "1",
  "message": "SELinux enforcing mode set successfully"
}
```

**Output (permission denied):**
```json
{
  "success": false,
  "mode": "1",
  "message": "Permission denied: requires root access"
}
```

### Invalid mode argument
```bash
./adbjson shell setenforce 2
```

**Output:**
```
Error: mode must be 0 (permissive) or 1 (enforcing)
```

## Response Fields

- **success** (boolean): Whether the setenforce operation was successful
- **mode** (string): The mode that was attempted (0 or 1)
- **message** (string): Description of the operation result

## Mode Values

- **0**: Set SELinux to permissive mode (logs denials but doesn't enforce)
- **1**: Set SELinux to enforcing mode (enforces security policies)

## Notes

- This command requires root access to succeed
- Use `adbjson root` to gain root privileges first
- SELinux (Security-Enhanced Linux) provides mandatory access control
- Permissive mode is useful for debugging security policy issues
- Most production devices will return "Permission denied"
- Common use cases:
  - Development: Temporarily disable restrictions for testing
  - Debugging: Identify security policy violations
  - Troubleshooting: Bypass SELinux restrictions temporarily

## Related Commands

- `adbjson root` - Gain root privileges
- `adbjson shell getenforce` - Check current SELinux status
- `adbjson unroot` - Restore non-root privileges
