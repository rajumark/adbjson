package main

import (
	"adbjson/cmd"
	"adbjson/internal/logger"
)

func main() {
	// Initialize logger with default INFO level
	// Debug mode will be set by cobra flag in cmd package
	logger.Init(logger.INFO)
	
	cmd.Execute()
}
