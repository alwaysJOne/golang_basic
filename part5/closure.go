package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	//클로저 함수 : 함수와 그 함수가 접근할 수 있는 변수를 함께 기억하는 특별한 함수
	//클로저는 자신이 만들어질 때의 변수 상태를 기억한다.

	// multiple := func(num int) int {
	// 	return num * 2
	// }

	// res1 := multiple(10)
	// fmt.Println(res1)

	num2 := 2
	multiple := func(num int) int {
		return num * num2
	}

	//함수가 선언된 시점에 외부에 있는 변수를 함수내에서 사용한다면 함수가 정의되는 순간 변수를 기억해서 사용한다.
	//
	fmt.Println(multiple(10))

	//num의 값이 계속해서 증가하는 것은
	//함수가 정의되는 시점에 변수의 정보를 기억하고 있기 때문이다.
	//함수를 호출할 때 값을 지속적으로 유지하기위해서 사용한다. -> 무분별하게 사용시 메모리를 과하게 사용할 수 있다.
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
	fmt.Println(makeCounter())
}
