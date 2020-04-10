/*
Package core contains the core basic transforms which are available in PipeScript.
It should be imported by default by basically all users of PipeScript
*/
package core

func Register() {
	Filter.Register()
	I.Register()
	Map.Register()
	Reduce.Register()
	While.Register()
}
