package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

var (
	confMap = make(map[string]string)
	Base    = Config{}
)

// 系统配置
type Server struct {
	Name string
	Port string
}

type Config struct {
	// 服务配置
	Server Server
	// 日志配置
	Log LogConfig
	// redis配置
	Redis RedisConfig
	// mongodb配置
	Mongo MongoConfig
	// mysql配置
	Mysql MysqlConfig
	// jwt配置
	Jwt JwtConfig
}

func GetConf(key string) string {
	return confMap[key]
}

func NewConfig(filePath string) {
	conf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
		return
	}
	initYml(conf)
}

// 加载yml配置文件
func initYml(byteArray []byte) {
	err := yaml.Unmarshal(byteArray, &Base)
	if err != nil {
		panic(err)
	}
}

// 加载键值对配置文件
func initConf(byteArray []byte) {
	confStr := string(byteArray)
	confStrSlice := strings.Split(confStr, "\n")
	for i := 0; i < len(confStrSlice); i++ {
		confStrSlice[i] = strings.Trim(confStrSlice[i], "\r")
		if strings.HasPrefix(confStrSlice[i], "//") || confStrSlice[i] == "" {
			continue
		}
		oneConf := strings.Split(confStrSlice[i], "=")
		if len(oneConf) == 2 {
			confMap[strings.Trim(oneConf[0], "\r")] = strings.Trim(oneConf[0], "\r")
		}
	}
}

// logger
type LogConfig struct {
	Dir     string
	Console bool
	Level   string
}

type MongoConfig struct {
	Host     string
	User     string
	DbName   string `yaml:"dbName"`
	Password string
	PoolSize int
}

type MysqlConfig struct {
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

type RedisConfig struct {
	PoolSize string
	Password string
	Host     string
	Hosts    []string
}

type JwtConfig struct {
	JwtSecret     string `yaml:"jwtSecret"`
	JwtExpireTime int    `yaml:"jwtExpireTime"`
	Issuer        string `yaml:"issuer"`
	Secret        []byte
}
