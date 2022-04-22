package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/pkg/database/mysql"
	"vending/app/infrastructure/pkg/database/redis"
	zapLog "vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util/jwt"
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
	Jwt *jwt.Config
}

var conf = flag.String("conf", "/Users/yapi/WorkSpace/GolandWorkSpace/Template/config.yml", "conf")

func NewConfig() *Config {
	var C Config
	flag.Parse()
	v := viper.New()
	v.SetConfigFile(*conf)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("%v unmarshal ReadInConfig error", err))
	}
	if err := v.Unmarshal(&C); err != nil {
		panic(fmt.Sprintf("%v unmarshal Unmarshal error", err))
	}
	Base = &C

	NewLog(Base)
	NewJwt(Base)
	NewMongo(Base)
	return Base
}

func NewLog(c *Config) {
	zapLog.New(c.Log)
}

func NewJwt(c *Config) {
	jwt.New(c.Jwt)
}

func NewMongo(c *Config) {
	mongo.New(c.Mongo)
}

func NewMysql(c *Config) {
	mysql.New(c.Mysql)
}

func NewRedis(c *Config) {
	redis.New(c.Redis)
}
