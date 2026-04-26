package errors

import "fmt"

// ErrorType represents the category of error
type ErrorType int

const (
	// ADBExecutionError - Error executing ADB command
	ADBExecutionError ErrorType = iota
	// ParseError - Error parsing ADB output
	ParseError
	// ValidationError - Error validating parsed result
	ValidationError
	// MarshalError - Error marshaling to JSON
	MarshalError
	// DeviceError - Device not connected or offline
	DeviceError
	// PermissionError - Insufficient permissions
	PermissionError
)

// AppError represents a structured application error
type AppError struct {
	Type    ErrorType
	Message string
	Err     error
	Context map[string]interface{}
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewADBExecutionError creates a new ADB execution error
func NewADBExecutionError(command string, err error) *AppError {
	return &AppError{
		Type:    ADBExecutionError,
		Message: fmt.Sprintf("failed to execute ADB command: %s", command),
		Err:     err,
		Context: map[string]interface{}{"command": command},
	}
}

// NewParseError creates a new parse error
func NewParseError(parser string, err error) *AppError {
	return &AppError{
		Type:    ParseError,
		Message: fmt.Sprintf("failed to parse output with parser: %s", parser),
		Err:     err,
		Context: map[string]interface{}{"parser": parser},
	}
}

// NewValidationError creates a new validation error
func NewValidationError(field string, reason string) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: fmt.Sprintf("validation failed for field: %s", field),
		Context: map[string]interface{}{"field": field, "reason": reason},
	}
}

// NewMarshalError creates a new marshal error
func NewMarshalError(err error) *AppError {
	return &AppError{
		Type:    MarshalError,
		Message: "failed to marshal JSON",
		Err:     err,
	}
}

// NewDeviceError creates a new device error
func NewDeviceError(message string) *AppError {
	return &AppError{
		Type:    DeviceError,
		Message: message,
	}
}

// NewPermissionError creates a new permission error
func NewPermissionError(message string) *AppError {
	return &AppError{
		Type:    PermissionError,
		Message: message,
	}
}

// GetType returns the error type
func (e *AppError) GetType() ErrorType {
	return e.Type
}

// GetContext returns the error context
func (e *AppError) GetContext() map[string]interface{} {
	return e.Context
}
