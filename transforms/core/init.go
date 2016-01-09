/*
Package core contains the core basic transforms which are available in PipeScript.
It should be imported by default by basically all users of PipeScript
*/
package core

func init() {
	first.Register()
	last.Register()

	//prev.Register()
	next.Register()

	count.Register()

	tTransform.Register()
	tshift.Register()

	split.Register()
}
