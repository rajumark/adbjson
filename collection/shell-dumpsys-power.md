# shell dumpsys power

Show power manager dump information.

## Command
```bash
adbjson shell dumpsys power
```

## Description
Executes `adb shell dumpsys power` and outputs the result as structured JSON.

## Examples

### Show power manager dump
```bash
./adbjson shell dumpsys power
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Power Manager State",
      "content": "mIsPowered=true\nmPlugType=2\nmBatteryLevel=85\nmBatteryLevelLow=false\n..."
    },
    {
      "name": "Wake Locks",
      "content": "size=3\nWake Lock #1: PARTIAL_WAKE_LOCK 'AudioMix'* (uid=1000, pid=312, ws=WorkSource{1000 null})\nWake Lock #2: SCREEN_DIM_WAKE_LOCK 'WindowManager'* (uid=1000, pid=312, ws=WorkSource{1000 null})\n..."
    },
    {
      "name": "Suspend Blockers",
      "content": "size=2\nSuspend Blocker #1: PowerManagerService.Display (uid=1000, pid=312)\nSuspend Blocker #2: PowerManagerService.Broadcasts (uid=1000, pid=312)\n..."
    },
    {
      "name": "Wakefulness Session Observer",
      "content": "default timeout: 300000\noverride timeout: -1\nWakefulness Session Power Group powerGroupId: 0\ncurrent wakefulness: 3\ncurrent user activity event: 0\ncurrent user activity duration: 198096\nprevious user activity event: 2\nprevious user activity duration: 229011\nis in override timeout: false\nmIsInteractive: false\ncurrent screen policy: 0\ncurrent screen policy duration: 198620\nprevious screen policy: 3\npast screen policy duration: 91213\n"
    },
    {
      "name": "FaceDownDetector",
      "content": "mFaceDown=false\nmActive=false\nmLastFlipTime=0\nmSensorMaxLatencyMicros=2000000\nmUserInteractionBackoffMillis=60000\nmPreviousResultTime=0\nmPreviousResultType=1\nmMillisSaved=0\nmZAccelerationThreshold=-9.5\nmAccelerationThreshold=0.2\nmTimeThreshold=PT1S\nmEnabledOverride=true"
    },
    {
      "name": "AmbientDisplaySuppressionController",
      "content": "ambientDisplaySuppressed=false\nmSuppressionTokens={}\n"
    },
    {
      "name": "Low Power Standby Controller",
      "content": "mIsActive=false\nmIsEnabled=false\nmSupportedConfig=false\nmEnabledByDefaultConfig=false\nmStandbyTimeoutConfig=0\nmEnableCustomPolicy=false\nAllowed UIDs=[]\n"
    },
    {
      "name": "ScreenTimeoutOverridePolicy",
      "content": "mScreenTimeoutOverrideConfig=-1\nmLastAutoReleaseReason=-1"
    },
    {
      "name": "PowerManagerFlags",
      "content": "enable_early_screen_timeout_detector: true (def:true)\nimprove_wakelock_latency: true (def:true)\nper_display_wake_by_touch: true (def:true)\nframework_wakelock_info: false (def:false)\nmove_wsc_logging_to_notifier: true (def:true)\nwakelock_attribution_via_workchain: true (def:true)\ndisable_frozen_process_wakelocks: false (def:false)\n"
    }
  ],
  "count": 27
}
```

## Response Fields

- **sections** (array): List of power manager sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "Power Manager State", "Wake Locks")
- **content** (string): Raw content of the section

## Common Sections

### Power Manager State
- Overall power management status
- Battery level and charging state
- System wakefulness and interactive state
- Power source and plug type information

### Wake Locks
- Active wake locks in the system
- Wake lock types and owners
- Process and user information
- Work source attribution

### Suspend Blockers
- System suspend blockers
- Process and user information
- Blocker types and states
- Power management coordination

### Wakefulness Session Observer
- User activity tracking
- Screen policy information
- Wakefulness session details
- Timeout and override settings

### FaceDownDetector
- Face-down detection status
- Sensor configuration and thresholds
- Power saving statistics
- Device orientation monitoring

### AmbientDisplaySuppressionController
- Ambient display suppression state
- Suppression tokens and policies
- Always-on display management
- Power optimization settings

### Low Power Standby Controller
- Low power standby status
- Configuration and support flags
- Allowed UIDs and policies
- Standby timeout settings

### ScreenTimeoutOverridePolicy
- Screen timeout override configuration
- Auto-release reasons and status
- Timeout policy management
- User interaction handling

### PowerManagerFlags
- Power manager feature flags
- System configuration options
- Debug and logging settings
- Performance optimization flags

## Notes

- Provides comprehensive power management debugging information
- Essential for debugging battery drain and power issues
- Shows detailed wake lock and suspend blocker information
- Contains sensitive system information for debugging
- Content varies based on power management features and Android version
- Useful for analyzing power consumption and battery optimization problems
- Can help identify wake lock leaks and power management issues

## Related Commands

- `adbjson shell dumpsys battery` - Battery service information
- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell dumpsys connectivity` - Connectivity service information
