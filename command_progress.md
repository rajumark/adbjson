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
- [ ] `adb get-devpath` - Get device path

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
- [ ] `adb shell pm list packages -3` - List third-party packages
- [ ] `adb shell pm list packages -s` - List system packages
- [ ] `adb shell pm path <package>` - Get package path
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
- [ ] `adb shell wm size` - Get screen size
- [ ] `adb shell wm density` - Get screen density
- [ ] `adb shell wm density <dpi>` - Set screen density
- [ ] `adb shell wm overscan <left,top,right,bottom>` - Set overscan

## Logcat

- [ ] `adb logcat` - View logs
- [ ] `adb logcat -c` - Clear logs
- [ ] `adb logcat -d` - Dump logs
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

## Shell Commands

- [ ] `adb shell` - Enter interactive shell
- [ ] `adb shell <command>` - Execute shell command
- [ ] `adb shell ps` - List processes
- [ ] `adb shell top` - Show running processes
- [ ] `adb shell kill <pid>` - Kill process
- [ ] `adb shell killall <name>` - Kill all processes
- [ ] `adb shell df` - Disk usage
- [ ] `adb shell du` - Directory usage
- [ ] `adb shell free` - Memory usage
- [ ] `adb shell uptime` - System uptime
- [ ] `adb shell date` - System date/time
- [ ] `adb shell reboot` - Reboot device
- [ ] `adb shell reboot recovery` - Reboot to recovery
- [ ] `adb shell reboot bootloader` - Reboot to bootloader
- [ ] `adb shell shutdown` - Shutdown device

## Network

- [ ] `adb shell netstat` - Network statistics
- [ ] `adb shell ifconfig` - Network interfaces
- [ ] `adb shell ip addr show` - IP addresses
- [ ] `adb shell ping <host>` - Ping host
- [ ] `adb shell netcfg` - Network configuration
- [ ] `adb shell svc wifi disable` - Disable WiFi
- [ ] `adb shell svc wifi enable` - Enable WiFi
- [ ] `adb shell svc data disable` - Disable mobile data
- [ ] `adb shell svc data enable` - Enable mobile data

## Battery & Power

- [ ] `adb shell dumpsys battery` - Battery info
- [ ] `adb shell dumpsys batterystats` - Battery statistics
- [ ] `adb shell settings put global airplane_mode_on 1` - Enable airplane mode
