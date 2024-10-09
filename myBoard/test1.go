package main

import (
	"fmt"
	"log"
	"net/http"

	m "myBoad/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 데이터베이스 연결 문자열
	dsn := "root:0000@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	// GORM을 사용하여 MySQL 데이터베이스 연결
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}
	fmt.Println("데이터베이스에 성공적으로 연결되었습니다.")

	// 데이터베이스 마이그레이션 (테이블 생성)
	if err := db.AutoMigrate(&m.Board{}); err != nil {
		log.Fatalf("테이블 생성 실패: %v", err)
	}

	fmt.Println("데이터베이스에 성공적으로 연결되었습니다.")

	r := gin.Default()

	r.Static("/resources", "./public/resources")
	r.StaticFile("/index.css", "./public/index.css")
	r.StaticFile("/index.js", "./public/index.js")

	r.LoadHTMLGlob("public/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":8081")

}
