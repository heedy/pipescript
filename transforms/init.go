/*
Package transforms contains the core basic transforms which are available in PipeScript.
It should be imported by default by basically all users of PipeScript
*/
package transforms

func init() {
	first.Register()
	last.Register()

	//prev.Register()
	next.Register()

	i.Register()
}
