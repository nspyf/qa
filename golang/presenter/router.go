package ctrl

import (
	"github.com/gin-gonic/gin"
	"mymod/util"
)

func Run() {
	e := gin.New()
	e.Use(gin.Logger(),gin.Recovery(),cors())

	user := e.Group("/user",tokenVerify())

	e.GET("/information",information())

	e.POST("/register",register())
	e.POST("/login",login())
	e.POST("/question",question())

	user.GET("/verification",verify())
	user.POST("/answer",answer())

	user.DELETE("/question",deleteQuestion())
	user.DELETE("/answer",deleteAnswer())

	err := e.Run(":"+config.Port)
	util.DoErr(err)
	return
}