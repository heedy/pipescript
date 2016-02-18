package interpolators

func Register() {
	// Register built-in interpolators
	after.Register()
	before.Register()
	closest.Register()
}
