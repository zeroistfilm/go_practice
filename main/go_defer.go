package main

import (
	"fmt"
	"sync"
)

func go_defer() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer func() {
			fmt.Println("worker 1 end defer first ")
		}()
		defer func() {
			fmt.Println("worker 1 end defer second")
		}()
		defer func() {
			fmt.Println("worker 1 end defer therd")
		}()
		test1()
		defer wg.Done()

	}()
	go func() {
		test2()
		defer wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i, "test")
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func test1() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "hello")
	}
}

func test2() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "world")
	}
}
