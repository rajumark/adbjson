# shell date

Show system date and time.

## Command
```bash
adbjson shell date
```

## Description
Executes `adb shell date` and outputs the result as structured JSON.

## Examples

### Show current date and time
```bash
./adbjson shell date
```

**Output:**
```json
{
  "datetime": "Sun Apr 26 17:40:13 IST 2026"
}
```

## Response Fields

- **datetime** (string): Current system date and time

## Date Format

The output follows the standard Unix date format:
- **Day of week** (Sun, Mon, Tue, Wed, Thu, Fri, Sat)
- **Month** (Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec)
- **Day** (1-31)
- **Time** (HH:MM:SS in 24-hour format)
- **Timezone** (System timezone abbreviation)
- **Year** (4-digit year)

## Notes

- Shows the device's current system date and time
- Uses the device's local timezone settings
- Format may vary slightly based on Android version and device configuration
- Useful for timestamping and time-related operations
- Can help verify device time synchronization
- Time is displayed in the device's locale format

## Related Commands

- `adbjson shell uptime` - Show system uptime
- `adbjson shell getprop ro.build.date` - Get build date
- `adbjson shell settings get global time_12_24` - Check time format setting
