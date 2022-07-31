package main

import "fmt"

func slice() {
	var a []int
	a = []int{1, 2, 3}
	b := []int{4, 5, 6}
	a[1] = 10
	fmt.Println(a, b)
	println(a, b)
	// fmt println vs println ?

	s := make([]int, 5, 10) //슬라이스 타입, 슬라이스 길이, 내부내열 최대 길이
	s[4] = 55555555556666666
	println(len(s), cap(s))
	fmt.Println(s)

	ss := []int{1, 2, 3, 4, 5}
	ss = ss[2:5]
	fmt.Println(ss)

	sss := []int{0, 1, 2, 3, 4, 5}
	s1 := sss[2:5] //sss와 주소값 다름
	s2 := sss      //sss와 주소값 같음
	s3 := sss[:]   //sss와 주소값 같ㄴ
	println(sss, s1, s2, s3)
	//[6/6]0xc000147ed8
	//[3/4]0xc000147ee8
	//[6/6]0xc000147ed8
	//[6/6]0xc000147ed8

	k := []int{1, 2}
	k = append(k, 3)
	k = append(k, 4, 5, 6, 7, 8)

	fmt.Println(k)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}
	sliceB = append(sliceB, sliceC...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceB)

	source := []int{4, 1, 6}
	fmt.Println(source)
	println(len(source), cap(source))

	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))
	//슬라이스는 실제 배열을 가르키는 포인터 정보만 가진다.
	//복사를 좀 더 명확히 표현하면, 소스 슬라이드가 갖는 배열의 데이터를 타겟 슬라이그가 가지는
	//배열로 복제하는 것이다.

	//슬라이스는 내부적으로 사용하는 배열의 부분영역인 세그먼트에 대한 정보를 가진다.
	//슬라이스는 3개의 필드를 가진다.

	println(source)
	println(source[1:2])
}
