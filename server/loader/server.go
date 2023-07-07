package loader

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitaepark/codedeploy-go/server/config"
)

type Server struct {
	config config.Config
	router *gin.Engine
}

func NewServer(config config.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/api/health-check", server.healthCheck)

	server.router = router
}

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": "OK" })
}