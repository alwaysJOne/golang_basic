// 자료형 : 포인터(1)
package main

import "fmt"

func ptest(p *int) {
	*p = 9999
}

func vtest(v int) {
	v = 9999
}

func main() {
	//포인터
	//메모리주소를 가르키는 것
	//포인터변수 -> 메모리 주소를 값으로 가지는 변수(참조변수)
	//불필요한 메모리낭비를 막을 수 있다.
	//포인터를 선언할 때 애스터리스크(*) 사용

	//포인터  선언
	var p1 *int
	var p2 *int = new(int)

	//참조하고있는 값이  없을 때는  nil을 반환함
	fmt.Println(p1)
	fmt.Println(p2, *p2)

	num1 := 10

	//&연산자 -> 주소값을 반환
	p1 = &num1
	p2 = &num1

	fmt.Println(p1, *p1)
	fmt.Println(p1, *p2)

	//p1이 가르키는 주소로 접근해서 값을 100  넣어줍니다.
	//p1 = 100 //컴파일에러
	//p1++ // 컴파일에러 포인터 연산x
	*p1 = 100

	fmt.Println(*p1, *p2, num1)

	//주소값을 반환받아 변수생성시 포인터변수로 생성됨
	var p3 = &num1
	p4 := &num1

	fmt.Println(*p1, *p2, *p3, *p4, num1)

	//함수에 인자를 전달할 때 포인터를 전달하면 원본에 영향을 준다
	num2, num3 := 10, 10

	ptest(&num2)
	vtest(num3)

	fmt.Println(num2, num3)

}
