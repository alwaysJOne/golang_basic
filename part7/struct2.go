// 구조체 심화(1)
package main

import "fmt"

type Human struct { // 사람 구조체 정의
	name string
	age  int
}

func (p *Human) greeting() { // greeting 함수 Human구조체에 연결
	fmt.Println("Hello")
}

func newHuman(name string, age int) *Human {
	return &Human{name, age} //구조체 인스턴스 생성 후 참조값 리턴
}

type Student struct {
	Human // 임베딩( Is-a )
	grade int
}

func main() {
	// type Product struct {
	// 	brand string
	// 	price int
	// 	inch  float64
	// }
	//구조체의 생성자 패턴추가
	var h1 *Human = new(Human) //new 함수로 초기화와 동시에 구제체 필드 값 할당은 불가능
	h1.name = "최지원"
	h1.age = 25

	h2 := &Human{
		name: "최지투",
		age:  30,
	}

	fmt.Println(h1, h2)

	fmt.Println("-----------------------------------")

	//구조체의 값들을 초기화하기위한 메서드 -> 생성자
	// func NewProduct(brand string, price int, inch float64) *Product {
	// 	return &Product{brand, price, inch} //구조체 인스턴스 생성 후 참조값 리턴
	// }

	h3 := newHuman("최지삼", 22)
	h4 := newHuman("최지사", 23)
	fmt.Println(h3, h4)

	//기존 클래스의 모든 필드와 메소드를 다른클래스에서 사용할 수 있게 해주는
	//상속이라는 개념을 구현  -> 구조체 임베딩(상속과같은 직접연결은 아니지만 유사하게 구현하는 방법)

	// type Student struct {
	// 	h      Human // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음( Has - a )
	// 	grade  int
	// }

	// var s1 Student
	// s1.h.greeting()

	//임베딩
	//필드명은 사용하지 않고, 구조체 타입만 지정했다. 이렇게 되면 구조체가 해당 타입을 포함한다.
	// type Student struct {
	// 	Person  // 임베딩( Is-a )
	// 	school string
	// 	grade  int
	// }

	var s Student
	s.name = "최지오"
	s.age = 17
	s.grade = 5

	s.greeting() // Hello

	// func (p *Student) greeting() {    // 이름이 같은 메서드를 정의하면 오버라이딩 됨
	//fmt.Println("Hello Students~")
	//}
}
