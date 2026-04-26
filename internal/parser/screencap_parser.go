package parser

import (
	"encoding/base64"
	"fmt"
	"strings"

	"adbjson/internal/model"
)

// ScreencapParser handles parsing of screencap command output
type ScreencapParser struct{}

// NewScreencapParser creates a new instance of ScreencapParser
func NewScreencapParser() *ScreencapParser {
	return &ScreencapParser{}
}

// Parse parses the raw output from "adb shell screencap" command
func (p *ScreencapParser) Parse(output string, filename string) (*model.ScreencapResponse, error) {
	// Determine format based on output characteristics
	format := "raw"
	if len(output) > 0 {
		// Check if output is PNG format by looking at PNG header
		if len(output) >= 8 && output[:8] == "\x89PNG\r\n\x1a\n" {
			format = "png"
		} else {
			// For binary data, we'll encode it as base64 for JSON compatibility
			format = "base64"
		}
	}
	
	// Encode binary data as base64 for JSON compatibility
	var data string
	if format == "base64" {
		data = base64.StdEncoding.EncodeToString([]byte(output))
	} else {
		// For PNG or other text formats, keep as is
		data = output
	}
	
	response := &model.ScreencapResponse{
		Data:     data,
		Format:   format,
		Size:     len(output),
		Filename: filename,
	}
	
	return response, nil
}

// Validate validates the parsed screencap response
func (p *ScreencapParser) Validate(response *model.ScreencapResponse) error {
	if response == nil {
		return fmt.Errorf("screencap response is nil")
	}
	
	if response.Size == 0 {
		return fmt.Errorf("screencap data is empty")
	}
	
	if response.Format == "" {
		return fmt.Errorf("screencap format is not specified")
	}
	
	// Validate base64 format if specified
	if response.Format == "base64" && response.Data != "" {
		_, err := base64.StdEncoding.DecodeString(response.Data)
		if err != nil {
			return fmt.Errorf("invalid base64 data: %v", err)
		}
	}
	
	// Validate PNG format if specified
	if response.Format == "png" && response.Data != "" {
		if !strings.HasPrefix(response.Data, "\x89PNG") {
			return fmt.Errorf("invalid PNG data: missing PNG header")
		}
	}
	
	return nil
}
