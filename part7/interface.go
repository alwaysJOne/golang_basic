// 인터페이스 기본(1)
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64      //면적
	perimeter() float64 //둘레
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showShapeInfo(s Shape) {
	fmt.Printf("타입: %T\n", s)
	fmt.Printf("면적: %.2f\n", s.area())
	fmt.Printf("둘레: %.2f\n", s.perimeter())
}

func showObject(s interface{}) {
	fmt.Println(s)
}

func main() {
	//인터페이스 -> 메서드들의 집합체
	//단순히 동작에 대한 방법만 표시 -> 실제 어떻게 이루어지는지 관심x

	/*
	  type 인터페이스명 interface {
	     메서드1() 반환 값(타입형)
	  	 메서드2() //반환 값 없을 경우
	   }
	*/

	//값을 할당해주지 않은 인터페이스인 경우 nil 리턴
	var s1, s2 Shape
	fmt.Println(s1, s2)

	//인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
	s1 = Rect{10, 10}
	fmt.Println("사각형 면적  : ", s1.area())
	fmt.Println("사각형 둘레 : ", s1.perimeter())

	s2 = Circle{6.4}
	fmt.Println("삼각형 면적  : ", s2.area())
	fmt.Println("삼각형 둘레 : ", s2.perimeter())

	//하나의 자료구조로 관리할 수도있음
	shapes := []Shape{
		Rect{10, 10},
		Circle{6.4},
		Rect{5, 7},
		Circle{3.2},
	}

	// for _, shape := range shapes {
	// 	fmt.Printf("타입: %T\n", shape)
	// 	fmt.Printf("면적: %.2f\n", shape.area())
	// 	fmt.Printf("둘레: %.2f\n", shape.perimeter())
	// 	fmt.Println()
	// }

	// func showShapeInfo(s  Shape){
	// 	fmt.Printf("타입: %T\n", shape)
	// 	fmt.Printf("면적: %.2f\n", shape.area())
	// 	fmt.Printf("둘레: %.2f\n", shape.perimeter())
	// }

	//덕타이핑 : 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다
	// ->  구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식

	for _, shape := range shapes {
		showShapeInfo(shape)
		fmt.Println()
	}

	//빈 인터페이스 : 어떤 값이든 허용 가능(유연성 극대화)
	// func showObject(s interface{}) {
	// 	fmt.Println(s)
	// }

	showObject(s1)
	showObject(s2)
	showObject(10)
	showObject("circle")
}
