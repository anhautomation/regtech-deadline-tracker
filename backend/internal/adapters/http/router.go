package http

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "regtech-backend/internal/docs" // swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *DeadlineHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://localhost:5173",
		},
		AllowOriginFunc: func(origin string) bool {
			return strings.HasSuffix(origin, ".vercel.app")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/deadlines", handler.ListDeadlines)
		v1.POST("/deadlines", handler.CreateDeadline)
		v1.PUT("/deadlines/:id/complete", handler.MarkCompleted)
	}

	return r
}
