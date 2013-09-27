package core

import (
	"strings"
)

type Properties []*Property

func (s Properties) String() string {
	output := []string{}
	for _, p := range s {
		output = append(output, p.String())
	}
	return strings.Join(output, "\n")
}
