package main

import (
	"github.com/nagexiucai/gosh/config"
	"github.com/nagexiucai/gosh/biz"
)

// Representational State Transfer
// 加州大学
// Roy Fielding
// 比SOAP/WSDL更合理更轻量

// 在HTTP站点设计中很流行

func main() {
	conf := config.GetConfig()
	conf.DB.IP = "192.168.200.46"
	conf.DB.Port = 9527
	conf.DB.Username = "root"
	conf.DB.Password = "wwwwww"
	conf.DB.Name = "basic"

	app := &biz.Biz{}
	app.Initialize(conf)
	app.Run(":9527")
}
