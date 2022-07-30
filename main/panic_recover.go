package main

import "fmt"

func panic_recover() {
	panicTest()
}

func panicTest() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			panicTest()
		}

	}()

	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	result := num1 / num2

	fmt.Println(result)
}
