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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Web!")
}

func weatherAPIData() (string, error) {
	// API의 기본 URL (초단기 실황 조회)
	baseURL := "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getUltraSrtNcst"

	// API 요청에 필요한 파라미터 설정
	params := url.Values{}
	params.Add("serviceKey", "8h1yVur/0C1jzB+QnTVFLYlJk9Gfckktb8MFl25lzXaujB8wU9qgmg5abLreCLJnE2J/OK4DW5bqq0p4FPIEqw==") // 서비스 키
	params.Add("numOfRows", "10")                                                                                        // 가져올 데이터 개수
	params.Add("pageNo", "1")                                                                                            // 페이지 번호
	params.Add("dataType", "JSON")                                                                                       // 응답 데이터 형식
	params.Add("base_date", time.Now().Format("20060102"))                                                               // 현재 날짜 (YYYYMMDD 형식)
	params.Add("base_time", "0600")                                                                                      // 6시 기준 예보 시간
	params.Add("nx", "60")                                                                                               // 서울의 X 좌표 (60)
	params.Add("ny", "127")                                                                                              // 서울의 Y 좌표 (127)

	// 최종 URL 생성
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// API 호출 (HTTP GET 요청)
	resp, err := http.Get(finalURL)
	if err != nil {
		return "", err // 에러 발생 시 빈 응답 반환
	}
	defer resp.Body.Close()

	// HTTP 응답 코드 출력 (200이면 정상)
	fmt.Println("Response code:", resp.StatusCode)

	// 응답 데이터를 문자열로 읽어들이기
	var responseData string
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF { // 데이터를 모두 읽으면 종료
				break
			}
			return "", err // 에러가 발생하면 반환
		}
		responseData += string(line) // 응답 데이터를 추가
	}

	return responseData, nil // 응답 데이터와 nil 반환
}

func main() {
	data, err := weatherAPIData() // API 데이터 호출
	if err != nil {
		log.Fatalf("Failed to fetch weather data: %v", err)
	}
	fmt.Println("Weather Data:", data) // API 응답 데이터 출력

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
