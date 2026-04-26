# JSON Schema Validation

## Overview
adbjson includes a validation system to ensure data structures are valid before output.

## Validator Features

### Data Structure Validation
Validates Go data structures including:
- Structs
- Slices/Arrays
- Maps
- Pointers

### JSON Validation
Validates JSON bytes against:
- JSON syntax validity
- Expected data structure type

## Usage

### Validate Data Structures
```go
import "adbjson/internal/validator"

validator := validator.NewValidator()

// Validate a struct
err := validator.Validate(myStruct)

// Validate a slice
err := validator.Validate(mySlice)

// Validate a map
err := validator.Validate(myMap)
```

### Validate JSON
```go
// Validate JSON syntax
err := validator.ValidateJSON(jsonBytes, nil)

// Validate JSON against expected type
expectedType := struct{ Name string }{}
err := validator.ValidateJSON(jsonBytes, expectedType)
```

## Validation Rules

### Nil Checks
- Rejects nil data
- Rejects nil pointers
- Validates nested nil pointers in structs

### Type Checking
- Ensures JSON matches expected structure
- Validates nested structures recursively
- Handles slices, arrays, and maps

## Error Handling

Validation errors include:
- Field name that failed validation
- Descriptive error message

```go
type ValidationError struct {
    Field   string
    Message string
}
```

## Integration with Parsers

Parsers can use the validator in their `Validate` method:

```go
func (p *DevicesParser) Validate(result *model.DevicesResponse) error {
    validator := validator.NewValidator()
    return validator.Validate(result)
}
```

## Testing

Run validator tests:
```bash
go test ./internal/validator/... -v
```

## Future Enhancements
- JSON Schema (Draft 7) support
- Custom validation rules
- Field-level validation tags
- Range and format validation
