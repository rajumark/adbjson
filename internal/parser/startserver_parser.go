package parser

import (
	"adbjson/internal/model"
	"strings"
)

// StartServerParser parses adb start-server output
type StartServerParser struct{}

// NewStartServerParser creates a new start-server parser
func NewStartServerParser() *StartServerParser {
	return &StartServerParser{}
}

// Parse parses the raw output from "adb start-server" command
func (p *StartServerParser) Parse(output string) (*model.ServerResponse, error) {
	output = strings.TrimSpace(output)
	
	// Empty output means server is already running
	if output == "" {
		return &model.ServerResponse{
			Success: true,
			Message: "ADB server is already running",
		}, nil
	}
	
	// Check if daemon started successfully
	success := strings.Contains(output, "daemon started successfully")
	message := "ADB server started successfully"
	
	if !success {
		// If server was already running
		if strings.Contains(output, "daemon is already running") {
			success = true
			message = "ADB server is already running"
		} else {
			message = output
		}
	}
	
	return &model.ServerResponse{
		Success: success,
		Message: message,
	}, nil
}
