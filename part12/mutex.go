package main

import (
	"fmt"
	"sync"
)

func increment(count *int, mutex *sync.Mutex) {
	mutex.Lock()   // Mutex 잠금 (다른 고루틴 접근 차단)
	*count += 1    // count 값을 증가
	mutex.Unlock() // Mutex 잠금 해제 (다른 고루틴 접근 허용)
}

func main() {
	//경쟁조건 : 여러 고루틴이 동시에 같은 메모리 공간에 접근하고 수정하려 할 때, 예기치 않은 결과가 발생하는 문제
	count := 0

	var mutex sync.Mutex

	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		//하나의 고루틴이 count를 증가시키기 전에 다른 고루틴이 개입하여 동일한 메모리 공간을 변경할 수 있음
		//일부 업데이트가 누락될 수 있음
		go func() {
			increment(&count, &mutex)
			done <- true
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	fmt.Println(count)
}

//성능 저하: 모든 고루틴이 크리티컬 섹션에 접근할 때 대기해야 하기 때문에, 성능이 저하될 수 있습니다.
//교착 상태(데드락): 잘못된 Mutex 사용으로 잠금이 해제되지 않으면 프로그램이 멈추거나 교착 상태가 발생할 수 있습니다.
