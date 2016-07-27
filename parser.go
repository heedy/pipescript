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

//line parser.y:23
type parserSymType struct {
	yys         int
	script      *Script
	sfunc       scriptFunc
	scriptArray []*Script
	objBuilder  map[string]*Script
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
const parserInitialStackSize = 16

//line parser.y:459

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
	-1, 56,
	1, 19,
	9, 19,
	10, 19,
	12, 19,
	13, 19,
	14, 19,
	20, 19,
	22, 19,
	26, 19,
	-2, 21,
}

const parserNprod = 43
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 229

var parserAct = [...]int{

	4, 60, 29, 30, 27, 28, 33, 23, 35, 36,
	38, 27, 28, 59, 23, 57, 58, 22, 22, 3,
	68, 23, 22, 71, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 42, 55, 43, 39, 56, 68,
	67, 2, 44, 61, 64, 63, 25, 26, 14, 24,
	31, 32, 29, 30, 27, 28, 65, 13, 66, 5,
	40, 41, 12, 6, 23, 10, 11, 1, 0, 75,
	0, 76, 25, 26, 0, 24, 31, 32, 29, 30,
	27, 28, 74, 0, 0, 0, 0, 73, 25, 26,
	23, 24, 31, 32, 29, 30, 27, 28, 70, 0,
	0, 72, 0, 0, 25, 26, 23, 24, 31, 32,
	29, 30, 27, 28, 70, 69, 0, 0, 0, 0,
	25, 0, 23, 24, 31, 32, 29, 30, 27, 28,
	15, 16, 17, 20, 34, 0, 0, 7, 23, 0,
	9, 0, 0, 0, 0, 0, 62, 18, 0, 19,
	0, 21, 15, 16, 17, 20, 34, 0, 0, 7,
	0, 0, 9, 0, 0, 0, 0, 0, 0, 18,
	0, 19, 0, 21, 24, 31, 32, 29, 30, 27,
	28, 15, 16, 17, 20, 8, 0, 0, 7, 23,
	0, 9, 0, 0, 0, 0, 0, 0, 18, 0,
	19, 0, 21, 15, 16, 17, 20, 34, 0, 0,
	7, 0, 0, 37, 31, 32, 29, 30, 27, 28,
	18, 0, 19, 0, 21, 0, 0, 0, 23,
}
var parserPact = [...]int{

	177, -1000, -8, -1000, 37, 148, -1000, 148, 199, 148,
	-1000, -1000, -1000, -1000, 32, -1000, -1000, -1000, 177, 177,
	13, -1000, 177, 148, 148, 148, 148, 148, 148, 148,
	148, 148, 148, 37, 21, 162, 37, 148, -20, -12,
	-4, -9, 126, 148, -1000, -1000, 201, 162, 111, -20,
	-20, -6, -6, -13, -13, 148, -20, 148, -1000, -1000,
	20, 95, -1000, 1, 79, -13, 63, -1000, 148, -1000,
	148, -1000, -1000, -1000, -1000, 37, 37,
}
var parserPgo = [...]int{

	0, 67, 41, 66, 0, 65, 19, 63, 62, 59,
	57, 1, 48,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 6, 6, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 7, 7, 7, 8, 8, 5, 5, 10,
	10, 10, 10, 10, 10, 10, 11, 11, 12, 12,
	3, 3, 3,
}
var parserR2 = [...]int{

	0, 1, 1, 3, 1, 1, 2, 2, 1, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 1, 1, 1, 3, 3, 1, 5, 4,
	4, 4, 4, 3, 1, 1, 3, 3, 1, 5,
	1, 1, 1,
}
var parserChk = [...]int{

	-1000, -1, -2, -6, -4, -9, -7, 11, 8, 14,
	-5, -3, -8, -10, -12, 4, 5, 6, 21, 23,
	7, 25, 26, 27, 12, 9, 10, 17, 18, 15,
	16, 13, 14, -4, 8, -4, -4, 14, -4, 5,
	-2, -2, 21, 23, -6, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, 14, -4, 27, 20, 22,
	-11, -4, 20, -11, -4, -4, -4, 20, 19, 20,
	19, 22, 22, 24, 19, -4, -4,
}
var parserDef = [...]int{

	0, -2, 1, 2, 4, 5, 8, 0, 35, 0,
	22, 23, 24, 27, 0, 40, 41, 42, 0, 0,
	34, 38, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 6, 35, 10, 7, 0, 21, 0,
	0, 0, 0, 0, 3, 9, 11, 12, 13, 14,
	15, 16, 17, 18, 20, 0, -2, 0, 25, 26,
	0, 0, 33, 0, 0, 19, 0, 29, 0, 31,
	0, 30, 32, 28, 39, 36, 37,
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
	lval  parserSymType
	stack [parserInitialStackSize]parserSymType
	char  int
}

func (p *parserParserImpl) Lookahead() int {
	return p.char
}

func parserNewParser() parserParser {
	return &parserParserImpl{}
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
	var parserVAL parserSymType
	var parserDollar []parserSymType
	_ = parserDollar // silence set and not used
	parserS := parserrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserrcvr.char = -1
	parsertoken := -1 // parserrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		parserstate = -1
		parserrcvr.char = -1
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
	if parserrcvr.char < 0 {
		parserrcvr.char, parsertoken = parserlex1(parserlex, &parserrcvr.lval)
	}
	parsern += parsertoken
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parsertoken { /* valid shift */
		parserrcvr.char = -1
		parsertoken = -1
		parserVAL = parserrcvr.lval
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
		if parserrcvr.char < 0 {
			parserrcvr.char, parsertoken = parserlex1(parserlex, &parserrcvr.lval)
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
			parserrcvr.char = -1
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
		//line parser.y:63
		{
			parserVAL.script = parserDollar[1].script
			parserlex.(*parserLex).output = parserVAL.script
		}
	case 3:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:78
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
		//line parser.y:98
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
		//line parser.y:111
		{
			parserVAL.sfunc.transform = parserDollar[1].sfunc.transform
			parserVAL.sfunc.args = append(parserDollar[1].sfunc.args, parserDollar[2].script)
		}
	case 7:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:117
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[2].script}
		}
	case 9:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:135
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
		//line parser.y:145
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
		//line parser.y:156
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
		//line parser.y:166
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
		//line parser.y:176
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
		//line parser.y:186
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
		//line parser.y:196
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
		//line parser.y:206
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
		//line parser.y:216
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
		//line parser.y:226
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
		//line parser.y:239
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
		//line parser.y:260
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
		//line parser.y:270
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
		//line parser.y:297
		{
			parserVAL.script = parserDollar[2].script
		}
	case 26:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:299
		{
			parserVAL.script = parserDollar[2].script
		}
	case 27:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:308
		{
			s, err := parserGetScript(parserDollar[1].sfunc)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			parserVAL.script = s

		}
	case 28:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line parser.y:320
		{
			if _, ok := parserDollar[1].objBuilder[parserDollar[2].strVal]; ok {
				parserlex.Error(fmt.Sprintf("Key %s found multiple times in json object", parserDollar[2].strVal))
				goto ret1
			}
			parserDollar[1].objBuilder[parserDollar[2].strVal] = parserDollar[4].script

			// Now generate the objectScript
			s, err := newObjectTransform(parserDollar[1].objBuilder)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			parserVAL.script = s
		}
	case 29:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:342
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 30:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:349
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = parserDollar[3].scriptArray
		}
	case 31:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:357
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 32:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:363
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{parserDollar[3].script}
		}
	case 33:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:369
		{
			// Allows calling as a function
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 34:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:376
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 35:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:382
		{
			parserVAL.sfunc.transform = parserDollar[1].strVal
			parserVAL.sfunc.args = []*Script{}
		}
	case 36:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:395
		{
			parserVAL.scriptArray = append(parserDollar[1].scriptArray, parserDollar[3].script)
		}
	case 37:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:400
		{
			parserVAL.scriptArray = []*Script{parserDollar[1].script, parserDollar[3].script}
		}
	case 38:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:412
		{
			parserVAL.objBuilder = make(map[string]*Script)
		}
	case 39:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line parser.y:417
		{
			if _, ok := parserDollar[1].objBuilder[parserDollar[2].strVal]; ok {
				parserlex.Error(fmt.Sprintf("Key %s found multiple times in json object", parserDollar[2].strVal))
				goto ret1
			}
			parserDollar[1].objBuilder[parserDollar[2].strVal] = parserDollar[4].script
			parserVAL.objBuilder = parserDollar[1].objBuilder
		}
	case 40:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:435
		{
			num, err := strconv.ParseFloat(parserDollar[1].strVal, 64)
			if err != nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			parserVAL.script = ConstantScript(num)
		}
	case 41:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:445
		{
			parserVAL.script = ConstantScript(parserDollar[1].strVal)
		}
	case 42:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:450
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
