# shell dumpsys input

Show input system dump information.

## Command
```bash
adbjson shell dumpsys input
```

## Description
Executes `adb shell dumpsys input` and outputs the result as structured JSON.

## Examples

### Show input system dump
```bash
./adbjson shell dumpsys input
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Input Manager State",
      "content": "mInteractive=true\nmSystemReady=true\nmInputDevices={...}\nmInputFilterEnabled=true\n..."
    },
    {
      "name": "Input Reader State",
      "content": "mConfiguration={...}\nmDeviceIds={...}\n..."
    },
    {
      "name": "Input Dispatcher State",
      "content": "mFocusedApplication={...}\n..."
    },
    {
      "name": "KeyCombination rules",
      "content": "KEYCODE_VOLUME_DOWN + KEYCODE_POWER\nKEYCODE_VOLUME_DOWN + KEYCODE_VOLUME_UP\nKEYCODE_VOLUME_UP + KEYCODE_POWER"
    },
    {
      "name": "AppLaunchShortcutManager",
      "content": "InputGestureData { trigger = KeyTrigger{mKeycode=KEYCODE_T, mModifierState=65536}, action = Action[keyGestureType=51, appLaunchData=ComponentData{mPackageName='com.google.android.talk', mClassName='com.google.android.talk.SigningInActivity'}] }\nInputGestureData { trigger = KeyTrigger{mKeycode=KEYCODE_E, mModifierState=65536}, action = Action[keyGestureType=51, appLaunchData=CategoryData{mCategory='android.intent.category.APP_EMAIL'}] }\n..."
    },
    {
      "name": "InputGestureManager",
      "content": ""
    },
    {
      "name": "System Shortcuts",
      "content": "InputGestureData { trigger = KeyTrigger{mKeycode=KEYCODE_A, mModifierState=65536}, action = Action[keyGestureType=5, appLaunchData=null] }\nInputGestureData { trigger = KeyTrigger{mKeycode=KEYCODE_TAB, mModifierState=65536}, action = Action[keyGestureType=2, appLaunchData=null] }\n..."
    },
    {
      "name": "Blocklisted Triggers",
      "content": "KeyTrigger{mKeycode=KEYCODE_E, mModifierState=65536}\nKeyTrigger{mKeycode=KEYCODE_C, mModifierState=4096}\nKeyTrigger{mKeycode=KEYCODE_B, mModifierState=65536}\n..."
    },
    {
      "name": "Custom Gestures",
      "content": "MOTO_INPUT_MONITOR_PID_MAP: {}\nMOTO_INPUT_MONITOR_UID_MAP: {}\n"
    }
  ],
  "count": 113
}
```

## Response Fields

- **sections** (array): List of input system sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "Input Manager State", "AppLaunchShortcutManager")
- **content** (string): Raw content of the section

## Common Sections

### Input Manager State
- Overall input manager status
- System readiness and interactivity
- Input device information
- Filter and configuration settings

### Input Reader State
- Input reader configuration
- Device ID mappings
- Input device states
- Reader thread information

### Input Dispatcher State
- Input dispatcher status
- Focused application information
- Event dispatching state
- Touch and key event handling

### KeyCombination Rules
- System key combinations
- Hardware button mappings
- Power and volume key combinations
- System-level shortcut definitions

### AppLaunchShortcutManager
- App launch gesture mappings
- Key triggers and actions
- Package and component information
- Category-based app shortcuts

### InputGestureManager
- Gesture management system
- System and custom gestures
- Gesture recognition settings
- Input gesture configuration

### System Shortcuts
- System-level key shortcuts
- Navigation and action shortcuts
- Modifier key combinations
- System gesture mappings

### Blocklisted Triggers
- Blocked key combinations
- Disabled input triggers
- Conflict prevention settings
- Input restriction policies

### Custom Gestures
- User-defined gestures
- Motorola-specific input monitoring
- Custom gesture configurations
- Device-specific input features

## Notes

- Provides comprehensive input system debugging information
- Essential for debugging input device and gesture issues
- Shows detailed key combination and shortcut mappings
- Contains sensitive system information for debugging
- Content varies based on input hardware and Android version
- Useful for analyzing input performance and gesture recognition problems
- Can help identify input device conflicts and gesture mapping issues

## Related Commands

- `adbjson shell dumpsys window` - Window manager information
- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell getevent` - Raw input device events
