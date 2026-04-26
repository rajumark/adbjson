package parser

import (
	"adbjson/internal/model"
	"strings"
)

// DevicesLParser parses adb devices -l output
type DevicesLParser struct{}

// NewDevicesLParser creates a new devices-l parser
func NewDevicesLParser() *DevicesLParser {
	return &DevicesLParser{}
}

// Parse parses the raw output from "adb devices -l" command
func (p *DevicesLParser) Parse(output string) (*model.DevicesResponse, error) {
	lines := strings.Split(output, "\n")
	devices := []model.Device{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Skip header line
		if strings.HasPrefix(line, "List of devices attached") {
			continue
		}
		
		// Split by whitespace
		parts := strings.Fields(line)
		
		// Handle malformed lines safely
		if len(parts) < 2 {
			continue
		}
		
		device := model.Device{
			ID:     parts[0],
			Status: parts[1],
		}
		
		// Parse key:value pairs
		for i := 2; i < len(parts); i++ {
			part := parts[i]
			if strings.Contains(part, ":") {
				kv := strings.SplitN(part, ":", 2)
				if len(kv) == 2 {
					key := strings.ToLower(kv[0])
					value := kv[1]
					
					switch key {
					case "usb":
						device.USB = value
					case "product":
						device.Product = value
					case "model":
						device.Model = value
					case "device":
						device.Device = value
					case "transport_id":
						device.TransportID = value
					}
				}
			}
		}
		
		devices = append(devices, device)
	}
	
	return &model.DevicesResponse{
		Devices: devices,
	}, nil
}
