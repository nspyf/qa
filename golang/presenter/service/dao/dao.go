package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mymod/model/db"
	"mymod/util"
)

type Config struct {
	Driver   	string 	`json:"driver"`
	User     	string 	`json:"user"`
	Password 	string 	`json:"password"`
	Host     	string 	`json:"host"`
	Port     	string 	`json:"port"`
	Dbname   	string 	`json:"db_name"`
	Param		string	`json:"param"`
}

type GormEngine struct {
	db *gorm.DB
}

func (e *GormEngine) Init(path string) {
	var c Config
	var err error
	err = util.ReadJSON(path,&c)
	util.DoErr(err)
	e.db, err = gorm.Open(c.Driver, c.User+":"+c.Password+"@tcp("+c.Host+":"+c.Port+")/"+c.Dbname+"?"+c.Param)
	util.DoErr(err)
	e.db.AutoMigrate(&db_mod.User{},&db_mod.Question{},&db_mod.Answer{})
	return
}

func (e *GormEngine) Creat(new interface{}) {
	e.db.Create(new)
	return
}

func (e *GormEngine) RetrieveByID(data interface{},ID uint) {
	e.db.First(data,ID)
	return
}

func (e *GormEngine) RetrieveByStruct(data interface{},tar interface{}) {
	e.db.Where(tar).First(data)
	return
}

func (e *GormEngine) RetrieveArrByStruct(data interface{},tar interface{}) {
	e.db.Where(tar).Find(data)
	return
}

func (e *GormEngine) Update(tar interface{},new interface{}) { // 注意：只更新非零
	e.db.Model(tar).Updates(new)
	return
}

func (e *GormEngine) Delete(tar interface{}) {
	e.db.Delete(tar)
	return
}