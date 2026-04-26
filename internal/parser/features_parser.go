package parser

import (
	"adbjson/internal/model"
	"strings"
)

// FeaturesParser parses adb shell pm list features output
type FeaturesParser struct {
	*BaseParser
}

// NewFeaturesParser creates a new features parser
func NewFeaturesParser() *FeaturesParser {
	return &FeaturesParser{
		BaseParser: NewBaseParser("features", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm list features" command
func (p *FeaturesParser) Parse(output string) (*model.FeaturesResponse, error) {
	lines := strings.Split(output, "\n")
	features := []model.Feature{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse feature line (format: feature:android.hardware.camera)
		if strings.HasPrefix(line, "feature:") {
			featureName := strings.TrimPrefix(line, "feature:")
			featureName = strings.TrimSpace(featureName)
			
			if featureName != "" {
				features = append(features, model.Feature{Name: featureName})
			}
		}
	}
	
	return &model.FeaturesResponse{
		Features: features,
		Count:    len(features),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *FeaturesParser) Validate(result *model.FeaturesResponse) error {
	// Features response is always valid even with empty list
	return nil
}
