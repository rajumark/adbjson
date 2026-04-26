package parser

import (
	"adbjson/internal/model"
	"strings"
)

// SerialNoParser parses adb get-serialno output
type SerialNoParser struct{}

// NewSerialNoParser creates a new serial number parser
func NewSerialNoParser() *SerialNoParser {
	return &SerialNoParser{}
}

// Parse parses the raw output from "adb get-serialno" command
func (p *SerialNoParser) Parse(output string) (*model.SerialNoResponse, error) {
	serialNo := strings.TrimSpace(output)
	
	return &model.SerialNoResponse{
		SerialNo: serialNo,
	}, nil
}
