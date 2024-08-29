package main

import "fmt"

func main() {
	//배열
	//동일한 데이터 타입을 가진 요소들이 일정한 순서로 나열된 고정된 크기의 데이터 구조

	//배열의 생성방법
	var arr1 [5]int              //기본적으로 값은 0 초기화
	fmt.Println(arr1, len(arr1)) //len() : 배열의 길이

	//배열의 값 할당방법 -> 인덱스를 이용한다
	arr1[0] = 1
	arr1[2] = 3
	fmt.Println(arr1, len(arr1))

	//배열의 접근방법 -> 인덱스를 이용한다
	fmt.Println(arr1[0], arr1[1])

	//배열을 선언과 동시에 초기화해서 사용할 수도 있다.
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr2, len(arr2))
	fmt.Println(arr3, len(arr3))
	fmt.Println(arr4, len(arr4))
	fmt.Println(arr5, len(arr5))
	fmt.Println(arr6, len(arr6))

	//여러가지 타입을 사용할 수 있다.
	arr7 := [...]string{"one", "two", "three"}

	fmt.Println(arr7, len(arr7))

	//배열 전체에 순차적으로 접근하는 방법
	//길이가 30인 배열에 1부터 30까지 순차적으로 저장

	//길이를 이용한 접근
	var arr8 [30]int
	for i := 0; i < len(arr8); i++ {
		arr8[i] = i + 1
	}

	fmt.Println(arr8, len(arr8))

	//for range를 이용한 방법
	for i, v := range arr8 {
		fmt.Println("index : ", i, "arr8[i] : ", v)
	}

	//값만 필요할 땐 생략해도 된다. or블랭크식별자
	for v := range arr8 {
		fmt.Println("arr8[i] : ", v)
	}

	//배열사용시 주의사항
	//go에서 배열은 참조타입이 아닌 값타입이다
	arr9 := [5]int{1, 10, 100, 1000, 10000}
	arr10 := arr9

	arr10[0] = 1000

	fmt.Println(arr9)
	fmt.Println(arr10)
}
