package core

import (
    "time"
)

// The definition for a generated event.
type Event struct {
    elementImpl
    events Events
    After []time.Duration
    Weight int
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
    str := "EVENT"
    if len(e.After) == 2 {
        if e.After[0] == e.After[1] {
            str += " AFTER " + formatDuration(e.After[0])
        } else {
            str += " AFTER " + formatDuration(e.After[0]) + " - " + formatDuration(e.After[1])
        }
    }
    str += " DO\n"
    str += lineStartRegex.ReplaceAllString(e.events.String(), "  ") + "\n"
    str += "END"
    return str
}
