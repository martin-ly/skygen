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
	// Normalize weights.
	weights := make([]struct {
		min float64
		max float64
	}, len(s))
	current, total := float64(0), float64(s.TotalWeight())

	for i, e := range s {
		weight := float64(e.Weight) / total
		weights[i].min = current
		weights[i].max = current + weight
		current += float64(weight)
	}

	// Determine a weighted random entry.
	pick := rand.Float64()
	var event *Event
	for i, e := range s {
		if pick >= weights[i].min && pick <= weights[i].max {
			event = e
			break
		}
	}

	// If an event was found then generate it.
	if event != nil {
		return event.Generate(stream, id, timestamp)
	}

	return nil
}

// The total weight of all the events
func (s Events) TotalWeight() int {
	sum := 0
	for _, e := range s {
		sum += e.Weight
	}
	return sum
}
