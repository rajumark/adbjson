# shell dumpsys

Get detailed system service information.

## Command
```bash
adbjson shell dumpsys [service] [options]
```

## Description
Executes `adb shell dumpsys` and outputs the result as structured JSON. Retrieves detailed information about Android system services.

## Sample Output
```json
{
  "service": "activity",
  "dump_time": "2024-01-15T10:30:45Z",
  "activities": [
    {
      "package": "com.android.settings",
      "activity": "com.android.settings.Settings",
      "state": "resumed",
      "process": "com.android.settings"
    }
  ],
  "processes": [
    {
      "pid": 1234,
      "process": "com.android.settings",
      "state": "foreground"
    }
  ],
  "task_stack_size": 3
}
```

## Services
Common system services that can be queried:
- `activity` - Activity manager service
- `package` - Package manager service  
- `window` - Window manager service
- `input` - Input manager service
- `power` - Power manager service
- `battery` - Battery service
- `connectivity` - Network connectivity service
- `wifi` - Wi-Fi service
- `telephony` - Telephony service
- `location` - Location service
- `meminfo` - Memory information
- `cpuinfo` - CPU information

## Examples
```bash
# Get activity manager info
adbjson shell dumpsys activity

# Get battery service info
adbjson shell dumpsys battery

# Get memory information
adbjson shell dumpsys meminfo

# List all available services
adbjson shell dumpsys -l
```

## Options
- `-l` - List all available services
- `-t TIMEOUT` - Timeout in seconds
- `--skip-services` - Skip services that take too long

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
- `--service SERVICE` - Specific service to dump
