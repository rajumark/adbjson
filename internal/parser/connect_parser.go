package parser

import (
	"adbjson/internal/model"
	"strings"
)

// ConnectParser parses adb connect output
type ConnectParser struct {
	*BaseParser
}

// NewConnectParser creates a new connect parser
func NewConnectParser() *ConnectParser {
	return &ConnectParser{
		BaseParser: NewBaseParser("connect", "1.0.0"),
	}
}

// Parse parses the raw output from "adb connect" command
func (p *ConnectParser) Parse(output string) (*model.ConnectResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.ConnectResponse{
		Connected: false,
		Message:   strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful connection
		if strings.Contains(line, "connected to") {
			response.Connected = true
			// Extract target from message like "connected to 127.0.0.1:5555"
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				response.Target = parts[2]
			}
			break
		}
		
		// Check for failed connection
		if strings.Contains(line, "failed to connect") || strings.Contains(line, "Connection refused") {
			response.Connected = false
			// Extract target from message like "failed to connect to '127.0.0.1:5555'"
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
func (p *ConnectParser) Validate(result *model.ConnectResponse) error {
	// Connect response is always valid
	return nil
}
