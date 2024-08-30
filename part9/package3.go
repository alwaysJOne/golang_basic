package main

import (
	math "realCode/utils" //모듈명/패키지명으로 가져옴
)

func main() {
	//사용자가 직접 작성하여 프로젝트에서 사용하는 패키지
	//프로젝트의 특정 기능을 구현 -> 다른 파일이나 프로젝트에서 재활용가능
	//모듈을 기준으로 접근해서 사용가능함
	//디렉터리명 = 패키지명

	//utils.PrintAdd(1, 3)
	math.PrintAdd(5, 10) //별칭사용가능

	//test.printMinus(10, 5) //접근불가
}
