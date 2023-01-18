package composites

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBComposite struct {
	db *mongo.Database
}


