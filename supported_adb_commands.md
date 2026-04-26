# Supported ADB Commands

This document lists all ADB commands currently supported by adbjson.

## Available Commands

### devices ✅
List all connected ADB devices.

**Usage:**
```bash
./adbjson devices
```

**Output:**
```json
{
  "devices": [
    {
      "id": "emulator-5554",
      "status": "device"
    }
  ]
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### adb-version ✅
Get ADB version information.

**Usage:**
```bash
./adbjson adb-version
```

**Output:**
```json
{
  "version": "1.0.41",
  "revision": "0"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### get-state ✅
Get device state.

**Usage:**
```bash
./adbjson get-state
```

**Output:**
```json
{
  "state": "device"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### get-serialno ✅
Get device serial number.

**Usage:**
```bash
./adbjson get-serialno
```

**Output:**
```json
{
  "serial_no": "ZD222XW5RL"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

## Planned Commands

The following commands are planned for future releases:

- `packages` - List installed packages on device
- `battery` - Get battery status information
- `logcat` - Stream logcat output
- `shell` - Execute arbitrary shell commands
- `install` - Install APK files
- `uninstall` - Uninstall packages
- `screenshot` - Capture device screenshot
- `screenrecord` - Record device screen

---

## Requesting New Commands

To request a new ADB command to be added to adbjson, please open an issue on GitHub with:
- The ADB command you want to add
- Expected JSON output format
- Use case description
