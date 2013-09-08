//line grammar.y:2
package parser

import __yyfmt__ "fmt"

//line grammar.y:3
import (
	"github.com/skydb/skygen/core"
	"time"
)

//line grammar.y:12
type yySymType struct {
	yys            int
	token          int
	integer        int
	duration       time.Duration
	duration_range []time.Duration
	str            string
	strs           []string
	script         *core.Script
	event          *core.Event
	events         core.Events
}

const TSTARTSCRIPT = 57346
const TEVENT = 57347
const TDO = 57348
const TEND = 57349
const TAFTER = 57350
const TWEIGHT = 57351
const TTRUE = 57352
const TFALSE = 57353
const TMINUS = 57354
const TIDENT = 57355
const TSTRING = 57356
const TDURATIONYEAR = 57357
const TDURATIONDAY = 57358
const TDURATIONHOUR = 57359
const TDURATIONMINUTE = 57360
const TDURATIONSECOND = 57361
const TINT = 57362

var yyToknames = []string{
	"TSTARTSCRIPT",
	"TEVENT",
	"TDO",
	"TEND",
	"TAFTER",
	"TWEIGHT",
	"TTRUE",
	"TFALSE",
	"TMINUS",
	"TIDENT",
	"TSTRING",
	"TDURATIONYEAR",
	"TDURATIONDAY",
	"TDURATIONHOUR",
	"TDURATIONMINUTE",
	"TDURATIONSECOND",
	"TINT",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line grammar.y:141

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 22
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 28

var yyAct = []int{

	15, 27, 25, 22, 11, 18, 13, 4, 16, 10,
	6, 8, 23, 14, 6, 2, 1, 26, 24, 21,
	17, 20, 19, 12, 7, 9, 5, 3,
}
var yyPact = []int{

	11, -1000, -1000, -1000, 9, -1000, 3, 0, -9, 7,
	-20, -4, -11, -1000, -1000, -1000, -9, -14, -1000, 5,
	-1000, -16, -1000, -1000, -18, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 27, 26, 7, 25, 24, 23, 20, 19, 18,
	17, 4, 16,
}
var yyR1 = []int{

	0, 12, 1, 3, 3, 2, 5, 5, 5, 4,
	4, 11, 6, 6, 7, 7, 8, 8, 9, 9,
	10, 10,
}
var yyR2 = []int{

	0, 2, 1, 0, 2, 6, 0, 4, 2, 0,
	2, 5, 0, 1, 0, 1, 0, 1, 0, 1,
	0, 1,
}
var yyChk = []int{

	-1000, -12, 4, -1, -3, -2, 5, -5, 8, -4,
	9, -11, -6, 15, 6, 20, 12, -7, 16, -3,
	-11, -8, 17, 7, -9, 18, -10, 19,
}
var yyDef = []int{

	0, -2, 3, 1, 2, 4, 6, 9, 12, 0,
	0, 8, 14, 13, 3, 10, 12, 16, 15, 0,
	7, 18, 17, 5, 20, 19, 11, 21,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line grammar.y:48
		{
			l := yylex.(*yylexer)
			l.script = yyS[yypt-0].script
		}
	case 2:
		//line grammar.y:56
		{
			yyVAL.script = core.NewScript()
			yyVAL.script.SetEvents(yyS[yypt-0].events)
		}
	case 3:
		//line grammar.y:64
		{
			yyVAL.events = make(core.Events, 0)
		}
	case 4:
		//line grammar.y:68
		{
			yyVAL.events = append(yyS[yypt-1].events, yyS[yypt-0].event)
		}
	case 5:
		//line grammar.y:75
		{
			yyVAL.event = core.NewEvent()
			yyVAL.event.After = yyS[yypt-4].duration_range
			yyVAL.event.Weight = yyS[yypt-3].integer
			yyVAL.event.SetEvents(yyS[yypt-1].events)
		}
	case 6:
		//line grammar.y:85
		{
			yyVAL.duration_range = []time.Duration{}
		}
	case 7:
		//line grammar.y:89
		{
			yyVAL.duration_range = []time.Duration{yyS[yypt-2].duration, yyS[yypt-0].duration}
		}
	case 8:
		//line grammar.y:93
		{
			yyVAL.duration_range = []time.Duration{yyS[yypt-0].duration, yyS[yypt-0].duration}
		}
	case 9:
		//line grammar.y:100
		{
			yyVAL.integer = 1
		}
	case 10:
		//line grammar.y:104
		{
			yyVAL.integer = yyS[yypt-0].integer
		}
	case 11:
		//line grammar.y:111
		{
			yyVAL.duration = yyS[yypt-4].duration + yyS[yypt-3].duration + yyS[yypt-2].duration + yyS[yypt-1].duration + yyS[yypt-0].duration
		}
	case 12:
		//line grammar.y:117
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 13:
		//line grammar.y:118
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer*24*365) * time.Hour
		}
	case 14:
		//line grammar.y:122
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 15:
		//line grammar.y:123
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer*24) * time.Hour
		}
	case 16:
		//line grammar.y:127
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 17:
		//line grammar.y:128
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Hour
		}
	case 18:
		//line grammar.y:132
		{
			yyVAL.duration = 0 * time.Minute
		}
	case 19:
		//line grammar.y:133
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Minute
		}
	case 20:
		//line grammar.y:137
		{
			yyVAL.duration = 0 * time.Second
		}
	case 21:
		//line grammar.y:138
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Second
		}
	}
	goto yystack /* stack new state and value */
}
