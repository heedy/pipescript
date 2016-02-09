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

%type <script> script pipescript constant algebraic simpletransform transform statement parensvalue
%type <sfunc> function simplefunction
%type <scriptArray> script_array
%token <strVal> pNUMBER  pSTRING  pBOOL pIDENTIFIER pIDENTIFIER_SPACE
%token <strVal> pAND pOR pNOT pCOMPARISON pPLUS pMINUS pMULTIPLY pDIVIDE pMODULO pPOW pCOMMA
%token <strVal> pRPARENS pLPARENS pRSQUARE pLSQUARE pRBRACKET pLBRACKET pPIPE pCOLON

%nonassoc pNOARGS
%nonassoc pLPARENS pLBRACKET pLSQUARE
%nonassoc pARGS
%left pCOMMA

/* Order of operations for algebraic expressions */
%left pOR
%left pAND
%left pCOMPARISON
%left pNOT
%left pPLUS pMINUS
%left pMULTIPLY pDIVIDE
%left pMODULO pPOW
%left pUMINUS      /*  supplies  precedence  for  unary  minus  */
%left pCOLON

//%nonassoc pSUPER



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
pipescript: transform
	|
	pipescript pPIPE transform
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
 Handle functions (transforms). transforms are bash-like arguments
 	Input: statement
	Output: algebraic
*************************************************************************************/
transform: algebraic	// Algebraic HAS to be first
	|
	function
		{
			s,err := parserGetScript($1)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			$$ = s
		}
	;

function:
	function algebraic %prec pARGS
		{
			$$.transform = $1.transform
			$$.args = append($1.args,$2)
		}
	|
	pIDENTIFIER_SPACE algebraic %prec pARGS
		{
			$$.transform = $1
			$$.args = []*Script{$2}
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
	algebraic pCOLON algebraic
		{
			err := $1.Append($3)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}
			$$ = $1
		}
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
	/* The parser has difficulty handling subtraction of IDENTIFIER_SPACE, since we have conflict with bash/function style transforms.
	pIDENTIFIER_SPACE pMINUS algebraic -> pIDENTIFIER_SPACE (pMINUS algebraic) by default (unary minus)
	We want it to work as normal subtraction. Resolve this here. */
	pIDENTIFIER_SPACE pMINUS algebraic //%prec pSUPER
	{
		// First get the script of this function
		sf := scriptFunc{
			transform: $1,
			args: []*Script{},
		}
		s,err := parserGetScript(sf)
		if err!=nil {
			parserlex.Error(err.Error())
			goto ret1
		}
		// Now subtract the two
		s,err = subtractScript(s,$3)
		if err!=nil {
			parserlex.Error(err.Error())
			goto ret1
		}
		$$ = s
	}
	|
	algebraic pMINUS algebraic %prec pMINUS
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

statement:
	simpletransform
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
simplefunction/transform combines script_array and identifier to form function: f(a,b,c,d)
*************************************************************************************/

simpletransform:
	simplefunction
		{
			s,err := parserGetScript($1)
			if err!=nil {
				parserlex.Error(err.Error())
				goto ret1
			}

			$$ = s

		}
	;

simplefunction:

	/* Set up the handlers of parentheses - we need to match correct type of parens */
	pIDENTIFIER pLPARENS script_array pRPARENS //%prec pARGS
		{
			$$.transform = $1
			$$.args = $3
		}
	|

	pIDENTIFIER pLSQUARE script_array pRSQUARE //%prec pARGS
		{
			$$.transform = $1
			$$.args = $3
		}
	|


	pIDENTIFIER pLPARENS algebraic pRPARENS //%prec pARGS
		{
			$$.transform = $1
			$$.args = []*Script{$3}
		}
	|
	pIDENTIFIER pLSQUARE algebraic pRSQUARE //%prec pARGS
		{
			$$.transform = $1
			$$.args = []*Script{$3}
		}
	|
	pIDENTIFIER pLPARENS pRPARENS
		{
			// Allows calling as a function
			$$.transform = $1
			$$.args = []*Script{}
		}
	|
	pIDENTIFIER %prec pNOARGS
		{
			$$.transform = $1
			$$.args = []*Script{}
		}
	|
	pIDENTIFIER_SPACE %prec pNOARGS
		{
			$$.transform = $1
			$$.args = []*Script{}
		}
	;


/*************************************************************************************
script_array allows us to prepare the args of a function f(a,b,c,d)
*************************************************************************************/

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
Prepare constant values. The lexed values are all strings, so convert to correct type
and convert into ConstantScript
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

func parserGetScript(sf scriptFunc) (*Script,error) {
	RegistryLock.RLock()
	v,ok := TransformRegistry[sf.transform]
	RegistryLock.RUnlock()
	if ok {
		s,err := v.Script(sf.args)
		if err!=nil {
			return nil, err
		}
		return s,nil
	}
	return nil, fmt.Errorf("Transform %s not found",sf.transform)
}
