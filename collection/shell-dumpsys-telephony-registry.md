# shell dumpsys telephony.registry

Show telephony registry dump information.

## Command
```bash
adbjson shell dumpsys telephony.registry
```

## Description
Executes `adb shell dumpsys telephony.registry` and outputs the result as structured JSON.

## Examples

### Show telephony registry dump
```bash
./adbjson shell dumpsys telephony.registry
```

**Output:**
```json
{
  "sections": [
    {
      "name": "PhoneStateListenerRegistry",
      "content": "PhoneStateListenerRegistry:\n  mRecords:\n    {callingPackage=*** callerUid=1000 binder=android.os.BinderProxy@b3ba9f0 callback=null onSubscriptionsChangedListenererCallback=null onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=com.android.internal.telephony.ICarrierPrivilegesCallback$Stub$Proxy@5ceab69 carrierConfigChangeListener=null satelliteStateChangeListener=null subId=-1 phoneId=0 events={}}\n    {callingPackage=*** callerUid=1000 binder=android.telephony.TelephonyCallback$IPhoneStateListenerStub@d6f29ee callback=android.telephony.TelephonyCallback$IPhoneStateListenerStub@d6f29ee onSubscriptionsChangedListenererCallback=null onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=null carrierConfigChangeListener=null satelliteStateChangeListener=null subId=1 phoneId=0 events=[9]}\n    {callingPackage=*** callerUid=10331 binder=android.os.BinderProxy@da084ab callback=com.android.internal.telephony.IPhoneStateListener$Stub$Proxy@d06c708 onSubscriptionsChangedListenererCallback=null onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=null carrierConfigChangeListener=null satelliteStateChangeListener=null subId=1 phoneId=0 events=[1, 17, 34, 21, 6, 7, 8, 9, 42, 43, 44, 45]}\n    ...\n"
    },
    {
      "name": "SubscriptionControllerRegistry",
      "content": "SubscriptionControllerRegistry:\n  mRecords:\n    {callingPackage=*** callerUid=10207 binder=android.os.BinderProxy@42c40d9 callback=null onSubscriptionsChangedListenererCallback=com.android.internal.telephony.IOnSubscriptionsChangedListener$Stub$Proxy@36fea9e onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=null carrierConfigChangeListener=null satelliteStateChangeListener=null subId=-1 phoneId=-1 events={}}\n    {callingPackage=*** callerUid=10206 binder=android.os.BinderProxy@afee9b callback=null onSubscriptionsChangedListenererCallback=com.android.internal.telephony.IOnSubscriptionsChangedListener$Stub$Proxy@c7e0238 onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=null carrierConfigChangeListener=null satelliteStateChangeListener=null subId=-1 phoneId=-1 events={}}\n    ...\n"
    },
    {
      "name": "CarrierConfigChangeListenerRegistry",
      "content": "CarrierConfigChangeListenerRegistry:\n  mRecords:\n    {callingPackage=*** callerUid=1000 binder=android.os.BinderProxy@b3ba9f0 callback=null onSubscriptionsChangedListenererCallback=null onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=com.android.internal.telephony.ICarrierPrivilegesCallback$Stub$Proxy@5ceab69 carrierConfigChangeListener=null satelliteStateChangeListener=null subId=-1 phoneId=0 events={}}\n    {callingPackage=*** callerUid=1001 binder=android.os.BinderProxy@29d9525 callback=null onSubscriptionsChangedListenererCallback=null onOpportunisticSubscriptionsChangedListenererCallback=null carrierPrivilegesCallback=com.android.internal.telephony.ICarrierPrivilegesCallback$Stub$Proxy@27f6fa carrierConfigChangeListener=null satelliteStateChangeListener=null subId=-1 phoneId=0 events={}}\n    ...\n"
    }
  ],
  "count": 3
}
```

## Response Fields

- **sections** (array): List of telephony registry sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "PhoneStateListenerRegistry", "SubscriptionControllerRegistry")
- **content** (string): Raw content of the section

## Common Sections

### PhoneStateListenerRegistry
- Phone state listener registrations
- Callback binder information
- Subscription and phone ID mappings
- Event listener configurations
- Package and UID information

### SubscriptionControllerRegistry
- Subscription change listeners
- Opportunistic subscription listeners
- Subscription state tracking
- User-specific subscription information
- Carrier privilege callbacks

### CarrierConfigChangeListenerRegistry
- Carrier configuration change listeners
- Carrier privilege callbacks
- Configuration update tracking
- Satellite state change listeners
- Multi-subscription support

## Registry Information

Each registry section contains detailed information about:

- **callingPackage**: Package name of the caller (masked as ***)
- **callerUid**: User ID of the calling process
- **binder**: Binder reference for IPC communication
- **callback**: Callback interface reference
- **subId**: Subscription ID (-1 for all subscriptions)
- **phoneId**: Phone ID for multi-SIM devices
- **events**: List of registered events for the listener

## Event Types

Common event types include:
- 1: Service state changes
- 6: Signal strength changes
- 7: Message waiting indicator
- 8: Call forwarding indicator
- 9: Data connection state
- 17: Oem hook raw
- 21: Precise call state
- 34: Data activity
- 42-45: Carrier network change events

## Notes

- Provides comprehensive telephony registry debugging information
- Essential for debugging telephony and cellular connectivity issues
- Shows detailed listener and callback registrations
- Contains sensitive system information for debugging
- Content varies based on telephony hardware and Android version
- Useful for analyzing telephony service interactions and state management
- Can help identify listener registration issues and callback problems

## Related Commands

- `adbjson shell dumpsys connectivity` - Connectivity service information
- `adbjson shell dumpsys activity` - Activity manager information
- `adbjson shell dumpsys wifi` - WiFi service information
