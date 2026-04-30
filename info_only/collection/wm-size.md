# shell wm size

Get the physical screen size of the device.

## Command
```bash
adbjson shell wm size
```

## Description
Executes `adb shell wm size` and outputs the result as structured JSON.

## Sample Output
```json
{
  "physical_size": "1080x2400"
}
```

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
