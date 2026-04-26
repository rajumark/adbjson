package model

// Package represents an Android package
type Package struct {
	Name string `json:"name"`
}

// PackagesResponse represents the response for package list commands
type PackagesResponse struct {
	Packages []Package `json:"packages"`
	Count    int       `json:"count"`
}

// Feature represents an Android feature
type Feature struct {
	Name string `json:"name"`
}

// FeaturesResponse represents the response for feature list command
type FeaturesResponse struct {
	Features []Feature `json:"features"`
	Count    int       `json:"count"`
}

// Library represents an Android library
type Library struct {
	Name string `json:"name"`
}

// LibrariesResponse represents the response for library list command
type LibrariesResponse struct {
	Libraries []Library `json:"libraries"`
	Count     int       `json:"count"`
}

// Instrumentation represents an Android instrumentation
type Instrumentation struct {
	Name string `json:"name"`
}

// InstrumentationResponse represents the response for instrumentation list command
type InstrumentationResponse struct {
	Instrumentations []Instrumentation `json:"instrumentations"`
	Count            int               `json:"count"`
}

// Permission represents an Android permission
type Permission struct {
	Name string `json:"name"`
}

// PermissionsResponse represents the response for permission list command
type PermissionsResponse struct {
	Permissions []Permission `json:"permissions"`
	Count       int           `json:"count"`
}
