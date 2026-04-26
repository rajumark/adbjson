# Shell Auto-Completion

## Overview
adbjson provides shell auto-completion support for bash, zsh, fish, and PowerShell via Cobra's built-in completion system.

## Installation

### Bash
```bash
# Add to ~/.bashrc or ~/.bash_profile
source <(adbjson completion bash)

# Or load completion for current session
eval "$(adbjson completion bash)"
```

### Zsh
```bash
# Add to ~/.zshrc
source <(adbjson completion zsh)

# Or load completion for current session
eval "$(adbjson completion zsh)"
```

### Fish
```bash
# Add to ~/.config/fish/completions/adbjson.fish
adbjson completion fish | source
```

### PowerShell
```powershell
# Add to PowerShell profile
adbjson completion powershell | Out-String | Invoke-Expression
```

## Usage

After installation, press `TAB` to autocomplete commands and flags:

```bash
./adbjson <TAB>
# Shows: adb-version, battery, devices, devices-l, get-devpath, get-serialno, get-state, kill-server, screencap, screensize, screendensity, start-server, version, wm-density, wm-size

./adbjson devices <TAB>
# Shows: --compact, --debug, --format, --help, --pretty

./adbjson devices --format <TAB>
# Shows: json, yaml
```

## Completion Features

- **Command completion**: Auto-complete all available commands
- **Flag completion**: Auto-complete command flags
- **Flag value completion**: Auto-complete flag values (e.g., --format json/yaml)
- **Description support**: Show command and flag descriptions in completion menu

## Generating Completion Scripts

### Bash
```bash
adbjson completion bash > /etc/bash_completion.d/adbjson
```

### Zsh
```bash
adbjson completion zsh > /usr/share/zsh/vendor-completions/_adbjson
```

### Fish
```bash
adbjson completion fish > ~/.config/fish/completions/adbjson.fish
```

### PowerShell
```powershell
adbjson completion powershell > adbjson.ps1
```

## Customizing Completion

Cobra's completion system is highly customizable. See [Cobra Shell Completions](https://github.com/spf13/cobra/blob/main/shell_completions.md) for advanced options.
