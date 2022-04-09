package mongo

import (
	"context"
	"fmt"
	"testing"
	"vending/common/util"
	"vending/common/util/snowflake"
)

var tConfig = &Config{
	Host:            "139.9.0.61:17017",
	PoolSize:        300,
	MaxConnIdleTime: 600,
	DbName:          defaultDb,
}

func TestMgo_InsertOne(t *testing.T) {
	Init(tConfig)
	OpCn("user").InsertOne(map[string]any{
		"_id":  snowflake.NextId(),
		"name": "name",
		"age":  uint8(2),
	})
}

func TestMgo_InsertMany(t *testing.T) {
	Init(tConfig)
	arr := make([]interface{}, 0)
	for i := 0; i < 5; i++ {
		arr = append(arr, map[string]any{
			"_id":     snowflake.NextId(),
			"name":    "name",
			"age":     i,
			"addTime": util.NowDateTimeFormat(),
		})
	}
	fmt.Println(OpCn("user").InsertMany(arr))
}

func TestMgo_Delete(t *testing.T) {
	Init(tConfig)
	OpCn("user").Delete(map[string]any{
		"age": map[string]any{
			"$in": []int{2},
		},
	})
}

func TestMgo_Update(t *testing.T) {
	Init(tConfig)
	OpCn("user").Update(map[string]any{
		"age": map[string]any{
			"$in": []int{2},
		},
	}, map[string]any{
		"$set": map[string]any{
			"isUpdate": true,
		},
	})
}

func TestMgo_FindOne(t *testing.T) {
	Init(tConfig)
	U := &User{}
	//U := make(map[string]any)
	OpCn("user").FindOne(map[string]interface{}{
		"age": 2,
	}).Decode(U)
	fmt.Println(U)
}

type User struct {
	Id       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Age      int    `bson:"age" json:"age"`
	IsUpdate bool   `bson:"isUpdate" json:"isUpdate"`
}

func TestMgo_Find(t *testing.T) {
	Init(tConfig)
	U := make([]*User, 0)
	OpCn("user").Find(B{
		"age": 0,
	}).All(context.TODO(), &U)
	fmt.Println(U)
}

func TestMgo_Count(t *testing.T) {
	Init(tConfig)
	fmt.Println(OpCn("user").Count())
}

func TestMgo_FindBy(t *testing.T) {
	Init(tConfig)
	uList := make([]User, 0)
	OpCn("user").FindBy(0, 2,
		B{"age": 1},
		B{"age": B{"$gte": 1}}).All(context.TODO(), &uList)
	fmt.Println(uList)
}
