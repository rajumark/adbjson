# adbjson

<div align="center">
  <h2>A cross-platform CLI tool that wraps Android Debug Bridge (ADB) commands and outputs structured JSON</h2>
  
  <p>
    <a href="https://github.com/rajumark/adbjson/releases">
      <img src="https://img.shields.io/github/release/rajumark/adbjson.svg" alt="Latest Release">
    </a>
    <a href="https://github.com/rajumark/adbjson/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/rajumark/adbjson.svg" alt="License">
    </a>
    <a href="https://goreportcard.com/report/github.com/rajumark/adbjson">
      <img src="https://goreportcard.com/badge/github.com/rajumark/adbjson" alt="Go Report Card">
    </a>
  </p>
</div>

## 📋 Navigation

<div style="display: flex; gap: 10px; margin-bottom: 20px;">
  <button onclick="showTab('user')" id="user-tab" style="padding: 10px 20px; background-color: #007bff; color: white; border: none; border-radius: 5px; cursor: pointer;">👤 User Guide</button>
  <button onclick="showTab('developer')" id="developer-tab" style="padding: 10px 20px; background-color: #6c757d; color: white; border: none; border-radius: 5px; cursor: pointer;">👨‍💻 Developer Guide</button>
</div>

<div id="user-content" style="display: block;">
  <h2>👤 User Guide</h2>
  
  <h3>📦 Installation</h3>
  
  <h4>Option 1: Download Pre-built Binary (Recommended)</h4>
  
  <p>Download the latest release from <a href="https://github.com/rajumark/adbjson/releases">GitHub Releases</a>:</p>
  
  <ul>
    <li><strong>macOS:</strong> <code>adbjson-macos-latest.tar.gz</code></li>
    <li><strong>Linux:</strong> <code>adbjson-ubuntu-latest.tar.gz</code></li>
    <li><strong>Windows:</strong> <code>adbjson-windows-latest.zip</code></li>
  </ul>
  
  <p>Extract the archive and place the <code>adbjson</code> executable in your PATH.</p>
  
  <h4>Option 2: Build from Source</h4>
  
  ```bash
  # Clone the repository
  git clone https://github.com/rajumark/adbjson.git
  cd adbjson
  
  # Download platform-tools for your platform
  # macOS
  curl -L -o platform-tools-darwin.zip https://dl.google.com/android/repository/platform-tools-latest-darwin.zip
  unzip -q platform-tools-darwin.zip -d platform-tools
  mv platform-tools/platform-tools platform-tools/platform-tools-darwin
  rm platform-tools-darwin.zip
  
  # Linux
  curl -L -o platform-tools-linux.zip https://dl.google.com/android/repository/platform-tools-latest-linux.zip
  unzip -q platform-tools-linux.zip -d platform-tools
  mv platform-tools/platform-tools platform-tools/platform-tools-linux
  rm platform-tools-linux.zip
  
  # Windows (PowerShell)
  Invoke-WebRequest -Uri "https://dl.google.com/android/repository/platform-tools-latest-windows.zip" -OutFile "platform-tools-windows.zip"
  Expand-Archive -Path "platform-tools-windows.zip" -DestinationPath "platform-tools"
  Move-Item "platform-tools\platform-tools" "platform-tools\platform-tools-windows"
  Remove-Item "platform-tools-windows.zip"
  
  # Build
  go mod tidy
  go build -o adbjson
  ```
  
  <h3>🚀 Usage</h3>
  
  ```bash
  # List connected devices
  ./adbjson devices
  
  # Get ADB version
  ./adbjson adb-version
  
  # Compact JSON output
  ./adbjson devices --compact
  
  # Show CLI version
  ./adbjson --version
  
  # Get help
  ./adbjson --help
  ```
  
  <h3>✨ Features</h3>
  
  <ul>
    <li>Execute ADB commands and get JSON output</li>
    <li>Bundled ADB support (no system installation required)</li>
    <li>Pretty-printed or compact JSON output</li>
    <li>Cross-platform (macOS, Linux, Windows)</li>
    <li>Easy to integrate with other tools and scripts</li>
  </ul>
  
  <h3>📚 Available Commands</h3>
  
  <p>See the <a href="collection/">collection/</a> directory for detailed documentation of all available commands.</p>
  
  <h3>🆘 Getting Help</h3>
  
  <ul>
    <li>Run <code>./adbjson --help</code> for command-line help</li>
    <li>Check <a href="collection/">individual command documentation</a></li>
    <li>View <a href="command_progress.md">implementation progress</a></li>
  </ul>
</div>

<div id="developer-content" style="display: none;">
  <h2>👨‍💻 Developer Guide</h2>
  
  <h3>🛠️ Development Setup</h3>
  
  <h4>Prerequisites</h4>
  
  <ul>
    <li>Go 1.21 or later</li>
    <li>Git</li>
    <li>Make (optional, for building)</li>
  </ul>
  
  <h4>Setup Steps</h4>
  
  ```bash
  # Clone the repository
  git clone https://github.com/rajumark/adbjson.git
  cd adbjson
  
  # Install dependencies
  go mod tidy
  
  # Run tests
  go test ./...
  
  # Build the project
  go build -o adbjson
  
  # Run the binary
  ./adbjson --version
  ```
  
  <h3>🏗️ Project Structure</h3>
  
  <pre>
adbjson/
├── cmd/              # Command implementations
├── internal/
│   ├── adb/         # ADB execution logic
│   ├── config/      # Configuration management
│   ├── errors/      # Error handling
│   └── formatter/   # JSON formatting
├── collection/      # Command documentation
├── docs/           # Additional documentation
└── main.go         # Entry point
  </pre>
  
  <h3>🤝 Contributing</h3>
  
  <h4>Adding New Commands</h4>
  
  <ol>
    <li>Create a new command file in <code>cmd/</code></li>
    <li>Add the command to the main CLI in <code>main.go</code></li>
    <li>Create documentation in <code>collection/</code></li>
    <li>Add tests for your command</li>
    <li>Update <code>command_progress.md</code></li>
  </ol>
  
  <h4>Code Style</h4>
  
  <ul>
    <li>Follow Go standard formatting (<code>go fmt</code>)</li>
    <li>Use meaningful variable and function names</li>
    <li>Add comments for public functions and complex logic</li>
    <li>Write tests for new functionality</li>
  </ul>
  
  <h4>Submitting Changes</h4>
  
  <ol>
    <li>Fork the repository</li>
    <li>Create a feature branch (<code>git checkout -b feature/amazing-feature</code>)</li>
    <li>Commit your changes (<code>git commit -m 'Add amazing feature'</code>)</li>
    <li>Push to the branch (<code>git push origin feature/amazing-feature</code>)</li>
    <li>Open a Pull Request</li>
  </ol>
  
  <h3>🧪 Testing</h3>
  
  ```bash
  # Run all tests
  go test ./...
  
  # Run tests with coverage
  go test -cover ./...
  
  # Run specific package tests
  go test ./internal/adb
  
  # Run benchmarks
  go test -bench=. ./...
  ```
  
  <h3>📦 Building Releases</h3>
  
  <p>Releases are automatically built using GitHub Actions. See <code>.github/workflows/build.yml</code> for the build configuration.</p>
  
  <h3>📖 Documentation</h3>
  
  <ul>
    <li><a href="PROJECT.md">PROJECT.md</a> - Detailed project documentation</li>
    <li><a href="command_progress.md">command_progress.md</a> - Implementation progress</li>
    <li><a href="documentation/">documentation/</a> - Architecture and design docs</li>
  </ul>
  
  <h3>🐛 Bug Reports</h3>
  
  <p>Please report bugs by opening an issue on GitHub. Include:</p>
  
  <ul>
    <li>Operating system and version</li>
    <li>Go version</li>
    <li>Steps to reproduce</li>
    <li>Expected vs actual behavior</li>
    <li>Any error messages</li>
  </ul>
</div>

<script>
function showTab(tabName) {
  // Hide all content
  document.getElementById('user-content').style.display = 'none';
  document.getElementById('developer-content').style.display = 'none';
  
  // Reset all tab buttons
  document.getElementById('user-tab').style.backgroundColor = '#6c757d';
  document.getElementById('developer-tab').style.backgroundColor = '#6c757d';
  
  // Show selected content and highlight tab
  if (tabName === 'user') {
    document.getElementById('user-content').style.display = 'block';
    document.getElementById('user-tab').style.backgroundColor = '#007bff';
  } else if (tabName === 'developer') {
    document.getElementById('developer-content').style.display = 'block';
    document.getElementById('developer-tab').style.backgroundColor = '#007bff';
  }
}

// Show user tab by default
showTab('user');
</script>
