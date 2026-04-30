package ir

import (
	"adbjson/internal/adb"
	"adbjson/internal/spec"
	"fmt"
)

// Engine is the new IR-based command execution engine
type Engine struct {
	adbExecutor *adb.Executor
	specRegistry *spec.SpecRegistry
	adapter     *Adapter
}

// NewEngine creates a new IR-based engine
func NewEngine() (*Engine, error) {
	// Initialize components
	adbExecutor := adb.NewExecutor()
	specRegistry := spec.NewSpecRegistry()
	adapter := NewAdapter()
	
	// Load command specs
	err := specRegistry.LoadSpecs("info_only/supported_commands/specs")
	if err != nil {
		return nil, fmt.Errorf("failed to load command specs: %w", err)
	}
	
	return &Engine{
		adbExecutor:  adbExecutor,
		specRegistry: specRegistry,
		adapter:      adapter,
	}, nil
}

// ExecuteCommand executes a command using the IR-based approach
func (e *Engine) ExecuteCommand(commandName string, args []string) ([]byte, error) {
	// Get command specification
	spec, err := e.specRegistry.GetSpec(commandName)
	if err != nil {
		return nil, err
	}
	
	// Build ADB command
	adbArgs := append(spec.Command, args...)
	
	// Execute ADB command
	rawOutput, err := e.adbExecutor.Execute(adbArgs...)
	if err != nil {
		return nil, fmt.Errorf("ADB execution failed: %w", err)
	}
	
	// Parse raw output into IR
	irDoc, err := e.parseToIR(spec, rawOutput)
	if err != nil {
		return nil, fmt.Errorf("IR parsing failed: %w", err)
	}
	
	// Convert IR to JSON
	jsonOutput, err := e.adapter.ToJSON(irDoc, spec.Output.Schema)
	if err != nil {
		return nil, fmt.Errorf("JSON conversion failed: %w", err)
	}
	
	return jsonOutput, nil
}

// parseToIR parses raw output into IR using the specified parser strategy
func (e *Engine) parseToIR(cmdSpec *spec.CommandSpec, rawOutput string) (*Document, error) {
	var strategy ParserStrategy
	
	switch cmdSpec.Parser.Type {
	case spec.ParserTypeTabular:
		strategy = NewTabularParser(
			cmdSpec.Name,
			cmdSpec.Parser.SkipHeader,
			cmdSpec.Parser.Delimiter,
		)
	case spec.ParserTypeKeyValue:
		strategy = NewKeyValueParser(
			cmdSpec.Name,
			cmdSpec.Parser.Delimiter,
		)
	case spec.ParserTypeRegex:
		strategy = NewRegexParser(
			cmdSpec.Name,
			cmdSpec.Parser.Pattern,
			cmdSpec.Parser.Matches,
		)
	case spec.ParserTypeSingleValue:
		strategy = NewSingleValueParser(
			cmdSpec.Name,
			cmdSpec.Parser.Key,
		)
	default:
		return nil, fmt.Errorf("unsupported parser type: %s", cmdSpec.Parser.Type)
	}
	
	return strategy.Parse(rawOutput)
}

// ListCommands returns all available command names
func (e *Engine) ListCommands() []string {
	return e.specRegistry.ListSpecs()
}
