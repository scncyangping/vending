package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

var client *redis.Client
var clusterClient *redis.ClusterClient

type Config struct {
	PoolSize string
	Password string
	Host     string
	Hosts    []string
}

func New(c *Config) {
	var err error
	if c.Hosts != nil && len(c.Hosts) > 0 {
		err = c.newRedisClusterClient()
	} else {
		err = c.newRedisClient()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func Client() *redis.Client {
	return client
}

func ClusterClient() *redis.ClusterClient {
	return clusterClient
}

// client for single
func (c *Config) newRedisClient() error {
	var (
		host         string
		password     string
		poolSize     string
		redisOptions *redis.Options
	)

	host = c.Host
	password = c.Password
	poolSize = c.PoolSize
	redisOptions = &redis.Options{}
	redisOptions.Addr = host
	redisOptions.DB = 0

	if password != "" {
		redisOptions.Password = password
	}
	if poolSize != "" {
		redisOptions.PoolSize, _ = strconv.Atoi(poolSize)
	}

	client = redis.NewClient(redisOptions)

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Redis Init Error: Host: %v, Error:%v ", host, err)
		return err
	}

	log.Fatalf("Redis Init Success, Host: %v", host)

	return nil
}

// client for cluster
func (c *Config) newRedisClusterClient() error {
	var (
		hosts          []string
		password       string
		poolSize       string
		clusterOptions *redis.ClusterOptions
	)

	hosts = c.Hosts
	password = c.Password
	poolSize = c.PoolSize

	clusterOptions = &redis.ClusterOptions{}

	clusterOptions.Addrs = hosts

	if password != "" {
		clusterOptions.Password = password
	}
	if poolSize != "" {
		clusterOptions.PoolSize, _ = strconv.Atoi(poolSize)
	}

	clusterClient = redis.NewClusterClient(clusterOptions)

	_, err := clusterClient.Ping().Result()
	if err != nil {
		log.Fatalf("Redis Init Error: Hosts:%v, Error:%v ", hosts, err)
		return err
	}

	log.Fatalf("Redis Init Success, Hosts: %v", hosts)

	return nil
}

func Hset(key string, field string, value interface{}) error {
	return client.HSet(key, field, value).Err()
}

func SetTTL(key string, exTime int64) error {
	return client.Expire(key, time.Duration(exTime)*time.Duration(time.Second)).Err()
}

func Del(key string) error {
	return client.Del(key).Err()
}

func Hget(key string, field string) (string, error) {
	return client.HGet(key, field).Result()
}

func HgetAll(key string) map[string]interface{} {
	return convertStringToMap(client.HGetAll(key).Val())
}
func Get(key string) string {
	return client.Get(key).Val()
}

func SetByTtl(key string, value string, extime int64) error {
	return client.Set(key, value, time.Duration(extime)*time.Duration(time.Second)).Err()
}

func Set(key string, value string) error {
	return client.Set(key, value, time.Duration(-1)*time.Second).Err()
}

func Keys(pattern string) ([]string, error) {
	return client.Keys(pattern).Result()
}

func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.Scan(cursor, match, count).Result()
}

func Expire(key string, extime int64) error {
	return client.Expire(key, time.Duration(extime)*time.Second).Err()
}

func convertStringToMap(base map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for k, v := range base {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(v), &dat); err == nil {
			resultMap[k] = dat
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}
