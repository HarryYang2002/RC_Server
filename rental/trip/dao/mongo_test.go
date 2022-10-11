package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/id"
	"server/shared/mongo/objid"
	mongotesting "server/shared/mongo/testing"
	"testing"
)

var mongoURI string

func TestCreateTrip(t *testing.T) {
	mongoURI = "mongodb://localhost:27017"
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("SZTURC"))
	acct := id.AccountID("account1")
	trip, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			PoiName: "startPoint",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			FeeCent:  100,
			KmDriven: 35,
			PoiName:  "endPoint",
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Errorf("cannot create trip: %v", err)
	}
	t.Errorf("inserted row %s with updatedat %v", trip.ID, trip.UpdatedAt)

	got, err := m.GetTrip(c, objid.ToTripID(trip.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", got)
	}
	t.Errorf("got trip %+v", got)
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
