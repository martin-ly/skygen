package parser

import (
	"bufio"
	"fmt"
	"github.com/skydb/skygen/core"
	"strconv"
)

type yylexer struct {
	src        *bufio.Reader
	buf        []byte
	empty      bool
	current    byte
	index      int
	lineidx    int
	charidx    int
	tlineidx   int
	tcharidx   int
	startToken int
	err        error
	script     *core.Script
}

func newLexer(src *bufio.Reader, startToken int) *yylexer {
	y := &yylexer{
		src:        src,
		startToken: startToken,
	}
	y.current, _ = src.ReadByte()
	return y
}

func (y *yylexer) getc() byte {
	var err error
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}

	if y.current, err = y.src.ReadByte(); err == nil {
		y.index++
		y.charidx++

		// Reset line and character index at "\n"
		if y.current == 10 {
			y.lineidx++
			y.charidx = 0
		}
	}
	return y.current
}

func (y *yylexer) Error(e string) {
	y.err = fmt.Errorf("Unexpected '%s' at line %d, char %d, %s", y.buf, y.tlineidx+1, y.tcharidx+1, e)
}

func (y *yylexer) Lex(yylval *yySymType) int {
	if y.startToken != 0 {
		token := y.startToken
		y.startToken = 0
		return token
	}
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.tlineidx, y.tcharidx = y.lineidx, y.charidx
	y.buf = y.buf[:0]

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate3
	case c == '#':
		goto yystate6
	case c == ',':
		goto yystate11
	case c == '-':
		goto yystate12
	case c == '=':
		goto yystate20
	case c == 'A':
		goto yystate21
	case c == 'E':
		goto yystate27
	case c == 'P':
		goto yystate37
	case c == 'S':
		goto yystate53
	case c == 'T':
		goto yystate61
	case c == 'W':
		goto yystate70
	case c == '\'':
		goto yystate8
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate13
	case c >= 'B' && c <= 'D' || c >= 'F' && c <= 'O' || c == 'Q' || c == 'R' || c == 'U' || c == 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate22
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule25
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate4
	case c == '\\':
		goto yystate5
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate3
	}

yystate4:
	c = y.getc()
	goto yyrule2

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate3
	}

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate7
	}

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate9
	case c == '\\':
		goto yystate10
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate8
	}

yystate9:
	c = y.getc()
	goto yyrule3

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate8
	}

yystate11:
	c = y.getc()
	goto yyrule21

yystate12:
	c = y.getc()
	goto yyrule23

yystate13:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c == '%':
		goto yystate14
	case c == 'd':
		goto yystate15
	case c == 'h':
		goto yystate16
	case c == 'm':
		goto yystate17
	case c == 's':
		goto yystate18
	case c == 'y':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate13
	}

yystate14:
	c = y.getc()
	goto yyrule9

yystate15:
	c = y.getc()
	goto yyrule5

yystate16:
	c = y.getc()
	goto yyrule6

yystate17:
	c = y.getc()
	goto yyrule7

yystate18:
	c = y.getc()
	goto yyrule8

yystate19:
	c = y.getc()
	goto yyrule4

yystate20:
	c = y.getc()
	goto yyrule22

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'F':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'R':
		goto yystate26
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'N':
		goto yystate28
	case c == 'V':
		goto yystate30
	case c == 'X':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'U' || c == 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'D':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'N':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'I':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule20
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'R':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'O':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate39:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'B':
		goto yystate40
	case c == 'P':
		goto yystate48
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate40:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'A':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate41:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'B':
		goto yystate42
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate42:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'I':
		goto yystate43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate43:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'L':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate44:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'I':
		goto yystate45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate45:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate46:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'Y':
		goto yystate47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate47:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate48:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate49:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'R':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate50:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate51:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'Y':
		goto yystate52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate52:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate53:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'C':
		goto yystate54
	case c == 'E':
		goto yystate59
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c == 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate54:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'H':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate55:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate56:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'M':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate57:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'A':
		goto yystate58
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate58:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate59:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate60:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate61:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'R':
		goto yystate62
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate62:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'A':
		goto yystate63
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate63:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'N':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate64:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'S':
		goto yystate65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate65:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'I':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate66:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate67:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'N':
		goto yystate68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate68:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate69:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate70:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'E':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate71:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'I':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate72:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'G':
		goto yystate73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate73:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'H':
		goto yystate74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate74:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == 'T':
		goto yystate75
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yystate75:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate22
	}

yyrule1: // "#"[^\n]+

	goto yystate0
yyrule2: // \"(\\.|[^\\"])*\"
	{
		return y.quotedstrtoken(yylval, TSTRING)
	}
yyrule3: // \'(\\.|[^\\'])*\'
	{
		return y.quotedstrtoken(yylval, TSTRING)
	}
yyrule4: // [0-9]+"y"
	{
		return y.intdurationtoken(yylval, TDURATIONYEAR)
	}
yyrule5: // [0-9]+"d"
	{
		return y.intdurationtoken(yylval, TDURATIONDAY)
	}
yyrule6: // [0-9]+"h"
	{
		return y.intdurationtoken(yylval, TDURATIONHOUR)
	}
yyrule7: // [0-9]+"m"
	{
		return y.intdurationtoken(yylval, TDURATIONMINUTE)
	}
yyrule8: // [0-9]+"s"
	{
		return y.intdurationtoken(yylval, TDURATIONSECOND)
	}
yyrule9: // [0-9]+"%"
	{
		return y.intpercenttoken(yylval, TPERCENT)
	}
yyrule10: // [0-9]+
	{
		return y.inttoken(yylval, TINT)
	}
yyrule11: // "EVENT"
	{
		return y.token(yylval, TEVENT)
	}
yyrule12: // "SET"
	{
		return y.token(yylval, TSET)
	}
yyrule13: // "END"
	{
		return y.token(yylval, TEND)
	}
yyrule14: // "AFTER"
	{
		return y.token(yylval, TAFTER)
	}
yyrule15: // "WEIGHT"
	{
		return y.token(yylval, TWEIGHT)
	}
yyrule16: // "PROBABILITY"
	{
		return y.token(yylval, TPROBABILITY)
	}
yyrule17: // "SCHEMA"
	{
		return y.token(yylval, TSCHEMA)
	}
yyrule18: // "PROPERTY"
	{
		return y.token(yylval, TPROPERTY)
	}
yyrule19: // "TRANSIENT"
	{
		return y.token(yylval, TTRANSIENT)
	}
yyrule20: // "EXIT"
	{
		return y.token(yylval, TEXIT)
	}
yyrule21: // ","
	{
		return y.token(yylval, TCOMMA)
	}
yyrule22: // "="
	{
		return y.token(yylval, TEQUALS)
	}
yyrule23: // "-"
	{
		return y.token(yylval, TMINUS)
	}
yyrule24: // [a-zA-Z_~][a-zA-Z0-9_]*
	{
		return y.strtoken(yylval, TIDENT)
	}
yyrule25: // [ \t\n\r]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}

// Saves the token to the parser value and returns the token.
func (y *yylexer) token(yylval *yySymType, tok int) int {
	yylval.token = tok
	return tok
}

// Saves the string in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) strtoken(yylval *yySymType, tok int) int {
	yylval.str = string(y.buf)
	return y.token(yylval, tok)
}

// Saves the quoted string in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) quotedstrtoken(yylval *yySymType, tok int) int {
	str := string(y.buf)
	yylval.str = str[1 : len(str)-1]
	return y.token(yylval, tok)
}

// Saves the integer in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) inttoken(yylval *yySymType, tok int) int {
	var err error
	if yylval.integer, err = strconv.Atoi(string(y.buf)); err != nil {
		panic("strconv failed: " + string(y.buf))
	}
	return y.token(yylval, tok)
}

// Saves the integer in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) intdurationtoken(yylval *yySymType, tok int) int {
	var err error
	if yylval.integer, err = strconv.Atoi(string(y.buf[0 : len(y.buf)-1])); err != nil {
		panic("strconv failed: " + string(y.buf))
	}
	return y.token(yylval, tok)
}

// Saves the integer in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) intpercenttoken(yylval *yySymType, tok int) int {
	var err error
	if yylval.integer, err = strconv.Atoi(string(y.buf[0 : len(y.buf)-1])); err != nil {
		panic("strconv failed: " + string(y.buf))
	}
	return y.token(yylval, tok)
}
