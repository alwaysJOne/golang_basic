package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// WeatherItem 구조체는 날씨 정보를 저장합니다.
// Category는 한글로 번역된 날씨 항목, ObsValue는 그에 해당하는 관측 값입니다.
type WeatherItem struct {
	Category string `json:"category"`  // 날씨 항목 (예: 기온, 습도 등)
	ObsValue string `json:"obsrValue"` // 관측 값 (예: 온도, 습도 값)
}

// ApiResponse 구조체는 API 응답 전체를 저장합니다.
// 응답 코드와 실제 데이터를 포함합니다.
type ApiResponse struct {
	ResponseCode int         `json:"responseCode"` // HTTP 응답 코드
	Data         interface{} `json:"data"`         // 날씨 데이터 (WeatherItem 목록)
}

// 날씨 항목 카테고리와 한글 번역을 위한 매핑
// API에서 받은 카테고리 코드를 한국어로 변환합니다.
var categoryTranslation = map[string]string{
	"PTY": "강수형태",    // PTY: 강수형태 (비, 눈 등)
	"REH": "습도",      // REH: 습도
	"RN1": "1시간 강수량", // RN1: 1시간 강수량
	"T1H": "기온",      // T1H: 기온
	"UUU": "동서바람성분",  // UUU: 동서바람성분
	"VEC": "풍향",      // VEC: 풍향
	"VVV": "남북바람성분",  // VVV: 남북바람성분
	"WSD": "풍속",      // WSD: 풍속
}

// API로부터 날씨 데이터를 가져오는 함수
// API에서 받은 JSON 데이터를 WeatherItem 구조체로 변환합니다.
func fetchAPIData() (ApiResponse, error) {
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
		return ApiResponse{}, err // 에러 발생 시 빈 응답 반환
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
			return ApiResponse{}, err
		}
		responseData += string(line)
	}

	// JSON 데이터를 파싱하여 구조체로 변환
	var rawResponse map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &rawResponse)
	if err != nil {
		return ApiResponse{}, err
	}

	fmt.Println(rawResponse) // 원시 응답 출력 (디버깅용)

	// 날씨 데이터를 가공하여 WeatherItem 구조체 목록으로 변환
	var weatherItems []WeatherItem
	if body, ok := rawResponse["response"].(map[string]interface{})["body"].(map[string]interface{})["items"].(map[string]interface{})["item"].([]interface{}); ok {
		for _, item := range body {
			if itemData, ok := item.(map[string]interface{}); ok {
				category := itemData["category"].(string)   // 카테고리 (예: PTY, REH 등)
				obsrValue := itemData["obsrValue"].(string) // 관측 값 (예: 12.8 등)

				// 한글로 번역된 카테고리 추가
				translatedCategory := categoryTranslation[category]

				// WeatherItem 구조체에 저장
				weatherItems = append(weatherItems, WeatherItem{
					Category: translatedCategory,
					ObsValue: obsrValue,
				})
			}
		}
	}

	// 최종 응답 데이터를 ApiResponse 구조체로 반환
	return ApiResponse{
		ResponseCode: resp.StatusCode,
		Data:         weatherItems,
	}, nil
}

// HTTP 요청을 처리하는 함수
// "/api" 경로로 요청이 들어왔을 때 호출됩니다.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// API 데이터를 가져옴
	apiResponse, err := fetchAPIData()
	if err != nil {
		// 에러 발생 시 HTTP 500 응답
		http.Error(w, "Failed to fetch API data", http.StatusInternalServerError)
		return
	}

	// JSON 형식으로 클라이언트에 응답
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

// 메인 함수 (프로그램 시작 지점)
// 8081 포트에서 HTTP 서버를 실행합니다.
func main() {
	// "/api" 경로에 대한 요청을 처리하는 핸들러 함수 등록
	http.HandleFunc("/api", apiHandler)

	// 서버 실행
	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil)) // 8081 포트에서 서버 실행
}
