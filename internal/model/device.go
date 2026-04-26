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

// SerialNoResponse wraps the device serial number
type SerialNoResponse struct {
	SerialNo string `json:"serial_no"`
}

// DevPathResponse wraps the device path
type DevPathResponse struct {
	DevPath string `json:"dev_path"`
}

// ScreenSizeResponse wraps the screen size
type ScreenSizeResponse struct {
	PhysicalSize string `json:"physical_size"`
}
