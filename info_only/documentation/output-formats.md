# Output Formats

## Overview
adbjson supports multiple output formats for flexibility in integration with different tools and workflows.

## Supported Formats

### JSON (default)
Structured JSON output, the primary format for adbjson.

```bash
./adbjson devices
./adbjson devices --format json
```

### YAML
Human-readable YAML format for configuration files and documentation.

```bash
./adbjson devices --format yaml
```

## Usage Examples

### JSON Output
```bash
./adbjson devices
```
Output:
```json
{
  "devices": []
}
```

### YAML Output
```bash
./adbjson devices --format yaml
```
Output:
```yaml
devices: []
```

### Compact JSON
```bash
./adbjson devices --compact
```
Output:
```json
{"devices":[]}
```

### Compact YAML
```bash
./adbjson devices --format yaml --compact
```
Note: YAML formatter ignores compact flag as YAML is already human-readable.

## Programmatic Usage

```go
import "adbjson/internal/formatter"

// Format as JSON
jsonOutput, err := formatter.FormatOutputString(data, formatter.JSONFormat, false)

// Format as YAML
yamlOutput, err := formatter.FormatOutputString(data, formatter.YAMLFormat, false)
```

## Future Formats
- XML
- CSV (for list commands)
- TOML
