// A parser for PipeScript
// Generate using:
//	go tool yacc -o parser.go -p parser parser.y

%{
package pipescript

import (
	"strconv"
	"fmt"
)

type scriptFunc struct {
	transform string
	args []*Script
}

%}



%union{
	script *Script
	sfunc scriptFunc
	scriptArray []*Script
	strVal string	// This is how variables are passed in: by their string value
}

%type <script> script pipescript constant logical algebra
%type <sfunc> function
%type <scriptArray> script_array
%token <strVal> pNUMBER  pSTRING  pBOOL pIDENTIFIER
%token <strVal> pAND pOR pNOT pCOMPARISON pPLUS pMINUS pMULTIPLY pDIVIDE pMODULO pPOW pCOMMA
%token <strVal> pRPARENS pLPARENS pRSQUARE pLSQUARE pRBRACKET pLBRACKET pPIPE pCOLON

/* Set up order of operations */
%left pPIPE
%left pIDENTIFIER pCOMMA
%left pAND pOR
%left pCOMPARISON
%left pNOT
%left pPLUS pMINUS
%left pMULTIPLY pDIVIDE
%left pMODULO pPOW
%left pCOLON
%left pUMINUS      /*  supplies  precedence  for  unary  minus  */
%%

pipescript:
	script
		{
			$$ = $1
			parserlex.(*parserLex).output = $$
		}
	;

script:
	/* Set up the handlers of parentheses */
	pLPARENS script pRPARENS { $$ = $2 }
	|
	pLBRACKET script pRBRACKET { $$ = $2 }
	|
	pLSQUARE script pRSQUARE { $$ = $2 }
	|
	script pCOLON script
		{
			/* The colon is a pipe operator with high prescedence */
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	algebra
	|
	logical
	|
	constant
	|
	function
		{
			v,ok := TransformRegistry[$1.transform]
			if ok {
				s,err := v.Script($1.args)
				if err!=nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				$$ = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found",$1.transform))
				goto ret1
			}
		}
	|

	script pPIPE script
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	;


algebra:
	script pMODULO script
		{
			s,err := modScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pPOW script
		{
			s,err := powScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pMULTIPLY script
		{
			s,err := mulScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pDIVIDE script
		{
			s,err := divScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pPLUS script
		{
			s,err := addScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pMINUS script
		{
			s,err := subtractScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	pMINUS script %prec pUMINUS
		{
			s,err := negativeScript($2)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	;

logical:
	pNOT script
		{
			s,err := notScript($2)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	/* Comparisons are: ==,>=,<=,>,<,!= */
	script pCOMPARISON script
		{
			s,err := comparisonScript($2,$1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pAND script
		{
			s,err := andScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	script pOR script
		{
			s,err := orScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	;

function:
	/* Set up the handlers of parentheses */
	pIDENTIFIER pLPARENS script_array pRPARENS
		{
			$$.transform = $1
			$$.args = $3
		}
	|
	pIDENTIFIER pLBRACKET script_array pRBRACKET
		{
			$$.transform = $1
			$$.args = $3
		}
	|
	pIDENTIFIER pLSQUARE script_array pRSQUARE
		{
			$$.transform = $1
			$$.args = $3
		}

	|
	pIDENTIFIER
		{
			$$.transform = $1
			$$.args = []*Script{}
		}
	|
	function script
		{
			$$.transform = $1.transform
			$$.args = append($1.args,$2)
		}
	;
script_array:
	script_array pCOMMA script
		{
			$$ = append($1,$3)
		}
	|
	script
		{
			$$ = []*Script{$1}
		}
	;

/* Convert constants to their corresponding scripts */
constant:
	pNUMBER
		{
			num, err := strconv.ParseFloat($1, 64)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = ConstantScript(num)
		}
	|
	pSTRING
		{
			$$ = ConstantScript($1)
		}
	|
	pBOOL
		{
			if $1=="true" {
				$$ = ConstantScript(true)
			} else {
				$$ = ConstantScript(false)
			}
		}
	;

%%
