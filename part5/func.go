package main

import "fmt"

//어디에서든 선언할 수 있다.
func multiple(num int) (res int) {
	// return (num * 2)
	res = (num * 2)
	return
}

func goStart() {
	fmt.Println("hellow go")
}

func multipleResult(n1 int, n2 int) (int, int) {
	return n1 * 2, n2 * 2
}

func printMultipe(num int, cal func(int) int) {
	res := cal(num)
	fmt.Println(res)
}

func totalCal(num ...int) int {
	total := 0
	for i, n := range num {
		fmt.Println(i, "번째 값 : ", n)
		total += n
	}

	return total
}

func main() {
	//함수
	//func 함수명(매개변수) (반환타입 or 반환 값 변수명){
	// 		함수내용
	//		return 결과값
	// }
	//매개변수 : 함수내에서 필요한 값을 전달하기위한 값
	//반환값 : 함수의 결과값
	res := multiple(5)
	fmt.Println(res)
	fmt.Println(multiple(10))

	//func 함수명() {
	//	실행코드
	//}
	goStart()

	//타 언어와 달리 반환 값(return value) 여러 개 가능
	res1, res2 := multipleResult(3, 5)
	//res3 := multipleResult(4, 7) //예외 발생 -> 리턴된 모든 수를 받아줘야한다.
	res3, _ := multipleResult(4, 0)
	fmt.Println(res1, res2, res3)

	//가변적으로 인자의 개수를 달리할 수 있다.
	fmt.Println("모두 더한 결과값 : ", totalCal(10, 20, 30, 40, 50))

	//함수를 인자로 전달할 수 있다.
	printMultipe(50, multiple)

	//함수를 변수에 할당할 수 있다.
	showMyInfo := func() {
		fmt.Println("제 이름은 최지원입니다.")
		fmt.Println("저는 많은 분들에게 고언어를 전파하고 싶습니다.")
	}

	showMyInfo()
	showMyInfo()
	showMyInfo()

	//익명함수
	//대체로 어떠한 값에 함수를 전달해야할 때 사용한다.
	func() {
		fmt.Println("제 이름은 최지원입니다.")
		fmt.Println("저는 많은 분들에게 고언어를 전파하고 싶습니다.")
	}()
}
