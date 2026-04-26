# devices-l Command

## Description
List connected devices with detailed information.

## Command
```bash
adbjson devices-l
```

## Equivalent ADB Command
```bash
adb devices -l
```

## Sample Output
```json
{
  "devices": [
    {
      "id": "ZD222XW5RL",
      "status": "device",
      "usb": "20-1",
      "product": "mumba_gp",
      "model": "moto_g57_power",
      "device": "mumba",
      "transport_id": "2"
    }
  ]
}
```

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
