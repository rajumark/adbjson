package model

// Property represents a system property
type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PropertiesResponse represents the response for getprop command
type PropertiesResponse struct {
	Properties []Property `json:"properties"`
	Count      int         `json:"count"`
}
