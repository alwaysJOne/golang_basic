package db

import (
	"fmt"
	"log"
	m "myBoad/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	gDB *gorm.DB
}

func ConnectGorm(dsn string) (*DBHandler, error) {
	// GORM을 사용하여 MySQL 데이터베이스 연결
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}
	fmt.Println("데이터베이스에 성공적으로 연결되었습니다.")

	// 데이터베이스 마이그레이션 (테이블 생성)
	if err := gormDB.AutoMigrate(&m.Board{}); err != nil {
		log.Fatalf("테이블 생성 실패: %v", err)
	}

	dbHandler := &DBHandler{
		gDB: gormDB,
	}

	return dbHandler, err
}
