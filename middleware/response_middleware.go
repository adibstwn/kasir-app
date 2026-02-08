package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kasir-app/constant"
	"kasir-app/dto"
	error2 "kasir-app/errors"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// --- ERROR PATH ---
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *error2.BaseError
			if errors.As(err, &appErr) {
				c.JSON(appErr.StatusCode, dto.BaseResponse{
					StatusCode:    appErr.StatusCode,
					StatusMessage: appErr.Code,
					ErrorMessage:  &appErr.Message,
				})
				return
			}

			//c.JSON(500, dto.BaseResponse{
			//	StatusCode:    500,
			//	StatusMessage: "INTERNAL_SERVER_ERROR",
			//})
			return
		}

		status := 200
		msg := "SUCCESS"

		if s, ok := c.Get(constant.ResponseStatusCode); ok {
			status = s.(int)
		}
		if m, ok := c.Get(constant.ResponseMessage); ok {
			msg = m.(string)
		}

		if data, ok := c.Get(constant.ResponseData); ok {
			c.JSON(status, dto.BaseResponse{
				StatusCode:    status,
				StatusMessage: msg,
				Data:          data,
			})
		}
	}
}
