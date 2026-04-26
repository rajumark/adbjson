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

// UserInfo represents user information
type UserInfo struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

// UserInfoResponse represents the response for id command
type UserInfoResponse struct {
	UserInfo UserInfo `json:"user_info"`
}
