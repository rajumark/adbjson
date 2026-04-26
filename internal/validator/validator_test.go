package validator

import (
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	validator := NewValidator()
	
	tests := []struct {
		name    string
		data    interface{}
		wantErr bool
	}{
		{
			name:    "nil data",
			data:    nil,
			wantErr: true,
		},
		{
			name:    "valid struct",
			data:    struct{ Name string }{Name: "test"},
			wantErr: false,
		},
		{
			name:    "valid slice",
			data:    []string{"a", "b"},
			wantErr: false,
		},
		{
			name:    "valid map",
			data:    map[string]string{"key": "value"},
			wantErr: false,
		},
		{
			name:    "nil pointer",
			data:    (*string)(nil),
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_ValidateJSON(t *testing.T) {
	validator := NewValidator()
	
	tests := []struct {
		name         string
		jsonBytes    []byte
		expectedType interface{}
		wantErr      bool
	}{
		{
			name:      "valid JSON",
			jsonBytes: []byte(`{"key":"value"}`),
			wantErr:   false,
		},
		{
			name:      "invalid JSON",
			jsonBytes: []byte(`{invalid}`),
			wantErr:   true,
		},
		{
			name:         "JSON matching type",
			jsonBytes:    []byte(`{"name":"test"}`),
			expectedType: struct{ Name string }{},
			wantErr:      false,
		},
		{
			name:         "JSON not matching type",
			jsonBytes:    []byte(`["array","item"]`),
			expectedType: struct{ Name string }{},
			wantErr:      true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateJSON(tt.jsonBytes, tt.expectedType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	err := &ValidationError{Field: "test", Message: "invalid"}
	expected := "validation failed for field 'test': invalid"
	if err.Error() != expected {
		t.Errorf("Error() = %v, want %v", err.Error(), expected)
	}
}
