package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/id"
	mgo "server/shared/mongo"
)

const (
	accountIDField      = "accountid"
	profileField        = "profile"
	identityStatusField = profileField + ".identitystatus"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/SZTURC?readPreference=primary&ssl=false"))
	if err != nil {
		panic(any(err))
	}
	col := mc.Database("SZTURC").Collection("profile")
	var prevState rentalpb.IdentityStatus
	prevState = rentalpb.IdentityStatus_UNSUBMITTED
	filter := bson.M{
		identityStatusField: prevState,
	}
	if prevState == rentalpb.IdentityStatus_UNSUBMITTED {
		filter = mgo.ZeroOrDoesNotExist(identityStatusField, prevState)
	}
	aid := id.AccountID("a1")
	p := &rentalpb.Profile{
		Identity: &rentalpb.Identity{
			Name: "abc",
		},
		IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
	}
	filter[accountIDField] = aid.String()
	res, err := col.UpdateOne(c, filter, mgo.Set(bson.M{
		accountIDField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))
	if err != nil {
		return
	}
	fmt.Println(res)
}

func findRows(c context.Context, col *mongo.Collection) {
	cur, err := col.Find(c, bson.M{})
	if err != nil {
		panic(any(err))
	}
	for cur.Next(c) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenID string             `bson:"open_id"`
		}
		err := cur.Decode(&row)
		if err != nil {
			panic(any(err))
		}
		fmt.Printf("%+v\n", row)
	}
}

func insertRows(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})
	if err != nil {
		panic(any(err))
	}

	fmt.Printf("%+v", res)
}
