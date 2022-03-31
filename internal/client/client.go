package client

import (
	"context"

	"github.com/ebelanja/go-micro/auth"
	"github.com/ebelanja/go-micro/client"
	"github.com/ebelanja/go-micro/client/grpc"
	"github.com/ebelanja/go-micro/metadata"
	"github.com/ebelanja/micro/client/cli/util"
	cliutil "github.com/ebelanja/micro/client/cli/util"
	"github.com/ebelanja/micro/internal/config"
	ccli "github.com/micro/cli/v2"
)

// New returns a wrapped grpc client which will inject the
// token found in config into each request
func New(ctx *ccli.Context) client.Client {
	env := cliutil.GetEnv(ctx)
	token, _ := config.Get("micro", "auth", env.Name, "token")
	return &wrapper{grpc.NewClient(), token, env.Name, ctx}
}

type wrapper struct {
	client.Client
	token string
	env   string
	ctx   *ccli.Context
}

func (a *wrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	if len(a.token) > 0 {
		ctx = metadata.Set(ctx, "Authorization", auth.BearerScheme+a.token)
	}
	if len(a.env) > 0 && !util.IsLocal(a.ctx) && !util.IsServer(a.ctx) {
		// @todo this is temporarily removed because multi tenancy is not there yet
		// and the moment core and non core services run in different environments, we
		// get issues. To test after `micro env add mine 127.0.0.1:8081` do,
		// `micro run github.com/ebelanja/services/logspammer` works but
		// `micro -env=mine run github.com/ebelanja/services/logspammer` is broken.
		// Related ticket https://github.com/micro/development/issues/193
		//
		// env := strings.ReplaceAll(a.env, "/", "-")
		// ctx = metadata.Set(ctx, "Micro-Namespace", env)
	}
	return a.Client.Call(ctx, req, rsp, opts...)
}
