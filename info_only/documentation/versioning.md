# Version Management

## Overview
adbjson uses semantic versioning and includes build information for traceability.

## Version Format
- **Format**: MAJOR.MINOR.PATCH (e.g., 1.0.0)
- **Semantic Versioning**: 
  - MAJOR: Breaking changes
  - MINOR: New features (backward compatible)
  - PATCH: Bug fixes (backward compatible)

## Version Commands

### Show CLI Version
```bash
./adbjson version
```

Output (JSON):
```json
{
  "version": "1.0.0",
  "commit": "unknown",
  "branch": "unknown",
  "build_date": "unknown"
}
```

Output (YAML):
```bash
./adbjson version --format yaml
```

### Show ADB Version
```bash
./adbjson adb-version
```

## Build Information

Build information can be set during compilation using ldflags:

```bash
go build -ldflags \
  "-X adbjson/internal/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
   -X adbjson/internal/version.GitCommit=$(git rev-parse HEAD) \
   -X adbjson/internal/version.GitBranch=$(git rev-parse --abbrev-ref HEAD)"
```

## Version Policy

### Backward Compatibility
- MAJOR version changes may break backward compatibility
- MINOR and PATCH versions maintain backward compatibility
- Deprecated features will be marked in documentation

### Deprecation Process
1. Mark feature as deprecated in documentation
2. Add deprecation warning in CLI output
3. Remove in next MAJOR version

## Programmatic Usage

```go
import "adbjson/internal/version"

// Get version string
version := version.GetVersion() // "1.0.0"

// Get full version with build info
fullVersion := version.GetFullVersion()

// Get build information as map
buildInfo := version.GetBuildInfo()
```

## Current Version
- **Version**: 1.0.0
- **Status**: Stable
