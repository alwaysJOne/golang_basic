package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func weatherAPIData() (string, error) {
	baseURL := "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getUltraSrtNcst"

	params := url.Values{}
	params.Add("serviceKey", "8h1yVur/0C1jzB+QnTVFLYlJk9Gfckktb8MFl25lzXaujB8wU9qgmg5abLreCLJnE2J/OK4DW5bqq0p4FPIEqw==")
	params.Add("numOfRows", "100")                         // 가져올 데이터 개수
	params.Add("pageNo", "1")                              // 페이지 번호
	params.Add("dataType", "JSON")                         // 응답 데이터 형식
	params.Add("base_date", time.Now().Format("20060102")) // 현재 날짜 (YYYYMMDD 형식)
	params.Add("base_time", "0600")                        // 6시 기준 예보 시간
	params.Add("nx", "60")                                 // 서울의 X 좌표 (60)
	params.Add("ny", "127")

	//최종 요청 url
	//포맷된 문자열을 생성하고 그 결과를 반환
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	fmt.Println("요청 전송")
	resp, err := http.Get(finalURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//응답 상태코드 : 200 -> 정상적으로 응답이 됨
	fmt.Println("요청 상태 : ", resp.StatusCode)

	//응답데이터를 문자열로 읽어들이기
	var responseData string
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		responseData += string(line)
	}

	return responseData, nil
}

func main() {

	apiResponse, err := weatherAPIData()
	if err != nil {
		log.Fatal("요청 실패:", err)
		return
	}
	fmt.Println(apiResponse)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, apiResponse)
	})

	// 요청을 처리하는 함수
	// "/" 경로로 요청이 들어오면 호출됩니다.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ResponseWriter는 HTTP 응답을 클라이언트에게 보낼 때 사용하는 인터페이스
		// 서버에서 생성한 응답 데이터를 클라이언트로 전송할 수 있습니다.

		// *http.Request는 클라이언트가 보낸 HTTP 요청을 나타내는 구조체.
		// 요청과 관련된 다양한 정보를 얻을 수 있음.
		// ex) 요청 메서드(GET, POST 등), 요청 URL, 헤더, 바디 등을 확인할 수 있습니다.
		fmt.Fprintf(w, "Hello, Web")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test page")
	})

	fmt.Println("서버 실행... port 8888")
	http.ListenAndServe(":8888", nil)
}
