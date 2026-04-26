package parser

import (
	"fmt"
	"regexp"
	"strings"

	"adbjson/internal/model"
)

// ContentQueryParser handles parsing of content query command output
type ContentQueryParser struct{}

// NewContentQueryParser creates a new instance of ContentQueryParser
func NewContentQueryParser() *ContentQueryParser {
	return &ContentQueryParser{}
}

// Parse parses the raw output from "adb shell content query" command
func (p *ContentQueryParser) Parse(output string, uri string) (*model.ContentQueryResponse, error) {
	
	// Parse rows from output
	rows := []model.ContentQueryRow{}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Parse row format: "Row: <number> _id=<id>, name=<name>, value=<value>, is_preserved_in_restore=<value>"
		if strings.HasPrefix(line, "Row:") {
			row := p.parseRow(line)
			if row != nil {
				rows = append(rows, *row)
			}
		}
	}
	
	response := &model.ContentQueryResponse{
		URI:    uri,
		Count:  len(rows),
		Rows:   rows,
		Fields: []string{"_id", "name", "value", "is_preserved_in_restore"},
	}
	
	return response, nil
}

// parseRow parses a single row line
func (p *ContentQueryParser) parseRow(line string) *model.ContentQueryRow {
	// Remove "Row: <number>" prefix
	re := regexp.MustCompile(`^Row:\s*\d+\s*`)
	line = re.ReplaceAllString(line, "")
	
	// Split by comma and parse key-value pairs
	parts := strings.Split(line, ",")
	row := &model.ContentQueryRow{}
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		
		// Split by "="
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			continue
		}
		
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		
		switch key {
		case "_id":
			row.ID = value
		case "name":
			row.Name = value
		case "value":
			row.Value = value
		case "is_preserved_in_restore":
			row.IsPreservedInRestore = value
		}
	}
	
	// Only return row if we have at least name and value
	if row.Name != "" {
		return row
	}
	
	return nil
}

// Validate validates the parsed content query response
func (p *ContentQueryParser) Validate(response *model.ContentQueryResponse) error {
	if response == nil {
		return fmt.Errorf("content query response is nil")
	}
	
	if response.URI == "" {
		return fmt.Errorf("URI is required")
	}
	
	if response.Count < 0 {
		return fmt.Errorf("count cannot be negative")
	}
	
	if len(response.Rows) != response.Count {
		return fmt.Errorf("row count mismatch: expected %d, got %d", response.Count, len(response.Rows))
	}
	
	// Validate each row
	for i, row := range response.Rows {
		if row.Name == "" {
			return fmt.Errorf("row %d: name is required", i)
		}
	}
	
	return nil
}
