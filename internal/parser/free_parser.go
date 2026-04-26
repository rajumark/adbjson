package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strconv"
	"strings"
)

// FreeParser parses adb shell free output
type FreeParser struct {
	*BaseParser
}

// NewFreeParser creates a new free parser
func NewFreeParser() *FreeParser {
	return &FreeParser{
		BaseParser: NewBaseParser("free", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell free" command
func (p *FreeParser) Parse(output string) (*model.FreeResponse, error) {
	lines := strings.Split(output, "\n")
	response := &model.FreeResponse{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Parse memory line
		if strings.HasPrefix(line, "Mem:") {
			// Handle case where values might be on the same line with multiple spaces
			parts := strings.Fields(line)
			if len(parts) >= 7 {
				response.Memory = model.MemoryInfo{
					Total:   parts[1],
					Used:    parts[2],
					Free:    parts[3],
					Shared:  parts[4],
					Buffers: parts[5],
					Cached:  parts[6],
				}
			} else if len(parts) == 6 {
				// Sometimes cached might be missing
				response.Memory = model.MemoryInfo{
					Total:   parts[1],
					Used:    parts[2],
					Free:    parts[3],
					Shared:  parts[4],
					Buffers: parts[5],
					Cached:  "0",
				}
			}
		}
		
		// Parse buffers/cache line
		if strings.HasPrefix(line, "-/+ buffers/cache:") {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				response.Buffers = model.MemoryInfo{
					Used:  parts[2],
					Free:  parts[3],
				}
			}
		}
		
		// Parse swap line
		if strings.HasPrefix(line, "Swap:") {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				response.Swap = model.MemoryInfo{
					Total: parts[1],
					Used:  parts[2],
					Free:  parts[3],
				}
			}
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *FreeParser) Validate(result *model.FreeResponse) error {
	// Check if memory info was parsed
	if result.Memory.Total == "" {
		return fmt.Errorf("memory information not found in output")
	}
	
	// Validate that total values are numeric
	if _, err := strconv.ParseInt(result.Memory.Total, 10, 64); err != nil {
		return fmt.Errorf("invalid memory total value: %s", result.Memory.Total)
	}
	
	return nil
}
