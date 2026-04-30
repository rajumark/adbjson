# shell input keyevent

Send key event to the device.

## Command
```bash
adbjson shell input keyevent <code>
```

## Description
Executes `adb shell input keyevent <code>` and outputs the result as structured JSON.

## Examples

### Send HOME key event
```bash
./adbjson shell input keyevent 3
```

**Output:**
```json
{
  "success": true,
  "keycode": "3",
  "message": ""
}
```

### Send BACK key event
```bash
./adbjson shell input keyevent 4
```

**Output:**
```json
{
  "success": true,
  "keycode": "4",
  "message": ""
}
```

### Send POWER key event
```bash
./adbjson shell input keyevent 26
```

**Output:**
```json
{
  "success": true,
  "keycode": "26",
  "message": ""
}
```

### Invalid key code
```bash
./adbjson shell input keyevent 999
```

**Output:**
```json
{
  "success": false,
  "keycode": "999",
  "message": "Invalid key code"
}
```

## Response Fields

- **success** (boolean): Whether the key event was sent successfully
- **keycode** (string): The key code that was sent
- **message** (string): Error message if any (empty on success)

## Common Key Codes

- **3**: HOME button
- **4**: BACK button  
- **26**: POWER button
- **82**: Menu button
- **24**: Volume up
- **25**: Volume down
- **164**: Mute
- **85**: Play/Pause
- **224**: Light up screen
- **223**: Turn off screen

## Notes

- Input commands typically have no output on success
- The device must be awake and unlocked for key events to work properly
- Some key events may require specific app contexts
- Invalid key codes will result in failure
- Use this for automated testing or remote control scenarios

## Related Commands

- `adbjson shell input tap <x> <y>` - Tap screen at coordinates
- `adbjson shell input swipe <x1> <y1> <x2> <y2>` - Swipe screen
