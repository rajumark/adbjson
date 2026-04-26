package model

// Device represents an ADB device
type Device struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// DevicesResponse wraps the list of devices
type DevicesResponse struct {
	Devices []Device `json:"devices"`
}

// VersionResponse wraps the ADB version information
type VersionResponse struct {
	Version  string `json:"version"`
	Revision string `json:"revision,omitempty"`
}

// StateResponse wraps the device state information
type StateResponse struct {
	State string `json:"state"`
}
