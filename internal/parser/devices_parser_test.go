package parser

import (
	"adbjson/internal/model"
	"testing"
)

func TestDevicesParser_Parse(t *testing.T) {
	parser := NewDevicesParser()
	
	tests := []struct {
		name    string
		input   string
		want    *model.DevicesResponse
		wantErr bool
	}{
		{
			name: "single device",
			input: "List of devices attached\nemulator-5554\tdevice",
			want: &model.DevicesResponse{
				Devices: []model.Device{
					{ID: "emulator-5554", Status: "device"},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple devices",
			input: "List of devices attached\nemulator-5554\tdevice\nZD222XW5RL\tdevice",
			want: &model.DevicesResponse{
				Devices: []model.Device{
					{ID: "emulator-5554", Status: "device"},
					{ID: "ZD222XW5RL", Status: "device"},
				},
			},
			wantErr: false,
		},
		{
			name: "no devices",
			input: "List of devices attached",
			want: &model.DevicesResponse{
				Devices: []model.Device{},
			},
			wantErr: false,
		},
		{
			name: "device offline",
			input: "List of devices attached\nemulator-5554\toffline",
			want: &model.DevicesResponse{
				Devices: []model.Device{
					{ID: "emulator-5554", Status: "offline"},
				},
			},
			wantErr: false,
		},
		{
			name: "empty input",
			input: "",
			want: &model.DevicesResponse{
				Devices: []model.Device{},
			},
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Devices) != len(tt.want.Devices) {
				t.Errorf("Parse() device count = %d, want %d", len(got.Devices), len(tt.want.Devices))
				return
			}
			for i, device := range got.Devices {
				if device.ID != tt.want.Devices[i].ID || device.Status != tt.want.Devices[i].Status {
					t.Errorf("Parse() device[%d] = %+v, want %+v", i, device, tt.want.Devices[i])
				}
			}
		})
	}
}

func TestDevicesParser_Validate(t *testing.T) {
	parser := NewDevicesParser()
	
	tests := []struct {
		name    string
		result  *model.DevicesResponse
		wantErr bool
	}{
		{
			name: "valid response",
			result: &model.DevicesResponse{
				Devices: []model.Device{
					{ID: "emulator-5554", Status: "device"},
				},
			},
			wantErr: false,
		},
		{
			name: "empty response",
			result: &model.DevicesResponse{
				Devices: []model.Device{},
			},
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.Validate(tt.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevicesParser_Name(t *testing.T) {
	parser := NewDevicesParser()
	if parser.Name() != "devices" {
		t.Errorf("Name() = %v, want devices", parser.Name())
	}
}

func TestDevicesParser_Version(t *testing.T) {
	parser := NewDevicesParser()
	if parser.Version() != "1.0.0" {
		t.Errorf("Version() = %v, want 1.0.0", parser.Version())
	}
}
