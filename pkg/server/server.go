package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lyyyuna/storage-manager/pkg/config"
	"github.com/lyyyuna/storage-manager/pkg/model"
	"github.com/lyyyuna/storage-manager/pkg/s3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	r *gin.Engine

	cfg *config.Config

	db *gorm.DB

	s3Svc *s3.Service
}

func NewServer(cfg *config.Config) *Server {
	var s = &Server{
		cfg: cfg,
	}

	s.initDB()
	s.initS3()
	s.initRouter()

	return s
}

func (s *Server) Run() {
	log.Fatal(s.r.Run(s.cfg.Bind))
}

func (s *Server) initDB() {
	var err error
	s.db, err = gorm.Open(postgres.Open(s.cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := s.db.AutoMigrate(model.File{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func (s *Server) initS3() {
	var err error
	s.s3Svc, err = s3.NewService(s.cfg)
	if err != nil {
		log.Fatalf("failed to initialize S3 service: %v", err)
	}
}

func (s *Server) initRouter() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(cors.Default())
	{
		v1.POST("/record", s.record)
		v1.GET("/download/:name", s.download)
		v1.OPTIONS("/stats", func(ctx *gin.Context) {
			ctx.Data(204, "", nil)
		})
	}

	s.r = r
}
