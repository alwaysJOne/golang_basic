package utils

import (
	"fmt"
)

var PI float64

// init : 패키지 로드 시에 가장 먼저 호출
// 가장 먼저 초기화 되는 작업 작성 시 유용하다.
func init() {
	fmt.Println("초기화함수 실행")
}

func init() {
	PI = 3.14
}
func init() {
	fmt.Println(PI)
}
