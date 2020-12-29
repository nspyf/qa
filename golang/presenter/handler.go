package ctrl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mymod/model/request"
	"mymod/presenter/service"
	"mymod/util"
	"net/http"
)

func getClaims(c *gin.Context) (*map[string]interface{},error) {
	claims,ok := c.Get("claims")
	if ok == false {
		return nil,errors.New("get claims from context failed")
	}

	claimsObj, ok := claims.(util.Claims)
	if ok == false {
		return nil,errors.New("claims transform failed")
	}

	publicObj, ok := claimsObj.Public.(map[string]interface{})
	if ok == false {
		return nil,errors.New("public of claims transform failed")
	}
	return &publicObj,nil
}

func register() gin.HandlerFunc {
	return func (c *gin.Context) {
		req := &req_mod.Register{}
		err := util.DecodeReader(c.Request.Body,&req)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": -1,
				"message": "JSON error",
			})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusForbidden,gin.H{
				"status": -1,
				"message": "username or password can't be blank",
			})
			return
		}

		err = serv.Register(req)
		if err != nil {
			c.JSON(http.StatusForbidden,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status": 1,
			"message": "register successfully",
		})
		return
	}
}

func login() gin.HandlerFunc {
	return func (c *gin.Context) {
		req := &req_mod.Login{}
		err := util.DecodeReader(c.Request.Body,&req)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": -1,
				"message": "JSON error",
			})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusForbidden,gin.H{
				"status": -1,
				"message": "username or password error",
			})
			return
		}

		token,err := serv.Login(req)
		if err != nil {
			c.JSON(http.StatusForbidden,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status": 1,
			"message": "login successfully",
			"data": gin.H{
				"token": token,
			},
		})
		return
	}
}

func information() gin.HandlerFunc {
	return func (c *gin.Context) {
		username := c.Query("user")

		data,err := serv.Information(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status": 1,
			"message": "get information successfully",
			"data": data,
		})
		return
	}
}

func question() gin.HandlerFunc {
	return func (c *gin.Context) {
		req := &req_mod.Question{}
		err := util.DecodeReader(c.Request.Body,&req)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": -1,
				"message": "JSON error",
			})
			return
		}

		err = serv.Question(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status": 1,
			"message": "post question successfully",
		})
		return
	}
}

func answer() gin.HandlerFunc {
	return func (c *gin.Context) {
		req := &req_mod.Answer{}
		err := util.DecodeReader(c.Request.Body,&req)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": -1,
				"message": "JSON error",
			})
			return
		}



		claims,ok := c.Get("claims")
		if ok == false {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": "get claims from context failed",
			})
			return
		}

		claimsObj, ok := claims.(util.Claims)
		if ok == false {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": "claims transform failed",
			})
			return
		}

		publicObj, ok := claimsObj.Public.(map[string]interface{})
		if ok == false {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": "public of claims transform failed",
			})
			return
		}



		err = serv.Answer(publicObj["ID"],req)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status": -1,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status": 1,
			"message": "post answer successfully",
		})
		return
	}
}