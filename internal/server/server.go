package server

import (
	"fmt"
	"net/http"

	"github.com/thiagohdeplima/krakend-auth-server/internal/usecase"
)

const TOKEN_PATH = "/token"

type HTTPServer interface {
	Run(http.ResponseWriter, *http.Request)
}

type Server struct {
	usecase usecase.IssueToken
	handler http.Handler
}

func NewServer(uc usecase.IssueToken, handler http.Handler) HTTPServer {
	return &Server{usecase: uc, handler: handler}
}

func (s *Server) Run(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != TOKEN_PATH {
		s.handler.ServeHTTP(w, req)
		return
	}

	s.serveToken(w, req)
}

func (s *Server) serveToken(w http.ResponseWriter, req *http.Request) {
	resp, err := s.usecase.Run(req.Context(), "abc123", "abc123")

	fmt.Fprintf(w, "%+v --> %+v", resp, err)
}
