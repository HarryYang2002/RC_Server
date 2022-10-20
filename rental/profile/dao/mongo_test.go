package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/id"
	mgo "server/shared/mongo"
	mongotesting "server/shared/mongo/testing"
	"testing"
)

func TestUpdateProfile(t *testing.T)  {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("SZTURC"))
	acct := id.AccountID("account1")
	mgo.NewObjID = primitive.NewObjectID
	p := &rentalpb.Profile{
		Identity: &rentalpb.Identity{
			Name: "abc",
		},
		IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
	}
	err = m.UpdateProfile(c,acct,rentalpb.IdentityStatus_PENDING,p)
	if err != nil{
		fmt.Println(err)
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
