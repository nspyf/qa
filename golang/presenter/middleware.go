package ctrl

import (
	"github.com/gin-gonic/gin"
	"mymod/presenter/service"
	"net/http"
)

func cors() gin.HandlerFunc { // 预检
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin) // 允许cookie不能使用*，必须有明确的origin
		c.Header("Access-Control-Allow-Credentials", "true") // 允许cookie
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Host, Token")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE, PUT, PATCH")
		c.Header("Access-Control-Max-Age", "3600") // 预检请求的有效期/秒
		c.Header("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // 204
			return
		}
		c.Next()
	}
}

func tokenVerify() gin.HandlerFunc { // token验证
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		claims, err := serv.JwtObj.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if claims == nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"status": -1,
				"message": "claims nil",
			})
			c.Abort()
			return
		}
		c.Set("claims", *claims)
		c.Next()
	}
}