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
func (p Person) MyInfo() {
	fmt.Printf("안녕, 내이름은 %s 이고 나는 %d 살이야.\n", p.name, p.age)
}

type Product struct {
	name   string
	price  int
	width  float64
	height float64
}

func NewProduct(name string, price int, width float64, height float64) *Product {
	return &Product{name, price, width, height} //구조체 인스턴스 생성 후 참조값 리턴
}

func main() {
	//구조체
	//상속, 객체, 클래스 개념 없음
	//데이터를 그룹화하는 데 사용되며, 필드의 집합으로 정의함.
	//필드는 갖지만 메소드는 갖지 않는다.
	//구조체 내에 메서드를 연결할 수 있습니다. 이러한 메서드는 특정 구조체 인스턴스에 바인딩되며,
	//해당 구조체의 동작을 구현함.
	//구조체 -> 구조체 선언 -> 구조체 메서드 연결

	//구조체 인스턴스 선언방법1 -> 참조값 전달
	var kim *Person = new(Person)
	kim.name = "김지우"
	kim.age = 20

	//구조체에 연결된 메서드 호출
	kim.MyInfo()

	kim2 := kim
	kim2.name = "김지연"
	kim.MyInfo()
	kim2.MyInfo()

	fmt.Println("-----------------------------------")

	choi := &Person{name: "최슬기", age: 34}
	fmt.Println(choi)

	choi2 := choi
	choi2.name = "최수지"
	choi.MyInfo()
	choi2.MyInfo()

	fmt.Println("--------------------------------")
	//선언방법2 -> 값전달
	jang := Person{name: "장지수", age: 15}
	//jang2 := jang
	jang2 := &jang

	jang2.name = "장보고"
	fmt.Println(jang)
	fmt.Println(jang2)

	fmt.Println("--------------------------------")

	//선언방법3 익명으로 선언
	tv := struct {
		brand string
		size  int
	}{
		brand: "삼성",
		size:  75,
	}
	fmt.Println(tv)

	fmt.Println("-----------------------------------")

	//구조체의 값들을 초기화하기위한 메서드 -> 생성자
	// p1 := Product{
	// 	name: "smartTV",
	// 	price: 1000000,
	// 	width: 300,
	// 	height: 200,
	// }

	p1 := NewProduct("smartTV", 1000000, 300, 200)
	p2 := NewProduct("TV", 500000, 200, 150)
	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}
