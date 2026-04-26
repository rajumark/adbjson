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

// ConnectResponse wraps the connect operation response
type ConnectResponse struct {
	Connected bool   `json:"connected"`
	Target    string `json:"target"`
	Message   string `json:"message"`
}

// DisconnectResponse wraps the disconnect operation response
type DisconnectResponse struct {
	Disconnected bool   `json:"disconnected"`
	Target       string `json:"target"`
	Message      string `json:"message"`
}

// RootResponse wraps the root operation response
type RootResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UnrootResponse wraps the unroot operation response
type UnrootResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SinglePropertyResponse wraps a single property value
type SinglePropertyResponse struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

// SetenforceResponse wraps the setenforce operation response
type SetenforceResponse struct {
	Success bool   `json:"success"`
	Mode    string `json:"mode"`
	Message string `json:"message"`
}

// TcpipResponse wraps the tcpip operation response
type TcpipResponse struct {
	Success bool   `json:"success"`
	Port    string `json:"port"`
	Message string `json:"message"`
}

// KeyeventResponse wraps the keyevent operation response
type KeyeventResponse struct {
	Success  bool   `json:"success"`
	Keycode  string `json:"keycode"`
	Message  string `json:"message"`
}

// WmSizeResetResponse wraps the wm size reset operation response
type WmSizeResetResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// WmDensityResetResponse wraps the wm density reset operation response
type WmDensityResetResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Setting represents a key-value setting pair
type Setting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SettingsListResponse wraps the settings list operation response
type SettingsListResponse struct {
	Namespace string    `json:"namespace"`
	Settings  []Setting `json:"settings"`
	Count     int       `json:"count"`
}

// ServerResponse wraps server operation response
type ServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
