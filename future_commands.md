# Future ADB Commands

This document lists 100 most commonly used ADB commands planned for future implementation in adbjson.

## Device Management

1. `adb devices` - List connected devices ✅
2. `adb devices -l` - List devices with detailed info
3. `adb connect <host:port>` - Connect to device via TCP/IP
4. `adb disconnect <host:port>` - Disconnect from device
5. `adb kill-server` - Kill ADB server
6. `adb start-server` - Start ADB server
7. `adb version` - Show ADB version ✅
8. `adb get-state` - Get device state
9. `adb get-serialno` - Get serial number
10. `adb get-devpath` - Get device path

## File Management

11. `adb push <local> <remote>` - Push file to device
12. `adb pull <remote> <local>` - Pull file from device
13. `adb sync <directory>` - Sync directory
14. `adb shell ls <path>` - List files
15. `adb shell cd <path>` - Change directory
16. `adb shell mkdir <path>` - Create directory
17. `adb shell rm <file>` - Remove file
18. `adb shell rmdir <dir>` - Remove directory
19. `adb shell cp <src> <dst>` - Copy file
20. `adb shell mv <src> <dst>` - Move file
21. `adb shell cat <file>` - View file content
22. `adb shell chmod <permissions> <file>` - Change permissions
23. `adb shell chown <owner> <file>` - Change owner

## Package Management

24. `adb shell pm list packages` - List all packages
25. `adb shell pm list packages -3` - List third-party packages
26. `adb shell pm list packages -s` - List system packages
27. `adb shell pm path <package>` - Get package path
28. `adb shell pm install <apk>` - Install package
29. `adb install <apk>` - Install APK
30. `adb install -r <apk>` - Reinstall package
31. `adb uninstall <package>` - Uninstall package
32. `adb shell pm uninstall <package>` - Uninstall package
33. `adb shell pm clear <package>` - Clear package data
34. `adb shell pm enable <package>` - Enable package
35. `adb shell pm disable <package>` - Disable package
36. `adb shell pm dump <package>` - Dump package info
37. `adb shell pm grant <package> <permission>` - Grant permission
38. `adb shell pm revoke <package> <permission>` - Revoke permission

## App Information

39. `adb shell dumpsys package <package>` - Dump package info
40. `adb shell dumpsys activity` - Dump activity info
41. `adb shell dumpsys battery` - Dump battery info
42. `adb shell dumpsys wifi` - Dump WiFi info
43. `adb shell dumpsys connectivity` - Dump connectivity info
44. `adb shell dumpsys telephony` - Dump telephony info
45. `adb shell dumpsys window` - Dump window info
46. `adb shell dumpsys meminfo` - Dump memory info
47. `adb shell dumpsys cpuinfo` - Dump CPU info
48. `adb shell dumpsys dbinfo` - Dump database info

## Screen & Input

49. `adb shell screencap <file>` - Capture screenshot
50. `adb shell screenrecord <file>` - Record screen
51. `adb shell input tap <x> <y>` - Tap screen
52. `adb shell input swipe <x1> <y1> <x2> <y2>` - Swipe screen
53. `adb shell input text <text>` - Type text
54. `adb shell input keyevent <code>` - Send key event
55. `adb shell wm size` - Get screen size
56. `adb shell wm density` - Get screen density
57. `adb shell wm density <dpi>` - Set screen density
58. `adb shell wm overscan <left,top,right,bottom>` - Set overscan

## Logcat

59. `adb logcat` - View logs
60. `adb logcat -c` - Clear logs
61. `adb logcat -d` - Dump logs
62. `adb logcat -f <file>` - Save logs to file
63. `adb logcat -v time` - Show timestamps
64. `adb logcat -v brief` - Brief format
65. `adb logcat -v process` - Process format
66. `adb logcat -v thread` - Thread format
67. `adb logcat -v raw` - Raw format
68. `adb logcat -v tag` - Tag format
69. `adb logcat -s <tag>` - Filter by tag
70. `adb logcat *:E` - Show only errors
71. `adb logcat *:W` - Show warnings and errors
72. `adb logcat *:I` - Show info and above
73. `adb logcat -t <count>` - Show last N lines

## Shell Commands

74. `adb shell` - Enter interactive shell
75. `adb shell <command>` - Execute shell command
76. `adb shell ps` - List processes
77. `adb shell top` - Show running processes
78. `adb shell kill <pid>` - Kill process
79. `adb shell killall <name>` - Kill all processes
80. `adb shell df` - Disk usage
81. `adb shell du` - Directory usage
82. `adb shell free` - Memory usage
83. `adb shell uptime` - System uptime
84. `adb shell date` - System date/time
85. `adb shell reboot` - Reboot device
86. `adb shell reboot recovery` - Reboot to recovery
87. `adb shell reboot bootloader` - Reboot to bootloader
88. `adb shell shutdown` - Shutdown device

## Network

89. `adb shell netstat` - Network statistics
90. `adb shell ifconfig` - Network interfaces
91. `adb shell ip addr show` - IP addresses
92. `adb shell ping <host>` - Ping host
93. `adb shell netcfg` - Network configuration
94. `adb shell svc wifi disable` - Disable WiFi
95. `adb shell svc wifi enable` - Enable WiFi
96. `adb shell svc data disable` - Disable mobile data
97. `adb shell svc data enable` - Enable mobile data

## Battery & Power

98. `adb shell dumpsys battery` - Battery info
99. `adb shell dumpsys batterystats` - Battery statistics
100. `adb shell settings put global airplane_mode_on 1` - Enable airplane mode
