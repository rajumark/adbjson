# Command Progress

This document tracks the implementation progress of 100 most commonly used ADB commands in adbjson.

## Device Management

- [x] `adb devices` - List connected devices ✅
- [x] `adb devices -l` - List devices with detailed info
- [ ] `adb connect <host:port>` - Connect to device via TCP/IP
- [ ] `adb disconnect <host:port>` - Disconnect from device
- [x] `adb kill-server` - Kill ADB server
- [x] `adb start-server` - Start ADB server
- [x] `adb version` - Show ADB version ✅
- [x] `adb get-state` - Get device state
- [x] `adb get-serialno` - Get serial number
- [x] `adb get-devpath` - Get device path
- [ ] `adb reconnect` - Reconnect device
- [ ] `adb reconnect device` - Reconnect device
- [ ] `adb reconnect offline` - Reconnect offline device
- [ ] `adb root` - Run adbd as root
- [ ] `adb unroot` - Restore adbd non-root privileges
- [ ] `adb -P <port> start-server` - Designated adb server network port
- [ ] `adb tcpip 5555` - Allow device to listen on TCP/IP port
- [ ] `adb pair ipaddr:port` - Pair device with pairing code (Android 11+)

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

- [x] `adb shell pm list packages` - List all packages
- [ ] `adb shell pm list packages -f` - List packages with APK path
- [ ] `adb shell pm list packages -3` - List third-party packages
- [ ] `adb shell pm list packages -s` - List system packages
- [ ] `adb shell pm list packages -d` - List disabled packages
- [ ] `adb shell pm list packages -e` - List enabled packages
- [ ] `adb shell pm list packages -i` - Show package installer
- [ ] `adb shell pm list packages -u` - Include uninstalled packages
- [ ] `adb shell pm list packages <FILTER>` - Filter by package name
- [x] `adb shell pm list features` - List features
- [ ] `adb shell pm list libraries` - List libraries
- [ ] `adb shell pm path <package>` - Get package path
- [ ] `adb shell pm path android` - Get android package path
- [x] `adb shell pm list instrumentation` - List instrumentation
- [x] `adb shell pm list permissions` - List permissions
- [ ] `adb shell pm install <apk>` - Install package
- [ ] `adb install <apk>` - Install APK
- [ ] `adb install -r <apk>` - Reinstall package
- [ ] `adb install -l <apk>` - Protect installation directory
- [ ] `adb install -t <apk>` - Install test-only apps
- [ ] `adb install -s <apk>` - Install to sdcard
- [ ] `adb install -d <apk>` - Allow downgrade
- [ ] `adb install -g <apk>` - Grant all runtime permissions
- [ ] `adb install --abi <abi> <apk>` - Force specific ABI
- [ ] `adb uninstall <package>` - Uninstall package
- [ ] `adb uninstall -k <package>` - Uninstall but keep data
- [ ] `adb shell pm uninstall <package>` - Uninstall package
- [ ] `adb shell pm clear <package>` - Clear package data
- [ ] `adb shell pm enable <package>` - Enable package
- [ ] `adb shell pm disable <package>` - Disable package
- [ ] `adb shell pm disable-user <package>` - Disable package (user)
- [ ] `adb shell pm dump <package>` - Dump package info
- [ ] `adb shell pm grant <package> <permission>` - Grant permission
- [ ] `adb shell pm revoke <package> <permission>` - Revoke permission

## App Information

- [ ] `adb shell dumpsys package <package>` - Dump package info
- [ ] `adb shell dumpsys activity` - Dump activity info
- [ ] `adb shell dumpsys activity activities` - Dump activity activities
- [ ] `adb shell dumpsys activity activities | grep mResumedActivity` - View Reception Activity
- [ ] `adb shell dumpsys activity services [<packagename>]` - View Running Services
- [ ] `adb shell dumpsys battery` - Dump battery info
- [ ] `adb shell dumpsys wifi` - Dump WiFi info
- [ ] `adb shell dumpsys connectivity` - Dump connectivity info
- [ ] `adb shell dumpsys telephony` - Dump telephony info
- [ ] `adb shell dumpsys window` - Dump window info
- [ ] `adb shell dumpsys window displays` - Dump window displays
- [ ] `adb shell dumpsys meminfo` - Dump memory info
- [ ] `adb shell dumpsys cpuinfo` - Dump CPU info
- [ ] `adb shell dumpsys dbinfo` - Dump database info
- [ ] `adb shell dumpsys iphonesubinfo` - Dump IMEI (Android 4.4 and below)

## Screen & Input

- [ ] `adb shell screencap <file>` - Capture screenshot
- [ ] `adb shell screencap -p <file>` - Capture screenshot (PNG format)
- [ ] `adb exec-out screencap -p > sc.png` - Capture screenshot to computer
- [ ] `adb shell screenrecord <file>` - Record screen
- [ ] `adb shell screenrecord --size WIDTHxHEIGHT <file>` - Record screen with dimensions
- [ ] `adb shell screenrecord --bit-rate RATE <file>` - Record screen with bit-rate
- [ ] `adb shell screenrecord --time-limit TIME <file>` - Record screen with time limit
- [ ] `adb shell input tap <x> <y>` - Tap screen
- [ ] `adb shell input swipe <x1> <y1> <x2> <y2>` - Swipe screen
- [ ] `adb shell input swipe <x1> <y1> <x2> <y2> [duration]` - Swipe with duration
- [ ] `adb shell input text <text>` - Type text
- [ ] `adb shell input keyevent <code>` - Send key event
- [ ] `adb shell input keyevent 3` - HOME button
- [ ] `adb shell input keyevent 4` - Return key
- [ ] `adb shell input keyevent 26` - Power button
- [ ] `adb shell input keyevent 82` - Menu button
- [ ] `adb shell input keyevent 24` - Volume up
- [ ] `adb shell input keyevent 25` - Volume down
- [ ] `adb shell input keyevent 164` - Mute
- [ ] `adb shell input keyevent 85` - Play/Pause
- [ ] `adb shell input keyevent 224` - Light up screen
- [ ] `adb shell input keyevent 223` - Turn off screen
- [x] `adb shell wm size` - Get screen size
- [ ] `adb shell wm size <WxH>` - Set screen resolution
- [ ] `adb shell wm size reset` - Reset screen resolution
- [x] `adb shell wm density` - Get screen density
- [ ] `adb shell wm density <dpi>` - Set screen density
- [ ] `adb shell wm density reset` - Reset screen density
- [ ] `adb shell wm overscan <left,top,right,bottom>` - Set overscan
- [ ] `adb shell wm overscan reset` - Reset overscan

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
- [ ] `adb logcat -v threadtime` - Threadtime format
- [ ] `adb logcat -v raw` - Raw format
- [ ] `adb logcat -v tag` - Tag format
- [ ] `adb logcat -v long` - Long format
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
- [ ] `adb shell top -m <num>` - Show max processes
- [ ] `adb shell top -d <num>` - Set refresh interval
- [ ] `adb shell top -s <col>` - Sort by column
- [ ] `adb shell top -t` - Display thread information
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
- [ ] `adb shell date -s <datetime>` - Set system date/time (requires root)
- [ ] `adb shell reboot` - Reboot device
- [ ] `adb shell reboot recovery` - Reboot to recovery
- [ ] `adb shell reboot bootloader` - Reboot to bootloader
- [ ] `adb shell shutdown` - Shutdown device
- [x] `adb shell getprop` - Get all system properties
- [ ] `adb shell getprop ro.build.id` - Get build ID
- [ ] `adb shell getprop ro.build.display.id` - Get display ID
- [ ] `adb shell getprop ro.build.version.release` - Get Android version
- [ ] `adb shell getprop ro.build.version.sdk` - Get SDK version
- [ ] `adb shell getprop ro.build.version.security_patch` - Get security patch level
- [ ] `adb shell getprop ro.product.model` - Get product model
- [ ] `adb shell getprop ro.product.brand` - Get product brand
- [ ] `adb shell getprop ro.product.name` - Get device name
- [ ] `adb shell getprop ro.product.board` - Get processor model
- [ ] `adb shell getprop ro.product.cpu.abilist` - Get CPU supported ABI list
- [ ] `adb shell getprop ro.product.manufacturer` - Get manufacturer
- [ ] `adb shell getprop ro.hardware` - Get hardware
- [ ] `adb shell getprop ro.serialno` - Get serial number
- [ ] `adb shell getprop dhcp.wlan0.ipaddress` - Get WiFi IP
- [ ] `adb shell getprop ro.sf.lcd_density` - Get screen density
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
- [ ] `adb shell cat /sys/class/net/wlan0/address` - Get MAC address
- [ ] `adb shell ls /proc` - List proc directory
- [ ] `adb shell ls /` - List root directory
- [ ] `adb shell ls /sdcard` - List sdcard
- [ ] `adb shell ls /system` - List system
- [ ] `adb shell ls /data` - List data
- [ ] `adb shell stat /sdcard` - File stats
- [ ] `adb shell mount` - Mount points
- [ ] `adb shell mount -o remount,rw` - Remount as writable (requires root)
- [ ] `adb shell service list` - List services
- [ ] `adb shell grep` - Filter output
- [ ] `adb shell settings list system` - List system settings
- [ ] `adb shell settings list secure` - List secure settings
- [ ] `adb shell settings list global` - List global settings
- [ ] `adb shell settings get secure android_id` - Get android_id
- [ ] `adb shell settings put global adb_enabled 0` - Turn off Android Debug
- [ ] `adb shell settings put global hidden_api_policy_pre_p_apps 1` - Allow non SDK API
- [ ] `adb shell settings put global hidden_api_policy_p_apps 1` - Allow non SDK API
- [ ] `adb shell settings delete global hidden_api_policy_pre_p_apps` - Forbid non SDK API
- [ ] `adb shell settings delete global hidden_api_policy_p_apps` - Forbid non SDK API
- [ ] `adb shell settings put global policy_control <key-values>` - Show/hide status bar or navigation bar
- [x] `adb shell getenforce` - SELinux status
- [ ] `adb shell setenforce 1` - Enable SELinux (requires root)
- [ ] `adb shell setenforce 0` - Disable SELinux (requires root)
- [x] `adb shell id` - User ID
- [x] `adb shell whoami` - Current user
- [x] `adb shell uname -a` - System information
- [ ] `adb shell echo $PATH` - PATH environment
- [ ] `adb shell printenv` - Environment variables
- [ ] `adb shell cat /proc/<pid>/status` - Get process status
- [ ] `adb shell cat /proc/<pid>/status | grep Uid` - Query process UID

## Network

- [ ] `adb shell netstat` - Network statistics
- [ ] `adb shell ifconfig` - Network interfaces
- [ ] `adb shell ifconfig | grep Mask` - Get IP address
- [ ] `adb shell ifconfig wlan0` - Get WiFi IP
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
- [ ] `adb shell settings put global airplane_mode_on 0` - Disable airplane mode

## Application Management (am)

- [ ] `adb shell am start [options] <INTENT>` - Start an Activity
- [ ] `adb shell am start -n <package>/<activity>` - Start specific activity
- [ ] `adb shell am startservice [options] <INTENT>` - Start a Service
- [ ] `adb shell am stopservice [options] <INTENT>` - Stop a Service
- [ ] `adb shell am broadcast [options] <INTENT>` - Send a broadcast
- [ ] `adb shell am force-stop <packagename>` - Force stop an application
- [ ] `adb shell am send-trim-memory <pid> <level>` - Trim memory
- [ ] `adb shell monkey -p <packagename> -c android.intent.category.LAUNCHER 1` - Launch app
- [ ] `adb shell monkey -p <packagename> -v 500` - Monkey stress testing

## Utility Functions

- [ ] `adb shell screencap -p /sdcard/sc.png` - Screenshot to device
- [ ] `adb pull /sdcard/sc.png` - Pull screenshot to computer
- [ ] `adb shell screenrecord /sdcard/filename.mp4` - Record screen
- [ ] `adb pull /sdcard/filename.mp4` - Pull recording to computer
- [ ] `adb shell mount -o remount,rw -t yaffs2 <device> <mountpoint>` - Remount as writable (requires root)
- [ ] `adb shell cat /data/misc/wifi/*.conf` - Check WiFi password (requires root)
- [ ] `adb shell cat /data/misc/wifi/WifiConfigStore.xml` - Check WiFi password (Android O+, requires root)
- [ ] `adb shell date -s <datetime>` - Set system date/time (requires root)
- [ ] `adb reboot` - Reboot device
- [ ] `adb shell su` - Check if device is rooted
- [ ] `adb shell monkey -p <packagename> -v 500` - Monkey stress testing
- [ ] `adb root` - Run as root (for WiFi commands)
- [ ] `adb shell svc wifi enable` - Enable WiFi (requires root)
- [ ] `adb shell svc wifi disable` - Disable WiFi (requires root)

## Flashing-Phone Related Commands

- [ ] `adb reboot recovery` - Restart to Recovery mode
- [ ] `adb reboot bootloader` - Restart to Fastboot mode
- [ ] `adb sideload <path-to-update.zip>` - Sideload system update

## Security-Related Commands

- [ ] `adb root` - Run as root
- [ ] `adb shell setenforce 1` - Enable SELinux (requires root)
- [ ] `adb shell setenforce 0` - Disable SELinux (requires root)
- [ ] `adb enable-verity` - Enable dm_verity (requires root)
- [ ] `adb disable-verity` - Disable dm_verity (requires root)

## Input Method

- [ ] `adb shell ime set com.android.adbkeyboard/.AdbIME` - Set input method to ADBKeyBoard
- [ ] `adb shell am broadcast -a ADB_INPUT_TEXT --es msg 'text'` - Input Chinese via ADBKeyBoard
