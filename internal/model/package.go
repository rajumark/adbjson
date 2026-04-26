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
