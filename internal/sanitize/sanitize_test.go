package sanitize

import "testing"

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normal string",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "string with null byte",
			input: "hello\x00world",
			want:  "helloworld",
		},
		{
			name:  "string with control characters",
			input: "hello\x01\x02world",
			want:  "helloworld",
		},
		{
			name:  "string with tab",
			input: "hello\tworld",
			want:  "hello\tworld",
		},
		{
			name:  "string with newline",
			input: "hello\nworld",
			want:  "hello\nworld",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeString(tt.input); got != tt.want {
				t.Errorf("SanitizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeFilePath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normal path",
			input: "sdcard/file.txt",
			want:  "sdcard/file.txt",
		},
		{
			name:  "path traversal",
			input: "../etc/passwd",
			want:  "etc/passwd",
		},
		{
			name:  "windows traversal",
			input: "..\\windows\\system32",
			want:  "windowssystem32",
		},
		{
			name:  "absolute path",
			input: "/etc/passwd",
			want:  "etc/passwd",
		},
		{
			name:  "mixed traversal",
			input: "./../file.txt",
			want:  "file.txt",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeFilePath(tt.input); got != tt.want {
				t.Errorf("SanitizeFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidCommand(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "safe command",
			input: "devices",
			want:  true,
		},
		{
			name:  "rm rf",
			input: "rm -rf /",
			want:  false,
		},
		{
			name:  "dd command",
			input: "dd if=/dev/zero",
			want:  false,
		},
		{
			name:  "fork bomb",
			input: ":(){:|:&};:",
			want:  false,
		},
		{
			name:  "dev write",
			input: "echo test > /dev/sda",
			want:  false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidCommand(tt.input); got != tt.want {
				t.Errorf("IsValidCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizePackageName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normal package",
			input: "com.example.app",
			want:  "com.example.app",
		},
		{
			name:  "package with spaces",
			input: "com.example app",
			want:  "com.exampleapp",
		},
		{
			name:  "package with special chars",
			input: "com.example$app",
			want:  "com.exampleapp",
		},
		{
			name:  "package with underscore",
			input: "com.example_app",
			want:  "com.example_app",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizePackageName(tt.input); got != tt.want {
				t.Errorf("SanitizePackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
