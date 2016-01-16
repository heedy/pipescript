//line parser.y:5
package pipescript

import __yyfmt__ "fmt"

//line parser.y:7
import (
	"fmt"
	"strconv"
)

type scriptFunc struct {
	transform string
	args      []*Script
}

//line parser.y:22
type parserSymType struct {
	yys         int
	script      *Script
	sfunc       scriptFunc
	scriptArray []*Script
	strVal      string // This is how variables are passed in: by their string value
}

const pNUMBER = 57346
const pSTRING = 57347
const pBOOL = 57348
const pIDENTIFIER = 57349
const pAND = 57350
const pOR = 57351
const pNOT = 57352
const pCOMPARISON = 57353
const pPLUS = 57354
const pMINUS = 57355
const pMULTIPLY = 57356
const pDIVIDE = 57357
const pMODULO = 57358
const pPOW = 57359
const pCOMMA = 57360
const pRPARENS = 57361
const pLPARENS = 57362
const pRSQUARE = 57363
const pLSQUARE = 57364
const pRBRACKET = 57365
const pLBRACKET = 57366
const pPIPE = 57367
const pCOLON = 57368
const pUMINUS = 57369

var parserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"pNUMBER",
	"pSTRING",
	"pBOOL",
	"pIDENTIFIER",
	"pAND",
	"pOR",
	"pNOT",
	"pCOMPARISON",
	"pPLUS",
	"pMINUS",
	"pMULTIPLY",
	"pDIVIDE",
	"pMODULO",
	"pPOW",
	"pCOMMA",
	"pRPARENS",
	"pLPARENS",
	"pRSQUARE",
	"pLSQUARE",
	"pRBRACKET",
	"pLBRACKET",
	"pPIPE",
	"pCOLON",
	"pUMINUS",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line parser.y:406

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 42
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 170

var parserAct = [...]int{

	3, 18, 13, 36, 33, 2, 30, 31, 60, 66,
	57, 68, 20, 67, 58, 37, 34, 35, 20, 38,
	39, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 56, 32, 20, 54, 53, 55, 20, 54, 59,
	70, 61, 63, 65, 14, 73, 34, 35, 4, 70,
	62, 64, 75, 70, 69, 22, 23, 8, 21, 28,
	29, 26, 27, 24, 25, 72, 9, 77, 78, 79,
	74, 80, 7, 81, 24, 25, 62, 64, 15, 16,
	17, 19, 1, 0, 5, 0, 0, 6, 28, 29,
	26, 27, 24, 25, 10, 0, 12, 0, 11, 15,
	16, 17, 19, 0, 0, 5, 0, 0, 6, 0,
	26, 27, 24, 25, 0, 40, 0, 42, 0, 41,
	22, 23, 0, 21, 28, 29, 26, 27, 24, 25,
	72, 22, 23, 76, 21, 28, 29, 26, 27, 24,
	25, 72, 71, 22, 23, 0, 21, 28, 29, 26,
	27, 24, 25, 22, 0, 0, 21, 28, 29, 26,
	27, 24, 25, 21, 28, 29, 26, 27, 24, 25,
}
var parserPact = [...]int{

	74, -1000, 8, 135, -1000, 74, 74, 6, -1000, -1000,
	74, 74, 74, -23, 74, -1000, -1000, -1000, 74, 95,
	74, 74, 74, 74, 74, 74, 74, 74, 74, 74,
	76, -1000, 29, 12, -13, -7, 29, 135, 135, 135,
	74, 74, 74, 135, 76, 152, 145, -1000, -1000, 58,
	58, 96, 96, -1000, -1000, -11, -1000, -1000, -1000, -1000,
	35, 123, 22, 47, 31, 112, 74, 74, 74, -1000,
	74, -1000, 74, -1000, -1000, -1000, -1000, 123, 47, 112,
	135, 135,
}
var parserPgo = [...]int{

	0, 82, 4, 72, 66, 0, 2, 57, 48, 44,
	1, 8,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 3, 7, 7, 6, 9,
	9, 9, 10, 10, 10, 10, 10, 10, 10, 11,
	11, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 8, 8, 8, 8, 8, 8, 4,
	4, 4,
}
var parserR2 = [...]int{

	0, 1, 1, 3, 3, 3, 1, 1, 1, 2,
	2, 2, 4, 4, 4, 4, 4, 4, 1, 3,
	3, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1, 1, 3, 3, 3, 1,
	1, 1,
}
var parserChk = [...]int{

	-1000, -1, -2, -5, -8, 10, 13, -3, -7, -4,
	20, 24, 22, -6, -9, 4, 5, 6, -10, 7,
	25, 11, 8, 9, 16, 17, 14, 15, 12, 13,
	-5, -5, 26, -2, -2, -2, 26, -5, -5, -5,
	20, 24, 22, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -6, -10, 7, 19, 23, 21, -6,
	-11, -5, -11, -5, -11, -5, 20, 24, 22, 19,
	18, 19, 18, 23, 23, 21, 21, -5, -5, -5,
	-5, -5,
}
var parserDef = [...]int{

	0, -2, 1, 2, 21, 0, 0, 33, 34, 35,
	0, 0, 0, 6, 7, 39, 40, 41, 8, 18,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	22, 32, 0, 0, 0, 0, 0, 10, 9, 11,
	0, 0, 0, 3, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 5, 8, 18, 36, 37, 38, 4,
	0, 2, 0, 2, 0, 2, 0, 0, 0, 12,
	0, 15, 0, 13, 16, 14, 17, 0, 0, 0,
	19, 20,
}
var parserTok1 = [...]int{

	1,
}
var parserTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27,
}
var parserTok3 = [...]int{
	0,
}

var parserErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	parserDebug        = 0
	parserErrorVerbose = false
)

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

type parserParser interface {
	Parse(parserLexer) int
	Lookahead() int
}

type parserParserImpl struct {
	lookahead func() int
}

func (p *parserParserImpl) Lookahead() int {
	return p.lookahead()
}

func parserNewParser() parserParser {
	p := &parserParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const parserFlag = -1000

func parserTokname(c int) string {
	if c >= 1 && c-1 < len(parserToknames) {
		if parserToknames[c-1] != "" {
			return parserToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !parserErrorVerbose {
		return "syntax error"
	}

	for _, e := range parserErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + parserTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := parserPact[state]
	for tok := TOKSTART; tok-1 < len(parserToknames); tok++ {
		if n := base + tok; n >= 0 && n < parserLast && parserChk[parserAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if parserDef[state] == -2 {
		i := 0
		for parserExca[i] != -1 || parserExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; parserExca[i] >= 0; i += 2 {
			tok := parserExca[i]
			if tok < TOKSTART || parserExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if parserExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += parserTokname(tok)
	}
	return res
}

func parserlex1(lex parserLexer, lval *parserSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		token = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			token = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		token = parserTok3[i+0]
		if token == char {
			token = parserTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(token), uint(char))
	}
	return char, token
}

func parserParse(parserlex parserLexer) int {
	return parserNewParser().Parse(parserlex)
}

func (parserrcvr *parserParserImpl) Parse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	var parserDollar []parserSymType
	_ = parserDollar // silence set and not used
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parsertoken := -1 // parserchar translated into internal numbering
	parserrcvr.lookahead = func() int { return parserchar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		parserstate = -1
		parserchar = -1
		parsertoken = -1
	}()
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parsertoken), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar, parsertoken = parserlex1(parserlex, &parserlval)
	}
	parsern += parsertoken
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parsertoken { /* valid shift */
		parserchar = -1
		parsertoken = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar, parsertoken = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parsertoken {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error(parserErrorMessage(parserstate, parsertoken))
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parsertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parsertoken))
			}
			if parsertoken == parserEofCode {
				goto ret1
			}
			parserchar = -1
			parsertoken = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	// parserp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if parserp+1 >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:53
		{
			parserVAL.script = parserDollar[1].script
			parserlex.(*parserLex).output = parserVAL.script
		}
	case 3:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:68
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 4:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:88
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 5:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:98
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 7:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:117
		{
			v, ok := TransformRegistry[parserDollar[1].sfunc.transform]
			if ok {
				s, err := v.Script(parserDollar[1].sfunc.args)
				if err != nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				parserVAL.script = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found", parserDollar[1].sfunc.transform))
				goto ret1
			}
		}
	case 8:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:135
		{
			v, ok := TransformRegistry[parserDollar[1].sfunc.transform]
			if ok {
				s, err := v.Script(parserDollar[1].sfunc.args)
				if err != nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				parserVAL.script = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found", parserDollar[1].sfunc.transform))
				goto ret1
			}
		}
	case 9:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:153
		{
			/* A simplefunction can have at most one argument if it is to be called in bash style */
			if len(parserDollar[1].sfunc.args) > 1 {
				parserlex.Error("Used both function f(x,y) and bash-style function calling at same time for a single function call. This probably means you made a syntax error")
				goto ret1
			}
			parserVAL.sfunc.transform = parserDollar[1].sfunc.transform
			parserVAL.sfunc.args = append(parserDollar[1].sfunc.args, parserDollar[2].script)
		}
	case 10:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:164
		{
			parserVAL.sfunc.transform = parserDollar[1].sfunc.transform
			parserVAL.sfunc.args = append(parserDollar[1].sfunc.args, parserDollar[2].script)
		}
	case 11:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:170
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[2].script}
		}
	case 12:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:179
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 13:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:185
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 14:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:191
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 15:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:198
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 16:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:204
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 17:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:210
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 18:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:217
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 19:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:225
		{
			parserVAL.scriptArray = append(parserDollar[1].scriptArray, parserDollar[3].script)
		}
	case 20:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:230
		{
			parserVAL.scriptArray = []*Script{parserDollar[1].script, parserDollar[3].script}
		}
	case 22:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:245
		{
			s, err := notScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 23:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:256
		{
			s, err := comparisonScript(parserDollar[2].strVal, parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 24:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:266
		{
			s, err := andScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 25:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:276
		{
			s, err := orScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 26:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:286
		{
			s, err := modScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 27:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:296
		{
			s, err := powScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 28:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:306
		{
			s, err := mulScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 29:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:316
		{
			s, err := divScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 30:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:326
		{
			s, err := addScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 31:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:336
		{
			s, err := subtractScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 32:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:346
		{
			s, err := negativeScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 36:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:368
		{
			parserVAL.script = parserDollar[2].script
		}
	case 37:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:370
		{
			parserVAL.script = parserDollar[2].script
		}
	case 38:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:372
		{
			parserVAL.script = parserDollar[2].script
		}
	case 39:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:382
		{
			num, err := strconv.ParseFloat(parserDollar[1].strVal, 64)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = ConstantScript(num)
		}
	case 40:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:392
		{
			parserVAL.script = ConstantScript(parserDollar[1].strVal)
		}
	case 41:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:397
		{
			if parserDollar[1].strVal == "true" {
				parserVAL.script = ConstantScript(true)
			} else {
				parserVAL.script = ConstantScript(false)
			}
		}
	}
	goto parserstack /* stack new state and value */
}
