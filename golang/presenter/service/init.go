package serv

import (
	"mymod/presenter/service/dao"
	"mymod/util"
)

var e dao.GormEngine
var JwtObj util.JWTObj

func Init() {
	e.Init("./config/db.json")
	JwtObj.Init("./config/jwt.json")
	return
}