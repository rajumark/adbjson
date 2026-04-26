# shell dumpsys wifi

Show WiFi service dump information.

## Command
```bash
adbjson shell dumpsys wifi
```

## Description
Executes `adb shell dumpsys wifi` and outputs the result as structured JSON.

## Examples

### Show WiFi service dump
```bash
./adbjson shell dumpsys wifi
```

**Output:**
```json
{
  "sections": [
    {
      "name": "AdaptiveConnectivityEnabledSettingObserver",
      "content": "mAdaptiveConnectivityEnabled=true"
    },
    {
      "name": "WifiGlobals",
      "content": "mPollRssiIntervalMillis=3000\nmIsPollRssiIntervalOverridden=false\nmPollRssiShortIntervalMillis=3000\nmPollRssiLongIntervalMillis=6000\nmIpReachabilityDisconnectEnabled=true\nmIsBluetoothConnected=false\nmIsWpa3SaeUpgradeOffloadEnabled=false\nmIsUsingExternalScorer=false\nmIsWepAllowed=true\nmDisableFirmwareRoamingInIdleMode=false\nIsD2dSupportedWhenInfraStaDisabled=false\nmIsWpa3SaeH2eSupported=true\ncarrierId=1915, eapFailureCode=16385, displayNotification=false, threshold=3, durationMs=60000\nmSendDhcpHostnameRestriction=0"
    },
    {
      "name": "SarManager",
      "content": "isSarSupported: false\nisSarVoiceCallSupported: false\nisSarSoftApSupported: false\n"
    },
    {
      "name": "SarInfo",
      "content": "Current values:\nVoice Call state is: false\nWifi Client state is: false\nWifi Soft AP state is: false\nWifi ScanOnly state is: true\nEarpiece state is : false\nLast reported values:\nSoft AP state is: false\nVoice Call state is: false\nEarpiece state is: false\nLast reported scenario: -2\nReported 1777206216 seconds ago\n"
    },
    {
      "name": "LastCallerInfoManager",
      "content": "API key=1 API name=ScanningEnabled: tid=4914 uid=1000 pid=2843 packageName=android toggleState=true\nAPI key=2 API name=WifiEnabled: tid=5706 uid=10331 pid=5051 packageName=com.android.systemui toggleState=false\nAPI key=3 API name=SoftAp: tid=12424 uid=1073 pid=5237 packageName=<unknown> toggleState=false\nAPI key=4 API name=TetheredHotspot: tid=5685 uid=1073 pid=5237 packageName=com.google.android.networkstack.tethering toggleState=true\nAPI key=13 API name=API_CONNECT_CONFIG: tid=12814 uid=1000 pid=0 packageName=com.android.settings toggleState=true\nAPI key=18 API name=API_SAVE: tid=12814 uid=1000 pid=0 packageName=com.android.settings toggleState=true\nAPI key=19 API name=API_START_SCAN: tid=5207 uid=10597 pid=2843 packageName=com.instagram.android toggleState=true\nAPI key=33 API name=API_WIFI_SCANNER_START_SCAN: tid=5769 uid=10206 pid=23509 packageName=com.google.android.gms toggleState=true\n"
    },
    {
      "name": "WifiNative",
      "content": "mIsLocationModeEnabled: true\nmLastLocationModeEnabledTimeMs: 60459"
    },
    {
      "name": "HostapdHal",
      "content": "AIDL service declared: true\nHIDL service declared: false\nInitialized: true\nImplementation: HostapdHalAidlImp\nAIDL interface version: 2\n\n"
    },
    {
      "name": "WifiResourceCache",
      "content": "WifiResourceCache - resource value Begin ----\nResource Name: config_wifiDisableFirmwareRoamingInIdleMode, value: false\nResource Name: config_wifi6ghzSupport, value: false\nResource Name: config_wifi5ghzSupport, value: true\nResource Name: config_wifi24ghzSupport, value: true\nResource Name: config_wifi_background_scan_support, value: true\n...\nWifiResourceCache - resource value End ----\nboundToExternalScorer=failure, lastScorerBindingState=-1\n"
    }
  ],
  "count": 47
}
```

## Response Fields

- **sections** (array): List of WiFi service sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "WifiGlobals", "SarManager")
- **content** (string): Raw content of the section

## Common Sections

### WifiGlobals
- Global WiFi configuration and state
- RSSI polling intervals and settings
- Bluetooth connectivity status
- WPA3 and security protocol support
- Carrier information and EAP settings

### SarManager & SarInfo
- Specific Absorption Rate (SAR) management
- Voice call, WiFi client, and soft AP states
- SAR reporting and compliance information

### LastCallerInfoManager
- API call tracking and permissions
- Package names and process IDs
- Toggle states for various WiFi features
- Recent WiFi API usage history

### WifiNative
- Native WiFi service state
- Location mode settings
- Hardware interface information

### HostapdHal
- Host AP daemon HAL interface
- AIDL/HIDL service declarations
- Implementation details and version

### WifiResourceCache
- Configuration resource values
- Feature support flags
- Hardware capability settings
- System-wide WiFi configuration

## Notes

- Provides comprehensive WiFi service debugging information
- Essential for debugging WiFi connectivity issues
- Shows detailed hardware and software configuration
- Contains sensitive system information for debugging
- Content varies based on WiFi hardware and Android version
- Useful for analyzing WiFi performance and configuration problems
- Can help identify hardware limitations and software bugs

## Related Commands

- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell ifconfig` - Network interface information
- `adbjson shell ip addr show` - IP address information
