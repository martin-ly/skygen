package parser

import (
	"regexp"
	"strings"
	"testing"
)

func TestParserNestedEvents(t *testing.T) {
	str := trim(`
        EVENT
          EVENT
          END
          EVENT
          END
        END
    `)
	script, err := New().ParseString(str)
	if err != nil {
		t.Fatal("Parse error:", err)
	}
	if script.String() != str {
		t.Fatal("Unexpected:", "'"+script.String()+"'")
	}
}

func TestParserAfter(t *testing.T) {
	str := trim(`
        EVENT
          EVENT AFTER 5y4d3h2m1s
          END
        END
    `)
	script, err := New().ParseString(str)
	if err != nil {
		t.Fatal("Parse error:", err)
	}
	if script.String() != str {
		t.Fatal("Unexpected:", "'"+script.String()+"'")
	}
}

func TestParserValueSets(t *testing.T) {
	str := trim(`
        EVENT EXIT PROBABILITY 30%
          SET foo = "bar"
        END
    `)
	script, err := New().ParseString(str)
	if err != nil {
		t.Fatal("Parse error:", err)
	}
	if script.String() != str {
		t.Fatal("Unexpected:", "'"+script.String()+"'")
	}
}

func TestParserSchema(t *testing.T) {
	str := trim(`
        SCHEMA
          PROPERTY name STRING
          PROPERTY count INTEGER TRANSIENT
        END
    `)
	script, err := New().ParseString(str)
	if err != nil {
		t.Fatal("Parse error:", err)
	}
	if script.String() != str {
		t.Fatal("Unexpected:", "'"+script.String()+"'")
	}
}

func trim(s string) string {
	return regexp.MustCompile(`(?m)^        `).ReplaceAllString(strings.TrimSpace(s), "")
}
