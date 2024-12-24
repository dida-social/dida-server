package main

import (
	"context"
	"github.com/dida-social/dida-server/internal/handler/user"
	"github.com/dida-social/dida-server/pkg/frame/app"
	"github.com/dida-social/dida-server/pkg/frame/router"
	"time"
)

func main() {

	// new app instance
	server, err := app.New("dida")
	if err != nil {
		panic(err)
	}

	// router module
	rt := router.NewRouter("router", 8081)
	// register user handler
	rt.RegisterHandler(user.NewHandler("user"))

	// register modules
	if err = server.RegisterModules(context.Background(), rt); err != nil {
		panic(err)
	}

	// start server
	ctx, sCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer sCancel()
	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
