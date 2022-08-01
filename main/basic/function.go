package basic

func Function() {
	say("hello")

	msg := "Hello"
	valueChange(&msg)
	println(msg)

	variablefunc("this", "is", "a", "variable func")
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	println(sumCount(1, 23, 4, 5, 6, 2345, 845, 6546))
	println(sumCountReturn(1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6))
}

func say(msg string) {
	println(msg)
}

func valueChange(msgPointer *string) {
	//전달되는 값은 주소일건데, 실제 사용하는 값은 스트링일 것이야.
	//그래서 내부에서 사용할때도 *을 사용해서 메모리에서 꺼내써야한다.
	println(*msgPointer)
	*msgPointer = "changed"
}

func variablefunc(itemlist ...string) {
	for _, s := range itemlist {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sumCount(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return s, count
}

func sumCountReturn(nums ...int) (sum int, count int) {
	println("infunc sum init", sum)
	println("infunc count init", count)
	for _, n := range nums {
		sum += n
		count++
	}
	count = len(nums)
	return
}
