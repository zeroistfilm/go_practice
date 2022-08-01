package basic

func Forloop() {

	n := 1
	for n < 200 {
		n *= 2
		println(n)
	}

	names := []string{"홍길동", "이순신", "강감찬"}
L1:
	for index, item := range names {
		println(index, item)
		if index == 1 {
			break L1
		}

	}
	println("l1 end")

}
