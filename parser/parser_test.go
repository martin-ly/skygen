package parser

import (
	"regexp"
	"strings"
	"testing"
)

func TestParserScript(t *testing.T) {
	str := trim(`
        EVENT DO
          EVENT AFTER 5y4d3h2m1s DO
            
          END
          EVENT AFTER 2s - 3s DO
            
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
