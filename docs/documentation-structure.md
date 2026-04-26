# Documentation Structure Rules

## Collection Folder Structure

All individual command documentation files are stored in the `collection/` folder at the root of the project.

## Command Documentation Format

Each command must have its own MD file in the `collection/` folder with the following structure:

```markdown
# [command-name]

## Description
[Brief description of what the command does]

## Command
```bash
adbjson [full-command]
```

## Equivalent ADB Command
```bash
adb [full-command]
```

## Sample JSON Output
```json
{
  [sample JSON output]
}
```

## Sample YAML Output (optional)
```bash
adbjson [command] --format yaml
```
```yaml
[sample YAML output]
```

## Flags
- `--flag`: Description
- ...

## Notes
- Additional notes about the command
```

## Strict Rules

1. **Command Naming**: Only replace `adb` with `adbjson`. Never alter the original ADB command structure.
   - Example: `adb shell pm list packages` → `adbjson shell pm list packages`
   - Example: `adb devices -l` → `adbjson devices -l`
   - Example: `adb shell dumpsys battery` → `adbjson battery` (only for simplified commands)

2. **Equivalent ADB Command**: Every documentation file MUST include the "Equivalent ADB Command" section showing the original ADB command that maps to the adbjson command.

3. **List Commands**: For list commands that return many items (e.g., packages with 400+ items), show maximum 5 items in the documentation examples. The actual output will contain all items.

4. **Dump Commands**: For dump commands (e.g., dumpsys), show full JSON in documentation examples.

5. **Implementation Order**: Implement commands one at a time following this workflow:
   - Create command implementation
   - Build and verify CLI works
   - Create/update documentation in collection/
   - Update command_progress.md
   - Commit and push changes
   - Move to next command

## File Naming Convention

Documentation files should be named using the command structure with hyphens:
- `shell-pm-list-packages.md` for `adbjson shell pm list packages`
- `devices-l.md` for `adbjson devices -l`
- `battery.md` for `adbjson battery`
