package datasource

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME   = "root"
	PASS_WORD   = "mysql"
	HOST        = "172.21.42.9"
	REMOTE_HOST = "172.17.0.2"
	PORT        = "10000"
	DATABASE    = "blog"
	CHARSET     = "utf8"
)

// 初始化链接
func init() {
	log.Println("mysql init")
	// 解析启动模式
	mode := flag.String("mode", "local", "get service start mode(exp: local, remote")
	flag.Parse()
	log.Println("mode", *mode)

	host := HOST
	if *mode == "remote" {
		host = REMOTE_HOST
	}

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, host, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

	log.Println("mysql inited")
}
