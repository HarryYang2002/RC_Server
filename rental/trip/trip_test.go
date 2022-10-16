package trip

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/trip/client/poi"
	"server/rental/trip/dao"
	"server/shared/auth"
	"server/shared/id"
	mgo "server/shared/mongo"
	mongotesting "server/shared/mongo/testing"
	"server/shared/server"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	pm := &profileManage{}
	cm := &carManage{}
	s := newService(c, t, pm, cm)

	nowFunc = func() int64 {
		return 1665829594
	}
	req := &rentalpb.CreateTripRequest{
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 113.2323,
		},
		CarId: "car1",
	}
	pm.iID = "identity1"
	golden := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":1665829594},"current":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":1665829594},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		accountID    string
		tripID       string
		profileErr   error
		carVerifyERR error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:      "normal_create",
			accountID: "a1",
			tripID:    "63460c3237970cf89b68d543",
			want:      fmt.Sprintf(golden, "a1"),
		},
		{
			name:       "profile_err",
			accountID:  "a2",
			tripID:     "63460c3237970cf89b68d544",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			accountID:    "a3",
			tripID:       "63460c3237970cf89b68d545",
			carVerifyERR: fmt.Errorf("car_verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			accountID:    "a4",
			tripID:       "63460c3237970cf89b68d546",
			carUnlockErr: fmt.Errorf("car_unlock"),
			want:         fmt.Sprintf(golden, "a4"),
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgo.NewObjIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyERR
			c := auth.ContextWithAccountID(context.Background(), id.AccountID(cc.accountID))
			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error; got none")
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("error creating trip: %v", err)
				return
			}
			if res.Id != cc.tripID {
				t.Errorf("incorrect id;want: %q,got: %q", cc.tripID, res.Id)
			}
			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("cannot marshall response: %v", err)
			}
			got := string(b)
			if cc.want != got {
				t.Errorf("incorrect response:want: %s,got:%s", cc.want, got)
			}
		})
	}
}

func TestTripLifecycle(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("account_for_lifecycle"))
	s := newService(c, t, &profileManage{}, &carManage{})

	tid := id.TripID("66660c3237970cf89b68d545")
	mgo.NewObjIDWithValue(tid)

	cases := []struct {
		name    string
		op      func() (*rentalpb.Trip, error)
		want    string
		now     int64
		wantErr bool
	}{
		{
			name: "create_trip",
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					Start: &rentalpb.Location{
						Latitude:  32.123,
						Longitude: 113.2323,
					},
					CarId: "car1",
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			now:  10000,
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":10000},"current":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  32.444,
						Longitude: 114.7789,
					},
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":10000},"current":{"location":{"latitude":32.444,"longitude":114.7789},"fee_cent":9527,"km_driven":1034,"poi_name":"故宫","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":10000},"current":{"location":{"latitude":32.444,"longitude":114.7789},"fee_cent":18643,"km_driven":1034,"poi_name":"故宫","timestamp_sec":30000},"end":{"location":{"latitude":32.444,"longitude":114.7789},"fee_cent":18643,"km_driven":1034,"poi_name":"故宫","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛","timestamp_sec":10000},"current":{"location":{"latitude":32.444,"longitude":114.7789},"fee_cent":18643,"km_driven":1034,"poi_name":"故宫","timestamp_sec":30000},"end":{"location":{"latitude":32.444,"longitude":114.7789},"fee_cent":18643,"km_driven":1034,"poi_name":"故宫","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "update_after_finished",
			now:  50000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
				})
			},
			wantErr: true,
		},
	}
	rand.Seed(1234)
	for _, cc := range cases {
		nowFunc = func() int64 {
			return cc.now
		}
		trip, err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s: failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s: incorrect response; want: %s, got: %s", cc.name, cc.want, got)
		}
	}
}

func newService(c context.Context, t *testing.T, pm ProfileManage, cm CarManager) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	db := mc.Database("SZTURC")
	mongotesting.SetupIndexes(c, db)

	return &Service{
		ProfileManage: pm,
		CarManager:    cm,
		PoiManager:    &poi.Manager{},
		DistanceCalc:  &distCalc{},
		Mongo:         dao.NewMongo(db),
		Logger:        logger,
	}
}

type profileManage struct {
	iID id.IdentityID
	err error
}

type carManage struct {
	verifyErr error
	unlockErr error
}

type poiManage struct {
	iID id.IdentityID
	err error
}

func (p *profileManage) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return p.iID, p.err
}

func (c *carManage) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return c.verifyErr
}

func (c *carManage) Unlock(context.Context, id.CarID) error {
	return c.unlockErr
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}

type distCalc struct {
}

func (*distCalc) DistanceKm(c context.Context, from *rentalpb.Location, to *rentalpb.Location) (float64, error) {
	if from.Latitude == to.Latitude && from.Longitude == to.Longitude {
		return 0, nil
	}
	return 1034, nil
}
