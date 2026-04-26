package parser

import (
	"adbjson/internal/model"
	"strings"
)

// SettingsListParser parses adb shell settings list output
type SettingsListParser struct {
	*BaseParser
}

// NewSettingsListParser creates a new settings list parser
func NewSettingsListParser() *SettingsListParser {
	return &SettingsListParser{
		BaseParser: NewBaseParser("settings_list", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell settings list" command
func (p *SettingsListParser) Parse(output string, namespace string) (*model.SettingsListResponse, error) {
	lines := strings.Split(output, "\n")
	settings := []model.Setting{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Split by = to get key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		
		setting := model.Setting{
			Key:   strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		}
		
		settings = append(settings, setting)
	}
	
	response := &model.SettingsListResponse{
		Namespace: namespace,
		Settings:  settings,
		Count:     len(settings),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *SettingsListParser) Validate(result *model.SettingsListResponse) error {
	// Settings list response is always valid
	return nil
}
