package api

import (
	m "myBoad/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Res string `json:"res"`
}

// CreateBoard는 새로운 게시글을 생성하는 API입니다.
func (apis *APIs) CreateBoard(c *gin.Context) {
	req := &m.Board{} // 요청으로 전달된 게시글 정보를 받을 구조체

	// 클라이언트가 보낸 데이터를 Board 구조체로 바인딩
	if err := c.ShouldBind(req); err != nil {
		// 요청 데이터가 잘못된 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 데이터베이스에 게시글 생성 요청
	res, err := apis.db.CreateBoard(req)
	if err != nil {
		// 게시글 생성 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 생성된 게시글을 응답으로 반환
	c.JSON(http.StatusOK, res)
}

// GetBoardList는 모든 게시글을 조회하는 API입니다.
func (apis *APIs) GetBoardList(c *gin.Context) {
	// 데이터베이스에서 모든 게시글 리스트를 조회
	res, err := apis.db.GetBoardList()
	if err != nil {
		// 서버에서 오류가 발생한 경우 (Internal Server Error)
		c.JSON(http.StatusInternalServerError, &Response{Res: "Server error"})
		return
	}

	// 조회된 게시글 리스트를 응답으로 반환
	c.JSON(http.StatusOK, res)
}

// GetBoardByID는 게시글 ID를 통해 특정 게시글을 조회하는 API입니다.
func (apis *APIs) GetBoardByID(c *gin.Context) {
	// 클라이언트가 보낸 게시글 ID를 문자열로 받음
	idStr := c.Param("id")

	// 문자열을 uint 타입으로 변환 (게시글 ID는 숫자여야 함)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		// 변환 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Invalid ID"})
		return
	}

	// 데이터베이스에서 해당 ID의 게시글 조회
	res, err := apis.db.GetBoardByID(uint(id))
	if err != nil {
		// 게시글 조회 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 조회된 게시글을 응답으로 반환
	c.JSON(http.StatusOK, res)
}

// UpdateBoard는 특정 게시글을 수정하는 API입니다.
func (apis *APIs) UpdateBoard(c *gin.Context) {
	// 클라이언트가 보낸 게시글 ID를 문자열로 받음
	idS := c.Param("id")
	req := &m.Board{} // 수정할 게시글 정보를 받을 구조체

	// 문자열을 uint로 변환 (게시글 ID는 숫자여야 함)
	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		// 변환 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 클라이언트가 보낸 데이터를 Board 구조체로 바인딩
	if err := c.ShouldBind(req); err != nil {
		// 요청 데이터가 잘못된 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 데이터베이스에서 게시글 업데이트
	res, err := apis.db.UpdateBoard(uint(id), req)
	if err != nil {
		// 게시글 수정 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 수정된 게시글을 응답으로 반환
	c.JSON(http.StatusOK, res)
}

// DeleteBoardByID는 게시글 ID를 통해 특정 게시글을 삭제하는 API입니다.
func (apis *APIs) DeleteBoardByID(c *gin.Context) {
	// 클라이언트가 보낸 게시글 ID를 문자열로 받음
	idStr := c.Param("id")

	// 문자열을 uint로 변환 (게시글 ID는 숫자여야 함)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		// 변환 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Invalid ID"})
		return
	}

	// 데이터베이스에서 해당 ID의 게시글 삭제
	err = apis.db.DeleteBoardByID(uint(id))
	if err != nil {
		// 게시글 삭제 실패한 경우 (Bad Request)
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	// 삭제 성공 시 응답
	c.JSON(http.StatusOK, &Response{Res: "Success"})
}
