package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	conn    *sql.DB
	Factory = make(map[string]*sql.DB)
)

type Config struct {
	// 连接用户名
	User string
	// 连接密码
	Password string
	// 连接地址
	Host string
	// 连接地址数组 集群连接时使用
	Hosts []string
	// 数据库名称
	DbName string `yaml:"dbName"`
	// 连接参数,配置编码、是否启用ssl等 charset=utf8
	ConnectInfo string
	// 用于设置最大打开的连接数，默认值为0表示不限制
	MaxOpenConn int
	// 用于设置闲置的连接数
	MaxIdleConn int
	// 连接名称 用于多个连接时区分
	ConnName string
}

func New(c *Config) {
	if Conn == nil {
		conn = c.new()
	}
}

func Conn() *sql.DB {
	return conn
}

// FactoryGet 连接工厂获取
func FactoryGet(name string) *sql.DB {
	if name == "" {
		return nil
	}
	return Factory[name]
}

func (c *Config) new() *sql.DB {
	if c.MaxOpenConn == 0 {
		c.MaxOpenConn = defaultMysqlMaxOpenConn
	}
	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = defaultMysqlMaxIdleConn
	}
	if c.ConnectInfo == "" {
		c.ConnectInfo = defConnectInfo
	}

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		c.User,
		c.Password,
		c.Host,
		c.DbName,
		c.ConnectInfo)

	dbCon, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatalf("Init Mysql Error, url:[%v], Error: [%v]", url, err)
		return nil
	}
	dbCon.SetMaxOpenConns(c.MaxOpenConn)
	dbCon.SetMaxIdleConns(c.MaxIdleConn)
	err = dbCon.Ping()
	if err != nil {
		log.Fatalf("Init Mysql Error, url:[%v], Error: [%v]", url, err)
		return nil
	}
	if c.ConnName != "" {
		Factory[c.ConnName] = conn
	}
	log.Fatalf("Init Mysql Success, url:[%v]", url)
	return conn
}
