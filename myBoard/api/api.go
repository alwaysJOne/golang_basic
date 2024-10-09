package api

import "myBoad/db"

type APIs struct {
	db *db.DBHandler
}

func NewAPI(b *db.DBHandler) *APIs {
	return &APIs{db: b}
}
