# shell top

Show running processes with resource usage.

## Command
```bash
adbjson shell top -n <iterations>
```

## Description
Executes `adb shell top` and outputs the result as structured JSON.

## Examples

### Show processes with resource usage (1 iteration)
```bash
./adbjson shell top -n 1
```

**Output:**
```json
{
  "processes": [
    {
      "user": "shell",
      "pid": "22540",
      "pr": "20",
      "ni": "0",
      "virt": "2.0G",
      "res": "4.5M",
      "shr": "3.3M",
      "s": "R",
      "cpu": "17.8",
      "mem": "0.0",
      "time": "0:00.06",
      "args": "top -n 1"
    },
    {
      "user": "u0_a331",
      "pid": "5051",
      "pr": "20",
      "ni": "0",
      "virt": "10G",
      "res": "244M",
      "shr": "121M",
      "s": "S",
      "cpu": "3.5",
      "mem": "3.2",
      "time": "460:12.42",
      "args": "com.android.systemui"
    }
  ],
  "count": 20,
  "summary": {
    "tasks": "Tasks: 764 total,   1 running, 763 sleeping,   0 stopped,   0 zombie",
    "cpu": "800%cpu  11%user   0%nice  21%sys 761%idle   0%iow   4%irq   0%sirq   4%host",
    "mem": "Mem:  7703080K total,  7022968K used,   680112K free,   7096K buffers",
    "swap": "Swap:  5777304K total,  3009028K used,  2768276K free,  3074508K cached"
  }
}
```

## Response Fields

- **processes** (array): List of running processes with resource usage
- **count** (number): Total number of processes shown
- **summary** (object): System resource summary

### Process Fields

- **user** (string): User/UID running the process
- **pid** (string): Process ID
- **pr** (string): Priority
- **ni** (string): Nice value
- **virt** (string): Virtual memory usage
- **res** (string): Resident memory usage
- **shr** (string): Shared memory usage
- **s** (string): Process status (R=running, S=sleeping, I=idle)
- **cpu** (string): CPU usage percentage
- **mem** (string): Memory usage percentage
- **time** (string): Total CPU time used
- **args** (string): Process command/arguments

### Summary Fields

- **tasks** (string): Task summary (total, running, sleeping, stopped, zombie)
- **cpu** (string): CPU usage breakdown
- **mem** (string): Memory usage summary
- **swap** (string): Swap usage summary

## Flags

- **-n** - Number of iterations to run (default: 1)

## Notes

- Shows processes sorted by CPU usage (highest first)
- Includes both user processes and kernel threads
- Memory values are shown in kilobytes (K)
- CPU usage is shown as percentage
- Use `-n 1` for a single snapshot (recommended for JSON output)
- Process names may be truncated with "+" suffix

## Related Commands

- `adbjson shell ps -A` - List all processes without resource usage
- `adbjson shell free` - Show memory usage summary
- `adbjson shell df` - Show disk usage
