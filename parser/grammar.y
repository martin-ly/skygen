%{

package parser

import (
    "github.com/skydb/skygen/core"
    "time"
)

%}

%union{
    token int
    integer int
    duration time.Duration
    duration_range []time.Duration
    str string
    strs []string
    script *core.Script
    event *core.Event
    events core.Events
}

%token <token> TSTARTSCRIPT
%token <token> TEVENT, TDO, TEND, TAFTER, TWEIGHT
%token <token> TTRUE, TFALSE
%token <token> TMINUS
%token <str> TIDENT, TSTRING
%token <integer> TDURATIONYEAR, TDURATIONDAY, TDURATIONHOUR
%token <integer> TDURATIONMINUTE, TDURATIONSECOND
%token <integer> TINT

%type <script> script
%type <event> event
%type <events> events
%type <integer> event_weight
%type <duration_range> event_after
%type <duration> duration_year, duration_day, duration_hour
%type <duration> duration_minute, duration_second
%type <duration> duration

%start start

%%

start :
    TSTARTSCRIPT script
    {
        l := yylex.(*yylexer)
        l.script = $2
    }
;

script :
    events
    {
        $$ = core.NewScript()
        $$.SetEvents($1)
    }
;

events :
    /* empty */
    {
        $$ = make(core.Events, 0)
    }
|   events event
    {
        $$ = append($1, $2)
    }
;

event :
    TEVENT event_after event_weight TDO events TEND
    {
        $$ = core.NewEvent()
        $$.After = $2
        $$.Weight = $3
        $$.SetEvents($5)
    }
;

event_after :
    /* empty */
    {
        $$ = []time.Duration{}
    }
|   TAFTER duration TMINUS duration
    {
        $$ = []time.Duration{$2, $4}
    }
|   TAFTER duration
    {
        $$ = []time.Duration{$2, $2}
    }
;

event_weight :
    /* empty */
    {
        $$ = 1
    }
|   TWEIGHT TINT
    {
        $$ = $2
    }
;

duration :
    duration_year duration_day duration_hour duration_minute duration_second
    {
        $$ = $1 + $2 + $3 + $4 + $5
    }
;

duration_year :
    /* empty */   { $$ = 0 * time.Hour }
|   TDURATIONYEAR { $$ = time.Duration($1 * 24 * 365) * time.Hour }
;

duration_day :
    /* empty */   { $$ = 0 * time.Hour }
|   TDURATIONDAY  { $$ = time.Duration($1 * 24) * time.Hour }
;

duration_hour :
    /* empty */    { $$ = 0 * time.Hour }
|   TDURATIONHOUR  { $$ = time.Duration($1) * time.Hour }
;

duration_minute :
    /* empty */      { $$ = 0 * time.Minute }
|   TDURATIONMINUTE  { $$ = time.Duration($1) * time.Minute }
;

duration_second :
    /* empty */      { $$ = 0 * time.Second }
|   TDURATIONSECOND  { $$ = time.Duration($1) * time.Second }
;

%%

