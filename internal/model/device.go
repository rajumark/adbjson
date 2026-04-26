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
