# shell input text

Input text string at current cursor position.

## Command
```bash
adbjson shell input text <string> [flags]
```

## Description
Executes `adb shell input text` and outputs the result as structured JSON. Simulates text input with support for multiple input sources and displays.

## Examples

### Basic text input
```bash
./adbjson shell input text "hello world"
```

**Output:**
```json
{
  "text": "hello world",
  "source": "keyboard",
  "display_id": 0,
  "success": true,
  "message": "Text input command executed successfully"
}
```

### Text input with specific source
```bash
./adbjson shell input text --source stylus "test message"
```

**Output:**
```json
{
  "text": "test message",
  "source": "stylus",
  "display_id": 0,
  "success": true,
  "message": "Text input command executed successfully"
}
```

### Text input on specific display
```bash
./adbjson shell input text --display 0 "multi display test"
```

**Output:**
```json
{
  "text": "multi display test",
  "source": "keyboard",
  "display_id": 0,
  "success": true,
  "message": "Text input command executed successfully"
}
```

### Text input with special characters
```bash
./adbjson shell input text "special chars: @#$%^&*()"
```

**Output:**
```json
{
  "text": "special chars: @#$%^&*()",
  "source": "keyboard",
  "display_id": 0,
  "success": true,
  "message": "Text input command executed successfully"
}
```

### Text input with numbers and symbols
```bash
./adbjson shell input text "Order #1234 - $99.99"
```

**Output:**
```json
{
  "text": "Order #1234 - $99.99",
  "source": "keyboard",
  "display_id": 0,
  "success": true,
  "message": "Text input command executed successfully"
}
```

## Response Fields

- **text** (string): The text that was input
- **source** (string): Input source type used for text input
- **display_id** (number): Display ID where text was input
- **success** (boolean): True if text input command executed successfully
- **message** (string): Status message or error description

## Flags

- `-d, --display`: Specify display ID to input text on (default: 0)
- `-s, --source`: Input source type (default: "keyboard")
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Input Sources

- **keyboard**: Physical or virtual keyboard input (default)
- **stylus**: Stylus pen input
- **mouse**: Mouse input (for text fields)
- **touchscreen**: Touch screen input
- **trackball**: Trackball input
- **joystick**: Joystick input
- **gamepad**: Gamepad input
- **dpad**: Directional pad input

## Use Cases

- **Automated Testing**: Fill forms and input fields for UI testing
- **Data Entry**: Automated data input for testing applications
- **Demo Automation**: Create automated demonstrations with text input
- **Accessibility**: Programmatic text input for accessibility features
- **Testing**: Test text field validation and input handling
- **Remote Control**: Control device text input remotely via scripts
- **Bulk Input**: Input large amounts of text programmatically

## Compatibility

- **Android Versions**: All Android versions supporting input command
- **Device Requirements**: Any Android device with text input capability
- **Display Support**: Multi-display devices with display ID specification
- **Input Sources**: Device-dependent support for various input types
- **Text Fields**: Works with any active text input field or keyboard

## Notes

- Text is input at the current cursor position or active text field
- If no text field is active, the command may fail silently
- Special characters and Unicode text are supported
- Multi-display support requires Android 10+ for proper display ID handling
- Some input sources may not be supported on all devices
- Text input requires an active text field or keyboard focus
- Long text input may be truncated by system limitations

## Text Input Behavior

- **Current Focus**: Text goes to currently focused text field
- **Keyboard State**: Requires virtual or physical keyboard availability
- **Character Encoding**: Supports UTF-8 and Unicode characters
- **Text Length**: Limited by target application's input field constraints
- **Special Keys**: Cannot directly input control keys (use keyevent for that)

## Common Applications

- **Web Browsers**: Fill form fields and search boxes
- **Messaging Apps**: Input text messages and chat
- **Email Clients**: Compose emails and fill recipient fields
- **Document Editors**: Input text in documents and notes
- **Search Fields**: Input search queries
- **Login Forms**: Fill usernames and passwords
- **Settings**: Input configuration values

## Related Commands

- `adbjson shell input tap` - Simulate touch events to focus text fields
- `adbjson shell input keyevent` - Simulate key press events
- `adbjson shell getevent` - Monitor device input events
- `adbjson shell screencap` - Capture device screenshots
- `adbjson shell wm size` - Get screen resolution information

## Troubleshooting

- **No Text Input**: Ensure a text field is currently focused
- **Display Errors**: Use display ID 0 for primary display on most devices
- **Source Errors**: Use keyboard source for most text input scenarios
- **Special Characters**: Ensure target application supports Unicode input
- **Command Failures**: Check if device screen is unlocked and accessible
