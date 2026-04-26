# shell dumpsys location

Show location service dump information.

## Command
```bash
adbjson shell dumpsys location
```

## Description
Executes `adb shell dumpsys location` and outputs the result as structured JSON.

## Examples

### Show location service dump
```bash
./adbjson shell dumpsys location
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Location Service State",
      "content": "mEnabled=true\nmNetworkLocationEnabled=true\nmGpsLocationEnabled=true\nmPassiveLocationEnabled=true\n..."
    },
    {
      "name": "Location Providers",
      "content": "gps provider: ProviderState[ENABLED, ProviderRequest[@+10m0s0ms, HIGH_ACCURACY, WorkSource{10257 com.google.android.apps.maps}]]\nnetwork provider: ProviderState[ENABLED, ProviderRequest[@+6h0m0s0ms, LOW_POWER, WorkSource{10179 com.google.android.as}]]\npassive provider: ProviderState[ENABLED, ProviderRequest[OFF]]\nfused provider: ProviderState[ENABLED, ProviderRequest[@+10m0s0ms, HIGH_ACCURACY, WorkSource{10257 com.google.android.apps.maps}]]\n..."
    },
    {
      "name": "Location Requests",
      "content": "10206/com.google.android.gms[fused_location_provider]/9B8CA2AE: ProviderRequest[@+10m0s0ms, HIGH_ACCURACY, WorkSource{10257 com.google.android.apps.maps}]\n10257/com.google.android.apps.maps/663F82EB: ProviderRequest[@+10m0s0ms, HIGH_ACCURACY, WorkSource{10257 com.google.android.apps.maps}]\n..."
    },
    {
      "name": "Location Events",
      "content": "04-26 17:23:16.633: passive provider delivered location[1] to 10206/com.google.android.gms[fused_location_provider]/55FB1C2D\n04-26 17:23:16.880: gps provider received location[1]\n04-26 17:23:16.883: gps provider delivered location[1] to 10206/com.google.android.gms[fused_location_provider]/9B8CA2AE\n..."
    },
    {
      "name": "GNSS Status",
      "content": "mEnabled=true\nmStarted=true\nmFixCount=1234\nmTimeToFirstFix=2500\nmTTFF=2500\n..."
    },
    {
      "name": "Geofence Registry",
      "content": "mGeofences=[Geofence{id=1, latitude=37.7749, longitude=-122.4194, radius=100.0, ...}]\n..."
    },
    {
      "name": "Location Listener Registry",
      "content": "mListeners=[LocationListener{uid=10206, pid=1234, packageName=com.google.android.gms, ...}]\n..."
    },
    {
      "name": "Location Permissions",
      "content": "mLocationPermissions={10206: [ACCESS_FINE_LOCATION, ACCESS_COARSE_LOCATION], 10257: [ACCESS_FINE_LOCATION], ...}\n..."
    }
  ],
  "count": 35
}
```

## Response Fields

- **sections** (array): List of location service sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "Location Service State", "Location Providers")
- **content** (string): Raw content of the section

## Common Sections

### Location Service State
- Overall location service status
- Enabled/disabled states for different location types
- Service configuration and settings
- System-wide location policies

### Location Providers
- GPS provider status and configuration
- Network provider status and settings
- Passive provider information
- Fused location provider details
- Provider request parameters and work sources

### Location Requests
- Active location requests from apps
- Request parameters and accuracy requirements
- Work source attribution and package information
- Request timeouts and intervals

### Location Events
- Location update events and timestamps
- Provider-to-provider location deliveries
- Location request and fulfillment logs
- Error events and provider status changes

### GNSS Status
- GPS/GNSS system status
- Satellite information and signal strength
- Time to first fix (TTFF) statistics
- GNSS configuration and settings

### Geofence Registry
- Active geofences and their parameters
- Geofence transition events
- Geofence monitoring status
- Location-based trigger information

### Location Listener Registry
- Registered location listeners
- Listener parameters and requirements
- Package and process information
- Listener lifecycle and status

### Location Permissions
- App location permissions
- Permission grants and denials
- Background location access
- Permission-related events and changes

## Notes

- Provides comprehensive location service debugging information
- Essential for debugging location accuracy and performance issues
- Shows detailed provider status and request information
- Contains sensitive location data and system information
- Content varies based on location hardware and Android version
- Useful for analyzing location provider behavior and app location usage
- Can help identify location permission issues and provider conflicts

## Related Commands

- `adbjson shell dumpsys connectivity` - Connectivity service information
- `adbjson shell dumpsys telephony.registry` - Telephony registry information
- `adbjson shell dumpsys activity` - Activity manager information
