package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var conn *mongo.Client

type Config struct {
	Host            string
	User            string
	DbName          string `yaml:"dbName"`
	Password        string
	PoolSize        int
	MaxConnIdleTime int `yaml:"maxConnIdleTime"`
}

func New(c *Config) {
	if conn == nil {
		c.new()
	}
}

func Conn() *mongo.Client {
	return conn
}

func (c *Config) new() *mongo.Client {
	opt := options.Client().ApplyURI(c.Host)
	if len(c.User) != 0 { // 部分连接不需要帐号密码
		opt.Auth = &options.Credential{
			Username: c.Host,
			Password: c.Password,
		}
	}
	//只使用与mongo操作耗时小于3秒的
	opt.SetLocalThreshold(3 * time.Second)
	//指定连接可以保持空闲的最大毫秒数
	opt.SetMaxConnIdleTime(time.Duration(c.MaxConnIdleTime) * time.Second)
	//使用最大的连接数
	opt.SetMaxPoolSize(uint64(c.PoolSize))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	return client
}

func Do() {
	conn.Database("").Collection("123")
}
