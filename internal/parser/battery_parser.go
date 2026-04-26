package parser

import (
	"adbjson/internal/model"
	"strconv"
	"strings"
)

// BatteryParser parses adb shell dumpsys battery output
type BatteryParser struct{}

// NewBatteryParser creates a new battery parser
func NewBatteryParser() *BatteryParser {
	return &BatteryParser{}
}

// Parse parses the raw output from "adb shell dumpsys battery" command
func (p *BatteryParser) Parse(output string) (*model.BatteryInfoResponse, error) {
	response := &model.BatteryInfoResponse{}
	
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Parse key-value pairs
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) < 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			
			switch key {
			case "AC powered":
				response.ACPowered = value == "true"
			case "USB powered":
				response.USBPowered = value == "true"
			case "Wireless powered":
				response.WirelessPowered = value == "true"
			case "Dock powered":
				response.DockPowered = value == "true"
			case "Max charging current":
				if val, err := strconv.Atoi(value); err == nil {
					response.MaxChargingCurrent = val
				}
			case "Charge counter":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargeCounter = val
				}
			case "status":
				if val, err := strconv.Atoi(value); err == nil {
					response.Status = val
				}
			case "health":
				if val, err := strconv.Atoi(value); err == nil {
					response.Health = val
				}
			case "present":
				response.Present = value == "true"
			case "level":
				if val, err := strconv.Atoi(value); err == nil {
					response.Level = val
				}
			case "scale":
				if val, err := strconv.Atoi(value); err == nil {
					response.Scale = val
				}
			case "voltage":
				if val, err := strconv.Atoi(value); err == nil {
					response.Voltage = val
				}
			case "temperature":
				if val, err := strconv.Atoi(value); err == nil {
					response.Temperature = val
				}
			case "technology":
				response.Technology = value
			case "Charging state":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargingState = val
				}
			case "Charging policy":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargingPolicy = val
				}
			case "Capacity level":
				if val, err := strconv.Atoi(value); err == nil {
					response.CapacityLevel = val
				}
			case "vbus state":
				response.VBusState = value == "true"
			case "charge watt":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargeWatt = val
				}
			case "charge watt design":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargeWattDesign = val
				}
			case "charge type":
				if val, err := strconv.Atoi(value); err == nil {
					response.ChargeType = val
				}
			case "cycle count":
				if val, err := strconv.Atoi(value); err == nil {
					response.CycleCount = val
				}
			case "Full capacity":
				if val, err := strconv.Atoi(value); err == nil {
					response.FullCapacity = val
				}
			case "Full design capacity":
				if val, err := strconv.Atoi(value); err == nil {
					response.FullDesignCapacity = val
				}
			}
		}
	}
	
	return response, nil
}
