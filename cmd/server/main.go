package main

import (
	"context"
	"github.com/dida-social/dida-server/pkg/frame/app"
	"time"
)

func main() {

	// new app instance
	server, err := app.New("dida-server")
	if err != nil {
		panic(err)
	}

	// register releases
	server.RegisterReleaseFunc()

	// start server
	ctx, sCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer sCancel()
	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
