package parser

import (
	"fmt"

	"adbjson/internal/model"
)

// InputTapParser handles parsing of input tap command output
type InputTapParser struct{}

// NewInputTapParser creates a new instance of InputTapParser
func NewInputTapParser() *InputTapParser {
	return &InputTapParser{}
}

// Parse parses the raw output from "adb shell input tap" command
func (p *InputTapParser) Parse(output string, x, y, source string, displayID int) (*model.InputTapResponse, error) {
	// Input tap commands typically produce no output on success
	// We'll create a response based on the input parameters
	success := true
	message := "Tap command executed successfully"
	
	// If there's any output, it might indicate an error
	if output != "" {
		success = false
		message = output
	}
	
	// Default source if not specified
	if source == "" {
		source = "touchscreen"
	}
	
	response := &model.InputTapResponse{
		X:         x,
		Y:         y,
		Source:    source,
		DisplayID: displayID,
		Success:   success,
		Message:   message,
	}
	
	return response, nil
}

// Validate validates the parsed input tap response
func (p *InputTapParser) Validate(response *model.InputTapResponse) error {
	if response == nil {
		return fmt.Errorf("input tap response is nil")
	}
	
	if response.X == "" {
		return fmt.Errorf("x coordinate is required")
	}
	
	if response.Y == "" {
		return fmt.Errorf("y coordinate is required")
	}
	
	if response.Source == "" {
		return fmt.Errorf("input source is required")
	}
	
	// Validate display ID (should be 0 or positive)
	if response.DisplayID < 0 {
		return fmt.Errorf("display ID must be 0 or positive")
	}
	
	return nil
}
