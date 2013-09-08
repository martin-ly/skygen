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
		goto yystate17
	case c == 'A':
		goto yystate18
	case c == 'B' || c == 'C' || c >= 'F' && c <= 'R' || c >= 'T' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate19
	case c == 'D':
		goto yystate24
	case c == 'E':
		goto yystate26
	case c == 'S':
		goto yystate33
	case c == 'W':
		goto yystate36
	case c == '\'':
		goto yystate6
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule19
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
	goto yyrule15

yystate10:
	c = y.getc()
	goto yyrule17

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyrule8
	case c == 'd':
		goto yystate12
	case c == 'h':
		goto yystate13
	case c == 'm':
		goto yystate14
	case c == 's':
		goto yystate15
	case c == 'y':
		goto yystate16
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate12:
	c = y.getc()
	goto yyrule4

yystate13:
	c = y.getc()
	goto yyrule5

yystate14:
	c = y.getc()
	goto yyrule6

yystate15:
	c = y.getc()
	goto yyrule7

yystate16:
	c = y.getc()
	goto yyrule3

yystate17:
	c = y.getc()
	goto yyrule16

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'F':
		goto yystate20
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'T':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'E':
		goto yystate22
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'R':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'O':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'N':
		goto yystate27
	case c == 'V':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'D':
		goto yystate28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'E':
		goto yystate30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'N':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'T':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyrule9
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'E':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'T':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'E':
		goto yystate37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'I':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'G':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate39:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'H':
		goto yystate40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate40:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'T':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
	}

yystate41:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate19
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
yyrule8: // [0-9]+
	{
		return y.inttoken(yylval, TINT)
	}
yyrule9: // "EVENT"
	{
		return y.token(yylval, TEVENT)
	}
yyrule10: // "DO"
	{
		return y.token(yylval, TDO)
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
yyrule15: // ","
	{
		return y.token(yylval, TCOMMA)
	}
yyrule16: // "="
	{
		return y.token(yylval, TEQUALS)
	}
yyrule17: // "-"
	{
		return y.token(yylval, TMINUS)
	}
yyrule18: // [a-zA-Z_~][a-zA-Z0-9_]*
	{
		return y.strtoken(yylval, TIDENT)
	}
yyrule19: // [ \t\n\r]+

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
