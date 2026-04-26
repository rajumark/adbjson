package platform

import (
	"fmt"
	"runtime"
)

// GetADBPath returns the path to the bundled ADB binary for the current platform
func GetADBPath(basePath string) (string, error) {
	var platformDir string
	
	switch runtime.GOOS {
	case "darwin":
		platformDir = "platform-tools-darwin"
	case "linux":
		platformDir = "platform-tools-linux"
	case "windows":
		platformDir = "platform-tools-windows"
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	
	var adbBinary string
	if runtime.GOOS == "windows" {
		adbBinary = "adb.exe"
	} else {
		adbBinary = "adb"
	}
	
	return fmt.Sprintf("%s/%s/%s", basePath, platformDir, adbBinary), nil
}
