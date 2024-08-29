// 함수 Defer(1)
package main

import "fmt"

func deferTest1() {
	defer deferTest2() //마지막에 호출
	fmt.Println("test1함수가 실행되었습니다.")
}

func deferTest2() {
	fmt.Println("test2함수가 실행되었습니다.")
}

func deferTest3() {
	defer func() {
		fmt.Println("test3함수가 이제 끝납니다.")
	}()
	fmt.Println("test3함수가 실행되었습니다.")
}

func main() {
	//Defer 함수 실행(지연)
	//Defer를 호출한 함수의 마지막에 호출 된다.
	//주로 리소스 반환 등에 사용

	deferTest1()
	deferTest3()
}
