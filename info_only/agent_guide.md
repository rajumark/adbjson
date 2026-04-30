# Agent Guide for Implementing New Features

When implementing a new ADB command feature or updating an existing one, follow this systematic flow:

## Implementation Flow

### 1. Run Pure ADB Command First
- Execute the raw ADB command to understand its behavior
- Ensure it runs successfully without errors
- Test with different device states/scenarios if applicable

### 2. Analyze Output
- Examine the raw output format and structure
- Ensure the output contains meaningful data
- Consider edge cases (empty output, errors, multiple devices)

### 3. Design for Scalability
- Think about data variability (can data be more/less than expected?)
- Determine the best JSON structure for the output
- Consider future extensibility and backward compatibility
- Plan for arrays vs single values, nested structures, etc.

### 4. Create JSON Plan
- Design the JSON schema that best represents the data
- Include metadata (command executed, timestamp, device info if relevant)
- Plan for error handling in the JSON structure
- Consider both success and error response formats

### 5. Implement JSON Parsing Functionality
- Create parser for the specific command output
- Implement validation logic for the parsed data
- Add proper error handling and logging
- Follow existing code patterns and naming conventions

### 6. Build and Test
- Build the project to ensure no compilation errors
- Run the new command with the implemented feature
- Verify the JSON output matches the expected structure
- Test with various scenarios (empty data, errors, etc.)

### 7. Create Reference JSON File
- Add a sample JSON file to `supported_commands/data/`
- Include both success and error examples if applicable
- Document the command usage and expected output format

## Key Principles

- **Test First**: Always validate the raw ADB command before coding
- **Data-Driven Design**: Let the actual output guide the JSON structure
- **Error Resilience**: Handle edge cases and errors gracefully
- **Consistency**: Follow existing patterns in the codebase
- **Documentation**: Provide clear examples and usage instructions

## Example Command Structure

Each new command should follow the existing pattern:
- Command definition in `cmd/` package
- Parser implementation in `internal/parser/`
- Proper error handling with `apperrors` package
- Structured logging with `logger` package
