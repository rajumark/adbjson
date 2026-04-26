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
