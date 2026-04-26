# wm-size Command

## Description
Get the physical screen size of the device.

## Command
```bash
adbjson wm-size
```

## Equivalent ADB Command
```bash
adb shell wm size
```

## Sample Output
```json
{
  "physical_size": "1080x2400"
}
```

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
