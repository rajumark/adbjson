# shell mount

Show mount points and file system information.

## Command
```bash
adbjson shell mount
```

## Description
Executes `adb shell mount` and outputs the result as structured JSON.

## Examples

### Show all mount points
```bash
./adbjson shell mount
```

**Output:**
```json
{
  "mount_points": [
    {
      "device": "/dev/block/dm-24",
      "mount_point": "/",
      "type": "ext4",
      "options": "ro,dirsync,seclabel,nodev,noatime"
    },
    {
      "device": "tmpfs",
      "mount_point": "/dev",
      "type": "tmpfs",
      "options": "rw,seclabel,nosuid,nodev,noexec,relatime"
    },
    {
      "device": "/dev/block/dm-76",
      "mount_point": "/data",
      "type": "f2fs",
      "options": "rw,lazytime,seclabel,nosuid,nodev,noatime,background_gc=on,nogc_merge,discard,discard_unit=block,user_xattr,inline_xattr,acl,inline_data,inline_dentry,flush_merge,barrier,extent_cache,mode=adaptive,active_logs=6,reserve_root=32768,resuid=0,resgid=1065,inlinecrypt,alloc_mode=default,checkpoint_merge,fsync_mode=nobarrier,memory=normal,errors=continue,lookup_mode=perf"
    },
    {
      "device": "/dev/fuse",
      "mount_point": "/storage/emulated",
      "type": "fuse",
      "options": "rw,lazytime,nosuid,nodev,noexec,noatime,user_id=0,group_id=0,allow_other"
    }
  ],
  "count": 166
}
```

## Response Fields

- **mount_points** (array): List of all mount points
- **count** (number): Total number of mount points

### Mount Point Fields

- **device** (string): Device or filesystem being mounted
- **mount_point** (string): Directory where the filesystem is mounted
- **type** (string): Filesystem type (ext4, f2fs, tmpfs, fuse, etc.)
- **options** (string): Mount options and flags

## Common Filesystem Types

- **ext4** - Extended filesystem 4 (Linux standard)
- **f2fs** - Flash-Friendly File System (optimized for flash storage)
- **tmpfs** - Temporary filesystem (RAM-based)
- **fuse** - Filesystem in Userspace
- **functionfs** - Function filesystem (USB)
- **incremental-fs** - Incremental filesystem (app updates)

## Common Mount Options

- **ro/rw** - Read-only/Read-write
- **seclabel** - SELinux security labels
- **nosuid/nodev/noexec** - Security restrictions
- **noatime** - Don't update access time
- **lazytime** - Lazy time updates
- **user_xattr/acl** - Extended attributes and ACLs

## Notes

- Shows all mounted filesystems including system and user data
- Includes both block device mounts and virtual filesystems
- Useful for debugging storage and filesystem issues
- Can help identify available storage locations
- Shows mount permissions and security settings
- Data partition (/data) is typically the largest and most important

## Related Commands

- `adbjson shell df` - Show disk usage by mount point
- `adbjson shell du` - Show directory usage
- `adbjson shell ls /` - List root directory contents
