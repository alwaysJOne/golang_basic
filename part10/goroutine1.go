// 고루틴(Goroutine)기초(1)
package main

import (
	"fmt"
	"runtime"
	"time"
)

func testGo1() {
	fmt.Println("exe1함수 실행")
	time.Sleep(1 * time.Second)
	fmt.Println("exe1함수 종료")
}

func testGo2() {
	fmt.Println("exe2함수 실행")
	time.Sleep(1 * time.Second)
	fmt.Println("exe2함수 종료")
}

func testGo3() {
	fmt.Println("exe3함수 실행")
	time.Sleep(1 * time.Second)
	fmt.Println("exe3함수 종료")
}

func printNumber(name string) {
	for i := 1; i <= 1000; i++ {
		fmt.Println(name, " : ", i)
	}
}

func main() {
	//고루틴
	//타 언어의 쓰레드와 비슷한 기능을 함
	//매우 작은 메모리 스택으로 시작하며 필요에따라 조정함
	//즉, 수많은 고루틴 동시 실행 가능

	//testGo1() //가정 먼저 실행(일반적인 실행 흐름)

	//go루틴을 실행할 때는 go키워드를 앞에 붙여줌
	fmt.Println("Main Routine Start : ", time.Now())
	//go testGo2()
	//go testGo3()

	///////------------------
	//GOMAXPROCS 환경 변수를 사용하여 CPU 코어 수만큼 병렬적으로 고루틴을 실행할 수 있임
	//멀티 코어(다중 CPU) 최대한 활용
	//고루틴이 여러 CPU 코어에서 병렬로 실행되면,
	//고루틴 간의 실행 순서는 더욱 예측하기 어려진다

	// 현 시스템의 CPU 코어 개수 반환 후 설정
	runtime.GOMAXPROCS(runtime.NumCPU())
	///////

	//스케줄러는 고루틴의 실행을 선점형 방식으로 관리하기 때문에,
	//고루틴이 항상 순차적으로 실행되지 않고 중간에 다른 고루틴으로 전환될 수 있다.
	go printNumber("funk1")
	go printNumber("funk2")
	time.Sleep(5 * time.Second) // 5초 대기 (주석처리하면 별도 고루틴 종료 : 메인함수 종료 시 모두 종료)
	fmt.Println("Main Routine End : ", time.Now())
}
