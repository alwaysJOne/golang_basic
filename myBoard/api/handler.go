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

func (apis *APIs) CreateBoard(c *gin.Context) {
	req := &m.Board{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	res, err := apis.db.CreateBoard(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoardList(c *gin.Context) {
	res, err := apis.db.GetBoardList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{Res: "Server error"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoardByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // 문자열을 uint로 변환
	if err != nil {
		// 변환 실패 시 에러 처리
		c.JSON(http.StatusBadRequest, &Response{Res: "Invalid ID"})
		return
	}

	res, err := apis.db.GetBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) UpdateBoard(c *gin.Context) {
	idS := c.Param("id")
	req := &m.Board{}

	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	res, err := apis.db.UpdateBoard(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) DeleteBoardByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // 문자열을 uint로 변환
	if err != nil {
		// 변환 실패 시 에러 처리
		c.JSON(http.StatusBadRequest, &Response{Res: "Invalid ID"})
		return
	}

	err = apis.db.DeleteBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, &Response{Res: "Success"})
}
