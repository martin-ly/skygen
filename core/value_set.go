package core

import (
    "fmt"
    "strconv"
    "strings"
)

// A set of values to be set for an event.
type ValueSet struct {
    elementImpl
    Values map[string]interface{}
    Probability float64
    Weight int
}

// Creates a new value set.
func NewValueSet() *ValueSet {
    return &ValueSet{}
}

// Converts the value set to a string-based representation.
func (s *ValueSet) String() string {
    str := "SET "

    output := []string{}
    for k, v := range s.Values {
        switch v := v.(type) {
        case string: 
            output = append(output, fmt.Sprintf("%s = %s", k, strconv.Quote(v)))
        default:
            output = append(output, fmt.Sprintf("%s = %v", k, v))
        }
    }
    str += strings.Join(output, ", ")

    return str
}
