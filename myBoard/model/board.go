package model

import "time"

type Board struct {
	ID        uint      `gorm:"primaryKey"`         // 게시글 번호 (기본키, 자동증가)
	Title     string    `gorm:"size:255;not null"`  // 제목
	Content   string    `gorm:"type:text;not null"` // 내용
	Author    string    `gorm:"size:100;not null"`  // 글쓴이
	CreatedAt time.Time `gorm:"autoCreateTime"`     // 작성일
}
