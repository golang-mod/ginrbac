package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/internal/app/service"
)

const (
	requestInput  = "request.input"
	requestOutput = "request.output"
	requestIp     = "request.ip"
)

// NewAccess 记录日志
func NewAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := c.GetString(requestInput)
		// 处理请求
		c.Next()
		output := c.GetString(requestOutput)
		ip := c.GetString(requestIp)
		kc := app.NewContext(c)
		if kc.User.AdminId != 0 {
			// 记录访问日志
			logService := &service.AdminOperationLogService{}
			_ = logService.Save(kc, &schema.AdminOperationLogStoreForm{
				Method: c.Request.Method,
				Path:   c.Request.RequestURI,
				Ip:     ip,
				UserId: kc.User.AdminId,
				Input:  input,
				Status: c.Writer.Status(),
				Output: output,
			})
		}
	}
}
