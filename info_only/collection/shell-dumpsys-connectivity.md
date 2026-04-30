# shell dumpsys connectivity

Show connectivity service dump information.

## Command
```bash
adbjson shell dumpsys connectivity
```

## Description
Executes `adb shell dumpsys connectivity` and outputs the result as structured JSON.

## Examples

### Show connectivity service dump
```bash
./adbjson shell dumpsys connectivity
```

**Output:**
```json
{
  "sections": [
    {
      "name": "NetworkRequestInfo",
      "content": "NetworkRequestInfo:\n  mRequests:\n    NetworkRequest [ TRACK_DEFAULT id=1, legacyType=NONE, requestor=android:1000, isRequest=NetworkRequest [ TRACK_DEFAULT id=2, legacyType=NONE, requestor=android:1000, isRequest=...\n  mLegacyTypeMap: {NONE=0, MOBILE=2, WIFI=1, VPN=4, BLUETOOTH=3, ETHERNET=5, TEST=7, VPN_REQUIRED=8}\n"
    },
    {
      "name": "SatelliteAccessController",
      "content": "SupportConstrainedDataSatelliteOptIn: true\nRole-Sms Uids: {0={10207}}\nOpt-In Uids: {10206, 10207, 10222, 10241, 10257, 10387, 10411, 10512, 10606, 1010206}"
    },
    {
      "name": "Log",
      "content": "2026-04-25T21:17:59.216636 - SmsRoleUids:{10207} Opt-InUids:{10206, 10207, 10222, 10241, 10257, 10387, 10411, 10512, 10606, 1010206}\n2026-04-25T21:17:59.178789 - SmsRoleUids:{10207} Opt-InUids:{10206, 10207, 10222, 10241, 10257, 10387, 10411, 10512, 10606}\n2026-04-25T19:47:42.291528 - SmsRoleUids:{10207} Opt-InUids:{10206, 10207, 10222, 10241, 10257, 10387, 10411, 10512, 10606, 1010206}\n...\n"
    },
    {
      "name": "Legacy network activity",
      "content": "mTrackMultiNetworkActivities=true\nmDefaultCellularDataInactivityTimeout=10\nmDefaultWifiDataInactivityTimeout=15\nmIsDefaultNetworkActive=true\nmDefaultNetwork=230"
    },
    {
      "name": "Idle timers",
      "content": ""
    },
    {
      "name": "229",
      "content": "timeout=10 type=0"
    },
    {
      "name": "230",
      "content": "timeout=10 type=0\nWiFi active networks: {}\nCellular active networks: {230}\n\nClose QUIC connection: true\nRegistered QUIC connection close information: 0\n\nMulticast routing supported: true\nBackground firewall chain enabled: true\nIngressToVpnAddressFiltering: true\n"
    }
  ],
  "count": 34
}
```

## Response Fields

- **sections** (array): List of connectivity service sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "NetworkRequestInfo", "SatelliteAccessController")
- **content** (string): Raw content of the section

## Common Sections

### NetworkRequestInfo
- Network request tracking information
- Legacy type mappings
- Requestor information and UIDs
- Network request states and priorities

### SatelliteAccessController
- Satellite communication support
- SMS role assignments and UIDs
- Opt-in user lists and permissions
- Satellite data access control

### Log
- Historical connectivity events
- Permission changes and user actions
- Network state transitions
- Timestamped activity logs

### Legacy network activity
- Multi-network activity tracking
- Data inactivity timeouts
- Default network status
- Network activity monitoring

### Idle timers
- Network timeout configurations
- Per-network timer settings
- Inactivity management
- Connection lifecycle control

### Network-specific sections
- Individual network configurations
- Active network listings
- Connection state information
- Network-specific settings

## Notes

- Provides comprehensive connectivity service debugging information
- Essential for debugging network connectivity issues
- Shows detailed network request and state information
- Contains sensitive system information for debugging
- Content varies based on network hardware and Android version
- Useful for analyzing network performance and configuration problems
- Can help identify network routing and permission issues

## Related Commands

- `adbjson shell dumpsys wifi` - WiFi service information
- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell ifconfig` - Network interface information
- `adbjson shell ip addr show` - IP address information
