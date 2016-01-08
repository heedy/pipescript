// A parser for PipeScript
// Generate using:
//	go tool yacc -o parser.go -p parser parser.y

%{
package pipescript

import (
	"strconv"
	"fmt"
)

%}


%union{
	script *Script
	scriptArray []*Script
	strVal string	// This is how variables are passed in: by their string value
}

%type <script> script pipescript constant logical algebra
%type <scriptArray> script_array
%token <strVal> pNUMBER  pSTRING  pBOOL pIDENTIFIER
%token <strVal> pAND pOR pNOT pCOMPARISON pPLUS pMINUS pMULTIPLY pDIVIDE pMODULO pPOW pPIPE pCOMMA
%token <strVal> pRPARENS pLPARENS pRSQUARE pLSQUARE pRBRACKET pLBRACKET

/* Set up order of operations */
%left pPIPE
%left pIDENTIFIER
%left pAND pOR
%left pCOMPARISON
%left pNOT
%left pPLUS pMINUS
%left pMULTIPLY pDIVIDE
%left pMODULO pPOW
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
	algebra
	|
	logical
	|
	constant
	|
	pIDENTIFIER pLPARENS pRPARENS
		{
			v,ok := TransformRegistry[$1]
			if ok {
				s,err := v.Script(nil)
				if err!=nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				$$ = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found",$1))
				goto ret1
			}
		}
	|
	pIDENTIFIER script_array
		{
			v,ok := TransformRegistry[$1]
			if ok {
				s,err := v.Script($2)
				if err!=nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				$$ = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found",$1))
				goto ret1
			}
		}
	|
	pIDENTIFIER
		{
			v,ok := TransformRegistry[$1]
			if ok {
				s,err := v.Script(nil)
				if err!=nil {
					parserlex.Error(err.Error())
					goto ret1
				}
				$$ = s
			} else {
				parserlex.Error(fmt.Sprintf("Transform %s not found",$1))
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

script_array:
	/* Set up the handlers of parentheses */
	pLPARENS script_array pRPARENS { $$ = $2 }
	|
	pLBRACKET script_array pRBRACKET { $$ = $2 }
	|
	pLSQUARE script_array pRSQUARE { $$ = $2 }

	|
	script
		{
			$$ = []*Script{$1}
		}
	|
	script_array pCOMMA script
		{
			/* Allows usage in function */
			$$ = append($1,$3)
		}
	|
	script_array script
		{
			/* Allows usage in bash-like syntax */
			$$ = append($1,$2)
		}

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
