package parser

import (
	"adbjson/internal/model"
	"strings"
)

// TopParser parses adb shell top output
type TopParser struct {
	*BaseParser
}

// NewTopParser creates a new top parser
func NewTopParser() *TopParser {
	return &TopParser{
		BaseParser: NewBaseParser("top", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell top" command
func (p *TopParser) Parse(output string) (*model.TopResponse, error) {
	lines := strings.Split(output, "\n")
	processes := []model.TopProcess{}
	summary := model.TopSummary{}
	
	// Parse summary information
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Tasks:") {
			summary.Tasks = line
		} else if strings.HasPrefix(line, "Mem:") {
			summary.Mem = line
		} else if strings.HasPrefix(line, "Swap:") {
			summary.Swap = line
		} else if strings.Contains(line, "%cpu") {
			summary.Cpu = line
		}
	}
	
	// Parse process information
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Skip header lines and summary lines
		if strings.Contains(line, "PID") || 
		   strings.Contains(line, "USER") || 
		   strings.Contains(line, "Tasks:") ||
		   strings.Contains(line, "Mem:") ||
		   strings.Contains(line, "Swap:") ||
		   strings.Contains(line, "%cpu") {
			continue
		}
		
		// Parse process line
		// Format: PID USER PR NI VIRT RES SHR S[%CPU] %MEM TIME+ ARGS
		parts := strings.Fields(line)
		if len(parts) >= 12 {
			// Extract the process name/args (everything after TIME+)
			args := strings.Join(parts[11:], " ")
			
			process := model.TopProcess{
				User: parts[1],
				PID:  parts[0],
				PR:   parts[2],
				NI:   parts[3],
				VIRT: parts[4],
				RES:  parts[5],
				SHR:  parts[6],
				S:    parts[7],
				CPU:  parts[8],
				MEM:  parts[9],
				Time: parts[10],
				Args: args,
			}
			processes = append(processes, process)
		}
	}
	
	response := &model.TopResponse{
		Processes: processes,
		Count:     len(processes),
		Summary:   summary,
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *TopParser) Validate(result *model.TopResponse) error {
	// Top response is always valid
	return nil
}
