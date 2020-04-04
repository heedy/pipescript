package pipescript

import (
	"regexp"
	"strings"
	"unicode"
)

/*
Hopefully go will let us do this automatically in the future - Joseph
*/

const (
	eof           = 0
	errorString   = "<ERROR>"
	eofString     = "<EOF>"
	logicals      = "and|or|not"
	booleans      = "true|false"
	numbers       = `[0-9]+(\.[0-9]+)?`
	comparisons   = `<=|>=|<|>|==|!=`
	stringregex   = `\"(\\["nrt\\]|.)*?\"|'(\\['nrt\\]|.)*?'`
	pipes         = `:|\|`
	brackets      = `\[|\]|\(|\)|{|}`
	mathoperators = `\-|\*|/|\+|%|\^`
	idents        = `([a-zA-Z_\$][a-zA-Z_0-9\$]*)`
	allregex      = logicals + "|" + numbers + "|" + comparisons + "|" + booleans + "|" +
		stringregex + "|" + pipes + "|" + mathoperators + "|" + idents + "|" + brackets + "|,"
)

var (
	tokenizer   = regexp.MustCompile(`^(` + allregex + `)`)
	numberRegex = regexp.MustCompile("^" + numbers + "$")
	stringRegex = regexp.MustCompile("^" + stringregex + "$")
	identRegex  = regexp.MustCompile("^" + idents + "$")
)

// parserLex is the parserLexer for transforms
type parserLex struct {
	input    string
	position int

	errorString string

	output *Pipe
}

// Are we at the end of file?
func (l *parserLex) AtEOF() bool {
	return l.position >= len(l.input)
}

// Return the next string for the parserLexer
func (l *parserLex) Next() string {
	var c rune = ' '

	// skip whitespace
	for unicode.IsSpace(c) {
		if l.AtEOF() {
			return eofString
		}
		c = rune(l.input[l.position])
		l.position++
	}

	l.position--

	rest := l.input[l.position:]

	token := tokenizer.FindString(rest)
	l.position += len(token)

	if token == "" {
		return errorString
	}

	return token
}

func (l *parserLex) Error(s string) {
	l.errorString = s
}

func (l *parserLex) Lex(lval *parserSymType) int {
	token := l.Next()
	lval.strVal = token
	switch token {
	case eofString:
		return 0
	case errorString:
		l.Error("Error, unknown token")
		return 0
	case ")":
		return pRPARENS
	case "(":
		return pLPARENS
	case "]":
		return pRSQUARE
	case "[":
		return pLSQUARE
	case "{":
		return pLBRACKET
	case "}":
		return pRBRACKET
	case "and":
		return pAND
	case "or":
		return pOR
	case "not":
		return pNOT
	case ">=", "<=", ">", "<", "==", "!=":
		return pCOMPARISON
	case "true", "false":
		return pBOOL
	case "|":
		return pPIPE
	case ":":
		return pCOLON
	case ",":
		return pCOMMA
	case "-":
		return pMINUS
	case "+":
		return pPLUS
	case "/":
		return pDIVIDE
	case "*":
		return pMULTIPLY
	case "%":
		return pMODULO
	case "^":
		return pPOW
	default:
		switch {
		case numberRegex.MatchString(token):
			return pNUMBER

		case stringRegex.MatchString(token):
			// unquote token
			strval := token[1 : len(token)-1]

			// replace escape characters
			strval = strings.Replace(strval, "\\n", "\n", -1)
			strval = strings.Replace(strval, "\\r", "\r", -1)
			strval = strings.Replace(strval, "\\t", "\t", -1)
			strval = strings.Replace(strval, "\\\\", "\\", -1)
			strval = strings.Replace(strval, "\\\"", "\"", -1)
			strval = strings.Replace(strval, "\\'", "'", -1)

			lval.strVal = strval
			return pSTRING

		default:
			// If there is a space after the identifier, it isn't being called as f(x)
			if l.position < len(l.input) && unicode.IsSpace(rune(l.input[l.position])) {
				return pIDENTIFIER_SPACE
			}
			return pIDENTIFIER

		}
	}
}
