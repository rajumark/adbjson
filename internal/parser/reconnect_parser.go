package parser

import (
	"adbjson/internal/model"
	"strings"
)

// ReconnectParser parses adb reconnect output
type ReconnectParser struct {
	*BaseParser
}

// NewReconnectParser creates a new reconnect parser
func NewReconnectParser() *ReconnectParser {
	return &ReconnectParser{
		BaseParser: NewBaseParser("reconnect", "1.0.0"),
	}
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

// Validate checks if the parsed result is valid
func (p *ReconnectParser) Validate(result *model.ServerResponse) error {
	// Reconnect response is always valid
	return nil
}
