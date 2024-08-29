package main

import "fmt"

func main() {
	//상수
	//const 생성하며 초기화, 한 번 선언 후 값 변경 금지 , 고정 된 값 관리 용
	const num int = 10

	fmt.Println(num)
	// num = 20 상수는 값을 변경할 수 없다.

	//여러개를 한번에 생성하는 방법
	// const (
	// 	num  int = 10
	// 	num2 int = 20
	// )
	const num2, num3 int = 10, 20

	//짧은 선언
	// name := 10 -> var
	const num4 = 40
	const num5, num6 = 50, 60

	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 증가시키는 방식

	// const (
	// 	n1 = 1
	// 	n2 = 2
	// 	n3 = 3
	// 	n4 = 4
	// 	n5 = 5
	// )

	// const (
	// 	n1 = iota
	// 	n2 = iota
	// 	n3 = iota
	// 	n4 = iota
	// 	n5 = iota
	// )

	const (
		n1 = 10 + (iota * 10)
		n2
		n3
		n4
		n5
	)
	fmt.Println(n1, n2, n3, n4, n5)
}
