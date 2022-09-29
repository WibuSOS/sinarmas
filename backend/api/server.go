package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
