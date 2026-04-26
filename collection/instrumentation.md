# instrumentation

## Description
Lists all instrumentation tests on the connected Android device.

## Command
```bash
adbjson instrumentation
```

## Equivalent ADB Command
```bash
adb shell pm list instrumentation
```

## Sample JSON Output
```json
{
  "instrumentations": [
    {
      "name": "com.example.test/androidx.test.runner.AndroidJUnitRunner"
    }
  ],
  "count": 1
}
```

## Sample YAML Output
```bash
adbjson instrumentation --format yaml
```
```yaml
instrumentations:
- name: com.example.test/androidx.test.runner.AndroidJUnitRunner
count: 1
```

## Flags
- `--compact`: Compact JSON output
- `--debug`: Enable debug logging
- `--format`: Output format (json, yaml)
- `--pretty`: Pretty print JSON output

## Notes
- Requires a connected device or emulator
- Lists all instrumentation test runners
- Output is limited to first 5 items in documentation (full list in actual output)
