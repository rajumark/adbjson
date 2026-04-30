# shell dumpsys activity activities

Show activity manager activities dump information.

## Command
```bash
adbjson shell dumpsys activity activities
```

## Description
Executes `adb shell dumpsys activity activities` and outputs the result as structured JSON.

## Examples

### Show activity manager activities dump
```bash
./adbjson shell dumpsys activity activities
```

**Output:**
```json
{
  "sections": [
    {
      "name": "ACTIVITY MANAGER ACTIVITIES",
      "content": "Task stack information:\n  TaskRecord{f123456 #123 A=com.motorola.launcher3/.Launcher U=0 StackId=1 sz=2}\n    Intent { act=android.intent.action.MAIN cat=[android.intent.category.HOME] flg=0x10000000 cmp=com.motorola.launcher3/.Launcher }\n    RealActivity: com.motorola.launcher3/.Launcher\n    Activities:\n      ActivityRecord{f789abc u0 com.motorola.launcher3/.Launcher t123}\n        Process: com.motorola.launcher3\n        State: RESUMED\n        Visible: true\n        FrontOfTask: true\n        ...\n      ActivityRecord{f456def u0 com.android.settings/.Settings t123}\n        Process: com.android.settings\n        State: STOPPED\n        Visible: false\n        FrontOfTask: false\n        ...\n    ...\n  ...\n\nWindow management:\n  Display #0 (state=ON):\n    Stack #1:\n      Task #123 type=home mode=fullscreen\n        Activities: 2\n        Bounds: Rect(0, 0 - 1080, 2400)\n        Configuration: {1.3 405mcc857mnc [en_IN] ldltr sw443dp w443dp h985dp 390dpi nrml long layoutCompatNeeded port night finger -keyb/v/h -nav/h}\n        ...\n    ...\n  ...\n\nInsets information:\n  InsetsSourceProvider\n    mSource=InsetsSource id=ee350000 #845d91b type=statusBars frame=[0,0][1080,115] visible=true flags= sideHint=TOP\n    mControl=InsetsSourceControl mId=ee350000 mType=statusBars\n    mInsetsHint=Insets{left=0, top=115, right=0, bottom=0}\n    ...\n\nKeyguard state:\n  KeyguardController:\n    mKeyguardShowing=false\n    mAodShowing=false\n    mKeyguardGoingAway=false\n    mDismissalRequested=false\n    ...\n\nTask management:\n  TaskOrganizerController:\n    (fullscreen) Task{b433713 #1 type=undefined}\n    (fullscreen) Task{5835bb2 #2 type=undefined}\n    (multi-window) Task{e1a84e0 #3 type=undefined}\n    (fullscreen) Task{5da8e3d #10620 type=standard A=10199:com.google.android.dialer}\n    (fullscreen) Task{ce5a419 #10626 type=standard A=10410:com.openai.chatgpt}\n    ...\n\nCurrent task mapping:\n  mCurTaskIdForUser={0=10654}\n  mUserRootTaskInFront={}\n  isHomeRecentsComponent=true\n\nActivity visibility:\n  VisibleActivityProcess:[ProcessRecord{f70e689 29554:com.motorola.launcher3/u0a297}]\n  NumNonAppVisibleWindowUidMap:[10331:3]\n  SleepTokens={}\n"
    }
  ],
  "count": 1
}
```

## Response Fields

- **sections** (array): List of activity manager activities sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "ACTIVITY MANAGER ACTIVITIES")
- **content** (string): Raw content of the activities section

## Content Information

The activities section contains detailed information about:

### Task Stack Management
- Task records and their properties
- Activity stacks and task organization
- Intent information for activities
- Activity states and visibility

### Window Management
- Display configurations and states
- Stack organization and bounds
- Window insets and system UI
- Multi-window support information

### Activity Lifecycle
- Activity records and processes
- Activity states (RESUMED, PAUSED, STOPPED)
- Visibility and front-of-task status
- Task and stack relationships

### System UI State
- Keyguard (lock screen) status
- Always-on display state
- System bar configurations
- Insets and safe areas

### Task Organization
- Task organizer information
- Multi-window task management
- User-specific task mapping
- Recent apps integration

## Notes

- Provides detailed activity manager debugging information
- Essential for debugging activity lifecycle issues
- Shows complete task stack and window hierarchy
- Useful for analyzing app navigation and state management
- Contains sensitive system information for debugging
- Content varies based on current app states and system configuration
- Can help identify memory leaks and activity management issues

## Related Commands

- `adbjson shell dumpsys activity` - General activity manager information
- `adbjson shell dumpsys activity services` - Service information
- `adbjson shell ps -A` - List all processes
- `adbjson shell top -n 1` - Show process resource usage
