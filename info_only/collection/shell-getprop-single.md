# shell getprop [property]

Get system properties from the device.

## Command
```bash
adbjson shell getprop [property]
```

## Description
Executes `adb shell getprop` and outputs the result as structured JSON. If property is specified, gets only that property; otherwise gets all system properties.

## Examples

### Get all system properties
```bash
./adbjson shell getprop
```

**Output:**
```json
{
  "properties": [
    {
      "key": "ro.build.version.release",
      "value": "[16]"
    },
    {
      "key": "ro.product.model",
      "value": "[moto g play (2024)]"
    }
  ],
  "count": 1206
}
```

### Get single property
```bash
./adbjson shell getprop ro.build.version.release
```

**Output:**
```json
{
  "property": "ro.build.version.release",
  "value": "16"
}
```

### Get non-existent property
```bash
./adbjson shell getprop non.existent.property
```

**Output:**
```json
{
  "property": "non.existent.property",
  "value": ""
}
```

## Response Fields

### All Properties Response
- **properties** (array): List of property key-value pairs
- **count** (number): Total number of properties

### Single Property Response
- **property** (string): The requested property name
- **value** (string): The property value (empty if not found)

## Common Properties

- `ro.build.version.release` - Android version
- `ro.product.model` - Device model
- `ro.product.brand` - Device brand
- `ro.product.manufacturer` - Device manufacturer
- `ro.hardware` - Hardware platform
- `ro.serialno` - Device serial number
- `ro.sf.lcd_density` - Screen density
- `dhcp.wlan0.ipaddress` - WiFi IP address

## Notes

- Use single property queries for faster results when you know the property name
- Property values are returned as strings
- Empty value indicates property doesn't exist or has no value
- All properties output includes system, vendor, and OEM properties
