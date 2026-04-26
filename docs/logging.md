# Structured Logging

## Overview
adbjson uses structured JSON logging for better observability and debugging.

## Log Levels
- **DEBUG**: Detailed diagnostic information
- **INFO**: General informational messages (default)
- **WARN**: Warning messages
- **ERROR**: Error messages

## Enabling Debug Logging

### Via CLI Flag
```bash
./adbjson devices --debug
```

### Via Environment Variable
```bash
export ADBJSON_DEBUG=true
./adbjson devices
```

## Log Format
Logs are output as JSON to stderr:

```json
{
  "timestamp": "2026-04-26T09:24:11Z",
  "level": "info",
  "message": "Starting devices command",
  "context": {
    "device_count": 1
  }
}
```

## Usage in Commands
```go
import "adbjson/internal/logger"

func runCommand(cmd *cobra.Command, args []string) error {
    log := logger.Get()
    
    log.Info("Starting command", nil)
    log.Debug("Detailed info", map[string]interface{}{"key": "value"})
    log.Error("Something went wrong", map[string]interface{}{"error": err.Error()})
    
    return nil
}
```

## Benefits
- Structured logs are machine-readable
- Easy to parse and analyze
- Contextual information for debugging
- Consistent format across all commands
