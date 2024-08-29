// 구조체 기본(1)
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	height float64
}

// 리시버 변수 사용
func (p Person) myInfo() {
	fmt.Printf("저는 %v 이며 %v살, 키는  %v입니다.\n", p.name, p.age, p.height)
}

func (p *Person) setName(name string) {
	p.name = name
}

func main() {
	// type Person struct {
	// 	name   string
	// 	age    int
	// 	height float64
	// }

	//구조체
	//여러필드타입을 하나로 묶은 것(사용자 정의 타입)

	//구조체 선언
	var p1 Person
	//각 자료형의 초기값인 공백과 0으로 설정된다.
	fmt.Println(p1)

	p2 := Person{name: "장지수", age: 15}
	fmt.Println(p2)

	p2.height = 175.5
	fmt.Printf("저는 %v 이며 %v살, 키는  %v입니다.\n", p2.name, p2.age, p2.height)

	//주소값전달을 통한 포인터 생성도 가능하다.
	p3 := &p2

	p3.name = "장보고"
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println("--------------------------------")

	//구조체 포인터 선언
	p4 := &Person{name: "최슬기", age: 34}
	fmt.Println(p4)

	p5 := p4
	p5.name = "최수지"
	fmt.Println(p4)
	fmt.Println(p5)

	var p6 *Person = new(Person)
	p6.name = "김지우"
	p6.age = 20
	fmt.Println(p6)

	p7 := p6
	p7.name = "김지연"
	fmt.Println(p4)
	fmt.Println(p7)

	fmt.Println("--------------------------------")

	//익명으로 선언
	tv := struct {
		brand string
		size  int
	}{
		brand: "삼성",
		size:  75,
	}
	fmt.Println(tv)

	//go언어에는 객체지향을 위한 클래스(필드, 메서드)와 객체의 개념이 없음
	//구조체 -> 필드는 갖지만 메소드는 갖지 않는다.
	//리시버 -> 구조체 내에 메서드를 연결할 수 있습니다. 이러한 메서드는 구조체 인스턴스에 바인딩되어 동작을 구현함

	// func (p Person) MyInfo() {
	// 	fmt.Printf("저는 %v 이며 %v살, 키는  %v입니다.\n", p.name, p.age, p.height)
	// }
	p8 := &Person{name: "김삼식", age: 20, height: 180}

	//구조체에 연결된 메서드 호출
	p8.myInfo()

	// func (p *Person) setName(name string) {
	// 	p.name = name
	// }
	//값을 변경해야하는 메서드를 연결할때는 포인터리시버를 사용
	p8.setName("김두식")
	p8.myInfo()
}
