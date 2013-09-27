package core

import (
	"strings"
)

type ValueSets []*ValueSet

func (s ValueSets) String() string {
	output := []string{}
	for _, v := range s {
		output = append(output, v.String())
	}
	return strings.Join(output, "\n")
}
