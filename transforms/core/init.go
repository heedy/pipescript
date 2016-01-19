/*
Package core contains the core basic transforms which are available in PipeScript.
It should be imported by default by basically all users of PipeScript
*/
package core

func Register() {
	Changed.Register()

	First.Register()
	Last.Register()

	//prev.Register()
	Next.Register()

	Count.Register()
	Sum.Register()
	Mean.Register()

	T.Register()
	Tshift.Register()

	Reduce.Register()
}
