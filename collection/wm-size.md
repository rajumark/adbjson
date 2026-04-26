# wm-size Command

## ADB Command
```bash
adb shell wm size
```

## Description
Get the physical screen size of the device.

## Sample Output
```json
{
  "physical_size": "1080x2400"
}
```

## adbjson Command
```bash
./adbjson wm-size
```

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
