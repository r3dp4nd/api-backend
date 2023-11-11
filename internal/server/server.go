package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}

func NewServer(port string) *Server {
	return &Server{
		httpAddr: fmt.Sprintf(":%s", port),
		engine:   gin.New(),
	}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server shut down", err)
	}
}
