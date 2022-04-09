package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

type B bson.M

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

// Conn 提供连接
func Conn() *mongo.Client {
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
		panic(err)
	}
	return client
}

type mgo struct {
	database   string
	collection string
}

func Op(database, collection string) *mgo {
	return &mgo{
		database,
		collection,
	}
}

func OpCn(defaultCol string) *mgo {
	return &mgo{
		defaultDb,
		defaultCol,
	}
}

// InsertOne 插入单个文档
func (m *mgo) InsertOne(value interface{}) string {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID.(string)
}

// InsertMany 插入多个文档
func (m *mgo) InsertMany(values []interface{}) int {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		log.Fatal(err)
	}
	return len(result.InsertedIDs)
}

// Delete 删除
func (m *mgo) Delete(b interface{}) int64 {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	count, err := collection.DeleteMany(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	return count.DeletedCount
}

// DeleteOne 删除满足条件的一条数据
func (m *mgo) DeleteOne(filter interface{}) int64 {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	count, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return count.DeletedCount
}

// Update 更新文档
func (m *mgo) Update(filter, update B) int64 {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.UpsertedCount
}

// UpdateOne 更新单个文档
func (m *mgo) UpdateOne(filter, update interface{}) int64 {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.UpsertedCount
}

// FindOne 查询单个文档
func (m *mgo) FindOne(b interface{}) *mongo.SingleResult {
	client := Conn()
	collection, e := client.Database(m.database).Collection(m.collection).Clone()
	if e != nil {
		log.Fatal(e)
	}
	singleResult := collection.FindOne(context.TODO(), b)
	return singleResult
}

// Find 查询文档
func (m *mgo) Find(filter interface{}) *mongo.Cursor {
	client := Conn()
	collection, e := client.Database(m.database).Collection(m.collection).Clone()
	if e != nil {
		log.Fatal(e)
	}
	cursor, _ := collection.Find(context.TODO(), filter)
	return cursor
}

// Count 查询集合里有多少数据
func (m *mgo) Count() int64 {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return size
}

// FindBy 按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (m *mgo) FindBy(skip, limit int64, sort, filter interface{}) *mongo.Cursor {
	client := Conn()
	collection := client.Database(m.database).Collection(m.collection)
	findOptions := options.Find().SetSort(sort).SetLimit(limit).SetSkip(skip)
	temp, _ := collection.Find(context.Background(), filter, findOptions)
	return temp
}
