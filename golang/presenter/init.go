package ctrl

import (
	"mymod/presenter/service"
	"mymod/util"
)

type Config struct {
	Port	string	`json:"port"`
}

var config Config

func Init() {

	serv.Init()

	err := util.ReadJSON("./config/controller.json",&config)
	util.DoErr(err)

	return
}