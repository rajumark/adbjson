# shell dumpsys activity

Show activity manager dump information.

## Command
```bash
adbjson shell dumpsys activity
```

## Description
Executes `adb shell dumpsys activity` and outputs the result as structured JSON.

## Examples

### Show activity manager dump
```bash
./adbjson shell dumpsys activity
```

**Output:**
```json
{
  "sections": [
    {
      "name": "ACTIVITY MANAGER PROCESSES",
      "content": "ProcessRecord{f70e689 29554:com.motorola.launcher3/u0a297}\n  Package: com.motorola.launcher3\n  Process name: com.motorola.launcher3\n  State: RUNNING\n  Importance: FOREGROUND\n  LRU: 0\n  PID: 29554\n  UID: 10297\n  Last activity time: 69118\n  ...\n"
    },
    {
      "name": "ACTIVITY MANAGER ACTIVITIES",
      "content": "TaskRecord{f123456 #123 A=com.motorola.launcher3/.Launcher U=0 StackId=1 sz=2}\n  Intent { act=android.intent.action.MAIN cat=[android.intent.category.HOME] flg=0x10000000 cmp=com.motorola.launcher3/.Launcher }\n  RealActivity: com.motorola.launcher3/.Launcher\n  ...\n"
    },
    {
      "name": "ACTIVITY MANAGER SERVICES",
      "content": "Active Services:\n  ServiceRecord{f789abc u0 com.android.systemui/.tuner.TunerService}\n    Process: com.android.systemui\n    Created: 69118\n    ...\n"
    },
    {
      "name": "ACTIVITY MANAGER PROVIDERS",
      "content": "Published Content Providers:\n  ContentProviderRecord{f456789 u0 com.android.providers.settings/.SettingsProvider}\n    Process: com.android.systemui\n    Published: 69118\n    ...\n"
    },
    {
      "name": "ACTIVITY MANAGER BROADCASTS",
      "content": "Broadcast Queue:\n  mParallelBroadcasts=[\n    BroadcastRecord{f234567 u0 android.intent.action.TIME_TICK}\n      Intent { act=android.intent.action.TIME_TICK flg=0x40000010 }\n      ...\n"
    },
    {
      "name": "ACTIVITY MANAGER PERMISSIONS",
      "content": "Runtime permissions:\n  Uid 0:\n    granted: android.permission.WRITE_EXTERNAL_STORAGE\n    denied: android.permission.CAMERA\n    ...\n"
    },
    {
      "name": "ACTIVITY MANAGER CONFIGURATIONS",
      "content": "mConfiguration={1.3 405mcc857mnc [en_IN] ldltr sw443dp w443dp h985dp 390dpi nrml long layoutCompatNeeded port night finger -keyb/v/h -nav/h winConfig={ mBounds=Rect(0, 0 - 1080, 2400) mAppBounds=Rect(0, 0 - 1080, 2400) mMaxBounds=Rect(0, 0 - 1080, 2400) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=undefined mAlwaysOnTop=undefined mRotation=ROTATION_0} as.20 s.23157 fontWeightAdjustment=0 spnjio gidffffffffffffffffffff}\n  ...\n"
    },
    {
      "name": "ACTIVITY MANAGER USERS",
      "content": "mStartedUsers:\nUser #0: state=RUNNING_UNLOCKED\nUser #10: state=RUNNING_LOCKED\nmStartedUserArray: [0, 10]\nmUserLru: [10, 0]\nmUserProfileGroupIds:\nUser #0 -> profile #0\nUser #10 -> profile #0\nmCurrentProfileIds:[0, 10]\nmCurrentUserId:0\nmTargetUserId:-10000\nmLastActiveUsersForDelayedLocking:[]\nmDelayUserDataLocking:false\nmAllowUserUnlocking:true\nisStopUserOnSwitchEnabled():false\nmStopUserOnSwitch:-1\nmMaxRunningUsers:3\nmBackgroundUserScheduledStopTimeSecs:1800\nmUserSwitchUiEnabled:true\nmInitialized:true\nmIsBroadcastSentForSystemUserStarted:true\nmIsBroadcastSentForSystemUserStarting:true\nmSwitchingFromUserMessage:{}\nSwitchingToUserMessage:{}\nmLastUserUnlockingUptime: 69118\n"
    },
    {
      "name": "ACTIVITY MANAGER COMPONENT-ALIAS",
      "content": "Enabled: false\nAliases:\n\n"
    }
  ],
  "count": 16
}
```

## Response Fields

- **sections** (array): List of activity manager sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "ACTIVITY MANAGER PROCESSES")
- **content** (string): Raw content of the section

## Common Sections

### ACTIVITY MANAGER PROCESSES
- Information about running processes
- Process states, PIDs, UIDs
- Memory usage and importance levels

### ACTIVITY MANAGER ACTIVITIES  
- Information about active activities
- Task stacks and activity records
- Intent information and activity states

### ACTIVITY MANAGER SERVICES
- Information about running services
- Service records and process associations
- Service lifecycle states

### ACTIVITY MANAGER PROVIDERS
- Information about content providers
- Provider records and permissions
- Published provider details

### ACTIVITY MANAGER BROADCASTS
- Information about broadcast queues
- Pending and active broadcasts
- Intent details and receivers

### ACTIVITY MANAGER USERS
- Information about user profiles
- User states and configurations
- Multi-user system information

## Notes

- Provides comprehensive activity manager debugging information
- Useful for analyzing app behavior and system state
- Contains detailed process, activity, and service information
- Content varies based on Android version and device state
- Essential for debugging app lifecycle and performance issues
- Can help identify memory leaks and process issues

## Related Commands

- `adbjson shell dumpsys battery` - Get battery information
- `adbjson shell ps -A` - List all processes
- `adbjson shell top -n 1` - Show process resource usage
