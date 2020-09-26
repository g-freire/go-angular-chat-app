package repository

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


type MongoRepository interface {
	GetClient(host string) *mongo.Client
	ReadLastN(collection *mongo.Collection) []bson.M
}

type MongoGenericRepository struct {
	Conn *mongo.Client
}


func (m *MongoGenericRepository)  GetClient(host string)  {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	m.Conn = client
	if err != nil {
		log.Fatal(err)
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//	//		panic(err)
	//	//	}
	//	//}()
}


func (m *MongoGenericRepository) ReadLastN(database string, collection string, n int8) []bson.M {

	dcollection := m.Conn.Database(database).Collection(collection)
	filter := bson.M{}
	options := options.Find()
	options.SetSort(bson.D{{"_id", -1}})
	options.SetLimit(2)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cur, err := dcollection.Find(ctx, filter, options)

	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)

	var resultCollection []bson.M
	// iterates and append resutls
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		resultCollection = append(resultCollection, result)
		fmt.Printf("Result: %+v\n", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return resultCollection
}

//func (m *MongoGenericRepository) ReadLast(database string, collection string) []bson.M {
//
//
//
//}

