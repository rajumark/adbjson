package parser

import (
	"adbjson/internal/model"
	"strings"
)

// DisconnectParser parses adb disconnect output
type DisconnectParser struct {
	*BaseParser
}

// NewDisconnectParser creates a new disconnect parser
func NewDisconnectParser() *DisconnectParser {
	return &DisconnectParser{
		BaseParser: NewBaseParser("disconnect", "1.0.0"),
	}
}

// Parse parses the raw output from "adb disconnect" command
func (p *DisconnectParser) Parse(output string) (*model.DisconnectResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.DisconnectResponse{
		Disconnected: false,
		Message:      strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful disconnection
		if strings.Contains(line, "disconnected everything") {
			response.Disconnected = true
			response.Target = "all"
			break
		}
		
		// Check for specific device disconnection
		if strings.Contains(line, "disconnected") {
			response.Disconnected = true
			// Extract target from message like "disconnected 192.168.1.100:5555"
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				response.Target = parts[1]
			}
			break
		}
		
		// Check for error (device not found)
		if strings.Contains(line, "error: no such device") {
			response.Disconnected = false
			// Extract target from message like "error: no such device '192.168.1.100:5555'"
			if strings.Contains(line, "'") {
				start := strings.Index(line, "'")
				end := strings.LastIndex(line, "'")
				if start != -1 && end != -1 && end > start {
					response.Target = line[start+1 : end]
				}
			}
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DisconnectParser) Validate(result *model.DisconnectResponse) error {
	// Disconnect response is always valid
	return nil
}
