package basic

func Closure() {
	// 클로저는 함수 바깥에 있는 변수를 참조하는 함수값
	// 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯 읽거나 쓸 수 있다.

	next := nextValue()
	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
	println(anotherNext())
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
