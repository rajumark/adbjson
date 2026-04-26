package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// VmstatParser parses adb shell vmstat output
type VmstatParser struct {
	*BaseParser
}

// NewVmstatParser creates a new vmstat parser
func NewVmstatParser() *VmstatParser {
	return &VmstatParser{
		BaseParser: NewBaseParser("vmstat", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell vmstat" command
func (p *VmstatParser) Parse(output string) (*model.VmstatResponse, error) {
	output = strings.TrimSpace(output)
	response := &model.VmstatResponse{}
	
	// Handle case where output might be on multiple lines or single line
	lines := strings.Split(output, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Skip header line
		if strings.Contains(line, "procs") || strings.Contains(line, "memory") || 
		   strings.Contains(line, "swap") || strings.Contains(line, "io") || 
		   strings.Contains(line, "system") || strings.Contains(line, "cpu") {
			continue
		}
		
		// Parse data line
		// Format: r  b    swpd    free   buff   cache   si   so    bi    bo   in   cs us sy id wa
		parts := strings.Fields(line)
		if len(parts) >= 16 {
			response.Processes = model.VmstatProcesses{
				Running: parts[0],
				Blocked: parts[1],
			}
			
			response.Memory = model.VmstatMemory{
				SwapUsed: parts[2],
				Free:     parts[3],
				Buffers:  parts[4],
				Cache:    parts[5],
			}
			
			response.Swap = model.VmstatSwap{
				SwappedIn:  parts[6],
				SwappedOut: parts[7],
			}
			
			response.IO = model.VmstatIO{
				BlocksIn:  parts[8],
				BlocksOut: parts[9],
			}
			
			response.System = model.VmstatSystem{
				Interrupts:     parts[10],
				ContextSwitches: parts[11],
			}
			
			response.CPU = model.VmstatCPU{
				User:   parts[12],
				System: parts[13],
				Idle:   parts[14],
				Wait:   parts[15],
			}
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *VmstatParser) Validate(result *model.VmstatResponse) error {
	// Check if memory data was parsed
	if result.Memory.Free == "" {
		return fmt.Errorf("memory data not found in output")
	}
	
	return nil
}
