package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PermissionsParser parses adb shell pm list permissions output
type PermissionsParser struct {
	*BaseParser
}

// NewPermissionsParser creates a new permissions parser
func NewPermissionsParser() *PermissionsParser {
	return &PermissionsParser{
		BaseParser: NewBaseParser("permissions", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm list permissions" command
func (p *PermissionsParser) Parse(output string) (*model.PermissionsResponse, error) {
	lines := strings.Split(output, "\n")
	permissions := []model.Permission{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse permission line (format: permission:android.permission.INTERNET)
		if strings.HasPrefix(line, "permission:") {
			permissionName := strings.TrimPrefix(line, "permission:")
			permissionName = strings.TrimSpace(permissionName)
			
			if permissionName != "" {
				permissions = append(permissions, model.Permission{Name: permissionName})
			}
		}
	}
	
	return &model.PermissionsResponse{
		Permissions: permissions,
		Count:       len(permissions),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PermissionsParser) Validate(result *model.PermissionsResponse) error {
	// Permissions response is always valid even with empty list
	return nil
}
