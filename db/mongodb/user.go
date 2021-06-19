package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"my-project-name/model"
	"time"
)

const userCollection = "user"

func FindUserByName(userName string) (user model.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = db.Collection(userCollection).FindOne(ctx, bson.M{"user_name": userName}).Decode(&user)
	return
}
