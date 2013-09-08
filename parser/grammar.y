%{

package parser

import (
    "github.com/skydb/skygen/core"
)

%}

%union{
    token int
    integer int
    str string
    strs []string
    script *core.Script
    event *core.Event
    events core.Events
}

%token <token> TSTARTSCRIPT
%token <token> TEVENT, TDO, TEND
%token <token> TTRUE, TFALSE
%token <str> TIDENT, TSTRING
%token <integer> TINT

%type <script> script
%type <event> event
%type <events> events

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
    TEVENT TDO events TEND
    {
        $$ = core.NewEvent()
        $$.SetEvents($3)
    }
;

%%

