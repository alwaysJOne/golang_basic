package main

import (
	"fmt"
)

func main() {
	//사용자로부터 입력을 받는 방법
	var name string
	var age int

	fmt.Print("이름과 나이를 입력하세요: ")
	// fmt.Scan(&name, &age) // 공백과 개행문자를 통해서 구분한다.
	// fmt.Scanln(&name, &age) // 공백을 통해서만 구분하고 개행문자를 만나면 종료한다
	fmt.Scanf("%s\n%d", &name, &age) //입력형식을 내 마음대로 설정할 수 있다.

	fmt.Printf("입력된 이름: %s\n", name)
	fmt.Printf("입력된 나이: %d\n", age)
}
