package parser

import (
	"adbjson/internal/model"
	"strings"
)

// StateParser parses adb get-state output
type StateParser struct{}

// NewStateParser creates a new state parser
func NewStateParser() *StateParser {
	return &StateParser{}
}

// Parse parses the raw output from "adb get-state" command
func (p *StateParser) Parse(output string) (*model.StateResponse, error) {
	state := strings.TrimSpace(output)
	
	return &model.StateResponse{
		State: state,
	}, nil
}
