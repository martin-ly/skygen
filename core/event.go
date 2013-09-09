package core

import (
    "encoding/json"
    "fmt"
    "github.com/skydb/sky.go"
    "math/rand"
    "time"
)

// The definition for a generated event.
type Event struct {
    elementImpl
    valueSets ValueSets
    events Events
    After []time.Duration
    Weight int
}

// Creates a new event definition.
func NewEvent() *Event {
    return &Event{}
}

// Returns a list of value sets.
func (e *Event) ValueSets() ValueSets {
    return e.valueSets
}

// Sets the event's value sets.
func (e *Event) SetValueSets(sets ValueSets) {
    for _, v := range e.valueSets {
        v.SetParent(nil)
    }
    e.valueSets = sets
    for _, v := range e.valueSets {
        v.SetParent(e)
    }
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

// Generates an event.
func (e *Event) Generate(stream *sky.EventStream, id string, timestamp time.Time) error {
    // Move timestamp forward.
    if len(e.After) == 2 {
        duration := e.After[0]
        diff := int64(e.After[1]) - int64(e.After[0])
        if diff > 0 {
            duration += time.Duration(rand.Int63n(diff))
        }
        timestamp = timestamp.Add(duration)
    }


    // Generate data for event.
    data := make(map[string]interface{})
    for _, valueSet := range e.valueSets {
        if valueSet.Probability >= 100 || (valueSet.Probability > 0 && valueSet.Probability < rand.Intn(100)) {
            for k, v := range valueSet.Values {
                data[k] = v
            }
        }
    }

    // Only create an event if we have data.
    if len(data) > 0 {
        if err := stream.AddEvent(id, sky.NewEvent(timestamp, data)); err != nil {
            return err
        }
    }

    // Show events as they're being generated.
    json, _ := json.Marshal(map[string]interface{}{
        "timestamp":timestamp.UTC().Format(time.RFC3339),
        "data":data,
    })
    fmt.Println(string(json))

    // Continue to generate events down the chain.
    return e.events.Generate(stream, id, timestamp)
}

// Converts the event to a string-based representation.
func (e *Event) String() string {
    var inner string
    str := "EVENT"
    if len(e.After) == 2 {
        if e.After[0] == e.After[1] {
            str += " AFTER " + formatDuration(e.After[0])
        } else {
            str += " AFTER " + formatDuration(e.After[0]) + " - " + formatDuration(e.After[1])
        }
    }
    if e.Weight != 1 {
        str += fmt.Sprintf(" WEIGHT %d", e.Weight)
    }
    str += "\n"

    inner = e.valueSets.String()
    if inner != "" {
        str += lineStartRegex.ReplaceAllString(inner, "  ") + "\n"
    }

    inner = e.events.String()
    if inner != "" {
        str += lineStartRegex.ReplaceAllString(inner, "  ") + "\n"
    }

    str += "END"
    return str
}
