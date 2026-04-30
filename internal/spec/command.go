package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParserType represents the type of parser to use
type ParserType string

const (
	ParserTypeTabular    ParserType = "tabular"
	ParserTypeKeyValue   ParserType = "key_value"
	ParserTypeRegex      ParserType = "regex"
	ParserTypeSingleValue ParserType = "single_value"
)

// CommandSpec defines a command specification
type CommandSpec struct {
	Name        string            `json:"name"`
	Command     []string          `json:"command"`
	Parser      ParserConfig      `json:"parser"`
	Output      OutputConfig      `json:"output"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// ParserConfig defines parser configuration
type ParserConfig struct {
	Type       ParserType `json:"type"`
	Delimiter  string     `json:"delimiter,omitempty"`
	SkipHeader bool       `json:"skip_header,omitempty"`
	Pattern    string     `json:"pattern,omitempty"`
	Matches    []string   `json:"matches,omitempty"`
	Key        string     `json:"key,omitempty"`
}

// OutputConfig defines output configuration
type OutputConfig struct {
	Format     string            `json:"format"`     // json, yaml
	Schema     string            `json:"schema"`     // output schema name
	Template   string            `json:"template"`   // custom template if needed
	Transforms map[string]string `json:"transforms,omitempty"` // field transformations
}

// SpecRegistry manages command specifications
type SpecRegistry struct {
	specs map[string]*CommandSpec
}

// NewSpecRegistry creates a new spec registry
func NewSpecRegistry() *SpecRegistry {
	return &SpecRegistry{
		specs: make(map[string]*CommandSpec),
	}
}

// LoadSpecs loads specifications from a directory
func (r *SpecRegistry) LoadSpecs(specsDir string) error {
	return filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !strings.HasSuffix(path, ".json") {
			return nil
		}
		
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		var spec CommandSpec
		if err := json.Unmarshal(data, &spec); err != nil {
			return fmt.Errorf("failed to parse spec %s: %w", path, err)
		}
		
		r.specs[spec.Name] = &spec
		return nil
	})
}

// GetSpec gets a command specification by name
func (r *SpecRegistry) GetSpec(name string) (*CommandSpec, error) {
	spec, exists := r.specs[name]
	if !exists {
		return nil, fmt.Errorf("command spec not found: %s", name)
	}
	return spec, nil
}

// ListSpecs returns all available command names
func (r *SpecRegistry) ListSpecs() []string {
	names := make([]string, 0, len(r.specs))
	for name := range r.specs {
		names = append(names, name)
	}
	return names
}

// RegisterSpec registers a command specification
func (r *SpecRegistry) RegisterSpec(spec *CommandSpec) {
	r.specs[spec.Name] = spec
}
