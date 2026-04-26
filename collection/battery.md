# battery Command

## Description
Get detailed battery information from the device.

## Command
```bash
./adbjson battery
```

## Sample Output
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

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
