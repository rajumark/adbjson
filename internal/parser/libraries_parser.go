package parser

import (
	"adbjson/internal/model"
	"strings"
)

// LibrariesParser parses adb shell pm list libraries output
type LibrariesParser struct {
	*BaseParser
}

// NewLibrariesParser creates a new libraries parser
func NewLibrariesParser() *LibrariesParser {
	return &LibrariesParser{
		BaseParser: NewBaseParser("libraries", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm list libraries" command
func (p *LibrariesParser) Parse(output string) (*model.LibrariesResponse, error) {
	lines := strings.Split(output, "\n")
	libraries := []model.Library{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse library line (format: library:android.hardware.camera)
		if strings.HasPrefix(line, "library:") {
			libraryName := strings.TrimPrefix(line, "library:")
			libraryName = strings.TrimSpace(libraryName)
			
			if libraryName != "" {
				libraries = append(libraries, model.Library{Name: libraryName})
			}
		}
	}
	
	return &model.LibrariesResponse{
		Libraries: libraries,
		Count:     len(libraries),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *LibrariesParser) Validate(result *model.LibrariesResponse) error {
	// Libraries response is always valid even with empty list
	return nil
}
