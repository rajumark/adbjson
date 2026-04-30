package ir

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

// Adapter converts IR to specific output formats
type Adapter struct {
	schemas map[string]*Schema
}

// Schema defines the output schema structure
type Schema struct {
	Name   string                 `json:"name"`
	Fields map[string]*FieldDef   `json:"fields"`
}

// FieldDef defines a field in the schema
type FieldDef struct {
	Type     string `json:"type"`
	Source   string `json:"source"`   // IR path to the data
	Required bool   `json:"required"`
	Default  string `json:"default,omitempty"`
}

// NewAdapter creates a new adapter
func NewAdapter() *Adapter {
	adapter := &Adapter{
		schemas: make(map[string]*Schema),
	}
	
	// Register built-in schemas
	adapter.registerBuiltinSchemas()
	
	return adapter
}

// registerBuiltinSchemas registers built-in output schemas
func (a *Adapter) registerBuiltinSchemas() {
	// Devices schema
	devicesSchema := &Schema{
		Name: "devices_response",
		Fields: map[string]*FieldDef{
			"devices": {
				Type:   "array",
				Source: "devices",
			},
		},
	}
	a.schemas["devices_response"] = devicesSchema
	
	// Device schema (for individual devices)
	deviceSchema := &Schema{
		Name: "device",
		Fields: map[string]*FieldDef{
			"id": {
				Type:     "string",
				Source:   "id",
				Required: true,
			},
			"status": {
				Type:     "string",
				Source:   "status",
				Required: true,
			},
			"usb": {
				Type:   "string",
				Source: "usb",
			},
			"product": {
				Type:   "string",
				Source: "product",
			},
		},
	}
	a.schemas["device"] = deviceSchema
	
	// Properties schema (for getprop-like commands)
	propertiesSchema := &Schema{
		Name: "properties_response",
		Fields: map[string]*FieldDef{
			"properties": {
				Type:   "array",
				Source: "properties",
			},
			"count": {
				Type:   "int",
				Source: "properties.length",
			},
		},
	}
	a.schemas["properties_response"] = propertiesSchema
	
	// Single property schema
	singlePropertySchema := &Schema{
		Name: "single_property",
		Fields: map[string]*FieldDef{
			"property": {
				Type:     "string",
				Source:   "key",
				Required: true,
			},
			"value": {
				Type:     "string",
				Source:   "value",
				Required: true,
			},
		},
	}
	a.schemas["single_property"] = singlePropertySchema
	
	// Version schema (for getprop version)
	versionSchema := &Schema{
		Name: "version_response",
		Fields: map[string]*FieldDef{
			"version": {
				Type:     "string",
				Source:   "version",
				Required: true,
			},
		},
	}
	a.schemas["version_response"] = versionSchema
	
	// Version info schema (for adb version)
	versionInfoSchema := &Schema{
		Name: "version_info_response",
		Fields: map[string]*FieldDef{
			"raw_output": {
				Type:     "string",
				Source:   "raw_output",
				Required: true,
			},
		},
	}
	a.schemas["version_info_response"] = versionInfoSchema
	
	// Serial number schema
	serialnoSchema := &Schema{
		Name: "serialno_response",
		Fields: map[string]*FieldDef{
			"serialno": {
				Type:     "string",
				Source:   "serialno",
				Required: true,
			},
		},
	}
	a.schemas["serialno_response"] = serialnoSchema
	
	// State schema
	stateSchema := &Schema{
		Name: "state_response",
		Fields: map[string]*FieldDef{
			"state": {
				Type:     "string",
				Source:   "state",
				Required: true,
			},
		},
	}
	a.schemas["state_response"] = stateSchema
}

// ToJSON converts IR document to JSON using the specified schema
func (a *Adapter) ToJSON(doc *Document, schemaName string) ([]byte, error) {
	return a.ToJSONWithOptions(doc, schemaName, false)
}

// ToJSONWithOptions converts IR document to JSON with formatting options
func (a *Adapter) ToJSONWithOptions(doc *Document, schemaName string, compact bool) ([]byte, error) {
	schema, exists := a.schemas[schemaName]
	if !exists {
		return nil, fmt.Errorf("schema not found: %s", schemaName)
	}
	
	data, err := a.transform(doc.Root, schema)
	if err != nil {
		return nil, err
	}
	
	// Pure parsing - return raw data only, no wrappers
	if compact {
		return json.Marshal(data)
	}
	return json.MarshalIndent(data, "", "  ")
}

// ToYAML converts IR document to YAML using the specified schema
func (a *Adapter) ToYAML(doc *Document, schemaName string, compact bool) ([]byte, error) {
	schema, exists := a.schemas[schemaName]
	if !exists {
		return nil, fmt.Errorf("schema not found: %s", schemaName)
	}
	
	data, err := a.transform(doc.Root, schema)
	if err != nil {
		return nil, err
	}
	
	// Pure parsing - return raw data only, no wrappers
	return yaml.Marshal(data)
}

// getCount extracts count from data based on its structure
func (a *Adapter) getCount(data interface{}) int {
	switch v := data.(type) {
	case map[string]interface{}:
		// For responses with arrays (devices, properties, etc.)
		if devices, ok := v["devices"]; ok {
			// Handle both []interface{} and []map[string]interface{} types
			switch arr := devices.(type) {
			case []interface{}:
				return len(arr)
			case []map[string]interface{}:
				return len(arr)
			}
		}
		if properties, ok := v["properties"]; ok {
			// Handle both []interface{} and []map[string]interface{} types
			switch arr := properties.(type) {
			case []interface{}:
				return len(arr)
			case []map[string]interface{}:
				return len(arr)
			}
		}
		// For single object responses (version, serialno, etc.), count is 1
		return 1
	case []interface{}:
		return len(v)
	case []map[string]interface{}:
		return len(v)
	default:
		return 1
	}
}


// transform transforms IR node according to schema
func (a *Adapter) transform(node *Node, schema *Schema) (interface{}, error) {
	switch schema.Name {
	case "devices_response":
		return a.transformDevicesResponse(node)
	case "device":
		return a.transformDevice(node)
	case "properties_response":
		return a.transformPropertiesResponse(node)
	case "single_property":
		return a.transformSingleProperty(node)
	case "version_response":
		return a.transformVersionResponse(node)
	case "version_info_response":
		return a.transformVersionInfoResponse(node)
	default:
		return a.transformGeneric(node, schema)
	}
}

// transformDevicesResponse transforms devices array to expected format
func (a *Adapter) transformDevicesResponse(node *Node) ([]map[string]interface{}, error) {
	if node.Type != NodeTypeArray {
		return nil, fmt.Errorf("expected array node for devices response")
	}
	
	devices := make([]map[string]interface{}, 0)
	
	for _, child := range node.Children {
		device, err := a.transformDevice(child)
		if err != nil {
			continue // Skip invalid devices
		}
		devices = append(devices, device)
	}
	
	return devices, nil
}

// transformDevice transforms single device node
func (a *Adapter) transformDevice(node *Node) (map[string]interface{}, error) {
	if node.Type != NodeTypeObject {
		return nil, fmt.Errorf("expected object node for device")
	}
	
	device := make(map[string]interface{})
	
	for _, child := range node.Children {
		switch child.Key {
		case "id":
			device["id"] = child.GetValue()
		case "status":
			device["status"] = child.GetValue()
		case "usb":
			device["usb"] = child.GetValue()
		case "product":
			device["product"] = child.GetValue()
		}
	}
	
	return device, nil
}

// transformPropertiesResponse transforms properties array
func (a *Adapter) transformPropertiesResponse(node *Node) (map[string]interface{}, error) {
	if node.Type != NodeTypeArray {
		return nil, fmt.Errorf("expected array node for properties response")
	}
	
	properties := make([]map[string]interface{}, 0)
	
	for _, child := range node.Children {
		prop := map[string]interface{}{
			"key":   child.GetChild("key").GetValue(),
			"value": child.GetChild("value").GetValue(),
		}
		properties = append(properties, prop)
	}
	
	return map[string]interface{}{
		"properties": properties,
		"count":      len(properties),
	}, nil
}

// transformSingleProperty transforms single property
func (a *Adapter) transformSingleProperty(node *Node) (map[string]interface{}, error) {
	if node.Type != NodeTypeObject {
		return nil, fmt.Errorf("expected object node for single property")
	}
	
	result := make(map[string]interface{})
	
	for _, child := range node.Children {
		result[child.Key] = child.GetValue()
	}
	
	return result, nil
}

// transformVersionResponse transforms version response
func (a *Adapter) transformVersionResponse(node *Node) (map[string]interface{}, error) {
	if node.Type != NodeTypeObject {
		return nil, fmt.Errorf("expected object node for version response")
	}
	
	version := ""
	for _, child := range node.Children {
		if child.Key == "version" {
			version = child.GetValue()
			break
		}
	}
	
	return map[string]interface{}{
		"version": version,
	}, nil
}

// transformVersionInfoResponse transforms version info response
func (a *Adapter) transformVersionInfoResponse(node *Node) (map[string]interface{}, error) {
	if node.Type != NodeTypeObject {
		return nil, fmt.Errorf("expected object node for version info response")
	}
	
	rawOutput := ""
	for _, child := range node.Children {
		if child.Key == "raw_output" {
			rawOutput = child.GetValue()
			break
		}
	}
	
	return map[string]interface{}{
		"raw_output": rawOutput,
	}, nil
}

// transformGeneric provides generic transformation
func (a *Adapter) transformGeneric(node *Node, schema *Schema) (interface{}, error) {
	switch node.Type {
	case NodeTypeString:
		return node.Value, nil
	case NodeTypeInt:
		return node.Value, nil
	case NodeTypeFloat:
		return node.Value, nil
	case NodeTypeBool:
		return node.Value, nil
	case NodeTypeArray:
		result := make([]interface{}, 0)
		for _, child := range node.Children {
			value, err := a.transformGeneric(child, schema)
			if err != nil {
				continue
			}
			result = append(result, value)
		}
		return result, nil
	case NodeTypeObject:
		result := make(map[string]interface{})
		for _, child := range node.Children {
			value, err := a.transformGeneric(child, schema)
			if err != nil {
				continue
			}
			result[child.Key] = value
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported node type: %s", node.Type)
	}
}

// RegisterSchema registers a custom schema
func (a *Adapter) RegisterSchema(schema *Schema) {
	a.schemas[schema.Name] = schema
}
