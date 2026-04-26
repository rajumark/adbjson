# Error Handling Strategy

## Overview
adbjson uses structured error types for consistent error handling across all commands.

## Error Types

```go
type ErrorType int

const (
    ADBExecutionError ErrorType = iota  // Error executing ADB command
    ParseError                           // Error parsing ADB output
    ValidationError                      // Error validating parsed result
    MarshalError                         // Error marshaling to JSON
    DeviceError                          // Device not connected or offline
    PermissionError                      // Insufficient permissions
)
```

## AppError Structure

```go
type AppError struct {
    Type    ErrorType
    Message string
    Err     error
    Context map[string]interface{}
}
```

## Usage Examples

### ADB Execution Error
```go
output, err := executor.Execute("devices")
if err != nil {
    return apperrors.NewADBExecutionError("devices", err)
}
```

### Parse Error
```go
response, err := parser.Parse(output)
if err != nil {
    return apperrors.NewParseError(parser.Name(), err)
}
```

### Validation Error
```go
if err := parser.Validate(response); err != nil {
    return apperrors.NewValidationError("devices", err.Error())
}
```

### Marshal Error
```go
jsonBytes, err := json.Marshal(response)
if err != nil {
    return apperrors.NewMarshalError(err)
}
```

### Device Error
```go
if len(devices) == 0 {
    return apperrors.NewDeviceError("no devices connected")
}
```

### Permission Error
```go
if !hasPermission {
    return apperrors.NewPermissionError("insufficient permissions")
}
```

## Error Context
All errors can include contextual information:

```go
return &AppError{
    Type:    ADBExecutionError,
    Message: "failed to execute ADB command",
    Err:     err,
    Context: map[string]interface{}{
        "command": "devices",
        "timeout": 30,
    },
}
```

## Benefits
- **Type Safety**: Categorized error types
- **Context**: Rich error context for debugging
- **Unwrapping**: Support for error unwrapping
- **Logging**: Structured errors work well with structured logging
- **Testing**: Easy to test error conditions
