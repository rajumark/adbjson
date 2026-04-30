# shell input tap

Simulate touch tap at specified screen coordinates.

## Command
```bash
adbjson shell input tap <x> <y> [flags]
```

## Description
Executes `adb shell input tap` and outputs the result as structured JSON. Simulates touch events at specified screen coordinates with support for multiple input sources and displays.

## Examples

### Basic tap at coordinates
```bash
./adbjson shell input tap 500 1000
```

**Output:**
```json
{
  "x": "500",
  "y": "1000",
  "source": "touchscreen",
  "display_id": 0,
  "success": true,
  "message": "Tap command executed successfully"
}
```

### Tap with specific input source
```bash
./adbjson shell input tap --source stylus 300 800
```

**Output:**
```json
{
  "x": "300",
  "y": "800",
  "source": "stylus",
  "display_id": 0,
  "success": true,
  "message": "Tap command executed successfully"
}
```

### Tap on specific display
```bash
./adbjson shell input tap --display 1 250 600
```

**Output:**
```json
{
  "x": "250",
  "y": "600",
  "source": "touchscreen",
  "display_id": 1,
  "success": true,
  "message": "Tap command executed successfully"
}
```

### Tap with custom source and display
```bash
./adbjson shell input tap --source mouse --display 0 100 200
```

**Output:**
```json
{
  "x": "100",
  "y": "200",
  "source": "mouse",
  "display_id": 0,
  "success": true,
  "message": "Tap command executed successfully"
}
```

## Response Fields

- **x** (string): X coordinate where tap was performed
- **y** (string): Y coordinate where tap was performed
- **source** (string): Input source type used for the tap
- **display_id** (number): Display ID where tap was performed
- **success** (boolean): True if tap command executed successfully
- **message** (string): Status message or error description

## Flags

- `-d, --display`: Specify display ID to tap on (default: 0)
- `-s, --source`: Input source type (default: "touchscreen")
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Input Sources

- **touchscreen**: Touch screen input (default)
- **stylus**: Stylus pen input
- **mouse**: Mouse input
- **trackball**: Trackball input
- **joystick**: Joystick input
- **gamepad**: Gamepad input
- **dpad**: Directional pad input
- **keyboard**: Keyboard input
- **touchpad**: Touchpad input
- **touchnavigation**: Touch navigation input
- **rotaryencoder**: Rotary encoder input

## Use Cases

- **Automated Testing**: Simulate user interactions for UI testing
- **Accessibility**: Programmatic navigation for accessibility features
- **Demo Automation**: Create automated demonstrations
- **Game Development**: Test touch-based game controls
- **App Testing**: Simulate user gestures and taps
- **Remote Control**: Control device remotely via scripts
- **Performance Testing**: Measure response times to touch events

## Compatibility

- **Android Versions**: All Android versions supporting input command
- **Device Requirements**: Any Android device with touch input capability
- **Display Support**: Multi-display devices with display ID specification
- **Input Sources**: Device-dependent support for various input types

## Notes

- Coordinates are in pixels relative to the specified display
- Display ID 0 typically refers to the primary display
- Not all input sources are supported on all devices
- Multi-display support requires Android 10+ for proper display ID handling
- Touch coordinates may be affected by screen density and scaling
- Some devices may require accessibility permissions for programmatic input

## Coordinate System

- **Origin**: Top-left corner of the display (0,0)
- **X-axis**: Horizontal, increasing to the right
- **Y-axis**: Vertical, increasing downward
- **Range**: Depends on screen resolution (e.g., 1080x2400)
- **Multi-display**: Each display has its own coordinate system

## Common Screen Resolutions

- **HD**: 1280x720 pixels
- **Full HD**: 1920x1080 pixels
- **Quad HD**: 2560x1440 pixels
- **4K**: 3840x2160 pixels
- **Mobile**: Varies (e.g., 1080x2400, 1440x3120)

## Related Commands

- `adbjson shell input keyevent` - Simulate key press events
- `adbjson shell screencap` - Capture device screenshots
- `adbjson shell wm size` - Get screen resolution information
- `adbjson shell wm density` - Get screen density information
- `adbjson shell getevent` - Monitor device input events
