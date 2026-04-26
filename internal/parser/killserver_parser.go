package parser

import (
	"adbjson/internal/model"
)

// KillServerParser parses adb kill-server output
type KillServerParser struct{}

// NewKillServerParser creates a new kill-server parser
func NewKillServerParser() *KillServerParser {
	return &KillServerParser{}
}

// Parse parses the raw output from "adb kill-server" command
func (p *KillServerParser) Parse(output string) (*model.ServerResponse, error) {
	// kill-server typically has no output on success
	// We consider it successful if there's no error
	return &model.ServerResponse{
		Success: true,
		Message: "ADB server killed successfully",
	}, nil
}
