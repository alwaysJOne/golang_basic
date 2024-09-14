package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 작업이 완료되면 Done 호출

	fmt.Printf("작업자 %d 시작\n", id)
	time.Sleep(time.Second) // 1초 대기, 실제 작업이 있다고 가정
	fmt.Printf("작업자 %d 완료\n", id)
}

func main() {
	//WaitGroup: 고루틴들이 모두 완료될 때까지 대기하는 데 사용되는 동기화 도구
	//Add(고루틴 추가), Done(작업 완료 알림), Wait(고루틴 종료시 까지 대기)
	//추가 된 고루틴 개수와 완료 알림의 개수가 같아야 정확하게 동작한다. (Add == Done)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) // 각 고루틴을 실행하기 전에 WaitGroup에 1 추가

		go task(i, &wg) // 고루틴 시작
	}

	wg.Wait() // 모든 고루틴이 완료될 때까지 대기
	fmt.Println("모든 작업 완료")
}
