package basic

import "fmt"

func Map_() {
	//map[key type]value type
	//var idmap map[int]string: 이 형식은 nil값을 가지며 데이터를 쓸 수 없다. 초기화가 필요하다.

	idmap := make(map[int]string)
	// make 함수의 첫번쨰 파라미터로 맵 키워드와 키,값 타입을 적고
	// 이때 make함수는 해시테이블 자료구조를 메모리에 생성, 그 위치를 가르키는 map value를 리턴한다.
	// runtime.hmap 구조체를 가르키는 포인터이다.

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"TSLA": "Tesla",
	}

	idmap[0] = "hello"
	idmap[1] = "world"
	idmap[100] = "apple"

	println(idmap)
	println(tickers)
	fmt.Println(idmap)
	fmt.Println(tickers)

	println(idmap[0])
	fmt.Println(idmap[0])
	println(idmap[200])
	fmt.Println(idmap[200])

	fmt.Println(idmap[100])
	delete(idmap, 100)
	fmt.Println(idmap[100])

	val, isExist := tickers["MSFT"]
	if !isExist {
		println("not exist")
	} else {
		println(val)
	}

	for key, val := range tickers {
		fmt.Println(key, val)
	}
}
