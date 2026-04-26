# Testing Framework

## Overview
adbjson uses Go's built-in testing framework with table-driven tests for comprehensive coverage.

## Running Tests

### Run All Tests
```bash
go test ./...
```

### Run Tests with Verbose Output
```bash
go test ./... -v
```

### Run Tests for Specific Package
```bash
go test ./internal/parser/...
```

### Run Tests with Coverage
```bash
go test ./... -cover
```

### Run Tests with Coverage Report
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Structure

### Parser Tests
Each parser should have tests for:
- Parse method with various inputs
- Validate method
- Name method
- Version method

```go
func TestDevicesParser_Parse(t *testing.T) {
    parser := NewDevicesParser()
    
    tests := []struct {
        name    string
        input   string
        want    *model.DevicesResponse
        wantErr bool
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := parser.Parse(tt.input)
            // assertions
        })
    }
}
```

## Test Coverage Goals
- **Minimum**: 80% coverage
- **Target**: 90% coverage
- **Ideal**: 95%+ coverage

## Test Categories

### Unit Tests
- Test individual functions and methods
- No external dependencies
- Fast execution

### Integration Tests
- Test command execution with real ADB
- Requires connected device
- Slower execution

### Parser Tests
- Test output parsing logic
- Test validation logic
- Test edge cases

## Best Practices

1. **Table-Driven Tests**: Use table-driven tests for multiple scenarios
2. **Descriptive Names**: Use descriptive test names
3. **Test Edge Cases**: Test empty inputs, malformed data, etc.
4. **Isolation**: Tests should not depend on each other
5. **Fast Tests**: Keep tests fast for quick feedback

## Example Test Case

```go
func TestDevicesParser_Parse(t *testing.T) {
    parser := NewDevicesParser()
    
    tests := []struct {
        name    string
        input   string
        want    *model.DevicesResponse
        wantErr bool
    }{
        {
            name: "single device",
            input: "List of devices attached\nemulator-5554\tdevice",
            want: &model.DevicesResponse{
                Devices: []model.Device{
                    {ID: "emulator-5554", Status: "device"},
                },
            },
            wantErr: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := parser.Parse(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // more assertions
        })
    }
}
```

## CI/CD Integration
Tests run automatically on:
- Pull requests
- Push to main branch
- Scheduled runs

## Future Enhancements
- Mock ADB server for integration tests
- Property-based testing
- Fuzz testing
- Benchmark tests
