package user

import (
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	// TODO: verify code
}

type LoginResp struct {
	Name  string `json:"name"`
	Id    int64  `json:"id"`
	Token string `json:"token"`
}

func (h *Handler) Login(ctx *gin.Context, req LoginReq) (resp LoginResp, err error) {
	token, id, err := h.userSrv.Login(req.Name, req.Password)
	if err != nil {
		// TODO: biz error
		return LoginResp{}, err
	}
	return LoginResp{Name: req.Name, Id: id, Token: token}, nil
}
