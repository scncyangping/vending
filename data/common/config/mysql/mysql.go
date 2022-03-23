package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"vending/data/common/config"
	"vending/data/common/config/log"
	"vending/data/common/constants"
)

const (
	// Mysql 最大空闲连接数
	defaultMysqlMaxIdleConn = 1000
	// Mysql 最大连接数
	defaultMysqlMaxOpenConn = 2000
	// 默认还是连接
	defConnectInfo = "charset=utf8"
)

var (
	Conn    *sql.DB
	Factory = make(map[string]*sql.DB)
)

func NewMySql() *sql.DB {
	mysql := &config.Base.Mysql
	return newMysql(mysql)
}

// 连接工厂获取
func FactoryGet(name string) *sql.DB {
	if name == constants.EmptyStr {
		return nil
	}
	return Factory[name]
}

func newMysql(mysql *config.MysqlConfig) *sql.DB {
	if mysql.MaxOpenConn == constants.ZERO {
		mysql.MaxOpenConn = defaultMysqlMaxOpenConn
	}
	if mysql.MaxIdleConn == constants.ZERO {
		mysql.MaxIdleConn = defaultMysqlMaxIdleConn
	}
	if mysql.ConnectInfo == constants.EmptyStr {
		mysql.ConnectInfo = defConnectInfo
	}
	return connect(mysql)
}

func connect(mysql *config.MysqlConfig) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		mysql.User,
		mysql.Password,
		mysql.Host,
		mysql.DbName,
		mysql.ConnectInfo)

	dbCon, err := sql.Open("mysql", url)
	if err != nil {
		log.ZapLogger.Errorf("Init Mysql Error, url:[%v], Error: [%v]", url, err)
		return nil
	}
	dbCon.SetMaxOpenConns(mysql.MaxOpenConn)
	dbCon.SetMaxIdleConns(mysql.MaxIdleConn)
	err = dbCon.Ping()
	if err != nil {
		log.ZapLogger.Errorf("Init Mysql Error, url:[%v], Error: [%v]", url, err)
		return nil
	}
	Conn = dbCon
	if mysql.ConnName != constants.EmptyStr {
		Factory[mysql.ConnName] = Conn
	}
	log.ZapLogger.Infof("Init Mysql Success, url:[%v]", url)
	return Conn
}
