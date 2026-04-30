# shell ls-root

List root directory contents.

## Command
```bash
adbjson shell ls-root
```

## Description
Executes `adb shell ls /` and outputs the result as structured JSON.

## Examples

### List root directory contents
```bash
./adbjson shell ls-root
```

**Output:**
```json
{
  "items": [
    "adb_keys",
    "apex",
    "bin",
    "bootstrap-apex",
    "bugreports",
    "cache",
    "config",
    "d",
    "data",
    "data_mirror",
    "debug_ramdisk",
    "dev",
    "etc",
    "init",
    "init.environ.rc",
    "linkerconfig",
    "metadata",
    "mnt",
    "odm",
    "odm_dlkm",
    "oem",
    "postinstall",
    "proc",
    "product",
    "sdcard",
    "second_stage_resources",
    "storage",
    "sys",
    "system",
    "system_dlkm",
    "system_ext",
    "tmp",
    "vendor",
    "vendor_dlkm"
  ],
  "count": 34
}
```

## Response Fields

- **items** (array): List of all items in root directory
- **count** (number): Total number of items

## Directory Structure

The Android root directory contains several key directories:

### System Directories
- **system/** - Main Android system files and applications
- **system_ext/** - System extensions
- **system_dlkm/** - System device kernel modules
- **vendor/** - Vendor-specific files and drivers
- **vendor_dlkm/** - Vendor device kernel modules
- **product/** - Product-specific files
- **odm/** - Original design manufacturer files
- **odm_dlkm/** - ODM device kernel modules
- **oem/** - Original equipment manufacturer files

### User Data Directories
- **data/** - User applications and data
- **data_mirror/** - Data mirror for incremental updates
- **sdcard/** - External storage mount point
- **storage/** - Storage mount points

### System Resources
- **proc/** - Process filesystem
- **sys/** - System filesystem
- **dev/** - Device files
- **mnt/** - Mount points
- **apex/** - Android Package Extensions
- **bootstrap-apex/** - Bootstrap APEX packages

### Configuration and Runtime
- **etc/** - System configuration files
- **init** - Init process
- **init.environ.rc** - Environment configuration
- **metadata/** - System metadata
- **config/** - Configuration files
- **linkerconfig/** - Dynamic linker configuration

### Temporary and Debug
- **tmp/** - Temporary files
- **cache/** - Cache directory
- **bugreports/** - Bug report files
- **debug_ramdisk/** - Debug ramdisk
- **adb_keys/** - ADB key files
- **bin/** - Binary executables
- **postinstall/** - Post-installation scripts

## Notes

- Shows the complete Android filesystem structure
- Essential for understanding device layout and storage
- Useful for file system exploration and debugging
- Directory structure follows Android filesystem standards
- Total count may vary based on Android version and device configuration

## Related Commands

- `adbjson shell ls /proc` - List process filesystem
- `adbjson shell ls /system` - List system directory
- `adbjson shell ls /data` - List user data directory
- `adbjson shell mount` - Show mount points
