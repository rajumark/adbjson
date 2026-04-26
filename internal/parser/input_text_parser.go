package parser

import (
	"fmt"

	"adbjson/internal/model"
)

// InputTextParser handles parsing of input text command output
type InputTextParser struct{}

// NewInputTextParser creates a new instance of InputTextParser
func NewInputTextParser() *InputTextParser {
	return &InputTextParser{}
}

// Parse parses the raw output from "adb shell input text" command
func (p *InputTextParser) Parse(output string, text, source string, displayID int) (*model.InputTextResponse, error) {
	// Input text commands typically produce no output on success
	// We'll create a response based on the input parameters
	success := true
	message := "Text input command executed successfully"
	
	// If there's any output, it might indicate an error
	if output != "" {
		success = false
		message = output
	}
	
	// Default source if not specified
	if source == "" {
		source = "keyboard"
	}
	
	response := &model.InputTextResponse{
		Text:      text,
		Source:    source,
		DisplayID: displayID,
		Success:   success,
		Message:   message,
	}
	
	return response, nil
}

// Validate validates the parsed input text response
func (p *InputTextParser) Validate(response *model.InputTextResponse) error {
	if response == nil {
		return fmt.Errorf("input text response is nil")
	}
	
	if response.Text == "" {
		return fmt.Errorf("text is required")
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
