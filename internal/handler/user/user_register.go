package user

import "github.com/gin-gonic/gin"

type NameReq struct {
	Id string `json:"id"`
}

type NameResp struct {
	Name string `json:"name"`
}

func (h *Handler) Register(ctx *gin.Context, req NameReq) (resp NameResp, err error) {
	return NameResp{Name: "张三"}, nil
}
