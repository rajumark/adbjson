package parser

import (
	"adbjson/internal/model"
	"strings"
)

// WhoamiParser parses adb shell whoami output
type WhoamiParser struct {
	*BaseParser
}

// NewWhoamiParser creates a new whoami parser
func NewWhoamiParser() *WhoamiParser {
	return &WhoamiParser{
		BaseParser: NewBaseParser("whoami", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell whoami" command
func (p *WhoamiParser) Parse(output string) (*model.CurrentUserResponse, error) {
	username := strings.TrimSpace(output)
	
	return &model.CurrentUserResponse{
		CurrentUser: model.CurrentUser{Username: username},
	}, nil
}

// Validate checks if the parsed result is valid
func (p *WhoamiParser) Validate(result *model.CurrentUserResponse) error {
	// Current user response is always valid
	return nil
}
