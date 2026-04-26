# shell dumpsys battery

Get comprehensive battery information from the connected Android device.

## Command
```bash
adbjson shell dumpsys battery
```

## Description
Executes `adb shell dumpsys battery` and outputs the result as structured JSON. Provides detailed battery status, charging information, health metrics, and power source details.

## Examples

### Get battery information
```bash
./adbjson shell dumpsys battery
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

## Response Fields

### Power Source Status
- **ac_powered** (boolean): True if device is powered by AC adapter
- **usb_powered** (boolean): True if device is powered by USB connection
- **wireless_powered** (boolean): True if device is using wireless charging
- **dock_powered** (boolean): True if device is in a powered dock

### Charging Information
- **max_charging_current** (number): Maximum charging current in microamperes (μA)
- **charge_counter** (number): Total charge counter in microampere-hours (μAh)
- **charging_state** (number): Current charging state (0=not charging, 1=charging, 2=discharging)
- **charging_policy** (number): Charging policy (0=default, 1=fast, 2=slow)
- **charge_type** (number): Charge type (0=unknown, 1=AC, 2=USB, 3=wireless)
- **charge_watt** (number): Current charging power in watts
- **charge_watt_design** (number): Design charging power in watts

### Battery Status
- **status** (number): Battery status (1=unknown, 2=charging, 3=discharging, 4=not charging, 5=full)
- **health** (number): Battery health (1=unknown, 2=good, 3=overheat, 4=dead, 5=over voltage, 6=unspecified failure)
- **present** (boolean): True if battery is present in device
- **level** (number): Current battery level as percentage (0-100)
- **scale** (number): Maximum battery level (usually 100)

### Battery Properties
- **voltage** (number): Current battery voltage in millivolts (mV)
- **temperature** (number): Battery temperature in tenths of degrees Celsius (divide by 10 for °C)
- **technology** (string): Battery technology type (e.g., "Li-ion", "NiMH")
- **capacity_level** (number): Capacity level (-1=unknown, 0=critical, 1=low, 2=normal, 3=high)
- **cycle_count** (number): Number of charge/discharge cycles
- **full_capacity** (number): Current full capacity in microampere-hours (μAh)
- **full_design_capacity** (number): Design capacity in microampere-hours (μAh)

### Hardware State
- **vbus_state** (boolean): True if VBUS (USB voltage) is present

## Flags

- `--compact`: Compact JSON output (overrides --pretty)
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml) (default: "json")
- `--pretty`: Pretty print JSON output (default: true)

## Use Cases

- **Battery Monitoring**: Track battery level and charging status
- **Power Analysis**: Understand power consumption patterns
- **Device Health**: Monitor battery health and cycle count
- **Charging Optimization**: Analyze charging efficiency and power sources

## Compatibility

- **Android Versions**: All Android versions supporting dumpsys battery
- **Device Requirements**: Any Android device with battery
- **Permissions**: No special permissions required

## Notes

- Temperature is reported in tenths of degrees Celsius (320 = 32.0°C)
- Charge counter accumulates over battery lifetime
- Cycle count may not be available on all devices
- Power source detection depends on device hardware capabilities

## Related Commands

- `adbjson shell dumpsys power` - Power management information
- `adbjson shell dumpsys batterystats` - Detailed battery statistics
- `adbjson shell acpi --batteries` - ACPI battery information
