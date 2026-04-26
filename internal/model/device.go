package model

// Device represents an ADB device
type Device struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	USB          string `json:"usb,omitempty"`
	Product      string `json:"product,omitempty"`
	Model        string `json:"model,omitempty"`
	Device       string `json:"device,omitempty"`
	TransportID  string `json:"transport_id,omitempty"`
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

// ScreenDensityResponse wraps the screen density
type ScreenDensityResponse struct {
	PhysicalDensity string `json:"physical_density"`
}

// BatteryInfoResponse wraps battery information
type BatteryInfoResponse struct {
	ACPowered        bool    `json:"ac_powered"`
	USBPowered       bool    `json:"usb_powered"`
	WirelessPowered  bool    `json:"wireless_powered"`
	DockPowered      bool    `json:"dock_powered"`
	MaxChargingCurrent int   `json:"max_charging_current"`
	ChargeCounter    int     `json:"charge_counter"`
	Status           int     `json:"status"`
	Health           int     `json:"health"`
	Present          bool    `json:"present"`
	Level            int     `json:"level"`
	Scale            int     `json:"scale"`
	Voltage          int     `json:"voltage"`
	Temperature      int     `json:"temperature"`
	Technology       string  `json:"technology"`
	ChargingState    int     `json:"charging_state"`
	ChargingPolicy   int     `json:"charging_policy"`
	CapacityLevel    int     `json:"capacity_level"`
	VBusState        bool    `json:"vbus_state"`
	ChargeWatt       int     `json:"charge_watt"`
	ChargeWattDesign int     `json:"charge_watt_design"`
	ChargeType       int     `json:"charge_type"`
	CycleCount       int     `json:"cycle_count"`
	FullCapacity     int     `json:"full_capacity"`
	FullDesignCapacity int `json:"full_design_capacity"`
}

// ServerResponse wraps server operation response
type ServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
