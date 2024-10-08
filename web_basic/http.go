package main

import (
	"fmt"
	"log"
	"net/http"
)

// 요청을 처리하는 함수
// "/" 경로로 요청이 들어오면 호출됩니다.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter는 HTTP 응답을 클라이언트에게 보낼 때 사용하는 인터페이스
	// 서버에서 생성한 응답 데이터를 클라이언트로 전송할 수 있습니다.

	// *http.Request는 클라이언트가 보낸 HTTP 요청을 나타내는 구조체.
	// 요청과 관련된 다양한 정보를 얻을 수 있음.
	// ex) 요청 메서드(GET, POST 등), 요청 URL, 헤더, 바디 등을 확인할 수 있습니다.
	fmt.Fprintf(w, "Hello, Web!")
}

func main() {
	//ip : 인터넷 상의 장치(컴퓨터, 서버 등)를 구별하는 고유한 주소
	//Port : IP 주소에서 여러 애플리케이션(프로그램)이 통신할 수 있도록 하는 출입구 역할

	// "/" 경로로 들어오는 요청을 helloHandler로 처리
	// ip:port/로 접속
	// localhost는 현재 컴퓨터를 나타내는 호스트 이름입니다.
	// localhost는 항상 127.0.0.1이라는 IP 주소를 가리키며, 자기 자신과 통신할 때 사용됨
	http.HandleFunc("/", helloHandler)

	// 서버 실행
	fmt.Println("Server is running on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil)) // 8080 포트에서 서버 실행
}
