package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyyyuna/storage-manager/pkg/log"
	"github.com/lyyyuna/storage-manager/pkg/model"
)

func (s *Server) download(ctx *gin.Context) {
	name := ctx.Param("name")

	if name == "" {
		log.Errorf("name is empty")
		badRequest(ctx, "bad request")
		return
	}

	file, err := model.GetFile(s.db, name)
	if err != nil {
		log.Errorf("get file error: %v", err)
		badRequest(ctx, "bad request")
		return
	}

	key := file.Url // 私有文件key

	request, err := s.s3Svc.GeneratePresignedURL(key, 5*time.Minute)

	log.Infof("generate presigned url: %v", request)

	ctx.Redirect(http.StatusMovedPermanently, request)
}
