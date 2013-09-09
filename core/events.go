package core

import (
    "github.com/skydb/sky.go"
    "math/rand"
    "strings"
    "time"
)

type Events []*Event

func (s Events) String() string {
    output := []string{}
    for _, e := range s {
        output = append(output, e.String())
    }
    return strings.Join(output, "\n")
}

// Selects one event based on weights and generate events based on it.
func (s Events) Generate(stream *sky.EventStream, id string, timestamp time.Time) error {
    // Generate numbers based on weights and pick the highest.
    var maxValue int = 0
    var event *Event
    for _, e := range s {
        if e.Weight > 0 {
            if value := rand.Intn(e.Weight) + 1; value > maxValue {
                maxValue = value
                event = e
            }
        }
    }

    // If an event was found then generate it.
    if event != nil {
        return event.Generate(stream, id, timestamp)
    }

    return nil
}
