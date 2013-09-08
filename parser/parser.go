package parser

import (
	"bufio"
	"bytes"
	"github.com/skydb/skygen/core"
	"io"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(r io.Reader) (*core.Script, error) {
	l := newLexer(bufio.NewReader(r), TSTARTSCRIPT)
	yyParse(l)
	return l.script, l.err
}

func (p *Parser) ParseString(s string) (*core.Script, error) {
	return p.Parse(bytes.NewBufferString(s))
}
