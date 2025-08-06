package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyyyuna/storage-manager/pkg/log"
	"github.com/lyyyuna/storage-manager/pkg/model"
)

type recordInput struct {
	Url string `json:"url"`
}

func (s *Server) record(ctx *gin.Context) {
	var input recordInput
	if err := ctx.BindJSON(&input); err != nil {
		log.Errorf("bind json error: %v", err)
		badRequest(ctx, "bad request")
		return
	}

	name, err := model.CreateFile(s.db, input.Url)
	if err != nil {
		log.Errorf("create file error: %v", err)
		badRequest(ctx, "bad request")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
