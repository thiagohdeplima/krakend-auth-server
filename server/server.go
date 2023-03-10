package server

import (
	"fmt"
	"net/http"
)

func Run(config map[string]interface{}, w http.ResponseWriter, req *http.Request, h http.Handler) {
	path, _ := config["path"].(string)

	if req.URL.Path != path {
		h.ServeHTTP(w, req)
		return
	}

	if req.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	fmt.Fprintf(w, "%q requested", req.URL.Path)
}
