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
    value_sets core.ValueSets
    value_set *core.ValueSet
    key_values map[string]interface{}
    key_value key_value
}

%token <token> TSTARTSCRIPT
%token <token> TEVENT, TDO, TEND, TAFTER, TWEIGHT, TSET, TPROBABILITY
%token <token> TTRUE, TFALSE
%token <token> TMINUS, TCOMMA, TEQUALS
%token <str> TIDENT, TSTRING
%token <integer> TDURATIONYEAR, TDURATIONDAY, TDURATIONHOUR
%token <integer> TDURATIONMINUTE, TDURATIONSECOND
%token <integer> TINT, TPERCENT

%type <script> script
%type <event> event
%type <events> events
%type <integer> event_weight
%type <duration_range> event_after
%type <duration> duration_year, duration_day, duration_hour
%type <duration> duration_minute, duration_second
%type <duration> duration

%type <value_sets> value_sets
%type <value_set> value_set
%type <key_values> key_values
%type <key_value> key_value

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
    TEVENT event_after event_weight TDO value_sets events TEND
    {
        $$ = core.NewEvent()
        $$.After = $2
        $$.Weight = $3
        $$.SetValueSets($5)
        $$.SetEvents($6)
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

value_sets :
    /* empty */
    {
        $$ = make(core.ValueSets, 0)
    }
|   value_sets value_set
    {
        $$ = append($1, $2)
    }
;

value_set :
    TSET key_values value_set_probability
    {
        $$ = core.NewValueSet()
        $$.Values = $2
    }
;

value_set_probability :
    /* empty */
    {
        $$ = 1.0
    }
|   TPROBABILITY TPERCENT
    {
        $$ = $2
    }
;

key_values :
    /* empty */
    {
        $$ = make(map[string]interface{})
    }
|   key_value
    {
        $$ = make(map[string]interface{})
        $$[$1.key] = $1.value
    }
|   key_values TCOMMA key_value
    {
        $1[$3.key] = $3.value
    }
;

key_value :
    TIDENT TEQUALS TSTRING
    {
        $$.key = $1
        $$.value = $3
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

type key_value struct {
    key string
    value interface{}
}
