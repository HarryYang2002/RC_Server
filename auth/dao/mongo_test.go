package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	mgo "server/shared/mongo"
	mongotesting "server/shared/mongo/testing"
	"testing"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("SZTURC"))
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			mgo.IDFieldName: mustObjID("6332a23e9adaf7a55fa0ab49"),
			openIDField:     "openid_1",
		},
		bson.M{
			mgo.IDFieldName: mustObjID("6332a23e9adaf7a55fa0ab70"),
			openIDField:     "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}

	mgo.NewObjID = func() primitive.ObjectID {
		return mustObjID("6332a23e9adaf7a55fa0ab78")
	}

	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "6332a23e9adaf7a55fa0ab49",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "6332a23e9adaf7a55fa0ab70",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "6332a23e9adaf7a55fa0ab78",
		},
	}
	//子测试
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("faild resolve account id for %q: %v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id: want: %q.got: %q", cc.want, id)
			}
		})
	}
}

func mustObjID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(any(err))
	}
	return objID
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
