package sanitize

import (
	"strings"
	"unicode"
)

// SanitizeString removes potentially dangerous characters from input strings
func SanitizeString(input string) string {
	// Remove null bytes and other control characters except tab, newline, carriage return
	var result strings.Builder
	for _, r := range input {
		if r == '\t' || r == '\n' || r == '\r' {
			result.WriteRune(r)
		} else if !unicode.IsControl(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SanitizeFilePath sanitizes file paths to prevent directory traversal
func SanitizeFilePath(path string) string {
	// Remove any path traversal attempts
	path = strings.ReplaceAll(path, "../", "")
	path = strings.ReplaceAll(path, "..\\", "")
	path = strings.ReplaceAll(path, "./", "")
	path = strings.ReplaceAll(path, ".\\", "")
	
	// Remove backslashes (Windows path separator)
	path = strings.ReplaceAll(path, "\\", "")
	
	// Remove leading slashes to prevent absolute paths
	path = strings.TrimLeft(path, "/")
	
	return path
}

// SanitizeCommandArgs sanitizes command arguments
func SanitizeCommandArgs(args []string) []string {
	sanitized := make([]string, len(args))
	for i, arg := range args {
		sanitized[i] = SanitizeString(arg)
	}
	return sanitized
}

// IsValidCommand checks if a command string is safe to execute
func IsValidCommand(command string) bool {
	// Block potentially dangerous commands
	dangerousCommands := []string{
		"rm -rf /",
		"dd if=",
		":(){:|:&};:",
		"> /dev/",
		"mkfs",
		"format",
	}
	
	lowerCommand := strings.ToLower(command)
	for _, dangerous := range dangerousCommands {
		if strings.Contains(lowerCommand, strings.ToLower(dangerous)) {
			return false
		}
	}
	
	return true
}

// SanitizePackageName sanitizes Android package names
func SanitizePackageName(pkg string) string {
	// Package names should only contain alphanumeric, dots, and underscores
	var result strings.Builder
	for _, r := range pkg {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '_' {
			result.WriteRune(r)
		}
	}
	return result.String()
}
