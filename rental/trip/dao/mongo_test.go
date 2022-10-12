package dao

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
	"os"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/id"
	mgo "server/shared/mongo"
	"server/shared/mongo/objid"
	mongotesting "server/shared/mongo/testing"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	db := mc.Database("SZTURC")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexes: %v", err)
	}
	m := NewMongo(db)

	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			tripID:     "63460ab7099d584ae96f4fe4",
			accountID:  "a1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "another_finished",
			tripID:     "63460ab7099d584ae96f4fe5",
			accountID:  "a1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "in_progress",
			tripID:     "63460ab7099d584ae96f4fe6",
			accountID:  "a1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "another_in_progress",
			tripID:     "63460ab7099d584ae96f4fe7",
			accountID:  "a1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account",
			tripID:     "63460ab7099d584ae96f4fe8",
			accountID:  "a2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	for _, cc := range cases {
		mgo.NewObjID = func() primitive.ObjectID {
			return objid.MustFromID(id.TripID(cc.tripID))
		}
		trip, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s error creating trip: %v", cc.name, err)
			continue
		}
		if trip.ID.Hex() != cc.tripID {
			t.Errorf("%s incorrect trip id; want: %q;got: %q", cc.name, cc.tripID, trip.ID.Hex())
		}
	}
}

func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("SZTURC"))
	acct := id.AccountID("account1")
	mgo.NewObjID = primitive.NewObjectID
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
		t.Fatalf("cannot create trip: %v", err)
	}

	got, err := m.GetTrip(c, objid.ToTripID(trip.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", got)
	}
	//got.Trip.Start.PoiName = "error start"
	if diff := cmp.Diff(trip, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs; -want + got: %s", diff)
	}
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
