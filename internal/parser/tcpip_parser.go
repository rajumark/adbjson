package parser

import (
	"adbjson/internal/model"
	"strings"
)

// TcpipParser parses adb tcpip output
type TcpipParser struct {
	*BaseParser
}

// NewTcpipParser creates a new tcpip parser
func NewTcpipParser() *TcpipParser {
	return &TcpipParser{
		BaseParser: NewBaseParser("tcpip", "1.0.0"),
	}
}

// Parse parses the raw output from "adb tcpip" command
func (p *TcpipParser) Parse(output string, port string) (*model.TcpipResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.TcpipResponse{
		Success: false,
		Port:    port,
		Message: strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful tcpip
		if strings.Contains(line, "restarting in TCP mode port") {
			response.Success = true
			break
		}
		
		// Check for error messages
		if strings.Contains(line, "error") || strings.Contains(line, "failed") {
			response.Success = false
			break
		}
		
		// Check for no device
		if strings.Contains(line, "no device") || strings.Contains(line, "device not found") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *TcpipParser) Validate(result *model.TcpipResponse) error {
	// Tcpip response is always valid
	return nil
}
