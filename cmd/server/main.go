package main

import (
	"context"
	userhandler "github.com/dida-social/dida-server/internal/handler/user"
	usersrv "github.com/dida-social/dida-server/internal/service/user"
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
	userSrv := usersrv.NewUserService()
	rt.RegisterHandler(userhandler.NewHandler("user", userSrv))

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
