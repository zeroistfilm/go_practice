package main

import (
	"fmt"
	"time"
)

func channel_test() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)
}

func unbufferedchannel_wait() {
	//Unbuffered Channel로서 이 채널에서는 하나의 수신자가 데이타를 받을 때까지
	//송신자가 데이타를 보내는 채널에 묶여 있게 된다.
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		done <- true
	}()
	<-done
}

func bufferedChannel1() {
	// 당장의 수신자가 없더라도 버퍼만큼은 데이터를 보낼 수 있다.
	c := make(chan int, 1)
	c <- 1
	println(<-c)

}

func bufferedChannel2() {
	done := make(chan bool, 10)
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
			done <- true
		}

	}()
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)
	println(<-done)

}

func sendReceiveChannel() {
	ch := make(chan string, 1)
	sendChan(ch, "hihihi")
	receiveChan(ch)

}

func sendChan(ch chan<- string, data string) {
	ch <- data
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func channelClose() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	close(ch)

	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("데이터 없음")
	}
}

func channelRange() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	close(ch)

	//for {
	//	if i, success := <-ch; success {
	//		println(i)
	//	} else {
	//		break
	//	}
	//}
	for i := range ch {
		println(i)
	}

}

func channelSelect() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	go run1(done1)
	go run2(done2)

EXIT:
	for {
		println("looping..")
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
