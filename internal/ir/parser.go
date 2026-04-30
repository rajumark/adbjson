package ir

import (
	"regexp"
	"strings"
)

// ParserStrategy defines the interface for parsing strategies
type ParserStrategy interface {
	Parse(rawOutput string) (*Document, error)
	Name() string
}

// TabularParser parses tabular data (like adb devices output)
type TabularParser struct {
	command    string
	skipHeader bool
	delimiter  string
}

// NewTabularParser creates a new tabular parser
func NewTabularParser(command string, skipHeader bool, delimiter string) *TabularParser {
	return &TabularParser{
		command:    command,
		skipHeader: skipHeader,
		delimiter:  delimiter,
	}
}

// Name returns the parser name
func (p *TabularParser) Name() string {
	return "tabular"
}

// Parse parses tabular output into IR
func (p *TabularParser) Parse(rawOutput string) (*Document, error) {
	doc := NewDocument(p.command, rawOutput)
	
	// Create root array for rows
	root := NewArrayNode("devices")
	doc.Root = root
	
	lines := strings.Split(rawOutput, "\n")
	
	for i, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Skip header if configured
		if p.skipHeader && i == 0 {
			continue
		}
		
		// Parse row
		rowNode := p.parseRow(line)
		if rowNode != nil {
			root.AddChild(rowNode)
		}
	}
	
	return doc, nil
}

// parseRow parses a single row into an object node
func (p *TabularParser) parseRow(line string) *Node {
	row := NewObjectNode("device")
	
	// Handle devices -l format specially
	if strings.Contains(line, "product:") || strings.Contains(line, "transport_id:") {
		return p.parseDevicesLLine(line)
	}
	
	// Split by delimiter
	parts := strings.Split(line, p.delimiter)
	
	// Handle different column counts
	switch len(parts) {
	case 2: // Standard adb devices: ID, Status
		row.AddChild(NewStringNode("id", strings.TrimSpace(parts[0])))
		row.AddChild(NewStringNode("status", strings.TrimSpace(parts[1])))
		
	case 4: // Simple 4-column format
		row.AddChild(NewStringNode("id", strings.TrimSpace(parts[0])))
		row.AddChild(NewStringNode("status", strings.TrimSpace(parts[1])))
		row.AddChild(NewStringNode("usb", strings.TrimSpace(parts[2])))
		row.AddChild(NewStringNode("product", strings.TrimSpace(parts[3])))
		
	default:
		// Fallback: treat as single value
		row.AddChild(NewStringNode("value", line))
	}
	
	return row
}

// parseDevicesLLine parses the complex adb devices -l format
func (p *TabularParser) parseDevicesLLine(line string) *Node {
	row := NewObjectNode("device")
	
	// Split by whitespace and parse key-value pairs
	fields := strings.Fields(line)
	if len(fields) < 2 {
		row.AddChild(NewStringNode("value", line))
		return row
	}
	
	// First two fields are always ID and status
	row.AddChild(NewStringNode("id", fields[0]))
	row.AddChild(NewStringNode("status", fields[1]))
	
	// Parse remaining fields as key:value pairs
	for i := 2; i < len(fields); i++ {
		field := fields[i]
		if strings.Contains(field, ":") {
			parts := strings.SplitN(field, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				row.AddChild(NewStringNode(key, value))
			}
		}
	}
	
	return row
}

// KeyValueParser parses key-value pair data (like getprop output)
type KeyValueParser struct {
	command   string
	delimiter string
}

// NewKeyValueParser creates a new key-value parser
func NewKeyValueParser(command, delimiter string) *KeyValueParser {
	return &KeyValueParser{
		command:   command,
		delimiter: delimiter,
	}
}

// Name returns the parser name
func (p *KeyValueParser) Name() string {
	return "key_value"
}

// Parse parses key-value output into IR
func (p *KeyValueParser) Parse(rawOutput string) (*Document, error) {
	doc := NewDocument(p.command, rawOutput)
	
	// Create root array for properties
	root := NewArrayNode("properties")
	doc.Root = root
	
	lines := strings.Split(rawOutput, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse key-value pair
		kvNode := p.parseKeyValue(line)
		if kvNode != nil {
			root.AddChild(kvNode)
		}
	}
	
	return doc, nil
}

// parseKeyValue parses a single key-value line
func (p *KeyValueParser) parseKeyValue(line string) *Node {
	// Split by delimiter
	parts := strings.SplitN(line, p.delimiter, 2)
	
	if len(parts) != 2 {
		return nil
	}
	
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	
	// Clean key (remove brackets if present)
	key = strings.Trim(key, "[]")
	
	if key == "" {
		return nil
	}
	
	property := NewObjectNode("property")
	property.AddChild(NewStringNode("key", key))
	property.AddChild(NewStringNode("value", value))
	
	return property
}

// RegexParser parses data using regular expressions
type RegexParser struct {
	command string
	pattern string
	matches []string // names for capture groups
}

// NewRegexParser creates a new regex parser
func NewRegexParser(command, pattern string, matches []string) *RegexParser {
	return &RegexParser{
		command: command,
		pattern: pattern,
		matches: matches,
	}
}

// Name returns the parser name
func (p *RegexParser) Name() string {
	return "regex"
}

// Parse parses output using regex into IR
func (p *RegexParser) Parse(rawOutput string) (*Document, error) {
	doc := NewDocument(p.command, rawOutput)
	
	// Compile regex
	re, err := regexp.Compile(p.pattern)
	if err != nil {
		return nil, err
	}
	
	// Create root array for matches
	root := NewArrayNode("matches")
	doc.Root = root
	
	lines := strings.Split(rawOutput, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Find matches
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			matchNode := p.createMatchNode(matches[1:])
			root.AddChild(matchNode)
		}
	}
	
	return doc, nil
}

// createMatchNode creates a node from regex matches
func (p *RegexParser) createMatchNode(values []string) *Node {
	match := NewObjectNode("match")
	
	for i, value := range values {
		if i < len(p.matches) {
			match.AddChild(NewStringNode(p.matches[i], value))
		}
	}
	
	return match
}

// SingleValueParser parses single value output
type SingleValueParser struct {
	command string
	key     string
}

// NewSingleValueParser creates a new single value parser
func NewSingleValueParser(command, key string) *SingleValueParser {
	return &SingleValueParser{
		command: command,
		key:     key,
	}
}

// Name returns the parser name
func (p *SingleValueParser) Name() string {
	return "single_value"
}

// Parse parses single value output into IR
func (p *SingleValueParser) Parse(rawOutput string) (*Document, error) {
	doc := NewDocument(p.command, rawOutput)
	
	// Create root object with single value
	root := NewObjectNode("result")
	root.AddChild(NewStringNode(p.key, strings.TrimSpace(rawOutput)))
	doc.Root = root
	
	return doc, nil
}
