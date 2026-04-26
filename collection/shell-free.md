# shell free

Show memory usage information.

## Command
```bash
adbjson shell free
```

## Description
Executes `adb shell free` and outputs the result as structured JSON.

## Examples

### Show memory usage
```bash
./adbjson shell free
```

**Output:**
```json
{
  "memory": {
    "total": "7887953920",
    "used": "7358853120",
    "free": "529100800",
    "shared": "415993856",
    "buffers": "7286784",
    "cached": "0"
  },
  "buffers": {
    "total": "",
    "used": "7351566336",
    "free": "536387584",
    "shared": "",
    "buffers": "",
    "cached": ""
  },
  "swap": {
    "total": "5915959296",
    "used": "3120803840",
    "free": "2795155456",
    "shared": "",
    "buffers": "",
    "cached": ""
  }
}
```

## Response Fields

- **memory** (object): Physical memory information
- **buffers** (object): Memory usage excluding buffers/cache
- **swap** (object): Swap memory information

### Memory Fields

- **total** (string): Total memory in bytes
- **used** (string): Used memory in bytes
- **free** (string): Free memory in bytes
- **shared** (string): Shared memory in bytes
- **buffers** (string): Buffer memory in bytes
- **cached** (string): Cached memory in bytes

### Buffers Fields

- **used** (string): Used memory excluding buffers/cache
- **free** (string): Free memory excluding buffers/cache

### Swap Fields

- **total** (string): Total swap space in bytes
- **used** (string): Used swap space in bytes
- **free** (string): Free swap space in bytes

## Notes

- Memory values are shown in bytes
- Buffers section shows actual used/free memory excluding buffers and cache
- Swap shows virtual memory usage
- Useful for monitoring memory pressure and available resources
- Can help identify memory leaks or high memory usage

## Related Commands

- `adbjson shell top -n 1` - Show processes with memory usage
- `adbjson shell df` - Show disk usage
- `adbjson shell uptime` - Show system uptime
