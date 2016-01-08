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

//line parser.y:298

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 32
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 169

var parserAct = [...]int{

	51, 2, 16, 50, 27, 28, 29, 9, 55, 6,
	30, 31, 32, 56, 55, 54, 7, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 25, 26,
	8, 24, 22, 23, 20, 21, 18, 19, 52, 53,
	1, 49, 18, 19, 0, 17, 16, 33, 0, 35,
	55, 34, 16, 57, 25, 26, 58, 24, 22, 23,
	20, 21, 18, 19, 0, 0, 0, 0, 0, 48,
	0, 17, 16, 25, 26, 0, 24, 22, 23, 20,
	21, 18, 19, 0, 47, 0, 0, 0, 0, 0,
	17, 16, 25, 26, 0, 24, 22, 23, 20, 21,
	18, 19, 0, 0, 0, 0, 0, 0, 0, 17,
	16, 25, 26, 0, 24, 22, 23, 20, 21, 18,
	19, 24, 22, 23, 20, 21, 18, 19, 0, 16,
	0, 12, 13, 14, 15, 0, 16, 11, 0, 0,
	10, 22, 23, 20, 21, 18, 19, 3, 0, 5,
	0, 4, 0, 0, 0, 16, 20, 21, 18, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 16,
}
var parserPact = [...]int{

	127, -1000, 84, 127, 127, 127, -1000, -1000, -1000, 127,
	127, 127, -1000, -1000, -1000, 27, 127, 127, 127, 127,
	127, 127, 127, 127, 127, 127, 127, 65, 46, 20,
	84, -1000, 129, 127, 127, 127, -1000, 103, -24, -24,
	26, 26, 142, 142, 129, 110, 110, -1000, -1000, -1000,
	-4, 84, -10, 32, -1000, 127, -1000, -1000, 84,
}
var parserPgo = [...]int{

	0, 0, 40, 30, 16, 9, 7, 3,
}
var parserR1 = [...]int{

	0, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 5, 5, 5, 5, 5, 5, 5, 4, 4,
	4, 4, 6, 6, 6, 6, 6, 7, 7, 3,
	3, 3,
}
var parserR2 = [...]int{

	0, 1, 3, 3, 3, 3, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 4, 4, 4, 1, 2, 3, 1, 1,
	1, 1,
}
var parserChk = [...]int{

	-1000, -2, -1, 20, 24, 22, -5, -4, -3, -6,
	13, 10, 4, 5, 6, 7, 26, 25, 16, 17,
	14, 15, 12, 13, 11, 8, 9, -1, -1, -1,
	-1, -1, -1, 20, 24, 22, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, 19, 23, 21,
	-7, -1, -7, -7, 19, 18, 23, 21, -1,
}
var parserDef = [...]int{

	0, -2, 1, 0, 0, 0, 6, 7, 8, 9,
	0, 0, 29, 30, 31, 25, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	26, 17, 18, 0, 0, 0, 5, 10, 11, 12,
	13, 14, 15, 16, 19, 20, 21, 2, 3, 4,
	0, 28, 0, 0, 22, 0, 23, 24, 27,
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
		//line parser.y:51
		{
			parserVAL.script = parserDollar[1].script
			parserlex.(*parserLex).output = parserVAL.script
		}
	case 2:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:59
		{
			parserVAL.script = parserDollar[2].script
		}
	case 3:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:61
		{
			parserVAL.script = parserDollar[2].script
		}
	case 4:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:63
		{
			parserVAL.script = parserDollar[2].script
		}
	case 5:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:66
		{
			/* The colon is a pipe operator with high prescedence */
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 9:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:83
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
	case 10:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:100
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 11:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:113
		{
			s, err := modScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 12:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:123
		{
			s, err := powScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 13:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:133
		{
			s, err := mulScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 14:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:143
		{
			s, err := divScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 15:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:153
		{
			s, err := addScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 16:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:163
		{
			s, err := subtractScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 17:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:173
		{
			s, err := negativeScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 18:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:185
		{
			s, err := notScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 19:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:196
		{
			s, err := comparisonScript(parserDollar[2].strVal, parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 20:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:206
		{
			s, err := andScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 21:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:216
		{
			s, err := orScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 22:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:229
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 23:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:235
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 24:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:241
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 25:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:248
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 26:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:254
		{
			parserVAL.sfunc.transform = parserDollar[1].sfunc.transform
			parserVAL.sfunc.args = append(parserDollar[1].sfunc.args, parserDollar[2].script)
		}
	case 27:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:261
		{
			parserVAL.scriptArray = append(parserDollar[1].scriptArray, parserDollar[3].script)
		}
	case 28:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:266
		{
			parserVAL.scriptArray = []*Script{parserDollar[1].script}
		}
	case 29:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:274
		{
			num, err := strconv.ParseFloat(parserDollar[1].strVal, 64)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = ConstantScript(num)
		}
	case 30:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:284
		{
			parserVAL.script = ConstantScript(parserDollar[1].strVal)
		}
	case 31:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:289
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
