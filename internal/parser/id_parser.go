package parser

import (
	"adbjson/internal/model"
	"regexp"
	"strings"
)

// IdParser parses adb shell id output
type IdParser struct {
	*BaseParser
}

// NewIdParser creates a new id parser
func NewIdParser() *IdParser {
	return &IdParser{
		BaseParser: NewBaseParser("id", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell id" command
func (p *IdParser) Parse(output string) (*model.UserInfoResponse, error) {
	output = strings.TrimSpace(output)
	
	// Parse output format: uid=0(root) gid=0(root) groups=0(root)
	uidRegex := regexp.MustCompile(`uid=(\d+)\(([^)]+)\)`)
	gidRegex := regexp.MustCompile(`gid=(\d+)\(([^)]+)\)`)
	
	uidMatch := uidRegex.FindStringSubmatch(output)
	gidMatch := gidRegex.FindStringSubmatch(output)
	
	userID := "unknown"
	groupID := "unknown"
	
	if len(uidMatch) >= 3 {
		userID = uidMatch[1] + "(" + uidMatch[2] + ")"
	}
	if len(gidMatch) >= 3 {
		groupID = gidMatch[1] + "(" + gidMatch[2] + ")"
	}
	
	return &model.UserInfoResponse{
		UserInfo: model.UserInfo{
			UserID:  userID,
			GroupID: groupID,
		},
	}, nil
}

// Validate checks if the parsed result is valid
func (p *IdParser) Validate(result *model.UserInfoResponse) error {
	// User info response is always valid
	return nil
}
