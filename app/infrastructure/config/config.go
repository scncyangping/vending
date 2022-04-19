package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/pkg/database/mysql"
	"vending/app/infrastructure/pkg/database/redis"
	zapLog "vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/tool"
)

var (
	Base *Config
)

// Server 系统配置
type Server struct {
	Addr         string `yaml:"addr"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}

type Config struct {
	// 服务配置
	Server *Server
	// 日志配置
	Log *zapLog.Config
	// redis配置
	Redis *redis.Config
	// mongodb配置
	Mongo *mongo.Config
	// mysql配置
	Mysql *mysql.Config
	// jwt配置
	Jwt *tool.JwtConfig
}

func NewConfig(filePath string) *Config {
	conf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	err = yaml.Unmarshal(conf, &Base)
	if err != nil {
		panic(err)
	}
	Base.initLog()
	Base.initJwt()
	return Base
}

func (c *Config) initLog() *Config {
	zapLog.NewLogger(c.Log)
	return c
}

func (c *Config) initJwt() *Config {
	tool.NewJwt(c.Jwt)
	return c
}

func (c *Config) InitMongo() *Config {
	mongo.Init(c.Mongo)
	return c
}

func (c *Config) InitMysql() *Config {
	mysql.New(c.Mysql)
	return c
}

func (c *Config) InitRedis() *Config {
	redis.NewRedis(c.Redis)
	return c
}
