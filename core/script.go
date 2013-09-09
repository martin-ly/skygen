package core

import (
    "fmt"
    "github.com/skydb/sky.go"
    "strings"
    "time"
)

type Script struct {
    schema *Schema
    events Events
}

func NewScript() *Script {
    return &Script{}
}

// Returns the script schema.
func (s *Script) Schema() *Schema {
    return s.schema
}

// Sets the script's schema.
func (s *Script) SetSchema(schema *Schema) {
    if s.schema != nil {
        s.schema.SetParent(nil)
    }
    s.schema = schema
    if s.schema != nil {
        s.schema.SetParent(nil)
    }
}

// Returns a list of root events.
func (s *Script) Events() Events {
    return s.events
}

// Sets the script's root events.
func (s *Script) SetEvents(events Events) {
    for _, event := range s.events {
        event.SetParent(nil)
    }
    s.events = events
    for _, event := range s.events {
        event.SetParent(s)
    }
}

func (s *Script) Parent() Element {
    return nil
}

func (s *Script) SetParent(e Element) {
}

func (s *Script) Script() *Script {
    return s
}

// Generates a timeline for an object.
func (s *Script) Generate(stream *sky.EventStream, id string) error {
    timestamp, err := time.Parse(time.RFC3339, "2000-01-01T00:00:00Z")
    if err != nil {
        return fmt.Errorf("Invalid timestamp: %s", "2000-01-01T00:00:00Z")
    }

    return s.events.Generate(stream, id, timestamp)
}

// Converts the script to a string-based representation.
func (s *Script) String() string {
    output := make([]string, 0)
    if s.schema != nil {
        output = append(output, s.schema.String())
    }
    if str := s.events.String(); len(str) > 0 {
        output = append(output, str)
    }
    return strings.Join(output, "\n")
}
