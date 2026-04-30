# shell dumpsys meminfo

Show memory information dump.

## Command
```bash
adbjson shell dumpsys meminfo
```

## Description
Executes `adb shell dumpsys meminfo` and outputs the result as structured JSON.

## Examples

### Show memory information dump
```bash
./adbjson shell dumpsys meminfo
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Applications Memory Usage (in Kilobytes)",
      "content": "Uptime: 123456789 Realtime: 987654321\n\n** MEMINFO in pid 1234 [com.example.app] **\n                   Pss  Private  Private  SwapPss     Heap     Heap     Heap\n                 Total    Dirty    Clean    Dirty     Size    Alloc     Free\n                ------   ------   ------   ------   ------   ------   ------\n  Native Heap     1024     1024        0        0     2048     1024     1024\n  Dalvik Heap     512      512        0        0     1024      512      512\n  Dalvik Other     256      256        0        0        0        0        0\n   Stack         128      128        0        0        0        0        0\n   Ashmem         64       64        0        0        0        0        0\nOther dev         32       32        0        0        0        0        0\n   .so mmap      1024      256      768        0        0        0        0\n  .jar mmap        128        0      128        0        0        0        0\n  .apk mmap        256        0      256        0        0        0        0\n  .ttf mmap         64        0       64        0        0        0        0\n  .dex mmap        512        0      512        0        0        0        0\n   .oat mmap       128        0      128        0        0        0        0\n  .art mmap         64        0       64        0        0        0        0\nOther mmap        128       32       96        0        0        0        0\n   EGL mtrack     256      256        0        0        0        0        0\n    GL mtrack      64       64        0        0        0        0        0\n      Unknown      128      128        0        0        0        0        0\n        TOTAL     4352     2752     1600        0     3072     1536     1536\n\n   Objects\n               Views:        0         ViewRootImpl:        0\n         AppContexts:        0           Activities:        0\n              Assets:        0        AssetManagers:        0\n       Global Sessions:        0\n        SQLiteCursors:        0           DATABASES:        0\n\n      SQL\n         MEMORY_USED:        0\n PAGECACHE_OVERFLOW:        0 MALLOC_SIZE:        0\n..."
    },
    {
      "name": "Total PSS by OOM adjustment",
      "content": "  403,981K: System\n      403,981K: system (pid 2843)\n  620,272K: Persistent\n      282,013K: com.android.systemui (pid 5051)\n      108,828K: com.motorola.dciservice (pid 7162)\n       55,322K: com.android.phone (pid 5415)\n...\n  613,258K: Foreground\n      442,617K: com.instagram.android (pid 28486 / activities)\n      170,641K: com.motorola.launcher3 (pid 29554 / activities)\n  739,505K: Visible\n      171,730K: com.google.android.tts (pid 12534)\n      131,096K: com.facebook.stella (pid 29747)\n...\n  3,188,544K: Cached\n      325,846K: com.google.android.youtube (pid 28808 / activities)\n      222,414K: com.google.android.apps.maps (pid 19854 / activities)\n..."
    },
    {
      "name": "Total PSS by process",
      "content": "  4,352K: com.example.app (pid 1234)\n  3,258K: com.google.android.youtube (pid 28808)\n  2,241K: com.google.android.apps.maps (pid 19854)\n..."
    },
    {
      "name": "Total PSS by category",
      "content": "  789,636K: .apk mmap\n  621,433K: Dalvik\n  583,582K: Native\n  499,279K: .dex mmap\n  244,141K: Unknown\n  238,086K: .so mmap\n  177,608K: EGL mtrack\n  150,920K: Dalvik Other\n   92,158K: .jar mmap\n   87,397K: Other mmap\n   60,804K: .art mmap\n   57,340K: Gfx dev\n   50,804K: Stack\n   16,368K: GL mtrack\n   13,252K: .oat mmap\n    9,825K: Ashmem\n    5,255K: .ttf mmap\n    3,426K: Other dev\n       0K: Cursor\n       0K: Other mtrack\n\nTotal RAM: 7,703,080K (status normal)\nFree RAM: 4,791,107K (3,201,003K cached pss + 1,342,944K cached kernel +   247,160K free)\nDMA-BUF:   306,728K (  123,920K mapped +   182,808K unmapped)\nDMA-BUF Heaps:   306,728K\nDMA-BUF Heaps pool:   239,928K\nGPU:   571,460K (  555,092K dmabuf +    16,368K private)\nUsed RAM: 5,045,181K (3,480,433K used pss + 1,564,748K kernel)\nLost RAM:   250,318K\nZRAM: 1,115,004K physical used for 3,100,900K in swap (5,777,304K total swap)\nTuning: 256 (large 512), oom   322,560K, restore limit   107,520K (high-end-gfx)\n"
    }
  ],
  "count": 7
}
```

## Response Fields

- **sections** (array): List of memory information sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "Applications Memory Usage", "Total PSS by OOM adjustment")
- **content** (string): Raw content of the section

## Common Sections

### Applications Memory Usage
- Detailed memory usage per application
- PSS (Proportional Set Size) breakdown
- Heap allocation information
- Native and Dalvik memory usage
- Memory mapping details

### Total PSS by OOM adjustment
- Memory usage grouped by OOM (Out Of Memory) adjustment categories
- System, Persistent, Foreground, Visible, Perceptible, and Cached apps
- Memory pressure and priority classifications
- Process lifecycle management

### Total PSS by process
- Memory usage sorted by individual processes
- Process IDs and memory consumption
- Application-specific memory allocation
- System process memory usage

### Total PSS by category
- Memory usage categorized by type (Dalvik, Native, APK, etc.)
- Graphics memory usage (EGL, GL)
- Memory mapping statistics
- System-wide memory distribution

## Memory Categories

- **PSS (Proportional Set Size)**: Memory that is unique to a process plus its share of shared memory
- **Native Heap**: Native code memory allocation
- **Dalvik Heap**: Java/Kotlin application memory
- **Graphics Memory**: GPU and display memory usage
- **Memory Mappings**: APK, DEX, SO, and other mapped files
- **Kernel Memory**: System kernel memory usage

## System Memory Information

- **Total RAM**: Total physical memory available
- **Free RAM**: Available memory including cached and free
- **Used RAM**: Memory currently in use
- **DMA-BUF**: Direct memory access buffers
- **GPU Memory**: Graphics processing unit memory
- **ZRAM**: Compressed swap memory
- **Lost RAM**: Memory that cannot be accounted for

## Notes

- Provides comprehensive memory usage analysis
- Essential for debugging memory leaks and performance issues
- Shows detailed memory allocation per application and category
- Contains sensitive system memory information
- Content varies based on device hardware and Android version
- Useful for analyzing memory pressure and optimization opportunities
- Can help identify memory-intensive applications and system bottlenecks

## Related Commands

- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell dumpsys power` - Power manager information
- `adbjson shell ps` - Process information
