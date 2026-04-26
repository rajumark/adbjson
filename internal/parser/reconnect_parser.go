package parser

import (
	"adbjson/internal/model"
	"strings"
)

// ReconnectParser parses adb reconnect output
type ReconnectParser struct{}

// NewReconnectParser creates a new reconnect parser
func NewReconnectParser() *ReconnectParser {
	return &ReconnectParser{}
}

// Parse parses the raw output from "adb reconnect" command
func (p *ReconnectParser) Parse(output string) (*model.ServerResponse, error) {
	output = strings.TrimSpace(output)
	
	success := strings.Contains(output, "reconnecting")
	message := output
	
	if output == "" {
		success = false
		message = "No device to reconnect"
	}
	
	// Handle no devices case
	if strings.Contains(output, "no devices") {
		success = false
	}
	
	return &model.ServerResponse{
		Success: success,
		Message: message,
	}, nil
}
