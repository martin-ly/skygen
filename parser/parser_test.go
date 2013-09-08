package parser

import (
	"regexp"
	"strings"
	"testing"
)

func TestParserNestedEvents(t *testing.T) {
	str := trim(`
        EVENT DO
          EVENT DO
          END
          EVENT DO
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
        EVENT DO
          EVENT AFTER 5y4d3h2m1s DO
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
        EVENT DO
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

func trim(s string) string {
	return regexp.MustCompile(`(?m)^        `).ReplaceAllString(strings.TrimSpace(s), "")
}
