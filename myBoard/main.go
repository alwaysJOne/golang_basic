package main

import (
	"fmt"
	"myBoad/api"
	"myBoad/db"
	"myBoad/router"
)

func main() {
	//RESTAPI서버
	//REST 서버는 클라이언트가 요청을 보내면, 그 요청을 처리하고 데이터를 응답으로 보내주는 서버

	//REST API의 약속
	//URL을 통한 자원 접근하고 행위는 HTTP메서드로 구분한다.
	// /api/board
	/*
		GET: 서버에서 데이터를 요청할 때 사용합니다. (예: 게시글 정보 가져오기)
		POST: 서버에 새 데이터를 추가할 때 사용합니다. (예: 게시글 작성)
		PUT: 서버에 있는 데이터를 수정할 때 사용합니다. (예: 게시글 정보 수정)
		DELETE: 서버에서 데이터를 삭제할 때 사용합니다. (예: 게시글 삭제)
	*/
	dbHandler, err := db.ConnectGorm("root:0000@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8081")

}
