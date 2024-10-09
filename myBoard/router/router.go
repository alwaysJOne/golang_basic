package router

import (
	"myBoad/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()

	r.Static("/resources", "./public/resources")
	r.StaticFile("/index.css", "./public/index.css")
	r.StaticFile("/index.js", "./public/index.js")

	r.LoadHTMLGlob("public/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/boardEnrollForm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "boardEnrollForm.html", gin.H{})
	})

	r.GET("/boardDetail/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "boardDetail.html", gin.H{})
	})

	crudRouter := r.Group("/api/board")

	// CRUD API 엔드포인트 추가
	crudRouter.POST("", apis.CreateBoard)           // 게시글 작성
	crudRouter.GET("", apis.GetBoardList)           // 게시글 조회
	crudRouter.GET("/:id", apis.GetBoardByID)       // 단일 게시글 조회
	crudRouter.PUT("/:id", apis.UpdateBoard)        // 게시글 수정
	crudRouter.DELETE("/:id", apis.DeleteBoardByID) // 게시글 삭제

	return r
}
