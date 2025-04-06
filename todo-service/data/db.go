package data

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Database
}

type Collection struct {
	collection *mongo.Collection
}

func Setup(uri, database string) (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	return &Database{db: client.Database(database)}, nil
}

func (db *Database) GetCollection(collectionName string) Collection {
	return Collection{collection: db.db.Collection(collectionName)}
}

func (collection *Collection) InsertOne(ctx context.Context, data interface{}) (*primitive.ObjectID, error) {
	result, err := collection.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	if ObjID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &ObjID, nil
	}

	return nil, errors.New("invalid object id")
}

func (collection *Collection) InsertMany(ctx context.Context, data []interface{}) error {
	_, err := collection.collection.InsertMany(ctx, data)
	return err
}

func (collection *Collection) FindAll(ctx context.Context, results interface{}) error {
	cur, err := collection.collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	err = cur.All(ctx, results)
	return err
}

func (collection *Collection) FindById(ctx context.Context, id string, result interface{}) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = collection.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(result)
	return err
}

func (collection *Collection) UpdateByIDs(ctx context.Context, id string, data interface{}) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := collection.collection.UpdateByID(ctx, objID, bson.M{"$set": data})
	return result, err
}
