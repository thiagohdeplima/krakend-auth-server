package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thiagohdeplima/krakend-auth-server/internal/usecase"
)

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
	switch req.URL.Path {
	case "/oauth/token":
		s.serveToken(w, req)

	default:
		w.WriteHeader(404)
		s.handler.ServeHTTP(w, req)
	}
}

func (s *Server) serveToken(w http.ResponseWriter, req *http.Request) {
	resp, _ := s.usecase.Run(req.Context(), "abc123", "abc123")

	encoded, _ := json.Marshal(resp)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(encoded))
}
