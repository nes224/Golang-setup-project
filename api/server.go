package api

import (
	db "github.com/abs/bestinter/db/sqlc"
	"github.com/abs/bestinter/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store: store,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	server.router =router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
