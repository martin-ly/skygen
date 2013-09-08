package core

import (
    "fmt"
    "github.com/skydb/sky.go"
)

type Schema struct {
    elementImpl
    Properties Properties
}

func NewSchema() *Schema {
    return &Schema{}
}

// Updates the table to contain the given properties.
func (s *Schema) Sync(t *sky.Table) error {
    // Retrieve a list of properties from the table.
    properties, err := t.GetProperties()
    if err != nil {
        return err
    }
    
    // Create lookup by name.
    lookup := map[string]*sky.Property{}
    for _, property := range properties {
        lookup[property.Name] = property
    }

    // Loop over schema properties and create any that are missing.
    for _, p := range s.Properties {
        property := lookup[p.Name]
        if property == nil {
            property = sky.NewProperty(p.Name, p.Transient, p.DataType)
            if err = t.CreateProperty(property); err != nil {
                return err
            }
            lookup[property.Name] = property
        } else if property.DataType != p.DataType {
            return fmt.Errorf("Conflicting schema data type for existing property: %s (%s != %s)", p.Name, property.DataType, p.DataType)
        }
    }

    return nil
}

// Converts the schema to a string-based representation.
func (s *Schema) String() string {
    str := "SCHEMA\n"
    if inner := s.Properties.String(); inner != "" {
        str += lineStartRegex.ReplaceAllString(inner, "  ") + "\n"
    }
    str += "END"
    return str
}
