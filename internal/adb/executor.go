package adb

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"adbjson/internal/platform"
)

// Executor handles ADB command execution
type Executor struct {
	platformToolsPath string
}

// NewExecutor creates a new ADB executor with bundled ADB support
func NewExecutor() *Executor {
	// Get the directory where the executable is running
	execPath, _ := os.Executable()
	baseDir := filepath.Dir(execPath)
	
	return &Executor{
		platformToolsPath: filepath.Join(baseDir, "platform-tools"),
	}
}

// Execute runs an ADB command and returns the output
func (e *Executor) Execute(args ...string) (string, error) {
	// Try to use bundled ADB first
	adbPath, err := platform.GetADBPath(e.platformToolsPath)
	if err != nil {
		return "", fmt.Errorf("failed to get ADB path: %w", err)
	}
	
	// Check if bundled ADB exists
	if _, err := os.Stat(adbPath); err == nil {
		// Use bundled ADB
		cmd := exec.Command(adbPath, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("adb command failed: %w", err)
		}
		return string(output), nil
	}
	
	// Fall back to system ADB
	cmd := exec.Command("adb", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "not found") || strings.Contains(err.Error(), "executable file not found") {
			return "", fmt.Errorf("adb is not installed and bundled ADB not found at %s", adbPath)
		}
		return "", fmt.Errorf("adb command failed: %w", err)
	}
	
	return string(output), nil
}
