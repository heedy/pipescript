/*
Package strings implements basic text processing transforms
*/
package strings

func Register() {
	Wc.Register()
	Sentiment.Register()
	Append.Register()
	Contains.Register()
	Regex.Register()
}
