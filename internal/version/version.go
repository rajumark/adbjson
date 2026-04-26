package version

// Version information
const (
	// Version is the current version of adbjson
	Version = "1.0.0"
	
	// Build information (set during build)
	BuildDate = "unknown"
	GitCommit = "unknown"
	GitBranch = "unknown"
)

// GetVersion returns the full version string
func GetVersion() string {
	return Version
}

// GetFullVersion returns version with build information
func GetFullVersion() string {
	return Version + " (commit: " + GitCommit + ", branch: " + GitBranch + ", built: " + BuildDate + ")"
}

// GetBuildInfo returns build information as a map
func GetBuildInfo() map[string]string {
	return map[string]string{
		"version":   Version,
		"commit":    GitCommit,
		"branch":    GitBranch,
		"build_date": BuildDate,
	}
}
