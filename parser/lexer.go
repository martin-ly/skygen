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
	case c == ',':
		goto yystate9
	case c == '-':
		goto yystate10
	case c == '=':
		goto yystate18
	case c == 'A':
		goto yystate19
	case c == 'E':
		goto yystate25
	case c == 'P':
		goto yystate35
	case c == 'S':
		goto yystate51
	case c == 'T':
		goto yystate59
	case c == 'W':
		goto yystate68
	case c == '\'':
		goto yystate6
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate11
	case c >= 'B' && c <= 'D' || c >= 'F' && c <= 'O' || c == 'Q' || c == 'R' || c == 'U' || c == 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate20
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule24
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
	goto yyrule1

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
	case c == '\'':
		goto yystate7
	case c == '\\':
		goto yystate8
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate6
	}

yystate7:
	c = y.getc()
	goto yyrule2

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate6
	}

yystate9:
	c = y.getc()
	goto yyrule20

yystate10:
	c = y.getc()
	goto yyrule22

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyrule9
	case c == '%':
		goto yystate12
	case c == 'd':
		goto yystate13
	case c == 'h':
		goto yystate14
	case c == 'm':
		goto yystate15
	case c == 's':
		goto yystate16
	case c == 'y':
		goto yystate17
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate12:
	c = y.getc()
	goto yyrule8

yystate13:
	c = y.getc()
	goto yyrule4

yystate14:
	c = y.getc()
	goto yyrule5

yystate15:
	c = y.getc()
	goto yyrule6

yystate16:
	c = y.getc()
	goto yyrule7

yystate17:
	c = y.getc()
	goto yyrule3

yystate18:
	c = y.getc()
	goto yyrule21

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'F':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate22
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'R':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'N':
		goto yystate26
	case c == 'V':
		goto yystate28
	case c == 'X':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'U' || c == 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'D':
		goto yystate27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'N':
		goto yystate30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'I':
		goto yystate33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'R':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'O':
		goto yystate37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'B':
		goto yystate38
	case c == 'P':
		goto yystate46
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'A':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate39:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'B':
		goto yystate40
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate40:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'I':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate41:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'L':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate42:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'I':
		goto yystate43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate43:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate44:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'Y':
		goto yystate45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate45:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate46:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate47:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'R':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate48:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate49:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'Y':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate50:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate51:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'C':
		goto yystate52
	case c == 'E':
		goto yystate57
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c == 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate52:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'H':
		goto yystate53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate53:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate54:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'M':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate55:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'A':
		goto yystate56
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate56:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate57:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate58:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate59:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'R':
		goto yystate60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate60:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'A':
		goto yystate61
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate61:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'N':
		goto yystate62
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate62:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'S':
		goto yystate63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate63:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'I':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate64:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate65:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'N':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate66:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate67:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate68:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'E':
		goto yystate69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate69:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'I':
		goto yystate70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate70:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'G':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate71:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'H':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate72:
	c = y.getc()
	switch {
	default:
		goto yyrule23
	case c == 'T':
		goto yystate73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate73:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yyrule1: // \"(\\.|[^\\"])*\"
	{
		return y.quotedstrtoken(yylval, TSTRING)
	}
yyrule2: // \'(\\.|[^\\'])*\'
	{
		return y.quotedstrtoken(yylval, TSTRING)
	}
yyrule3: // [0-9]+"y"
	{
		return y.intdurationtoken(yylval, TDURATIONYEAR)
	}
yyrule4: // [0-9]+"d"
	{
		return y.intdurationtoken(yylval, TDURATIONDAY)
	}
yyrule5: // [0-9]+"h"
	{
		return y.intdurationtoken(yylval, TDURATIONHOUR)
	}
yyrule6: // [0-9]+"m"
	{
		return y.intdurationtoken(yylval, TDURATIONMINUTE)
	}
yyrule7: // [0-9]+"s"
	{
		return y.intdurationtoken(yylval, TDURATIONSECOND)
	}
yyrule8: // [0-9]+"%"
	{
		return y.intpercenttoken(yylval, TPERCENT)
	}
yyrule9: // [0-9]+
	{
		return y.inttoken(yylval, TINT)
	}
yyrule10: // "EVENT"
	{
		return y.token(yylval, TEVENT)
	}
yyrule11: // "SET"
	{
		return y.token(yylval, TSET)
	}
yyrule12: // "END"
	{
		return y.token(yylval, TEND)
	}
yyrule13: // "AFTER"
	{
		return y.token(yylval, TAFTER)
	}
yyrule14: // "WEIGHT"
	{
		return y.token(yylval, TWEIGHT)
	}
yyrule15: // "PROBABILITY"
	{
		return y.token(yylval, TPROBABILITY)
	}
yyrule16: // "SCHEMA"
	{
		return y.token(yylval, TSCHEMA)
	}
yyrule17: // "PROPERTY"
	{
		return y.token(yylval, TPROPERTY)
	}
yyrule18: // "TRANSIENT"
	{
		return y.token(yylval, TTRANSIENT)
	}
yyrule19: // "EXIT"
	{
		return y.token(yylval, TEXIT)
	}
yyrule20: // ","
	{
		return y.token(yylval, TCOMMA)
	}
yyrule21: // "="
	{
		return y.token(yylval, TEQUALS)
	}
yyrule22: // "-"
	{
		return y.token(yylval, TMINUS)
	}
yyrule23: // [a-zA-Z_~][a-zA-Z0-9_]*
	{
		return y.strtoken(yylval, TIDENT)
	}
yyrule24: // [ \t\n\r]+

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
