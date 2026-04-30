# shell ps

List processes running on the device.

## Command
```bash
adbjson shell ps [-A]
```

## Description
Executes `adb shell ps` and outputs the result as structured JSON.

## Examples

### List processes
```bash
./adbjson shell ps
```

**Output:**
```json
{
  "processes": [
    {
      "user": "root",
      "pid": "1",
      "ppid": "0",
      "name": "init"
    },
    {
      "user": "root",
      "pid": "2",
      "ppid": "1",
      "name": "[kthreadd]"
    }
  ],
  "count": 2
}
```

### List all processes (including system processes)
```bash
./adbjson shell ps -A
```

**Output:**
```json
{
  "processes": [
    {
      "user": "root",
      "pid": "1",
      "ppid": "0",
      "name": "init"
    },
    {
      "user": "u0_a206",
      "pid": "23509",
      "ppid": "1580",
      "name": "com.google.android.gms.persistent"
    }
  ],
  "count": 765
}
```

## Response Fields

- **processes** (array): List of running processes
- **count** (number): Total number of processes

### Process Fields

- **user** (string): User/UID running the process
- **pid** (string): Process ID
- **ppid** (string): Parent process ID
- **name** (string): Process name

## Flags

- **-A** - Show all processes (including system processes)

## Notes

- Use `-A` flag to see all processes including system processes
- Regular `ps` shows only processes for the current user
- `ps -A` shows processes for all users including system processes
- Output includes both user processes and kernel threads
- Process names may include package names for Android apps

## Related Commands

- `adbjson shell top -n 1` - Show running processes with resource usage
- `adbjson shell free` - Show memory usage
- `adbjson shell uptime` - Show system uptime
