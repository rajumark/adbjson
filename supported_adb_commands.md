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

### get-devpath ✅
Get device path.

**Usage:**
```bash
./adbjson get-devpath
```

**Output:**
```json
{
  "dev_path": "usb:20-1"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### wm-size ✅
Get screen size.

**Usage:**
```bash
./adbjson wm-size
```

**Output:**
```json
{
  "physical_size": "1080x2400"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### wm-density ✅
Get screen density.

**Usage:**
```bash
./adbjson wm-density
```

**Output:**
```json
{
  "physical_density": "390"
}
```

**Flags:**
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output

---

### battery ✅
Get battery information.

**Usage:**
```bash
./adbjson battery
```

**Output:**
```json
{
  "ac_powered": false,
  "usb_powered": true,
  "wireless_powered": false,
  "dock_powered": false,
  "max_charging_current": 900000,
  "charge_counter": 5211121,
  "status": 2,
  "health": 2,
  "present": true,
  "level": 75,
  "scale": 100,
  "voltage": 4187,
  "temperature": 320,
  "technology": "Li-ion",
  "charging_state": 0,
  "charging_policy": 0,
  "capacity_level": -1,
  "vbus_state": true,
  "charge_watt": 2,
  "charge_watt_design": 30,
  "charge_type": 1,
  "cycle_count": 29,
  "full_capacity": 6977000,
  "full_design_capacity": 7000000
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
