package basic

import "fmt"

func Logical() {
	i := 1
	max := 10
	if val := i * 2; val < max {
		println("in loop")
	}

	var name string
	var category = 1

	switch category {
	case 1:
		name = "paperbook"
	case 2:
		name = "ebook"
	case 3, 4:
		name = "blog"
	default:
		name = "other"

	}
	println(name)

	category = 101 //101
	switch x := category >> 2; x - 1 {
	default:
		println(x)
	}

	typecheck(1)

}

func typecheck(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int", v)
	case bool:
		fmt.Printf("%v is bool", v)
	case string:
		fmt.Printf("%v string", v)

	}

}
