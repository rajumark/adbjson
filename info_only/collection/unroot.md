# unroot

Restore adbd non-root privileges.

## Command
```bash
adbjson unroot
```

## Description
Executes `adb unroot` and outputs the result as structured JSON.

## Examples

### Unroot successful (from rooted state)
```bash
./adbjson unroot
```

**Output:**
```json
{
  "success": true,
  "message": "restarting adbd as non root"
}
```

### Unroot successful (already not root)
```bash
./adbjson unroot
```

**Output:**
```json
{
  "success": true,
  "message": "adbd not running as root"
}
```

### Unroot failed
```bash
./adbjson unroot
```

**Output:**
```json
{
  "success": false,
  "message": "failed to restart adbd as non root"
}
```

## Response Fields

- **success** (boolean): Whether the unroot operation was successful
- **message** (string): The raw message from ADB command

## Notes

- This command restores ADB daemon to normal (non-root) operation
- If adbd is not running as root, it will still return success
- Use this after `adbjson root` to return to normal operation
- Common scenarios:
  - Device was rooted: Successfully restarts adbd as non-root
  - Device not rooted: Returns success (already in desired state)
  - Permission issues: May fail to restart daemon

## Related Commands

- `adbjson root` - Restart adbd as root
- `adbjson get-state` - Check device connection state
