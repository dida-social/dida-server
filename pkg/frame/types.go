package frame

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
)

type Starter interface {
	Run(ctx context.Context) error
}

type ModuleRegister interface {
	RegisterModules(ctx context.Context, modules ...Module) error
}

type ReleaseFunc func() error
type Releaser interface {
	RegisterReleaseFunc(releaseFunctions ...ReleaseFunc) error
}

// AppModule app module
type AppModule interface {
	Name() string
	Starter
	ModuleRegister
	Releaser
}

type RouterModule interface {
	Module
	RegisterHandler(h Handler)
}

type Module interface {
	Name() string
	Starter
	io.Closer
}

type Handler interface {
	RegisterRouter(eg *gin.Engine)
}
