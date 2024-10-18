package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 웹 페이지 요청
	res, err := http.Get("https://alwaysjone.github.io/scraping_sample/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// HTTP 상태 코드 확인
	if res.StatusCode != 200 {
		log.Fatalf("Failed to load page, status code: %d", res.StatusCode)
	}

	// goquery로 HTML 문서 파싱
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 상품 정보 저장을 위한 슬라이스
	var products []struct {
		name  string
		price int
		count int
	}

	// 상품명, 가격, 수량 추출
	doc.Find("#product-body tr").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".product-name").Text()
		priceStr := s.Find(".product-price").Text()
		countStr := s.Find(".product-count").Text()

		// 가격과 수량을 정수로 변환
		priceStr = strings.ReplaceAll(priceStr, "원", "") // "원" 제거
		priceStr = strings.TrimSpace(priceStr)           // 공백 제거
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			log.Println("가격 변환 오류:", err)
			return
		}

		count, err := strconv.Atoi(countStr)
		if err != nil {
			log.Println("수량 변환 오류:", err)
			return
		}

		// 상품 정보 저장
		products = append(products, struct {
			name  string
			price int
			count int
		}{name, price, count})

		fmt.Printf("상품명: %s, 가격: %d원, 수량: %d\n", name, price, count)
	})

	// 총 상품 갯수, 총 가격, 평균 가격 계산
	totalCount := len(products)
	totalPrice := 0

	for _, product := range products {
		totalPrice += product.price
	}

	averagePrice := 0
	if totalCount > 0 {
		averagePrice = totalPrice / totalCount
	}

	fmt.Printf("총 상품 갯수: %d\n", totalCount)
	fmt.Printf("총 가격: %d원\n", totalPrice)
	fmt.Printf("평균 가격: %d원\n", averagePrice)
}
