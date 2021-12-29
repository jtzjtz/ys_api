package mongo

import (
	"fmt"
	"github.com/jtzjtz/kit/conn/mongo_pool"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoConn() *mongo.Client {
	c, err := mongo_pool.Dbconnect()
	if err != nil {
		fmt.Println(err)
	}
	return c
}
