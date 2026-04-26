package formatter

import (
	"fmt"
	"strings"
	"testing"
)

func TestJSONFormatter_Format(t *testing.T) {
	formatter := NewJSONFormatter()
	data := map[string]string{"key": "value"}
	
	// Test compact
	compact, err := formatter.Format(data, true)
	if err != nil {
		t.Fatalf("Format() compact error = %v", err)
	}
	if string(compact) != `{"key":"value"}` {
		t.Errorf("Format() compact = %v, want %v", string(compact), `{"key":"value"}`)
	}
	
	// Test pretty
	pretty, err := formatter.Format(data, false)
	if err != nil {
		t.Fatalf("Format() pretty error = %v", err)
	}
	if string(pretty) != "{\n  \"key\": \"value\"\n}" {
		t.Errorf("Format() pretty = %v, want formatted JSON", string(pretty))
	}
}

func TestYAMLFormatter_Format(t *testing.T) {
	formatter := NewYAMLFormatter()
	data := map[string]string{"key": "value"}
	
	output, err := formatter.Format(data, false)
	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}
	
	yamlStr := string(output)
	if yamlStr == "" {
		t.Error("Format() returned empty string")
	}
}

func TestParseFormat(t *testing.T) {
	tests := []struct {
		input string
		want  OutputFormat
	}{
		{"json", JSONFormat},
		{"yaml", YAMLFormat},
		{"yml", YAMLFormat},
		{"invalid", JSONFormat}, // defaults to JSON
		{"", JSONFormat},        // defaults to JSON
	}
	
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ParseFormat(tt.input); got != tt.want {
				t.Errorf("ParseFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFormatter(t *testing.T) {
	tests := []struct {
		format OutputFormat
		want   string
	}{
		{JSONFormat, "JSONFormatter"},
		{YAMLFormat, "YAMLFormatter"},
	}
	
	for _, tt := range tests {
		t.Run(string(tt.format), func(t *testing.T) {
			formatter := GetFormatter(tt.format)
			formatterType := fmt.Sprintf("%T", formatter)
			if !strings.Contains(formatterType, tt.want) {
				t.Errorf("GetFormatter() = %v, want %v", formatterType, tt.want)
			}
		})
	}
}

func TestSupportedFormats(t *testing.T) {
	formats := SupportedFormats()
	if len(formats) != 2 {
		t.Errorf("SupportedFormats() = %v, want 2 formats", len(formats))
	}
}

func TestIsValidFormat(t *testing.T) {
	if !IsValidFormat("json") {
		t.Error("IsValidFormat(json) = false, want true")
	}
	if !IsValidFormat("yaml") {
		t.Error("IsValidFormat(yaml) = false, want true")
	}
	if !IsValidFormat("invalid") {
		t.Error("IsValidFormat(invalid) = false, want true (defaults to JSON)")
	}
}
