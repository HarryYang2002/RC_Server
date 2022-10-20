package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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


type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

type ProfileRecord struct {
	AccountID   string            `bson:"accountid"`
	Profile     *rentalpb.Profile `bson:"profile"`
}

func (m *Mongo) GetProfile(c context.Context, aid id.AccountID) (*rentalpb.Profile, error) {
	res := m.col.FindOne(c, byAccountID(aid))
	if err := res.Err(); err != nil {
		return nil, err
	}
	var pr ProfileRecord
	err := res.Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode profile record: %v", err)
	}
	return pr.Profile, nil
}

func byAccountID(aid id.AccountID) bson.M {
	return bson.M{
		accountIDField: aid.String(),
	}
}

func (m *Mongo) UpdateProfile(c context.Context, aid id.AccountID, prevState rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	filter := bson.M{
		identityStatusField: prevState,
	}
	if prevState == rentalpb.IdentityStatus_UNSUBMITTED {
		filter = mgo.ZeroOrDoesNotExist(identityStatusField, prevState)
	}
	filter[accountIDField] = aid.String()
	_, err := m.col.UpdateOne(c, filter, mgo.Set(bson.M{
		accountIDField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))
	return err
}
