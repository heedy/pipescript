package interpolators

func Register() {
	// Register built-in interpolators
	After.Register()
	Before.Register()
	Closest.Register()

	Count.Register()
	Sum.Register()
	Mean.Register()
}
