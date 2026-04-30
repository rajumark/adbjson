# shell cmd deviceidle get

Get device idle controller state information.

## Command
```bash
adbjson shell cmd deviceidle get [light|deep|force|screen|charging|network]
```

## Description
Executes `adb shell cmd deviceidle get` and outputs the result as structured JSON. Retrieves the current state of various device idle modes.

## Sample Output
```json
{
  "light_idle": false,
  "deep_idle": false,
  "force_idle": false,
  "screen_on": true,
  "charging": false,
  "network_connected": true
}
```

## Parameters
- `light` - Get light idle state
- `deep` - Get deep idle state  
- `force` - Get force idle state
- `screen` - Get screen state
- `charging` - Get charging state
- `network` - Get network connectivity state

## Examples
```bash
# Get all device idle states
adbjson shell cmd deviceidle get light deep force screen charging network

# Get specific state
adbjson shell cmd deviceidle get deep
```

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
