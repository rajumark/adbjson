# Command Progress

This document tracks the implementation progress of 100 most commonly used ADB commands in adbjson.

## Device Management

- [x] `adb devices` - List connected devices ✅
- [ ] `adb devices -l` - List devices with detailed info
- [ ] `adb connect <host:port>` - Connect to device via TCP/IP
- [ ] `adb disconnect <host:port>` - Disconnect from device
- [ ] `adb kill-server` - Kill ADB server
- [ ] `adb start-server` - Start ADB server
- [x] `adb version` - Show ADB version ✅
- [x] `adb get-state` - Get device state
- [x] `adb get-serialno` - Get serial number
- [x] `adb get-devpath` - Get device path
- [ ] `adb reconnect` - Reconnect device
- [ ] `adb reconnect device` - Reconnect device
- [ ] `adb reconnect offline` - Reconnect offline device

## File Management

- [ ] `adb push <local> <remote>` - Push file to device
- [ ] `adb pull <remote> <local>` - Pull file from device
- [ ] `adb sync <directory>` - Sync directory
- [ ] `adb shell ls <path>` - List files
- [ ] `adb shell cd <path>` - Change directory
- [ ] `adb shell mkdir <path>` - Create directory
- [ ] `adb shell rm <file>` - Remove file
- [ ] `adb shell rmdir <dir>` - Remove directory
- [ ] `adb shell cp <src> <dst>` - Copy file
- [ ] `adb shell mv <src> <dst>` - Move file
- [ ] `adb shell cat <file>` - View file content
- [ ] `adb shell chmod <permissions> <file>` - Change permissions
- [ ] `adb shell chown <owner> <file>` - Change owner

## Package Management

- [ ] `adb shell pm list packages` - List all packages
- [ ] `adb shell pm list packages -f` - List packages with APK path
- [ ] `adb shell pm list packages -3` - List third-party packages
- [ ] `adb shell pm list packages -s` - List system packages
- [ ] `adb shell pm list features` - List features
- [ ] `adb shell pm list libraries` - List libraries
- [ ] `adb shell pm path <package>` - Get package path
- [ ] `adb shell pm path android` - Get android package path
- [ ] `adb shell pm list instrumentation` - List instrumentation
- [ ] `adb shell pm list permissions` - List permissions
- [ ] `adb shell pm install <apk>` - Install package
- [ ] `adb install <apk>` - Install APK
- [ ] `adb install -r <apk>` - Reinstall package
- [ ] `adb uninstall <package>` - Uninstall package
- [ ] `adb shell pm uninstall <package>` - Uninstall package
- [ ] `adb shell pm clear <package>` - Clear package data
- [ ] `adb shell pm enable <package>` - Enable package
- [ ] `adb shell pm disable <package>` - Disable package
- [ ] `adb shell pm dump <package>` - Dump package info
- [ ] `adb shell pm grant <package> <permission>` - Grant permission
- [ ] `adb shell pm revoke <package> <permission>` - Revoke permission

## App Information

- [ ] `adb shell dumpsys package <package>` - Dump package info
- [ ] `adb shell dumpsys activity` - Dump activity info
- [ ] `adb shell dumpsys battery` - Dump battery info
- [ ] `adb shell dumpsys wifi` - Dump WiFi info
- [ ] `adb shell dumpsys connectivity` - Dump connectivity info
- [ ] `adb shell dumpsys telephony` - Dump telephony info
- [ ] `adb shell dumpsys window` - Dump window info
- [ ] `adb shell dumpsys meminfo` - Dump memory info
- [ ] `adb shell dumpsys cpuinfo` - Dump CPU info
- [ ] `adb shell dumpsys dbinfo` - Dump database info

## Screen & Input

- [ ] `adb shell screencap <file>` - Capture screenshot
- [ ] `adb shell screenrecord <file>` - Record screen
- [ ] `adb shell input tap <x> <y>` - Tap screen
- [ ] `adb shell input swipe <x1> <y1> <x2> <y2>` - Swipe screen
- [ ] `adb shell input text <text>` - Type text
- [ ] `adb shell input keyevent <code>` - Send key event
- [x] `adb shell wm size` - Get screen size
- [x] `adb shell wm density` - Get screen density
- [ ] `adb shell wm density <dpi>` - Set screen density
- [ ] `adb shell wm overscan <left,top,right,bottom>` - Set overscan

## Logcat

- [ ] `adb logcat` - View logs
- [ ] `adb logcat -c` - Clear logs
- [ ] `adb logcat -d` - Dump logs
- [ ] `adb logcat -b main -d` - Dump main buffer
- [ ] `adb logcat -b system -d` - Dump system buffer
- [ ] `adb logcat -b events -d` - Dump events buffer
- [ ] `adb logcat -f <file>` - Save logs to file
- [ ] `adb logcat -v time` - Show timestamps
- [ ] `adb logcat -v brief` - Brief format
- [ ] `adb logcat -v process` - Process format
- [ ] `adb logcat -v thread` - Thread format
- [ ] `adb logcat -v raw` - Raw format
- [ ] `adb logcat -v tag` - Tag format
- [ ] `adb logcat -s <tag>` - Filter by tag
- [ ] `adb logcat *:E` - Show only errors
- [ ] `adb logcat *:W` - Show warnings and errors
- [ ] `adb logcat *:I` - Show info and above
- [ ] `adb logcat -t <count>` - Show last N lines
- [ ] `adb shell logcat -g` - Get log size
- [ ] `adb shell logcat -S` - Statistics
- [ ] `adb shell logcat -b crash -d` - Dump crash buffer
- [ ] `adb shell dmesg` - Kernel messages

## Shell Commands

- [ ] `adb shell` - Enter interactive shell
- [ ] `adb shell <command>` - Execute shell command
- [ ] `adb shell ps` - List processes
- [ ] `adb shell ps -A` - List all processes
- [ ] `adb shell top` - Show running processes
- [ ] `adb shell top -n 1` - Show running processes (1 iteration)
- [ ] `adb shell kill <pid>` - Kill process
- [ ] `adb shell killall <name>` - Kill all processes
- [ ] `adb shell df` - Disk usage
- [ ] `adb shell df -h` - Disk usage (human readable)
- [ ] `adb shell du` - Directory usage
- [ ] `adb shell du /sdcard` - Directory usage for sdcard
- [ ] `adb shell free` - Memory usage
- [ ] `adb shell vmstat` - Virtual memory statistics
- [ ] `adb shell uptime` - System uptime
- [ ] `adb shell date` - System date/time
- [ ] `adb shell reboot` - Reboot device
- [ ] `adb shell reboot recovery` - Reboot to recovery
- [ ] `adb shell reboot bootloader` - Reboot to bootloader
- [ ] `adb shell shutdown` - Shutdown device
- [ ] `adb shell getprop` - Get all system properties
- [ ] `adb shell getprop ro.build.id` - Get build ID
- [ ] `adb shell getprop ro.build.display.id` - Get display ID
- [ ] `adb shell getprop ro.build.version.release` - Get Android version
- [ ] `adb shell getprop ro.build.version.sdk` - Get SDK version
- [ ] `adb shell getprop ro.product.model` - Get product model
- [ ] `adb shell getprop ro.product.brand` - Get product brand
- [ ] `adb shell getprop ro.product.manufacturer` - Get manufacturer
- [ ] `adb shell getprop ro.hardware` - Get hardware
- [ ] `adb shell getprop ro.serialno` - Get serial number
- [ ] `adb shell getprop dhcp.wlan0.ipaddress` - Get WiFi IP
- [ ] `adb shell cat /proc/cpuinfo` - CPU information
- [ ] `adb shell cat /proc/meminfo` - Memory information
- [ ] `adb shell cat /proc/loadavg` - Load average
- [ ] `adb shell cat /proc/version` - Kernel version
- [ ] `adb shell cat /proc/partitions` - Partition information
- [ ] `adb shell cat /proc/uptime` - Uptime
- [ ] `adb shell cat /proc/filesystems` - File systems
- [ ] `adb shell cat /proc/net/dev` - Network devices
- [ ] `adb shell cat /proc/net/tcp` - TCP connections
- [ ] `adb shell cat /proc/net/udp` - UDP connections
- [ ] `adb shell cat /proc/net/wireless` - Wireless info
- [ ] `adb shell cat /proc/sys/kernel/hostname` - Hostname
- [ ] `adb shell cat /proc/mounts` - Mount points
- [ ] `adb shell cat /system/build.prop` - Build properties
- [ ] `adb shell ls /proc` - List proc directory
- [ ] `adb shell ls /` - List root directory
- [ ] `adb shell ls /sdcard` - List sdcard
- [ ] `adb shell ls /system` - List system
- [ ] `adb shell ls /data` - List data
- [ ] `adb shell stat /sdcard` - File stats
- [ ] `adb shell mount` - Mount points
- [ ] `adb shell service list` - List services
- [ ] `adb shell settings list system` - List system settings
- [ ] `adb shell settings list secure` - List secure settings
- [ ] `adb shell settings list global` - List global settings
- [ ] `adb shell getenforce` - SELinux status
- [ ] `adb shell id` - User ID
- [ ] `adb shell whoami` - Current user
- [ ] `adb shell uname -a` - System information
- [ ] `adb shell echo $PATH` - PATH environment
- [ ] `adb shell printenv` - Environment variables

## Network

- [ ] `adb shell netstat` - Network statistics
- [ ] `adb shell ifconfig` - Network interfaces
- [ ] `adb shell ip addr show` - IP addresses
- [ ] `adb shell ip route` - IP routes
- [ ] `adb shell ping <host>` - Ping host
- [ ] `adb shell ping -c 1 8.8.8.8` - Ping Google DNS
- [ ] `adb shell netcfg` - Network configuration
- [ ] `adb shell svc wifi disable` - Disable WiFi
- [ ] `adb shell svc wifi enable` - Enable WiFi
- [ ] `adb shell svc data disable` - Disable mobile data
- [ ] `adb shell svc data enable` - Enable mobile data

## Battery & Power

- [x] `adb shell dumpsys battery` - Battery info
- [ ] `adb shell dumpsys batterystats` - Battery statistics
- [ ] `adb shell settings put global airplane_mode_on 1` - Enable airplane mode
