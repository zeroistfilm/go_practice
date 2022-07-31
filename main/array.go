package main

func array() {
	//basicArray()
	arrayinit()
}

func basicArray() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	print(a[2])
}

func arrayinit() {
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3} //배열크기 자동으로

	for idx, itm := range a1 {
		println(idx, itm)
	}
	for idx, itm := range a2 {
		println(idx, itm)
	}

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10
	println(multiArray[0][1][2])

	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(a[1][2])
}
