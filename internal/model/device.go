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

// Process represents a process from ps output
type Process struct {
	User     string `json:"user"`
	PID      string `json:"pid"`
	PPID     string `json:"ppid"`
	Name     string `json:"name"`
}

// PsResponse wraps the ps operation response
type PsResponse struct {
	Processes []Process `json:"processes"`
	Count     int       `json:"count"`
}

// TopProcess represents a process from top output
type TopProcess struct {
	User     string  `json:"user"`
	PID      string  `json:"pid"`
	PR       string  `json:"pr"`
	NI       string  `json:"ni"`
	VIRT     string  `json:"virt"`
	RES      string  `json:"res"`
	SHR      string  `json:"shr"`
	S        string  `json:"s"`
	CPU      string  `json:"cpu"`
	MEM      string  `json:"mem"`
	Time     string  `json:"time"`
	Args     string  `json:"args"`
}

// TopResponse wraps the top operation response
type TopResponse struct {
	Processes []TopProcess `json:"processes"`
	Count     int          `json:"count"`
	Summary   TopSummary   `json:"summary"`
}

// TopSummary represents system summary from top
type TopSummary struct {
	Tasks  string `json:"tasks"`
	Cpu    string `json:"cpu"`
	Mem    string `json:"mem"`
	Swap   string `json:"swap"`
}

// Filesystem represents a filesystem from df output
type Filesystem struct {
	Filesystem string `json:"filesystem"`
	Blocks     string `json:"blocks"`
	Used       string `json:"used"`
	Available  string `json:"available"`
	UsePercent string `json:"use_percent"`
	MountedOn  string `json:"mounted_on"`
}

// DfResponse wraps the df operation response
type DfResponse struct {
	Filesystems []Filesystem `json:"filesystems"`
	Count       int           `json:"count"`
}

// MemoryInfo represents memory information from free output
type MemoryInfo struct {
	Total    string `json:"total"`
	Used     string `json:"used"`
	Free     string `json:"free"`
	Shared   string `json:"shared"`
	Buffers  string `json:"buffers"`
	Cached   string `json:"cached"`
}

// FreeResponse wraps the free operation response
type FreeResponse struct {
	Memory      MemoryInfo `json:"memory"`
	Buffers     MemoryInfo `json:"buffers"`
	Swap        MemoryInfo `json:"swap"`
}

// UptimeResponse wraps the uptime operation response
type UptimeResponse struct {
	CurrentTime string `json:"current_time"`
	Uptime      string `json:"uptime"`
	Users       string `json:"users"`
	LoadAverage string `json:"load_average"`
}

// MountPoint represents a mount point from mount output
type MountPoint struct {
	Device     string `json:"device"`
	MountPoint string `json:"mount_point"`
	Type       string `json:"type"`
	Options    string `json:"options"`
}

// MountResponse wraps the mount operation response
type MountResponse struct {
	MountPoints []MountPoint `json:"mount_points"`
	Count       int           `json:"count"`
}

// VmstatResponse wraps the vmstat operation response
type VmstatResponse struct {
	Processes VmstatProcesses `json:"processes"`
	Memory    VmstatMemory    `json:"memory"`
	Swap      VmstatSwap      `json:"swap"`
	IO        VmstatIO        `json:"io"`
	System    VmstatSystem    `json:"system"`
	CPU       VmstatCPU       `json:"cpu"`
}

// VmstatProcesses represents process statistics
type VmstatProcesses struct {
	Running  string `json:"running"`
	Blocked  string `json:"blocked"`
}

// VmstatMemory represents memory statistics
type VmstatMemory struct {
	SwapUsed string `json:"swap_used"`
	Free     string `json:"free"`
	Buffers  string `json:"buffers"`
	Cache    string `json:"cache"`
}

// VmstatSwap represents swap statistics
type VmstatSwap struct {
	SwappedIn  string `json:"swapped_in"`
	SwappedOut string `json:"swapped_out"`
}

// VmstatIO represents I/O statistics
type VmstatIO struct {
	BlocksIn  string `json:"blocks_in"`
	BlocksOut string `json:"blocks_out"`
}

// VmstatSystem represents system statistics
type VmstatSystem struct {
	Interrupts string `json:"interrupts"`
	ContextSwitches string `json:"context_switches"`
}

// VmstatCPU represents CPU statistics
type VmstatCPU struct {
	User   string `json:"user"`
	System string `json:"system"`
	Idle   string `json:"idle"`
	Wait   string `json:"wait"`
}

// DateResponse wraps the date operation response
type DateResponse struct {
	DateTime string `json:"datetime"`
}

// LsProcResponse wraps the ls /proc operation response
type LsProcResponse struct {
	Items []string `json:"items"`
	Count int      `json:"count"`
}

// LsRootResponse wraps the ls / operation response
type LsRootResponse struct {
	Items []string `json:"items"`
	Count int      `json:"count"`
}

// DumpsysActivityResponse wraps the dumpsys activity operation response
type DumpsysActivityResponse struct {
	Sections []DumpsysActivitySection `json:"sections"`
	Count    int                      `json:"count"`
}

// DumpsysActivitySection represents a section in dumpsys activity output
type DumpsysActivitySection struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// DumpsysWifiResponse wraps the dumpsys wifi operation response
type DumpsysWifiResponse struct {
	Sections []DumpsysWifiSection `json:"sections"`
	Count    int                  `json:"count"`
}

// DumpsysWifiSection represents a section in dumpsys wifi output
type DumpsysWifiSection struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// ServerResponse wraps server operation response
type ServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
