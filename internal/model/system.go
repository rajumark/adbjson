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

// SELinuxStatus represents SELinux status
type SELinuxStatus struct {
	Status string `json:"status"`
}

// SELinuxStatusResponse represents the response for getenforce command
type SELinuxStatusResponse struct {
	SELinuxStatus SELinuxStatus `json:"selinux_status"`
}
