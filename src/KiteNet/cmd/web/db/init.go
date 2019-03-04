package db

import (
	"KiteNet/log"
	"KiteNet/utils"
	"database/sql"
	"fmt"
	"net/url"

	//注册mysql驱动组件,手动导入勿删
	_ "github.com/go-sql-driver/mysql"
)

//Config -DB配置
type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	DBName   string `json:"dbName"`
}

//DB -mysql数据库单例
var DB *sql.DB

//Init 初始化DB
func Init(conf *Config) {
	s := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&loc=%s&parseTime=true",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DBName,
		url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", s)

	if !utils.CheckNilAndErr(db, err){
		glog.Info("db init success :"+s)
		DB = db
	}
}

//Close 关闭数据库连接
func Close() {
	e := DB.Close()
	utils.CheckErr(e)
}
