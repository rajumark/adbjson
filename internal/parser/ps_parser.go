package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PsParser parses adb shell ps output
type PsParser struct {
	*BaseParser
}

// NewPsParser creates a new ps parser
func NewPsParser() *PsParser {
	return &PsParser{
		BaseParser: NewBaseParser("ps", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell ps" command
func (p *PsParser) Parse(output string) (*model.PsResponse, error) {
	lines := strings.Split(output, "\n")
	processes := []model.Process{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Skip header lines
		if strings.Contains(line, "USER") && strings.Contains(line, "PID") {
			continue
		}
		
		// Parse process line
		// ps format: USER PID PPID NAME
		// ps -A format: USER PID PPID VSIZE RSS WCHAN PC ADDR NAME
		parts := strings.Fields(line)
		if len(parts) >= 4 {
			var name string
			if len(parts) > 8 {
				// ps -A format - name starts from 9th column
				name = strings.Join(parts[8:], " ")
			} else {
				// ps format - name starts from 4th column
				name = strings.Join(parts[3:], " ")
			}
			
			process := model.Process{
				User: parts[0],
				PID:  parts[1],
				PPID: parts[2],
				Name: name,
			}
			processes = append(processes, process)
		}
	}
	
	response := &model.PsResponse{
		Processes: processes,
		Count:     len(processes),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *PsParser) Validate(result *model.PsResponse) error {
	// Ps response is always valid
	return nil
}
