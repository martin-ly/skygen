package core

import (
)

// The definition for a generated event.
type Event struct {
    elementImpl
    events Events
}

// Creates a new event definition.
func NewEvent() *Event {
    return &Event{}
}

// Returns a list of child events.
func (e *Event) Events() Events {
    return e.events
}

// Sets the event's child events.
func (e *Event) SetEvents(events Events) {
    for _, event := range e.events {
        event.SetParent(nil)
    }
    e.events = events
    for _, event := range e.events {
        event.SetParent(e)
    }
}

// Converts the event to a string-based representation.
func (e *Event) String() string {
    str := "EVENT DO\n"
    str += lineStartRegex.ReplaceAllString(e.events.String(), "  ") + "\n"
    str += "END"
    return str
}
