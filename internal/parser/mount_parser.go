package parser

import (
	"adbjson/internal/model"
	"strings"
)

// MountParser parses adb shell mount output
type MountParser struct {
	*BaseParser
}

// NewMountParser creates a new mount parser
func NewMountParser() *MountParser {
	return &MountParser{
		BaseParser: NewBaseParser("mount", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell mount" command
func (p *MountParser) Parse(output string) (*model.MountResponse, error) {
	lines := strings.Split(output, "\n")
	mountPoints := []model.MountPoint{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Parse mount line
		// Format: /dev/block/dm-24 on / type ext4 (ro,dirsync,seclabel,nodev,noatime)
		// Alternative format: tmpfs /dev tmpfs rw,seclabel,nosuid,nodev,noexec,relatime 0 0
		if strings.Contains(line, " on ") {
			// Parse format with "on"
			parts := strings.Split(line, " on ")
			if len(parts) >= 2 {
				device := strings.TrimSpace(parts[0])
				rest := strings.Join(parts[1:], " on ")
				
				restParts := strings.Split(rest, " type ")
				if len(restParts) >= 2 {
					mountPointPath := strings.TrimSpace(restParts[0])
					typeAndOptions := strings.Join(restParts[1:], " type ")
					
					typeAndOptionsParts := strings.Split(typeAndOptions, " (")
					if len(typeAndOptionsParts) >= 2 {
						fsType := strings.TrimSpace(typeAndOptionsParts[0])
						options := strings.TrimSuffix(strings.Join(typeAndOptionsParts[1:], " ("), ")")
						
						mountPt := model.MountPoint{
							Device:     device,
							MountPoint: mountPointPath,
							Type:       fsType,
							Options:    options,
						}
						mountPoints = append(mountPoints, mountPt)
					}
				}
			}
		} else {
			// Parse alternative format without "on"
			// Format: tmpfs /dev tmpfs rw,seclabel,nosuid,nodev,noexec,relatime 0 0
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				device := parts[0]
				mountPointPath := parts[1]
				fsType := parts[2]
				options := parts[3]
				
				mountPt := model.MountPoint{
					Device:     device,
					MountPoint: mountPointPath,
					Type:       fsType,
					Options:    options,
				}
				mountPoints = append(mountPoints, mountPt)
			}
		}
	}
	
	response := &model.MountResponse{
		MountPoints: mountPoints,
		Count:       len(mountPoints),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *MountParser) Validate(result *model.MountResponse) error {
	// Mount response is always valid
	return nil
}
