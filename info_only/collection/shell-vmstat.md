# shell vmstat

Show virtual memory statistics.

## Command
```bash
adbjson shell vmstat
```

## Description
Executes `adb shell vmstat` and outputs the result as structured JSON.

## Examples

### Show virtual memory statistics
```bash
./adbjson shell vmstat
```

**Output:**
```json
{
  "processes": {
    "running": "0",
    "blocked": "0"
  },
  "memory": {
    "swap_used": "3003196",
    "free": "378544",
    "buffers": "7200",
    "cache": "3514364"
  },
  "swap": {
    "swapped_in": "532",
    "swapped_out": "1216"
  },
  "io": {
    "blocks_in": "2311",
    "blocks_out": "1571"
  },
  "system": {
    "interrupts": "686",
    "context_switches": "1101"
  },
  "cpu": {
    "user": "11",
    "system": "10",
    "idle": "79",
    "wait": "0"
  }
}
```

## Response Fields

- **processes** (object): Process-related statistics
- **memory** (object): Memory-related statistics
- **swap** (object): Swap memory statistics
- **io** (object): I/O statistics
- **system** (object): System statistics
- **cpu** (object): CPU statistics

### Process Fields

- **running** (string): Number of processes running
- **blocked** (string): Number of processes blocked

### Memory Fields

- **swap_used** (string): Amount of swap memory used (KB)
- **free** (string): Amount of free memory (KB)
- **buffers** (string): Amount of buffer memory (KB)
- **cache** (string): Amount of cache memory (KB)

### Swap Fields

- **swapped_in** (string): Pages swapped in from disk
- **swapped_out** (string): Pages swapped out to disk

### I/O Fields

- **blocks_in** (string): Blocks received from block device
- **blocks_out** (string): Blocks sent to block device

### System Fields

- **interrupts** (string): Number of interrupts
- **context_switches** (string): Number of context switches

### CPU Fields

- **user** (string): CPU time spent in user space (%)
- **system** (string): CPU time spent in system space (%)
- **idle** (string): CPU time spent idle (%)
- **wait** (string): CPU time spent waiting for I/O (%)

## Notes

- All memory values are shown in kilobytes (KB)
- CPU percentages represent time spent in each state
- Higher blocked processes may indicate I/O bottlenecks
- High swap usage may indicate memory pressure
- Context switches show system activity level
- Useful for monitoring system performance and resource usage

## Related Commands

- `adbjson shell free` - Show memory usage summary
- `adbjson shell top -n 1` - Show processes with resource usage
- `adbjson shell ps -A` - List all processes
