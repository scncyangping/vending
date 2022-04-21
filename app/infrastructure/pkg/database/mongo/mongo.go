package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	defaultDb  = "vending"
	defaultUri = "mongodb://%s"
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

func Init(c *Config) {
	if conn == nil {
		conn = c.new()
	}
}

// connect 提供连接
func connect() *mongo.Client {
	return conn
}

func (c *Config) new() *mongo.Client {
	opt := options.Client().ApplyURI(fmt.Sprintf(defaultUri, c.Host))
	if len(c.User) != 0 { // 部分连接不需要帐号密码
		opt.Auth = &options.Credential{
			Username: c.User,
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
		panic(err.(interface{}))
	}
	return client
}

type MgoV struct {
	database   string
	collection string
}

func Op(database, collection string) *MgoV {
	return &MgoV{
		database,
		collection,
	}
}

func OpCn(defaultCol string) *MgoV {
	return &MgoV{
		defaultDb,
		defaultCol,
	}
}

// InsertOne 插入单个文档
func (m *MgoV) InsertOne(value interface{}) string {
	collection := getCollection(m)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID.(string)
}

func getCollection(m *MgoV) *mongo.Collection {
	client := connect()
	collection := client.Database(m.database).Collection(m.collection)
	return collection
}

// InsertMany 插入多个文档
func (m *MgoV) InsertMany(values []interface{}) int {
	collection := getCollection(m)
	result, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		log.Fatal(err)
	}
	return len(result.InsertedIDs)
}

// Delete 删除
func (m *MgoV) Delete(b interface{}) int64 {
	collection := getCollection(m)
	count, err := collection.DeleteMany(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	return count.DeletedCount
}

// DeleteOne 删除满足条件的一条数据
func (m *MgoV) DeleteOne(filter interface{}) int64 {
	collection := getCollection(m)
	count, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return count.DeletedCount
}

// Update 更新文档
func (m *MgoV) Update(filter, update interface{}) int64 {
	collection := getCollection(m)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.UpsertedCount
}

// UpdateOne 更新单个文档
func (m *MgoV) UpdateOne(filter, update interface{}) int64 {
	collection := getCollection(m)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.UpsertedCount
}

// FindOne 查询单个文档
func (m *MgoV) FindOne(b interface{}, target interface{}) error {
	var err error
	collection := getCollection(m)
	singleResult := collection.FindOne(context.TODO(), b)
	if singleResult.Err() != nil {
		err = singleResult.Err()
	} else {
		err = singleResult.Decode(target)
	}
	return err
}

// Find 查询文档
func (m *MgoV) Find(filter interface{}, tSlice interface{}) error {
	var err error

	collection := getCollection(m)
	if cursor, er := collection.Find(context.TODO(), filter); er == nil {
		err = cursor.All(context.TODO(), tSlice)
	} else {
		err = er
	}
	return err
}

// Count 查询集合里有多少数据
func (m *MgoV) Count() int64 {
	collection := getCollection(m)
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return size
}

// FindBy 按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (m *MgoV) FindBy(skip, limit int64, sort, filter interface{}, tSlice interface{}) error {
	var err error

	collection := getCollection(m)
	findOptions := options.Find().SetSort(sort).SetLimit(limit).SetSkip(skip)

	if temp, er := collection.Find(context.Background(), filter, findOptions); er == nil {
		err = temp.All(context.TODO(), tSlice)
	} else {
		err = er
	}
	return err
}
