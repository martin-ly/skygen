package parser

import (
	"regexp"
	"strings"
	"testing"
)

func TestParserScript(t *testing.T) {
	str := trim(`
        EVENT DO
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

func trim(s string) string {
	return regexp.MustCompile(`(?m)^        `).ReplaceAllString(strings.TrimSpace(s), "")
}
