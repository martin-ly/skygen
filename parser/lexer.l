%{
package parser

import (
    "bufio"
    "fmt"
    "github.com/skydb/skygen/core"
    "strconv"
)

type yylexer struct{
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
        src: src,
        startToken:startToken,
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
%}

%yyc c
%yyn c = y.getc()

%%
    y.tlineidx, y.tcharidx = y.lineidx, y.charidx
    y.buf = y.buf[:0]

"#"[^\n]*

\"(\\.|[^\\"])*\"          return y.quotedstrtoken(yylval, TSTRING);
\'(\\.|[^\\'])*\'          return y.quotedstrtoken(yylval, TSTRING);
[0-9]+"y"                  return y.intdurationtoken(yylval, TDURATIONYEAR);
[0-9]+"d"                  return y.intdurationtoken(yylval, TDURATIONDAY);
[0-9]+"h"                  return y.intdurationtoken(yylval, TDURATIONHOUR);
[0-9]+"m"                  return y.intdurationtoken(yylval, TDURATIONMINUTE);
[0-9]+"s"                  return y.intdurationtoken(yylval, TDURATIONSECOND);
[0-9]+"%"                  return y.intpercenttoken(yylval, TPERCENT);
[0-9]+                     return y.inttoken(yylval, TINT);

"EVENT"                    return y.token(yylval, TEVENT);
"SET"                      return y.token(yylval, TSET);
"END"                      return y.token(yylval, TEND);
"AFTER"                    return y.token(yylval, TAFTER);
"WEIGHT"                   return y.token(yylval, TWEIGHT);
"PROBABILITY"              return y.token(yylval, TPROBABILITY);
"SCHEMA"                   return y.token(yylval, TSCHEMA);
"PROPERTY"                 return y.token(yylval, TPROPERTY);
"TRANSIENT"                return y.token(yylval, TTRANSIENT);
"EXIT"                     return y.token(yylval, TEXIT);
","                        return y.token(yylval, TCOMMA);
"="                        return y.token(yylval, TEQUALS);
"-"                        return y.token(yylval, TMINUS);

[a-zA-Z_~][a-zA-Z0-9_]*    return y.strtoken(yylval, TIDENT);
[ \t\n\r]+

%%
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
    yylval.str = str[1:len(str)-1]
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
    if yylval.integer, err = strconv.Atoi(string(y.buf[0:len(y.buf)-1])); err != nil {
        panic("strconv failed: " + string(y.buf))
    }
    return y.token(yylval, tok)
}

// Saves the integer in the buffer and the token to the parser value
// and returns the token.
func (y *yylexer) intpercenttoken(yylval *yySymType, tok int) int {
    var err error
    if yylval.integer, err = strconv.Atoi(string(y.buf[0:len(y.buf)-1])); err != nil {
        panic("strconv failed: " + string(y.buf))
    }
    return y.token(yylval, tok)
}
