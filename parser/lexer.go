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
	case c == '-':
		goto yystate9
	case c == 'A':
		goto yystate16
	case c == 'B' || c == 'C' || c >= 'F' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '~':
		goto yystate17
	case c == 'D':
		goto yystate22
	case c == 'E':
		goto yystate24
	case c == 'W':
		goto yystate31
	case c == '\'':
		goto yystate6
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate10
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule16
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
	goto yyrule14

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyrule8
	case c == 'd':
		goto yystate11
	case c == 'h':
		goto yystate12
	case c == 'm':
		goto yystate13
	case c == 's':
		goto yystate14
	case c == 'y':
		goto yystate15
	case c >= '0' && c <= '9':
		goto yystate10
	}

yystate11:
	c = y.getc()
	goto yyrule4

yystate12:
	c = y.getc()
	goto yyrule5

yystate13:
	c = y.getc()
	goto yyrule6

yystate14:
	c = y.getc()
	goto yyrule7

yystate15:
	c = y.getc()
	goto yyrule3

yystate16:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'F':
		goto yystate18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate17:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'T':
		goto yystate19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'E':
		goto yystate20
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'R':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'O':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'N':
		goto yystate25
	case c == 'V':
		goto yystate27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'D':
		goto yystate26
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'E':
		goto yystate28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'N':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'T':
		goto yystate30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule9
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'E':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'I':
		goto yystate33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'G':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'H':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == 'T':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate17
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
yyrule11: // "END"
	{
		return y.token(yylval, TEND)
	}
yyrule12: // "AFTER"
	{
		return y.token(yylval, TAFTER)
	}
yyrule13: // "WEIGHT"
	{
		return y.token(yylval, TWEIGHT)
	}
yyrule14: // "-"
	{
		return y.token(yylval, TMINUS)
	}
yyrule15: // [a-zA-Z_~][a-zA-Z0-9_]*
	{
		return y.strtoken(yylval, TIDENT)
	}
yyrule16: // [ \t\n\r]+

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
