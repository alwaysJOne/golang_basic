package main

//패키지명을 선언하는 것
//코드를 모듈화해서 재사용할 수록 만들어 두는 것
//패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행프로그램으로 만든다.
// main 패키지 안의 main() 함수가 프로그램의 시작점, 즉 Entry Point가 된다.
//패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안됨

//패키지에 작성된 코드를 해당 프로그램에서 사용하기위해서는
//import해줘야함
import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Web Server!")
}

func main() {
	//패키지 이름 = 디렉터리 이름
	//패키지명 = 소스파일명

	//1.내장패키지 : GO에서 기본적으로 지원하는 패키지
	//GOPATH에 src폴더에 위치함

	/*
		//fmt : 입출력패키지로 입출력을 처리하는 기능제공
		var name string

		fmt.Print("이름 :  ")    //출력기능
		fmt.Scanf("%s", &name) //입력기능

		//os.Stdout : 운영체제 표준콘솔출력
		fmt.Fprintf(os.Stdout, "Hi %s\n", name)

		//os : 운영체제와의 상호작용을 위한 다양한 기능을 제공
		// 파일 생성
		file, err := os.Create("newfile.txt")
		if err != nil {
			fmt.Println("파일 생성 오류:", err)
			return
		}
		defer file.Close()
		fmt.Println("파일이 생성되었습니다.")

		// 파일 삭제
		err = os.Remove("newfile.txt")
		if err != nil {
			fmt.Println("파일 삭제 오류:", err)
			return
		}
		fmt.Println("파일이 삭제되었습니다.")

	*/

	//math : 수학관련 기능들을 제공하는 패키지
	fmt.Println("PI : ", math.Pi)

	fmt.Println("2의 3제곱:", math.Pow(2, 3))

	//strings : 문자열 처리와 관련된 기능을 제공하는 패키지
	// 문자열 포함 여부 확인
	fmt.Println(strings.Contains("Hello, World!", "World"))

	// 문자열을 대문자로 변환
	fmt.Println(strings.ToUpper("go language"))

	// 문자열을 특정 구분자로 분할
	words := strings.Split("apple,banana,cherry", ",")
	fmt.Println(words)

	//net/http : 웹관련 기능을 제공하는 패키지
	http.HandleFunc("/", handler) // 루트 경로에 대한 요청을 처리할 핸들러 등록
	fmt.Println("웹 서버가 :5000에서 시작됩니다...")
	log.Fatal(http.ListenAndServe(":5000", nil)) // 웹 서버 시작
}
