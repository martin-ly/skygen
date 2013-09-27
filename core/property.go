package core

import (
	"fmt"
	"strings"
)

// A property definition for a schema.
type Property struct {
	Name      string
	Transient bool
	DataType  string
}

// Creates a new property definition.
func NewProperty() *Property {
	return &Property{}
}

// Converts the property to a string-based representation.
func (p *Property) String() string {
	str := fmt.Sprintf("PROPERTY %s %s", p.Name, strings.ToUpper(p.DataType))
	if p.Transient {
		str += " TRANSIENT"
	}
	return str
}
