/*
Package numeric contains basic statistical transforms (such as mean)
*/
package numeric

func Register() {
	Sum.Register()
	Count.Register()
	Bucket.Register()

	Mean.Register()

	Max.Register()
	Min.Register()
	/*
		Percent.Register()
	*/
}
