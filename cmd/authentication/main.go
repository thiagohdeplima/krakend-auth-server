package main

import (
	"context"
	"net/http"

	"github.com/thiagohdeplima/krakend-auth-server/cmd/authentication/server"

	"github.com/thiagohdeplima/krakend-auth-server/internal/repo/fake"
	"github.com/thiagohdeplima/krakend-auth-server/internal/usecase"

	"github.com/thiagohdeplima/krakend-auth-server/internal/auth"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
)

type registerer string

var (
	pluginName        = "authentication"
	HandlerRegisterer = registerer(pluginName)
)

func (r registerer) RegisterHandlers(f func(
	string,
	func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	var rep = fake.NewFakeRepository()

	var val = auth.NewAuthenticator(rep)
	var iss = issuer.NewTokenEmissor(rep)

	var uc = usecase.NewTokenIssuer(val, iss)
	var srv = server.NewServer(uc, handler)

	return http.HandlerFunc(srv.Run), nil
}

func main() {}
