package repository

import (
	"context"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	mgo *mongo.Client
	rds *redis.Client
}

type Repository struct {
}

func NewRepository(mgo *mongo.Client, r *redis.Client) *Repository {
	//var a = repository{mgo, r}
	//return &Repository{
	//	&AuthToken{a},
	//}
	return nil
}

func (r *repository) Close() {
	if r.mgo != nil {
		_ = r.mgo.Disconnect(context.Background())
	}
	if r.rds != nil {
		_ = r.rds.Close()
	}
}
