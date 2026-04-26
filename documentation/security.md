# Security - Input Sanitization

## Overview
adbjson includes input sanitization to prevent security vulnerabilities like command injection and path traversal attacks.

## Sanitization Functions

### SanitizeString
Removes control characters from strings while preserving tabs, newlines, and carriage returns.

```go
sanitized := sanitize.SanitizeString(input)
```

### SanitizeFilePath
Prevents directory traversal attacks by removing:
- `../` and `..\\` (path traversal)
- `./` and `.\\` (current directory references)
- Leading slashes (absolute paths)
- Backslashes (Windows path separators)

```go
safePath := sanitize.SanitizeFilePath(userInput)
```

### SanitizeCommandArgs
Sanitizes an array of command arguments.

```go
safeArgs := sanitize.SanitizeCommandArgs(args)
```

### IsValidCommand
Checks if a command string is safe to execute by blocking dangerous commands:
- `rm -rf /`
- `dd if=`
- Fork bombs
- Device writes

```go
if !sanitize.IsValidCommand(command) {
    return errors.New("dangerous command blocked")
}
```

### SanitizePackageName
Sanitizes Android package names to only allow alphanumeric, dots, and underscores.

```go
safePkg := sanitize.SanitizePackageName(userInput)
```

## Security Best Practices

1. **Always sanitize user input** before using it in ADB commands
2. **Validate file paths** to prevent directory traversal
3. **Block dangerous commands** that could harm the system
4. **Use package name sanitization** for package-related operations

## Example Usage

```go
import "adbjson/internal/sanitize"

func installPackage(pkgName string) error {
    // Sanitize package name
    safePkg := sanitize.SanitizePackageName(pkgName)
    
    // Validate command
    if !sanitize.IsValidCommand("install") {
        return errors.New("command not allowed")
    }
    
    // Execute with sanitized input
    return executor.Execute("install", safePkg)
}
```

## Testing
All sanitization functions have comprehensive unit tests:

```bash
go test ./internal/sanitize/... -v
```

## Threats Mitigated
- **Command Injection**: Blocks dangerous shell commands
- **Path Traversal**: Prevents `../` attacks
- **Null Byte Injection**: Removes null bytes
- **Control Character Injection**: Removes control characters
