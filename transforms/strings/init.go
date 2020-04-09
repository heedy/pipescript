/*
Package strings implements basic text processing transforms
*/
package strings

func Register() {

	Wc.Register()
	Regex.Register()
	Contains.Register()
	Startswith.Register()
	Endswith.Register()
	Urldomain.Register()
}
