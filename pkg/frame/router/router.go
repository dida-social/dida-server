package router

import (
	"context"
	"github.com/dida-social/dida-server/pkg/frame"
	"github.com/gin-gonic/gin"
)

type Router struct {
	eg *gin.Engine
}

func (r *Router) RegisterRouter(method string, path string, handler gin.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

//type Handler[Req, Resp any] func(req Req) (resp Resp)
//
//func (r *Router) RegisterRouter(method string, path string, handler Handler) {
//	r.eg.Handle(method, path, func(c *gin.Context) {
//
//	})
//}

func (r *Router) Name() string {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Close() error {
	//TODO implement me
	panic("implement me")
}

var _ frame.RouterModule = (*Router)(nil)

func NewRouter() *Router {
	return &Router{
		eg: gin.Default(),
	}
}
