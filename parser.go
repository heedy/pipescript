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
const pIDENTIFIER_SPACE = 57350
const pAND = 57351
const pOR = 57352
const pNOT = 57353
const pCOMPARISON = 57354
const pPLUS = 57355
const pMINUS = 57356
const pMULTIPLY = 57357
const pDIVIDE = 57358
const pMODULO = 57359
const pPOW = 57360
const pCOMMA = 57361
const pRPARENS = 57362
const pLPARENS = 57363
const pRSQUARE = 57364
const pLSQUARE = 57365
const pRBRACKET = 57366
const pLBRACKET = 57367
const pPIPE = 57368
const pCOLON = 57369
const pNOARGS = 57370
const pARGS = 57371
const pUMINUS = 57372

var parserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"pNUMBER",
	"pSTRING",
	"pBOOL",
	"pIDENTIFIER",
	"pIDENTIFIER_SPACE",
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
	"pNOARGS",
	"pARGS",
	"pUMINUS",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line parser.y:419

func parserGetScript(sf scriptFunc) (*Script, error) {
	RegistryLock.RLock()
	v, ok := TransformRegistry[sf.transform]
	RegistryLock.RUnlock()
	if ok {
		s, err := v.Script(sf.args)
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return nil, fmt.Errorf("Transform %s not found", sf.transform)
}

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 55,
	1, 19,
	9, 19,
	10, 19,
	12, 19,
	13, 19,
	14, 19,
	20, 19,
	22, 19,
	24, 19,
	26, 19,
	-2, 21,
}

const parserNprod = 41
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 213

var parserAct = [...]int{

	4, 59, 28, 29, 26, 27, 32, 22, 34, 35,
	37, 26, 27, 58, 22, 21, 57, 21, 21, 3,
	41, 22, 42, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 56, 54, 66, 13, 55, 69, 21,
	5, 43, 60, 63, 62, 66, 65, 12, 6, 10,
	11, 1, 0, 24, 25, 64, 23, 30, 31, 28,
	29, 26, 27, 68, 0, 0, 70, 71, 0, 72,
	0, 22, 24, 25, 0, 23, 30, 31, 28, 29,
	26, 27, 68, 67, 0, 0, 0, 0, 24, 25,
	22, 23, 30, 31, 28, 29, 26, 27, 0, 0,
	0, 0, 0, 0, 24, 0, 22, 23, 30, 31,
	28, 29, 26, 27, 14, 15, 16, 20, 33, 0,
	0, 7, 22, 0, 9, 0, 0, 0, 0, 2,
	61, 17, 0, 19, 0, 18, 14, 15, 16, 20,
	33, 0, 0, 7, 0, 0, 9, 38, 39, 40,
	0, 0, 0, 17, 0, 19, 0, 18, 23, 30,
	31, 28, 29, 26, 27, 14, 15, 16, 20, 8,
	0, 0, 7, 22, 0, 9, 0, 0, 0, 0,
	0, 0, 17, 0, 19, 0, 18, 14, 15, 16,
	20, 33, 0, 0, 7, 0, 0, 36, 30, 31,
	28, 29, 26, 27, 17, 0, 19, 0, 18, 0,
	0, 0, 22,
}
var parserPact = [...]int{

	161, -1000, -11, -1000, 79, 132, -1000, 132, 183, 132,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 161, 161, 161,
	-1, 161, 132, 132, 132, 132, 132, 132, 132, 132,
	132, 132, 79, 20, 185, 79, 132, -20, 13, -8,
	-9, 110, 132, -1000, -1000, 185, 146, 95, -20, -20,
	-6, -6, -13, -13, 132, -20, -1000, -1000, -1000, 26,
	63, -1000, 16, 44, -13, -1000, 132, -1000, 132, -1000,
	-1000, 79, 79,
}
var parserPgo = [...]int{

	0, 51, 129, 50, 0, 49, 19, 48, 47, 40,
	36, 1,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 6, 6, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 7, 7, 7, 8, 8, 8, 5, 10,
	10, 10, 10, 10, 10, 10, 11, 11, 3, 3,
	3,
}
var parserR2 = [...]int{

	0, 1, 1, 3, 1, 1, 2, 2, 1, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 1, 1, 1, 3, 3, 3, 1, 4,
	4, 4, 4, 3, 1, 1, 3, 3, 1, 1,
	1,
}
var parserChk = [...]int{

	-1000, -1, -2, -6, -4, -9, -7, 11, 8, 14,
	-5, -3, -8, -10, 4, 5, 6, 21, 25, 23,
	7, 26, 27, 12, 9, 10, 17, 18, 15, 16,
	13, 14, -4, 8, -4, -4, 14, -4, -2, -2,
	-2, 21, 23, -6, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, 14, -4, 20, 24, 22, -11,
	-4, 20, -11, -4, -4, 20, 19, 20, 19, 22,
	22, -4, -4,
}
var parserDef = [...]int{

	0, -2, 1, 2, 4, 5, 8, 0, 35, 0,
	22, 23, 24, 28, 38, 39, 40, 0, 0, 0,
	34, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 6, 35, 10, 7, 0, 21, 0, 0,
	0, 0, 0, 3, 9, 11, 12, 13, 14, 15,
	16, 17, 18, 20, 0, -2, 25, 26, 27, 0,
	0, 33, 0, 0, 19, 29, 0, 31, 0, 30,
	32, 36, 37,
}
var parserTok1 = [...]int{

	1,
}
var parserTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		//line parser.y:60
		{
			parserVAL.script = parserDollar[1].script
			parserlex.(*parserLex).output = parserVAL.script
		}
	case 3:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:75
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 5:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:95
		{
			s, err := parserGetScript(parserDollar[1].sfunc)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			parserVAL.script = s
		}
	case 6:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:108
		{
			parserVAL.sfunc.transform = parserDollar[1].sfunc.transform
			parserVAL.sfunc.args = append(parserDollar[1].sfunc.args, parserDollar[2].script)
		}
	case 7:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:114
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[2].script}
		}
	case 9:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:132
		{
			err := parserDollar[1].script.Append(parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = parserDollar[1].script
		}
	case 10:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:142
		{
			s, err := notScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 11:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:153
		{
			s, err := comparisonScript(parserDollar[2].strVal, parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 12:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:163
		{
			s, err := andScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 13:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:173
		{
			s, err := orScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 14:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:183
		{
			s, err := modScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 15:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:193
		{
			s, err := powScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 16:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:203
		{
			s, err := mulScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 17:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:213
		{
			s, err := divScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 18:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:223
		{
			s, err := addScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 19:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:236
		{
			// First get the script of this function
			sf := scriptFunc{
				transform: parserDollar[1].strVal,
				args:      []*Script{},
			}
			s, err := parserGetScript(sf)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			// Now subtract the two
			s, err = subtractScript(s, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 20:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:257
		{
			s, err := subtractScript(parserDollar[1].script, parserDollar[3].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 21:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:267
		{
			s, err := negativeScript(parserDollar[2].script)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = s
		}
	case 25:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:294
		{
			parserVAL.script = parserDollar[2].script
		}
	case 26:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:296
		{
			parserVAL.script = parserDollar[2].script
		}
	case 27:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:298
		{
			parserVAL.script = parserDollar[2].script
		}
	case 28:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:307
		{
			s, err := parserGetScript(parserDollar[1].sfunc)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			parserVAL.script = s

		}
	case 29:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:323
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 30:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:330
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 31:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:338
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 32:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:344
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 33:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:350
		{
			// Allows calling as a function
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 34:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:357
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 35:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:363
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 36:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:376
		{
			parserVAL.scriptArray = append(parserDollar[1].scriptArray, parserDollar[3].script)
		}
	case 37:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:381
		{
			parserVAL.scriptArray = []*Script{parserDollar[1].script, parserDollar[3].script}
		}
	case 38:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:395
		{
			num, err := strconv.ParseFloat(parserDollar[1].strVal, 64)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = ConstantScript(num)
		}
	case 39:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:405
		{
			parserVAL.script = ConstantScript(parserDollar[1].strVal)
		}
	case 40:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:410
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
