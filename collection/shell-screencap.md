# shell screencap

Capture device screenshot with various options.

## Command
```bash
adbjson shell screencap [filename] [flags]
```

## Description
Executes `adb shell screencap` and outputs the result as structured JSON. Captures the device screen and returns image data in base64 format or file information.

## Examples

### Basic screenshot capture
```bash
./adbjson shell screencap
```

**Output:**
```json
{
  "data": "iVBORw0KGgoAAAANSUhEUgAA...",
  "format": "base64",
  "size": 10368016,
  "filename": ""
}
```

### Capture in PNG format
```bash
./adbjson shell screencap --png
```

**Output:**
```json
{
  "data": "iVBORw0KGgoAAAANSUhEUgAA...",
  "format": "png",
  "size": 10368016,
  "filename": ""
}
```

### Save screenshot to file
```bash
./adbjson shell screencap --output /storage/emulated/0/Pictures/screenshot.png
```

**Output:**
```json
{
  "data": "iVBORw0KGgoAAAANSUhEUgAA...",
  "format": "base64",
  "size": 10368016,
  "filename": "/storage/emulated/0/Pictures/screenshot.png"
}
```

### Capture all displays
```bash
./adbjson shell screencap --all
```

**Output:**
```json
{
  "data": "iVBORw0KGgoAAAANSUhEUgAA...",
  "format": "base64",
  "size": 10368016,
  "filename": ""
}
```

### Capture specific display
```bash
./adbjson shell screencap --display 4630946747577212034
```

**Output:**
```json
{
  "data": "iVBORw0KGgoAAAANSUhEUgAA...",
  "format": "base64",
  "size": 10368016,
  "filename": ""
}
```

## Response Fields

- **data** (string): Base64 encoded image data or PNG data
- **format** (string): Output format ("base64", "png", or "raw")
- **size** (number): Size of captured data in bytes
- **filename** (string): Output filename if specified (optional)

## Flags

- `-a, --all`: Capture all active displays
- `-d, --display`: Specify display ID to capture (default: 4630946747577212034)
- `-p, --png`: Output in PNG format
- `-o, --output`: Save screenshot to file
- `--seamless`: Use seamless hint path
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Use Cases

- **Automated Testing**: Capture screenshots for visual regression testing
- **Documentation**: Generate screenshots for user manuals and documentation
- **Bug Reporting**: Capture device state for bug reports
- **Monitoring**: Periodic screen capture for monitoring applications
- **Remote Support**: Capture current screen for remote assistance

## Compatibility

- **Android Versions**: All Android versions supporting screencap command
- **Device Requirements**: Any Android device with display capability
- **Permissions**: May require storage permissions for file output
- **Display Support**: Supports multiple displays on compatible devices

## Notes

- Image data is returned as base64 encoded for JSON compatibility
- PNG format provides better compression and compatibility
- Large screenshots may produce substantial base64 data
- Multiple display capture appends integer postfix to filename
- Display IDs can be found using `dumpsys SurfaceFlinger --display-id`
- File saving requires appropriate storage permissions

## Related Commands

- `adbjson shell input tap` - Simulate touch events on screen
- `adbjson shell input keyevent` - Simulate key press events
- `adbjson shell wm size` - Get screen resolution information
- `adbjson shell wm density` - Get screen density information
