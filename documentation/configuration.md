# Configuration Management

## Overview
adbjson supports configuration through environment variables and (in the future) config files.

## Environment Variables

### ADB Configuration
- `ADBJSON_ADB_PATH` - Custom path to ADB executable (default: bundled ADB)
- `ADBJSON_TIMEOUT` - Default timeout for ADB commands in seconds (default: 30)

### Logging Configuration
- `ADBJSON_LOG_LEVEL` - Logging level: debug, info, warn, error (default: info)
- `ADBJSON_DEBUG` - Enable debug mode (true/false) (default: false)

### Output Configuration
- `ADBJSON_OUTPUT_FORMAT` - Default JSON output format: pretty, compact (default: pretty)

## Usage Examples

### Set Custom ADB Path
```bash
export ADBJSON_ADB_PATH=/usr/local/bin/adb
./adbjson devices
```

### Set Log Level
```bash
export ADBJSON_LOG_LEVEL=debug
./adbjson devices
```

### Set Timeout
```bash
export ADBJSON_TIMEOUT=60
./adbjson devices
```

### Enable Debug Mode
```bash
export ADBJSON_DEBUG=true
./adbjson devices
```

## Config File Location
Config files are stored in:
- `~/.adbjson/config.yaml` (TODO: Implement YAML support)

## Programmatic Usage

```go
import "adbjson/internal/config"

// Load configuration
cfg := config.Load()

// Access configuration
adbPath := cfg.ADBPath
logLevel := cfg.LogLevel
timeout := cfg.Timeout
```

## Default Values
| Setting | Default |
|---------|---------|
| ADB Path | Bundled ADB |
| Log Level | info |
| Output Format | pretty |
| Timeout | 30 seconds |
| Platform | Auto-detected |

## Future Enhancements
- YAML config file support
- Profile-based configurations
- Config validation
- Config hot-reloading
