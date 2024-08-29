package main

import "fmt"

func main() {
	//슬라이스
	//배열을 기반으로 하지만 가변적으로 길이와 용량을 관리할 수 있다.

	//슬라이스의 생성방법1 -> 배열과 유사하게
	var slice1 []int // 길이와 용량이 가변적이므로 정해주지 않아도된다. 길이 용량 0
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1, len(slice1), cap(slice1)) //cap() : 용량을 확인
	fmt.Println(slice2, len(slice2), cap(slice2))

	//슬라이스의 값 할당
	//slice1[0] = 1  //예외 발생 -> 가변형이기 때문에 아직 메모리공간을 가지고있지 않음
	slice2[0] = 10 //값 수정 가능

	//길이를 동적으로 늘리고 싶을 때는 append사용
	slice1 = append(slice1, 1, 2, 3)

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))

	//슬라이스의 접근방법 -> 인덱스를 이용한다
	fmt.Println(slice2[0], slice2[1])

	//슬라이스의 생성방법2-> make함수를 이용한다
	//make(자료형, 길이, 용량) //용량은 생략시 길이와 동일
	var slice3 []int = make([]int, 5)
	var slice4 = make([]int, 5)

	//용량만큼 메모리를 확보하고있음. 해당용량까지는 메모리 재할당없이 사용가능
	//효율적인 메모리관리 가능
	slice5 := make([]int, 5, 10)
	slice5 = append(slice5, 1, 2, 3, 4, 5, 6)

	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))
	fmt.Println(slice5, len(slice5), cap(slice5))
	fmt.Println("-----------------------------------------------------------------------------")

	//슬라이스는 배열과같은 값타입이 아닌 참조타입이다.

	//배열 : 값타입
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	arr1[0] = 1000

	fmt.Println(arr1)
	fmt.Println(arr2)

	//슬라이스 참조타입
	sliceRef1 := []int{1, 10, 100, 1000, 10000}
	sliceRef2 := sliceRef1

	sliceRef1[0] = 1000

	fmt.Println(sliceRef1)
	fmt.Println(sliceRef2)

	var sliceRef3 []int //nil 슬라이스 (길이와 용량 0)
	fmt.Println(sliceRef3 == nil)

	//슬라이스 복사방법
	//기존에 복사하고자하는 slice와 같은 크기의 slice를 생성 후 직접 값을 복사해 준다.
	sliceRef4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRef5 := make([]int, len(sliceRef4))

	for i, v := range sliceRef4 {
		sliceRef5[i] = v
	}

	//copy(sliceRef5, sliceRef4)

	fmt.Println("ex1 : ", sliceRef4)
	fmt.Println("ex1 : ", sliceRef5)

	fmt.Println()
}
