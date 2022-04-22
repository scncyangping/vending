package mongo

import (
	"fmt"
	"testing"
)

var tConfig = &Config{
	Host:            "139.9.0.61:17017",
	PoolSize:        300,
	MaxConnIdleTime: 600,
	DbName:          defaultDb,
}

func TestMgo_InsertOne(t *testing.T) {
	Init(tConfig)
	OpCn("user").InsertOne(map[string]interface{}{
		"_id":  "123",
		"name": "name",
		"age":  uint8(2),
	})
}

func TestMgo_InsertMany(t *testing.T) {
	Init(tConfig)
	arr := make([]interface{}, 0)
	for i := 0; i < 5; i++ {
		arr = append(arr, map[string]interface{}{
			"_id":  "123",
			"name": "name",
			"age":  i,
		})
	}
	fmt.Println(OpCn("user").InsertMany(arr))
}

func TestMgo_Delete(t *testing.T) {
	Init(tConfig)
	OpCn("user").Delete(map[string]interface{}{
		"age": map[string]interface{}{
			"$in": []int{2},
		},
	})
}

func TestMgo_Update(t *testing.T) {
	Init(tConfig)
	OpCn("user").Update(map[string]interface{}{
		"age": map[string]interface{}{
			"$in": []int{2},
		},
	}, map[string]interface{}{
		"$set": map[string]interface{}{
			"isUpdate": true,
		},
	})
}

func TestMgo_FindOne(t *testing.T) {
	Init(tConfig)
	//U := &User{}
	var U User
	//U := new(User)
	//U.Age = 1111
	//(*U).Age = 2222
	fmt.Println(&U)
	fmt.Println(U)
	//U := make(map[string]interface{})
	e := OpCn("user").FindOne(map[string]interface{}{
		"age": 1,
	}, &U)
	fmt.Println(U)
	fmt.Println(e)

}

type User struct {
	Id       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Age      int    `bson:"age" json:"age"`
	IsUpdate bool   `bson:"isUpdate" json:"isUpdate"`
}

func TestMgo_Find(t *testing.T) {
	Init(tConfig)
	//U := make([]User, 0)
	U := new([]*User)
	OpCn("user").Find(map[string]interface{}{
		"age": 1,
	}, U)
	fmt.Printf("%v", U)
	for _, v := range *U {
		fmt.Println(v)
	}
}

func TestMgo_Count(t *testing.T) {
	Init(tConfig)
	fmt.Println(OpCn("user").Count())
}

func TestMgo_FindBy(t *testing.T) {
	Init(tConfig)
	var uList []*User
	OpCn("user").FindBy(0, 2,
		map[string]interface{}{"age": 1},
		map[string]interface{}{"age": map[string]interface{}{"$gte": 1}}, &uList)
	fmt.Println(uList)
}
