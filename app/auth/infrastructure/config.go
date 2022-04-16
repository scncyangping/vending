package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"vending/config/jwt"
	"vending/config/log"
	"vending/config/mongo"
	"vending/config/mysql"
	"vending/config/redis"
)

var (
	Base *Config
)

type Config struct {
	// 服务配置
	Server *Server
	// 日志配置
	Log *log.Config
	// redis配置
	Redis *redis.Config
	// mongodb配置
	Mongo *mongo.Config
	// mysql配置
	Mysql *mysql.Config
	// jwt配置
	Jwt *jwt.Config
}

func NewConfig(filePath string) {
	conf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.(any))
		return
	}
	initYml(conf)
}

// 加载yml配置文件
func initYml(byteArray []byte) {
	err := yaml.Unmarshal(byteArray, &Base)
	if err != nil {
		panic(err.(any))
	}
	log.New(Base.Log)
	jwt.New(Base.Jwt)
	//Base.Mongo.New()
	//Base.Jwt.New()
	//Base.Redis.New()
}

// 系统配置
type Server struct {
	Name string
	Port string
}
