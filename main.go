package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thiagohdeplima/krakend-auth-server/server"
)

type registerer string

var (
	pluginName        = "authorization-server"
	HandlerRegisterer = registerer(pluginName)
)

func (r registerer) RegisterHandlers(f func(
	string,
	func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	if err := r.validateConfig(extra); err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		config := extra[string(r)].(map[string]interface{})

		server.Run(config, w, req, h)
	}), nil
}

func (r registerer) validateConfig(extra map[string]interface{}) error {
	var required = []string{"path"}
	var config, ok = extra[string(r)].(map[string]interface{})

	if !ok {
		return fmt.Errorf("plugin isn't loaded")
	}

	for _, key := range required {
		_, ok := config[key].(string)

		if !ok {
			return fmt.Errorf("missing config %s", key)
		}
	}

	return nil
}

func main() {}
