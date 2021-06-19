package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Aggregate(collectionName string, totalPipeline, pipeline interface{}, opts ...*options.AggregateOptions) (total int64, docs []map[string]interface{}, err error) {
	docs = make([]map[string]interface{}, 0)
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	coll := db.Collection(collectionName)
	cur0, err := coll.Aggregate(ctx, totalPipeline)
	if err != nil {
		return
	}
	defer cur0.Close(ctx)
	var result0 = struct {
		Count int64
	}{}
	for cur0.Next(ctx) {
		if err = cur0.Decode(&result0); err != nil {
			return
		}
	}
	if err = cur0.Err(); err != nil {
		return
	}

	total = result0.Count

	if total == 0 {
		return 0, docs, nil
	}
	cur, err := coll.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		doc := make(map[string]interface{})
		if err = cur.Decode(&doc); err != nil {
			return
		}
		docs = append(docs, doc)
	}
	if err = cur.Err(); err != nil {
		return
	}
	return
}

func FindOneByID(collectionName, id string) (map[string]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	doc := make(map[string]interface{})
	err = db.Collection(collectionName).FindOne(ctx, filter).Decode(&doc)
	return doc, err
}

func FindOne(collectionName string, filter interface{}) (map[string]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	doc := make(map[string]interface{})
	err := db.Collection(collectionName).FindOne(ctx, filter).Decode(&doc)
	return doc, err
}

func Find(collectionName string, filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error) {
	docs := make([]map[string]interface{}, 0)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := db.Collection(collectionName).Find(ctx, filter, opts...)
	if err != nil {
		return docs, nil
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		doc := make(map[string]interface{})
		if err = cur.Decode(&doc); err != nil {
			return docs, nil
		}
		docs = append(docs, doc)
	}
	if err = cur.Err(); err != nil {
		return docs, nil
	}
	return docs, nil
}

func DeleteManyByIDs(collectionName string, ids []primitive.ObjectID) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": bson.M{"$in": ids}}
	_, err := db.Collection(collectionName).DeleteMany(ctx, filter)
	return err
}

func DeleteMany(collectionName string, filter interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection(collectionName).DeleteMany(ctx, filter)
	return err
}

func InsertOne(collectionName string, doc interface{}, opts ...*options.InsertOneOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection(collectionName).InsertOne(ctx, doc, opts...)
	return err
}

func InsertMany(collectionName string, docs []interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection(collectionName).InsertMany(ctx, docs)
	return err
}

func UpdateOneByID(collectionName string, id *primitive.ObjectID, doc interface{}, opts ...*options.UpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": doc}
	_, err := db.Collection(collectionName).UpdateOne(ctx, filter, update, opts...)
	return err
}

func UpdateManyByIDs(collectionName string, ids []*primitive.ObjectID, doc interface{}, opts ...*options.UpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": doc}
	_, err := db.Collection(collectionName).UpdateMany(ctx, filter, update, opts...)
	return err
}

func UpdateMany(collectionName string, filter, doc interface{}, opts ...*options.UpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	update := bson.M{"$set": doc}
	_, err := db.Collection(collectionName).UpdateMany(ctx, filter, update, opts...)
	return err
}

func FindOneAndUpdate(collectionName string, filter, doc interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	update := bson.D{{"$set", doc}}
	return db.Collection(collectionName).FindOneAndUpdate(ctx, filter, update, opts...).Err()
}

func CountDocuments(collectionName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return db.Collection(collectionName).CountDocuments(ctx, filter, opts...)
}
