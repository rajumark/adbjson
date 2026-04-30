package ir

import (
	"time"
)

// NodeType represents the type of IR node
type NodeType string

const (
	NodeTypeString   NodeType = "string"
	NodeTypeInt      NodeType = "int"
	NodeTypeFloat    NodeType = "float"
	NodeTypeBool     NodeType = "bool"
	NodeTypeArray    NodeType = "array"
	NodeTypeObject   NodeType = "object"
	NodeTypeNull     NodeType = "null"
)

// Node represents a single piece of data in the IR
type Node struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	Type      NodeType    `json:"type"`
	Metadata  *Metadata   `json:"metadata,omitempty"`
	Children  []*Node     `json:"children,omitempty"`
}

// Metadata provides additional context about the node
type Metadata struct {
	Source     string            `json:"source,omitempty"`     // e.g., "adb", "parser"
	Timestamp  time.Time         `json:"timestamp,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// Document represents the root IR document
type Document struct {
	Command   string    `json:"command"`
	RawOutput string    `json:"raw_output,omitempty"`
	Root      *Node     `json:"root"`
	Metadata  *Metadata `json:"metadata,omitempty"`
}

// NewDocument creates a new IR document
func NewDocument(command, rawOutput string) *Document {
	return &Document{
		Command:   command,
		RawOutput: rawOutput,
		Metadata: &Metadata{
			Source:    "adbjson",
			Timestamp: time.Now(),
		},
	}
}

// NewStringNode creates a string node
func NewStringNode(key, value string) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Type:  NodeTypeString,
	}
}

// NewObjectNode creates an object node
func NewObjectNode(key string) *Node {
	return &Node{
		Key:      key,
		Type:     NodeTypeObject,
		Children: []*Node{},
	}
}

// NewArrayNode creates an array node
func NewArrayNode(key string) *Node {
	return &Node{
		Key:      key,
		Type:     NodeTypeArray,
		Children: []*Node{},
	}
}

// AddChild adds a child node to an object or array node
func (n *Node) AddChild(child *Node) {
	if n.Children == nil {
		n.Children = []*Node{}
	}
	n.Children = append(n.Children, child)
}

// GetChild gets a child by key (for object nodes)
func (n *Node) GetChild(key string) *Node {
	if n.Children == nil {
		return nil
	}
	for _, child := range n.Children {
		if child.Key == key {
			return child
		}
	}
	return nil
}

// GetValue safely gets the value as a string
func (n *Node) GetValue() string {
	if n.Value == nil {
		return ""
	}
	if str, ok := n.Value.(string); ok {
		return str
	}
	return ""
}
