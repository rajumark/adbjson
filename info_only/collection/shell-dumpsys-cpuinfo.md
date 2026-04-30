# shell dumpsys cpuinfo

Show CPU information dump.

## Command
```bash
adbjson shell dumpsys cpuinfo
```

## Description
Executes `adb shell dumpsys cpuinfo` and outputs the result as structured JSON.

## Examples

### Show CPU information dump
```bash
./adbjson shell dumpsys cpuinfo
```

**Output:**
```json
{
  "sections": [
    {
      "name": "CPU Load",
      "content": "Load: 12.3 / 15.6 / 18.9\n  8% 28486/com.instagram.android: 5% user + 3% kernel / faults: 1234 minor 567 major\n  6% 29554/com.motorola.launcher3: 4% user + 2% kernel / faults: 890 minor 123 major\n  4% 12534/com.google.android.tts: 2% user + 2% kernel / faults: 456 minor 89 major\n  3% 29747/com.facebook.stella: 2% user + 1% kernel / faults: 234 minor 45 major\n  2% 7527/com.google.android.inputmethod.latin: 1% user + 1% kernel / faults: 123 minor 23 major\n  2% 23509/com.google.android.gms.persistent: 1% user + 1% kernel / faults: 345 minor 67 major\n  1% 24442/com.google.android.as: 1% user + 0% kernel / faults: 678 minor 89 major\n  1% 5802/com.google.android.ext.services: 1% user + 0% kernel / faults: 234 minor 34 major\n  1% 16752/com.google.android.googlequicksearchbox:interactor: 1% user + 0% kernel / faults: 567 minor 78 major\n  1% 18251/com.google.android.permissioncontroller: 1% user + 0% kernel / faults: 890 minor 123 major\n  0% 28486/system_server: 0% user + 0% kernel / faults: 1234 minor 567 major\n  0% 1/init: 0% user + 0% kernel\n  0% 2/kthreadd: 0% user + 0% kernel\n  0% 3/ksoftirqd/0: 0% user + 0% kernel\n..."
    }
  ],
  "count": 1
}
```

## Response Fields

- **sections** (array): List of CPU information sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "CPU Load")
- **content** (string): Raw content of the section

## Common Sections

### CPU Load
- System load averages (1min, 5min, 15min)
- Per-process CPU usage breakdown
- User vs kernel CPU time distribution
- Process fault statistics (minor/major page faults)
- Real-time CPU usage monitoring

## CPU Usage Information

Each process line shows:
- **CPU Percentage**: Total CPU usage percentage
- **Process ID and Name**: PID and process identifier
- **User Time**: Percentage of CPU time spent in user space
- **Kernel Time**: Percentage of CPU time spent in kernel space
- **Faults**: Memory page fault statistics (minor/major)

## Load Averages

- **1-minute load**: Average system load over the past minute
- **5-minute load**: Average system load over the past 5 minutes  
- **15-minute load**: Average system load over the past 15 minutes

## Process Categories

- **Applications**: User-installed and system apps
- **System Processes**: Core Android system processes
- **Kernel Threads**: Kernel background threads
- **Services**: Background service processes

## Notes

- Provides real-time CPU usage analysis
- Essential for debugging performance issues and CPU bottlenecks
- Shows detailed per-process CPU consumption
- Contains sensitive system performance information
- Content varies based on CPU architecture and system load
- Useful for identifying CPU-intensive applications and system optimization
- Can help diagnose performance problems and resource contention

## Related Commands

- `adbjson shell dumpsys meminfo` - Memory information
- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell top` - Real-time process monitoring
