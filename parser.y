// A parser for PipeScript
// Generate using:
//	go tool yacc -o parser.go -p parser parser.y

%{
package pipescript

import (
	"fmt"
	"strconv"
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

%type <script> script pipescript colonscript constant algebraic simpletransform transform statement parensvalue
%type <sfunc> function simplefunction
%type <scriptArray> script_array
%token <strVal> pNUMBER  pSTRING  pBOOL pIDENTIFIER
%token <strVal> pAND pOR pNOT pCOMPARISON pPLUS pMINUS pMULTIPLY pDIVIDE pMODULO pPOW pCOMMA
%token <strVal> pRPARENS pLPARENS pRSQUARE pLSQUARE pRBRACKET pLBRACKET pPIPE pCOLON

%left pCOMMA
%left pCOLON

/* Order of operations for algebraic expressions */
%left pOR
%left pAND
%left pCOMPARISON
%left pNOT
%left pPLUS pMINUS
%left pMULTIPLY pDIVIDE
%left pMODULO pPOW
%left pUMINUS      /*  supplies  precedence  for  unary  minus  */

%%

script:
	pipescript
		{
			$$ = $1
			parserlex.(*parserLex).output = $$
		}
	;


/*************************************************************************************
Set up the scripts that are separated by pipe. Pipescript uses full transforms as its elements
 	Input: transform | transform | transform
	Output: pipescript
*************************************************************************************/
pipescript: algebraic
	|
	pipescript pPIPE algebraic
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	;

/*************************************************************************************
Create the colon-script. It only accepts simpletransforms (since bash-like calling would
cause issues)
 	Input: simpletransform : simpletransform
	Output: colonscript
*************************************************************************************/

colonscript:
	/* Only create a colonscript if there is a colon! */
	simpletransform pCOLON simpletransform
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	parensvalue pCOLON simpletransform
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	parensvalue pCOLON parensvalue
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	simpletransform pCOLON parensvalue
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	colonscript pCOLON simpletransform
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	|
	colonscript pCOLON parensvalue
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
	;


/*************************************************************************************
 Handle functions (transforms). transforms are all transforms (including bash args)
 and simpletransforms are transforms with function style f(x,y,z)
 	Input: statement
	Output: algebraic
*************************************************************************************/
transform: simpletransform
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
	;

simpletransform:
	simplefunction
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
	;

function:
	simplefunction algebraic
		{
			/* A simplefunction can have at most one argument if it is to be called in bash style */
			if len($1.args) > 1 {
				parserlex.Error("Used both function f(x,y) and bash-style function calling at same time for a single function call. This probably means you made a syntax error")
				goto ret1
			}
			$$.transform = $1.transform
			$$.args = append($1.args,$2)
		}
	|
	function algebraic
		{
			$$.transform = $1.transform
			$$.args = append($1.args,$2)
		}
	|
	pIDENTIFIER algebraic
		{
			$$.transform = $1
			$$.args = []*Script{$2}
		}
	;

simplefunction:
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
	pIDENTIFIER pLPARENS algebraic pRPARENS
		{
			$$.transform = $1
			$$.args = []*Script{$3}
		}
	|
	pIDENTIFIER pLBRACKET algebraic pRBRACKET
		{
			$$.transform = $1
			$$.args = []*Script{$3}
		}
	|
	pIDENTIFIER pLSQUARE algebraic pRSQUARE
		{
			$$.transform = $1
			$$.args = []*Script{$3}
		}

	|
	pIDENTIFIER
		{
			$$.transform = $1
			$$.args = []*Script{}
		}
	;

script_array:
	script_array pCOMMA algebraic
		{
			$$ = append($1,$3)
		}
	|
	algebraic pCOMMA algebraic
		{
			$$ = []*Script{$1,$3}
		}
	;

/*************************************************************************************
 Handle algebra and comparisons. Note that order of operations is defined above by
 prescedence.
 	Input: statement
	Output: algebraic
*************************************************************************************/

algebraic: statement
	|
	pNOT algebraic
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
	algebraic pCOMPARISON algebraic
		{
			s,err := comparisonScript($2,$1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pAND algebraic
		{
			s,err := andScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pOR algebraic
		{
			s,err := orScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pMODULO algebraic
		{
			s,err := modScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pPOW algebraic
		{
			s,err := powScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pMULTIPLY algebraic
		{
			s,err := mulScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pDIVIDE algebraic
		{
			s,err := divScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pPLUS algebraic
		{
			s,err := addScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	algebraic pMINUS algebraic
		{
			s,err := subtractScript($1,$3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	|
	pMINUS algebraic %prec pUMINUS
		{
			s,err := negativeScript($2)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = s
		}
	;

/*************************************************************************************
Create the statement!
	Output: statement
*************************************************************************************/

statement: colonscript
	|
	transform
	|
	constant
	|
	/* Set up the handlers of parentheses */
	parensvalue
	;

parensvalue:
	/* Set up the handlers of parentheses */
	pLPARENS pipescript pRPARENS { $$ = $2 }
	|
	pLBRACKET pipescript pRBRACKET { $$ = $2 }
	|
	pLSQUARE pipescript pRSQUARE { $$ = $2 }
	;

/*************************************************************************************
Prepare constant values
 	Input: lexed values
	Output: constant
*************************************************************************************/
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
