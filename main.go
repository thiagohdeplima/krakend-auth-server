package main

import (
	"context"
	"net/http"

	"github.com/thiagohdeplima/krakend-auth-server/internal/auth"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
	"github.com/thiagohdeplima/krakend-auth-server/internal/repo"
	"github.com/thiagohdeplima/krakend-auth-server/internal/server"
	"github.com/thiagohdeplima/krakend-auth-server/internal/usecase"
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

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	var rep = repo.NewFakeRepository()

	var val = auth.NewAuthenticator(rep)
	var iss = issuer.NewTokenEmissor(rep)

	var uc = usecase.NewTokenIssuer(val, iss)
	var srv = server.NewServer(uc, handler)

	return http.HandlerFunc(srv.Run), nil
}

func main() {}
