package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
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
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolve account id for 123: %v", err)
	} else {
		want := "6332a23e9adaf7a55fa0ab49"
		if id != want {
			t.Errorf("resolve account id: want: %q.got: %q", want, id)
		}
	}
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
