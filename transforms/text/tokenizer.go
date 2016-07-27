package text

import (
	"regexp"
	"strings"
)

var (
	spacereplace = regexp.MustCompile("[.,\\/#!$%\\^&\\*;:{}=\\-_`~()]+")
)

// Tokenize gets an array of the words. It is a direct reimplementation
// of https://github.com/thisandagain/sentiment/blob/develop/lib/tokenize.js
func Tokenize(txt string) []string {
	txt = strings.ToLower(txt)
	txt = spacereplace.ReplaceAllString(txt, " ")
	return strings.Fields(txt)
}
