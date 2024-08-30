package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	//외부패키지 : 다른 개발자들이 만들어놓은 기능 -> 라이브러리
	//go get명령어를 통해 모듈에 등록 후 사용가능
	//go get을 통해 모듈에 등록하면 GOPATH/pkg폴더에 다운로드가 됨

	//go 모듈 : 프로젝트의 종속성을 관리하기 위한 시스템
	//go.mod파일을 사용하여 프로젝트의 모듈이름과 종속성을 정의하고 go.sum파일을 사용하여 버전을 관리함

	//go mod tidy를 사용하여 필요없는 패키지를 제거할 수 있음

	//uuid : 고유한 식별자를 생성하는 기능을 제공해주는 패키지
	// 새로운 UUID 생성
	newUUID := uuid.New()
	fmt.Printf("%T : %v\n", newUUID, newUUID)

	// UUID 문자열로 변환
	uuidStr := newUUID.String()
	fmt.Printf("%T : %v\n", uuidStr, uuidStr)
}
