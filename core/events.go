package core

import (
    "strings"
)

type Events []*Event

func (s Events) String() string {
    output := []string{}
    for _, e := range s {
        output = append(output, e.String())
    }
    return strings.Join(output, "\n")
}
