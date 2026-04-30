# shell uptime

Show system uptime and load information.

## Command
```bash
adbjson shell uptime
```

## Description
Executes `adb shell uptime` and outputs the result as structured JSON.

## Examples

### Show system uptime
```bash
./adbjson shell uptime
```

**Output:**
```json
{
  "current_time": "17:34:12",
  "uptime": "28 days",
  "users": "9:57",
  "load_average": "load average: 3.12"
}
```

## Response Fields

- **current_time** (string): Current system time
- **uptime** (string): System uptime (days, hours, minutes)
- **users** (string): User session time
- **load_average** (string): System load average information

## Notes

- Shows how long the system has been running
- Includes current time from the device
- Load average shows system load over 1, 5, and 15 minute intervals
- Useful for monitoring system stability and performance
- Uptime format varies (days, hours, minutes)
- Users field shows user session information

## Related Commands

- `adbjson shell top -n 1` - Show current system load and processes
- `adbjson shell free` - Show memory usage
- `adbjson shell df` - Show disk usage
