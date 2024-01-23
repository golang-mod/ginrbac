package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goutils/response"
	"log"
	"net/http"
	"strconv"
)

func NewRbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		userID = strconv.Itoa(int(c.GetInt64(app.AdminUserId)))
		path := c.Request.URL.Path
		method := c.Request.Method

		rbac := driver.Rbac
		// 验证策略规则
		result := rbac.CheckPolicy(userID, path, method)
		log.Println(rbac.UserPrefix+userID, path, method, result)
		if result == false {
			c.JSON(http.StatusForbidden, response.Fail("抱歉，您没有权限访问", 10002))
			c.Abort()
			return
		}
		c.Next()
	}
}
