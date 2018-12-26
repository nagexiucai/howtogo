package main

import (
	"database/sql"
	"net/http"
	"html/template"
)

/*
// 初始化SQL
DROP TABLE IF EXISTS `employees-scratch`;
CREATE TABLE `employees-scratch` (
  `id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `city` varchar(24),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

// 查看表创建信息
SHOW CREATE TABLE `employees-scratch`;

// 查看表信息
show table status;
*/

// 定义Employee结构体，包含Id、Name、City属性
type Employee struct {
	Id int `json: id`
	Name string `json: name`
	City string `json: city`
}

func dbConnect() (db *sql.DB) {
	dbHost := "192.168.200.46"
	dbPort := 9527
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "wwwwww"
	dbName := "basic"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbHost+":"+string(dbPort)+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var form = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
}
