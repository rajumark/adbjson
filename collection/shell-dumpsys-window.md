# shell dumpsys window

Show window manager dump information.

## Command
```bash
adbjson shell dumpsys window
```

## Description
Executes `adb shell dumpsys window` and outputs the result as structured JSON.

## Examples

### Show window manager dump
```bash
./adbjson shell dumpsys window
```

**Output:**
```json
{
  "sections": [
    {
      "name": "WINDOW MANAGER DISPLAY CONTENTS",
      "content": "Display #0 (name=Built-in screen, type=INTERNAL)\n  mBaseDisplayInfo=DisplayInfo{...}\n  mDisplayInfo={...}\n  DisplayMetrics={density=3.0, densityDpi=480, width=1080, height=2400}\n  ...\n  Application tokens in display# 0:\n    AppWindowToken{...} token=Token{...}\n      allAppWindows=[Window{...}]\n      numWindows=1\n      isRelayoutNeeded=false\n      ...\n  Window tokens in display# 0:\n    WindowToken{...} token=Token{...}\n      windows=[Window{...}]\n      ...\n"
    },
    {
      "name": "WINDOW MANAGER POLICY STATE",
      "content": "mSystemReady=true\nmDisplayEnabled=true\nmKeyguardShowing=false\nmKeyguardOccluded=false\nmShowingDream=false\nmDreamingLockscreenShown=false\nmStatusBarVisible=true\nmNavigationBarVisible=true\n...\n"
    },
    {
      "name": "WINDOW MANAGER WINDOWS",
      "content": "Window #1 Window{74628cc u0 com.motorola.launcher3/com.android.launcher3.CustomizationPanelLauncher}\n  mDisplayId=0\n  mSession=Session{...}\n  mSurface=Surface(name=Surface(name=74628cc)/@0x...)\n  ...\nWindow #2 Window{59cb40b u0 StatusBar}\n  ...\n"
    },
    {
      "name": "WINDOW MANAGER ANIMATIONS",
      "content": "mAnimationState=ANIMATION_STATE_FINISHED\nmCurrentAnimation=null\nmPendingAnimations=[]\n...\n"
    },
    {
      "name": "WINDOW MANAGER TRACE",
      "content": "Status: Disabled\nLog level: 1\nBuffer size: 10485760 bytes\nBuffer usage: 0 bytes\nElements in the buffer: 0"
    },
    {
      "name": "WINDOW MANAGER LOGGING",
      "content": "Deprecated legacy command. Use Perfetto commands instead."
    },
    {
      "name": "WINDOW MANAGER HIGH REFRESH RATE BLACKLIST",
      "content": "High Refresh Rate Denylist\nPackages:\nCli support: false\nWhite list valid: true\nWhite Packages:\n..."
    },
    {
      "name": "INSTALLED PACKAGES HAVING APP-SPECIFIC CONFIGURATIONS",
      "content": "Current user ID : 0\n\nPackageName : com.arlosoft.macrodroid\nNightMode : null\nLocales : [en_IN]\n\nPackageName : com.Slack\nNightMode : null\nLocales : [en_US]\n..."
    },
    {
      "name": "WINDOW MANAGER CONSTANTS",
      "content": "system_gesture_exclusion_log_debounce_millis=1500\nsystem_gesture_exclusion_limit_dp=200\nsystem_gestures_excluded_by_pre_q_sticky_immersive=false\nignore_activity_orientation_request_screens=none\n"
    }
  ],
  "count": 152
}
```

## Response Fields

- **sections** (array): List of window manager sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "WINDOW MANAGER DISPLAY CONTENTS", "WINDOW MANAGER POLICY STATE")
- **content** (string): Raw content of the section

## Common Sections

### WINDOW MANAGER DISPLAY CONTENTS
- Display information and metrics
- Application and window tokens
- Surface and rendering information
- Display hierarchy and relationships

### WINDOW MANAGER POLICY STATE
- System readiness and display state
- Keyguard and lock screen status
- Status bar and navigation bar visibility
- System UI state information

### WINDOW MANAGER WINDOWS
- Detailed window information
- Window tokens and sessions
- Surface and rendering details
- Window hierarchy and relationships

### WINDOW MANAGER ANIMATIONS
- Animation state and status
- Current and pending animations
- Animation timing and transitions
- Window transition information

### WINDOW MANAGER TRACE
- Debug trace configuration
- Buffer status and usage
- Logging level and settings
- Performance monitoring data

### WINDOW MANAGER HIGH REFRESH RATE BLACKLIST
- High refresh rate management
- Package whitelist/blacklist
- Display refresh rate control
- Performance optimization settings

### INSTALLED PACKAGES HAVING APP-SPECIFIC CONFIGURATIONS
- Per-package configuration settings
- Night mode preferences
- Locale settings per package
- User-specific configurations

### WINDOW MANAGER CONSTANTS
- System gesture exclusion settings
- Display and interaction constants
- Performance and optimization parameters
- System-wide window manager settings

## Notes

- Provides comprehensive window manager debugging information
- Essential for debugging UI rendering and display issues
- Shows detailed window hierarchy and state information
- Contains sensitive system information for debugging
- Content varies based on display hardware and Android version
- Useful for analyzing window management and performance issues
- Can help identify rendering problems and animation issues

## Related Commands

- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell dumpsys activity activities` - Activities information
- `adbjson shell dumpsys connectivity` - Connectivity service information
