package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// UptimeParser parses adb shell uptime output
type UptimeParser struct {
	*BaseParser
}

// NewUptimeParser creates a new uptime parser
func NewUptimeParser() *UptimeParser {
	return &UptimeParser{
		BaseParser: NewBaseParser("uptime", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell uptime" command
func (p *UptimeParser) Parse(output string) (*model.UptimeResponse, error) {
	output = strings.TrimSpace(output)
	response := &model.UptimeResponse{}
	
	// Parse uptime output format: "17:33:40 up 28 days,  9:57,  0 users,  load average: 3.67, 2.90, 2.88"
	parts := strings.Split(output, ",")
	
	if len(parts) >= 4 {
		// Extract current time and uptime (first part)
		timeAndUptime := strings.TrimSpace(parts[0])
		timeUptimeParts := strings.Fields(timeAndUptime)
		if len(timeUptimeParts) >= 4 {
			response.CurrentTime = timeUptimeParts[0]
			response.Uptime = strings.Join(timeUptimeParts[2:], " ")
		}
		
		// Extract users (second part)
		response.Users = strings.TrimSpace(parts[1])
		
		// Extract load average (last part)
		response.LoadAverage = strings.TrimSpace(parts[3])
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *UptimeParser) Validate(result *model.UptimeResponse) error {
	// Check if current time was parsed
	if result.CurrentTime == "" {
		return fmt.Errorf("current time not found in output")
	}
	
	// Check if uptime was parsed
	if result.Uptime == "" {
		return fmt.Errorf("uptime not found in output")
	}
	
	return nil
}
