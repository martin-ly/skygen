package core

import (
    "fmt"
    "github.com/skydb/sky.go"
    "time"
)

type Script struct {
    events Events
}

func NewScript() *Script {
    return &Script{}
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
func (s *Script) Generate(t *sky.Table, id string) error {
    timestamp, err := time.Parse(time.RFC3339, "2000-01-01T00:00:00Z")
    if err != nil {
        return fmt.Errorf("Invalid timestamp: %s", "2000-01-01T00:00:00Z")
    }

    return s.events.Generate(t, id, timestamp)
}

// Converts the script to a string-based representation.
func (s *Script) String() string {
    return s.events.String()
}
