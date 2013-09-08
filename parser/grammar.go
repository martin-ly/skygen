//line grammar.y:2
package parser

import __yyfmt__ "fmt"

//line grammar.y:3
import (
	"github.com/skydb/skygen/core"
	"strings"
	"time"
)

//line grammar.y:13
type yySymType struct {
	yys            int
	token          int
	integer        int
	boolean        bool
	duration       time.Duration
	duration_range []time.Duration
	str            string
	strs           []string
	script         *core.Script
	schema         *core.Schema
	property       *core.Property
	properties     core.Properties
	event          *core.Event
	events         core.Events
	value_sets     core.ValueSets
	value_set      *core.ValueSet
	key_values     map[string]interface{}
	key_value      key_value
}

const TSTARTSCRIPT = 57346
const TEVENT = 57347
const TDO = 57348
const TEND = 57349
const TAFTER = 57350
const TWEIGHT = 57351
const TSET = 57352
const TPROBABILITY = 57353
const TSCHEMA = 57354
const TPROPERTY = 57355
const TTRANSIENT = 57356
const TTRUE = 57357
const TFALSE = 57358
const TMINUS = 57359
const TCOMMA = 57360
const TEQUALS = 57361
const TIDENT = 57362
const TSTRING = 57363
const TDURATIONYEAR = 57364
const TDURATIONDAY = 57365
const TDURATIONHOUR = 57366
const TDURATIONMINUTE = 57367
const TDURATIONSECOND = 57368
const TINT = 57369
const TPERCENT = 57370

var yyToknames = []string{
	"TSTARTSCRIPT",
	"TEVENT",
	"TDO",
	"TEND",
	"TAFTER",
	"TWEIGHT",
	"TSET",
	"TPROBABILITY",
	"TSCHEMA",
	"TPROPERTY",
	"TTRANSIENT",
	"TTRUE",
	"TFALSE",
	"TMINUS",
	"TCOMMA",
	"TEQUALS",
	"TIDENT",
	"TSTRING",
	"TDURATIONYEAR",
	"TDURATIONDAY",
	"TDURATIONHOUR",
	"TDURATIONMINUTE",
	"TDURATIONSECOND",
	"TINT",
	"TPERCENT",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line grammar.y:257
type key_value struct {
	key   string
	value interface{}
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 38
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 51

var yyAct = []int{

	40, 6, 49, 23, 43, 37, 32, 18, 26, 20,
	50, 41, 21, 15, 47, 46, 24, 28, 10, 5,
	35, 17, 45, 9, 12, 38, 22, 14, 9, 2,
	1, 33, 30, 39, 44, 34, 29, 42, 36, 31,
	25, 19, 13, 16, 8, 27, 48, 7, 11, 4,
	3,
}
var yyPact = []int{

	25, -1000, 7, -1000, -1000, -1000, 23, 11, -1000, 19,
	-1000, -1000, -7, 12, -13, -8, 20, -24, -1, -15,
	-1000, 3, -1000, -1000, -13, -18, -1000, -1000, -1000, 10,
	-1000, -20, -1000, 18, -1000, -9, -22, -1000, -1000, 4,
	-1000, -5, -1000, -1000, -1000, -9, -26, -11, -1000, -1000,
	-1000,
}
var yyPgo = []int{

	0, 50, 49, 48, 47, 45, 44, 1, 43, 42,
	41, 40, 39, 38, 37, 7, 36, 35, 34, 33,
	0, 30,
}
var yyR1 = []int{

	0, 21, 1, 2, 2, 4, 4, 3, 5, 5,
	7, 7, 6, 9, 9, 9, 8, 8, 16, 16,
	17, 18, 18, 19, 19, 19, 20, 15, 10, 10,
	11, 11, 12, 12, 13, 13, 14, 14,
}
var yyR2 = []int{

	0, 2, 2, 0, 3, 0, 2, 4, 0, 1,
	0, 2, 7, 0, 4, 2, 0, 2, 0, 2,
	3, 0, 2, 0, 1, 3, 3, 5, 0, 1,
	0, 1, 0, 1, 0, 1, 0, 1,
}
var yyChk = []int{

	-1000, -21, 4, -1, -2, 12, -7, -4, -6, 5,
	7, -3, 13, -9, 8, 20, -8, 9, -15, -10,
	22, 20, 6, 27, 17, -11, 23, -5, 14, -16,
	-15, -12, 24, -7, -17, 10, -13, 25, 7, -19,
	-20, 20, -14, 26, -18, 18, 11, 19, -20, 28,
	21,
}
var yyDef = []int{

	0, -2, 3, 1, 10, 5, 2, 0, 11, 13,
	4, 6, 0, 16, 28, 0, 0, 0, 15, 30,
	29, 8, 18, 17, 28, 32, 31, 7, 9, 10,
	14, 34, 33, 0, 19, 23, 36, 35, 12, 21,
	24, 0, 27, 37, 20, 0, 0, 0, 25, 22,
	26,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28,
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
		//line grammar.y:68
		{
			l := yylex.(*yylexer)
			l.script = yyS[yypt-0].script
		}
	case 2:
		//line grammar.y:76
		{
			yyVAL.script = core.NewScript()
			yyVAL.script.SetSchema(yyS[yypt-1].schema)
			yyVAL.script.SetEvents(yyS[yypt-0].events)
		}
	case 3:
		//line grammar.y:85
		{
			yyVAL.schema = nil
		}
	case 4:
		//line grammar.y:89
		{
			yyVAL.schema = core.NewSchema()
			yyVAL.schema.Properties = yyS[yypt-1].properties
		}
	case 5:
		//line grammar.y:97
		{
			yyVAL.properties = make(core.Properties, 0)
		}
	case 6:
		//line grammar.y:101
		{
			yyVAL.properties = append(yyS[yypt-1].properties, yyS[yypt-0].property)
		}
	case 7:
		//line grammar.y:108
		{
			yyVAL.property = core.NewProperty()
			yyVAL.property.Name = yyS[yypt-2].str
			yyVAL.property.DataType = strings.ToLower(yyS[yypt-1].str)
			yyVAL.property.Transient = yyS[yypt-0].boolean
		}
	case 8:
		//line grammar.y:117
		{
			yyVAL.boolean = false
		}
	case 9:
		//line grammar.y:118
		{
			yyVAL.boolean = true
		}
	case 10:
		//line grammar.y:123
		{
			yyVAL.events = make(core.Events, 0)
		}
	case 11:
		//line grammar.y:127
		{
			yyVAL.events = append(yyS[yypt-1].events, yyS[yypt-0].event)
		}
	case 12:
		//line grammar.y:134
		{
			yyVAL.event = core.NewEvent()
			yyVAL.event.After = yyS[yypt-5].duration_range
			yyVAL.event.Weight = yyS[yypt-4].integer
			yyVAL.event.SetValueSets(yyS[yypt-2].value_sets)
			yyVAL.event.SetEvents(yyS[yypt-1].events)
		}
	case 13:
		//line grammar.y:145
		{
			yyVAL.duration_range = []time.Duration{}
		}
	case 14:
		//line grammar.y:149
		{
			yyVAL.duration_range = []time.Duration{yyS[yypt-2].duration, yyS[yypt-0].duration}
		}
	case 15:
		//line grammar.y:153
		{
			yyVAL.duration_range = []time.Duration{yyS[yypt-0].duration, yyS[yypt-0].duration}
		}
	case 16:
		//line grammar.y:160
		{
			yyVAL.integer = 1
		}
	case 17:
		//line grammar.y:164
		{
			yyVAL.integer = yyS[yypt-0].integer
		}
	case 18:
		//line grammar.y:171
		{
			yyVAL.value_sets = make(core.ValueSets, 0)
		}
	case 19:
		//line grammar.y:175
		{
			yyVAL.value_sets = append(yyS[yypt-1].value_sets, yyS[yypt-0].value_set)
		}
	case 20:
		//line grammar.y:182
		{
			yyVAL.value_set = core.NewValueSet()
			yyVAL.value_set.Values = yyS[yypt-1].key_values
			yyVAL.value_set.Probability = yyS[yypt-0].integer
		}
	case 21:
		//line grammar.y:191
		{
			yyVAL.integer = 1.0
		}
	case 22:
		//line grammar.y:195
		{
			yyVAL.integer = yyS[yypt-0].integer
		}
	case 23:
		//line grammar.y:202
		{
			yyVAL.key_values = make(map[string]interface{})
		}
	case 24:
		//line grammar.y:206
		{
			yyVAL.key_values = make(map[string]interface{})
			yyVAL.key_values[yyS[yypt-0].key_value.key] = yyS[yypt-0].key_value.value
		}
	case 25:
		//line grammar.y:211
		{
			yyS[yypt-2].key_values[yyS[yypt-0].key_value.key] = yyS[yypt-0].key_value.value
		}
	case 26:
		//line grammar.y:218
		{
			yyVAL.key_value.key = yyS[yypt-2].str
			yyVAL.key_value.value = yyS[yypt-0].str
		}
	case 27:
		//line grammar.y:227
		{
			yyVAL.duration = yyS[yypt-4].duration + yyS[yypt-3].duration + yyS[yypt-2].duration + yyS[yypt-1].duration + yyS[yypt-0].duration
		}
	case 28:
		//line grammar.y:233
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 29:
		//line grammar.y:234
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer*24*365) * time.Hour
		}
	case 30:
		//line grammar.y:238
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 31:
		//line grammar.y:239
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer*24) * time.Hour
		}
	case 32:
		//line grammar.y:243
		{
			yyVAL.duration = 0 * time.Hour
		}
	case 33:
		//line grammar.y:244
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Hour
		}
	case 34:
		//line grammar.y:248
		{
			yyVAL.duration = 0 * time.Minute
		}
	case 35:
		//line grammar.y:249
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Minute
		}
	case 36:
		//line grammar.y:253
		{
			yyVAL.duration = 0 * time.Second
		}
	case 37:
		//line grammar.y:254
		{
			yyVAL.duration = time.Duration(yyS[yypt-0].integer) * time.Second
		}
	}
	goto yystack /* stack new state and value */
}
