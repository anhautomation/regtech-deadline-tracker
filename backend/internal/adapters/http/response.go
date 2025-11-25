package http

import (
	"regtech-backend/internal/core/contract"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, contract.Response{
		Code:    contract.SUCCESS,
		Message: contract.ErrorMessage[contract.SUCCESS],
		Data:    data,
	})
}

func Fail(c *gin.Context, code string, messageOverride ...string) {
	msg := contract.ErrorMessage[code]
	if len(messageOverride) > 0 && messageOverride[0] != "" {
		msg = messageOverride[0]
	}

	c.JSON(200, contract.Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
