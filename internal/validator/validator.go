package validator

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Validator validates data structures
type Validator struct{}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks if a data structure is valid
func (v *Validator) Validate(data interface{}) error {
	// Check for nil
	if data == nil {
		return &ValidationError{Field: "root", Message: "data is nil"}
	}

	// Use reflection to validate
	val := reflect.ValueOf(data)
	
	// If it's a pointer, get the underlying value
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return &ValidationError{Field: "root", Message: "data is nil pointer"}
		}
		val = val.Elem()
	}

	// Validate based on type
	switch val.Kind() {
	case reflect.Struct:
		return v.validateStruct(val)
	case reflect.Slice, reflect.Array:
		return v.validateSlice(val)
	case reflect.Map:
		return v.validateMap(val)
	}

	return nil
}

// validateStruct validates a struct
func (v *Validator) validateStruct(val reflect.Value) error {
	typ := val.Type()
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}
		
		// Check for nil pointers in struct
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			if err := v.Validate(field.Interface()); err != nil {
				return &ValidationError{
					Field:   fieldType.Name,
					Message: err.Error(),
				}
			}
		}
	}
	
	return nil
}

// validateSlice validates a slice
func (v *Validator) validateSlice(val reflect.Value) error {
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		if elem.CanInterface() {
			if err := v.Validate(elem.Interface()); err != nil {
				return &ValidationError{
					Field:   fmt.Sprintf("[%d]", i),
					Message: err.Error(),
				}
			}
		}
	}
	return nil
}

// validateMap validates a map
func (v *Validator) validateMap(val reflect.Value) error {
	for _, key := range val.MapKeys() {
		elem := val.MapIndex(key)
		if elem.CanInterface() {
			if err := v.Validate(elem.Interface()); err != nil {
				return &ValidationError{
					Field:   fmt.Sprintf("%v", key.Interface()),
					Message: err.Error(),
				}
			}
		}
	}
	return nil
}

// ValidateJSON validates JSON bytes against expected structure
func (v *Validator) ValidateJSON(jsonBytes []byte, expectedType interface{}) error {
	// Unmarshal to check if JSON is valid
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return &ValidationError{Field: "json", Message: "invalid JSON: " + err.Error()}
	}
	
	// If expected type is provided, unmarshal to that type
	if expectedType != nil {
		expectedValue := reflect.New(reflect.TypeOf(expectedType))
		if err := json.Unmarshal(jsonBytes, expectedValue.Interface()); err != nil {
			return &ValidationError{Field: "json", Message: "JSON does not match expected type: " + err.Error()}
		}
	}
	
	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Message)
}
