package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralError struct {
	Err string `json:"error"`
}

func (e *GeneralError) Error() string {
	return e.Err
}

// okResponse
func okResponse(c *gin.Context) {
	okResponseWithCustomMsg(c, "success")
}

func okResponseWithCustomMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// errInternal 内部错误（连不上第三方服务、数据库错误）
func errInternal(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err,
	})
}

// badRequest 参数错误
func badRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

// itemSliceResponse 通用返回列表方法
func itemSliceResponse[T any](c *gin.Context, items []T) {
	cnt := len(items)

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": cnt,
	})
}

// itemResponse 通用返回单个元素的方法
func itemResponse[T any](c *gin.Context, item T) {
	c.JSON(http.StatusOK, item)
}
