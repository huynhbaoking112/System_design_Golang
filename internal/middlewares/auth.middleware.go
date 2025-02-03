package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/package/response"
)

// Trong Gin, c.Abort() được sử dụng để dừng ngay lập tức việc xử lý các middleware tiếp theo và handler chính. Khi gọi c.Abort(), Gin sẽ không thực thi bất kỳ middleware hay handler nào sau đó.

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
