// Code generated by goyacc -o promql/generated_parser.y.go promql/generated_parser.y. DO NOT EDIT.

//line promql/generated_parser.y:15
package promql

import __yyfmt__ "fmt"

//line promql/generated_parser.y:15

import (
	"github.com/prometheus/prometheus/pkg/labels"
)

//line promql/generated_parser.y:22
type yySymType struct {
	yys      int
	node     Node
	item     Item
	matchers []*labels.Matcher
	matcher  *labels.Matcher
}

const ERROR = 57346
const EOF = 57347
const COMMENT = 57348
const IDENTIFIER = 57349
const METRIC_IDENTIFIER = 57350
const LEFT_PAREN = 57351
const RIGHT_PAREN = 57352
const LEFT_BRACE = 57353
const RIGHT_BRACE = 57354
const LEFT_BRACKET = 57355
const RIGHT_BRACKET = 57356
const COMMA = 57357
const ASSIGN = 57358
const COLON = 57359
const SEMICOLON = 57360
const STRING = 57361
const NUMBER = 57362
const DURATION = 57363
const BLANK = 57364
const TIMES = 57365
const SPACE = 57366
const operatorsStart = 57367
const SUB = 57368
const ADD = 57369
const MUL = 57370
const MOD = 57371
const DIV = 57372
const LAND = 57373
const LOR = 57374
const LUNLESS = 57375
const EQL = 57376
const NEQ = 57377
const LTE = 57378
const LSS = 57379
const GTE = 57380
const GTR = 57381
const EQL_REGEX = 57382
const NEQ_REGEX = 57383
const POW = 57384
const operatorsEnd = 57385
const aggregatorsStart = 57386
const AVG = 57387
const COUNT = 57388
const SUM = 57389
const MIN = 57390
const MAX = 57391
const STDDEV = 57392
const STDVAR = 57393
const TOPK = 57394
const BOTTOMK = 57395
const COUNT_VALUES = 57396
const QUANTILE = 57397
const aggregatorsEnd = 57398
const keywordsStart = 57399
const OFFSET = 57400
const BY = 57401
const WITHOUT = 57402
const ON = 57403
const IGNORING = 57404
const GROUP_LEFT = 57405
const GROUP_RIGHT = 57406
const BOOL = 57407
const keywordsEnd = 57408
const startSymbolsStart = 57409
const START_LABELS = 57410
const startSymbolsEnd = 57411

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ERROR",
	"EOF",
	"COMMENT",
	"IDENTIFIER",
	"METRIC_IDENTIFIER",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"COMMA",
	"ASSIGN",
	"COLON",
	"SEMICOLON",
	"STRING",
	"NUMBER",
	"DURATION",
	"BLANK",
	"TIMES",
	"SPACE",
	"operatorsStart",
	"SUB",
	"ADD",
	"MUL",
	"MOD",
	"DIV",
	"LAND",
	"LOR",
	"LUNLESS",
	"EQL",
	"NEQ",
	"LTE",
	"LSS",
	"GTE",
	"GTR",
	"EQL_REGEX",
	"NEQ_REGEX",
	"POW",
	"operatorsEnd",
	"aggregatorsStart",
	"AVG",
	"COUNT",
	"SUM",
	"MIN",
	"MAX",
	"STDDEV",
	"STDVAR",
	"TOPK",
	"BOTTOMK",
	"COUNT_VALUES",
	"QUANTILE",
	"aggregatorsEnd",
	"keywordsStart",
	"OFFSET",
	"BY",
	"WITHOUT",
	"ON",
	"IGNORING",
	"GROUP_LEFT",
	"GROUP_RIGHT",
	"BOOL",
	"keywordsEnd",
	"startSymbolsStart",
	"START_LABELS",
	"startSymbolsEnd",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line promql/generated_parser.y:155

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 67

var yyAct = [...]int{

	3, 17, 20, 11, 9, 5, 10, 8, 9, 7,
	1, 12, 6, 4, 0, 0, 0, 0, 18, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 14, 0, 0, 0, 0, 15,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 2,
}
var yyPact = [...]int{

	-2, -1000, -6, -1000, -1000, -3, -9, -1000, -1000, -1,
	1, -1000, 0, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 13, 12, 7, 11, 10,
}
var yyR1 = [...]int{

	0, 5, 5, 2, 2, 1, 1, 3, 3, 4,
	4, 4, 4, 4,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 3, 2, 3, 3, 1,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -5, 68, 2, -1, 11, -2, 12, -3, 7,
	15, 12, -4, 34, 35, 40, 41, 2, -3, 19,
	2,
}
var yyDef = [...]int{

	0, -2, 0, 2, 1, 0, 0, 6, 4, 0,
	0, 5, 0, 9, 10, 11, 12, 13, 3, 7,
	8,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:117
		{
			yylex.(*parser).generatedParserResult.(*VectorSelector).LabelMatchers = yyDollar[2].matchers
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:119
		{
			yylex.(*parser).errorf("unknown syntax error after parsing %v", yylex.(*parser).token.desc())
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:125
		{
			yyVAL.matchers = append(yyDollar[1].matchers, yyDollar[3].matcher)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:127
		{
			yyVAL.matchers = []*labels.Matcher{yyDollar[1].matcher}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:132
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:134
		{
			yyVAL.matchers = []*labels.Matcher{}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:140
		{
			yyVAL.matcher = yylex.(*parser).newLabelMatcher(yyDollar[1].item, yyDollar[2].item, yyDollar[3].item)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:142
		{
			yylex.(*parser).errorf("unexpected %v in label matching, expected string", yylex.(*parser).token.desc())
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:146
		{
			yyVAL.item = yyDollar[1].item
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:147
		{
			yyVAL.item = yyDollar[1].item
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:148
		{
			yyVAL.item = yyDollar[1].item
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:149
		{
			yyVAL.item = yyDollar[1].item
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:151
		{
			yylex.(*parser).errorf("expected label matching operator but got %s", yylex.(*parser).token.Val)
		}
	}
	goto yystack /* stack new state and value */
}
