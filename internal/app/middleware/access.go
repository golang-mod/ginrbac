package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/forms"
	"github.com/zhiniuer/goadmin/internal/app/service"
)

const ()

// NewAccess 记录日志
func NewAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		input := c.GetString(app.AdminInput)
		output := c.GetString(app.AdminOutput)
		ip := c.GetString(app.AdminIp)
		userId := c.GetInt64(app.AdminUserId)

		if userId != 0 {
			// 记录访问日志
			logService := &service.AdminOperationLogService{}
			_ = logService.Save(&forms.AdminOperationLogStoreForm{
				Method: c.Request.Method,
				Path:   c.Request.RequestURI,
				Ip:     ip,
				UserId: int64(userId),
				Input:  input,
				Status: c.Writer.Status(),
				Output: output,
			})
		}
	}
}
