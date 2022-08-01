package basic

func Variable_type() {

	var a int
	var f float32 = 11.

	a = 10
	f = 12.0
	println(a, f)

	var i, j, k int = 1, 2, 3
	println(i, j, k)

	var ii = 1
	var s = "hi"

	println(ii, s)

	iii := 123 //Short Assignment Statement
	println(iii)

	const ci int = 10
	const cs string = "HI"

	const cii = 10
	const css = "Hi"

	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	println(ci, cs, cii, css, Visa, Master, Amex)
}
