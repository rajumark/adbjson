package parser

import (
	"adbjson/internal/model"
	"strings"
)

// UnameParser parses adb shell uname -a output
type UnameParser struct {
	*BaseParser
}

// NewUnameParser creates a new uname parser
func NewUnameParser() *UnameParser {
	return &UnameParser{
		BaseParser: NewBaseParser("uname", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell uname -a" command
func (p *UnameParser) Parse(output string) (*model.SystemInfoResponse, error) {
	output = strings.TrimSpace(output)
	
	// Parse output format: Linux localhost 4.14.0-android-g5f8c9a9 #1 SMP PREEMPT aarch64
	parts := strings.Fields(output)
	
	systemInfo := model.SystemInfo{
		KernelName:    "unknown",
		NodeName:      "unknown",
		KernelRelease: "unknown",
		KernelVersion: "unknown",
		Machine:       "unknown",
		Processor:     "unknown",
		Hardware:      "unknown",
		OS:            "unknown",
	}
	
	if len(parts) >= 1 {
		systemInfo.KernelName = parts[0]
	}
	if len(parts) >= 2 {
		systemInfo.NodeName = parts[1]
	}
	if len(parts) >= 3 {
		systemInfo.KernelRelease = parts[2]
	}
	if len(parts) >= 4 {
		systemInfo.KernelVersion = parts[3]
	}
	if len(parts) >= 5 {
		systemInfo.Machine = parts[4]
	}
	if len(parts) >= 6 {
		systemInfo.Processor = parts[5]
	}
	if len(parts) >= 7 {
		systemInfo.Hardware = parts[6]
	}
	if len(parts) >= 8 {
		systemInfo.OS = parts[7]
	}
	
	return &model.SystemInfoResponse{
		SystemInfo: systemInfo,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *UnameParser) Validate(result *model.SystemInfoResponse) error {
	// System info response is always valid
	return nil
}
