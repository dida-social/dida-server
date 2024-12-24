package user

import (
	"github.com/dida-social/dida-server/pkg/frame"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	name string
}

func HandlerFuncWrap[T any, P any](fn func(ctx *gin.Context, req T) (resp P, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBind(&req); err != nil {
			return
		}
		resp, err := fn(c, req)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func (h *Handler) RegisterRouter(eg *gin.Engine) {
	rg := eg.Group("/v1/" + h.name)
	rg.GET("name", h.Login)
}

func (h *Handler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": h.name,
	})
}

var _ frame.Handler = (*Handler)(nil)

func NewHandler(name string) *Handler {
	return &Handler{
		name: name,
	}
}
