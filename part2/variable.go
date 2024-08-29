package main

import "fmt"

func main() {
	//변수를 만드는 방법
	//키워드 변수명 타입
	var num int

	//변수명 : 시작은 문자,_만 가능, 대소문자구분, 키워드사용x, 문자와 숫자 _만 가능
	//카멜케이스 : 소문자로 시작해서 단어가 바뀔때 첫글자를 대문자 ex) firstNum
	//스네이크케이스 : 소문자로 작성하되 단어가 바뀔 때 _로 구분 ex) first_num

	//변수에 값을 대입하는 방법
	num = 10

	//변수에 있는 값을 가져오는 방법
	fmt.Println("숫자 :", num)

	num = 20
	fmt.Println("숫자 :", num)

	//변수를 선언과 동시에 값을 초기화
	var num2 int = 20
	fmt.Println("숫자2 :", num2)

	//값을 바로 초기화할 때 짧은 선언
	// 변수명 := 초기값
	num3 := 30

	fmt.Println("숫자3 :", num3)

	//여러개를 한번에 선언
	//같은 타입인 경우
	var (
		num4 int
		num5 int = 50
	)

	num6, num7 := 60, 70
	fmt.Println(num4, num5, num6, num7)
}
