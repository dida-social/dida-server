package handler

import (
	"errors"
	"github.com/dida-social/dida-server/pkg/frame/xerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleFuncWrap[T any, P any](fn func(ctx *gin.Context, req T) (resp P, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBind(&req); err != nil {
			// TODO: internal error
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
		resp, err := fn(c, req)
		if err != nil {
			var xerr *xerror.XError
			if errors.As(err, &xerr) {
				c.JSON(http.StatusOK, gin.H{
					"code": xerr.Code,
					"msg":  xerr.Msg,
				})
			} else {
				// should internal error appear to customer ??
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": resp,
		})
	}
}
