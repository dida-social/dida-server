package user

import (
	"github.com/dida-social/dida-server/internal/service/user"
	"github.com/dida-social/dida-server/pkg/frame"
	"github.com/dida-social/dida-server/pkg/frame/handler"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	name    string
	userSrv user.Service
}

func (h *Handler) RegisterRouter(eg *gin.Engine) {
	rg := eg.Group("/v1/" + h.name)
	rg.GET("register", handler.HandleFuncWrap(h.Register))
	rg.GET("login", handler.HandleFuncWrap(h.Login))
}

var _ frame.Handler = (*Handler)(nil)

func NewHandler(name string, userSrv user.Service) *Handler {
	return &Handler{
		name:    name,
		userSrv: userSrv,
	}
}
