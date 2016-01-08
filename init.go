package pipescript

func init() {
	// Register the built-in transforms
	identityTransform.Register() // The identity is registered by default
}
