package trip

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"server/rental/trip/dao"
	"server/shared/auth"
	mgo "server/shared/mongo"
	"server/shared/server"

	rentalpb "server/rental/api/gen/v1"
	"server/rental/trip/client/poi"
	"server/shared/id"
	mongotesting "server/shared/mongo/testing"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := auth.ContextWithAccountID(context.Background(), id.AccountID("a1"))
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}
	pm := &profileManage{}
	cm := &carManage{}

	s := &Service{
		ProfileManage: pm,
		CarManager:    cm,
		PoiManager:    &poi.Manager{},
		Mongo:         dao.NewMongo(mc.Database("SZTURC")),
		Logger:        logger,
	}

	req := &rentalpb.CreateTripRequest{
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 113.2323,
		},
		CarId: "car1",
	}
	pm.iID = "identity1"
	golden := `{"account_id":"a1","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛"},"current":{"location":{"latitude":32.123,"longitude":113.2323},"poi_name":"天坛"},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		tripID       string
		profileErr   error
		carVerifyERR error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:   "normal_create",
			tripID: "63460c3237970cf89b68d543",
			want:   golden,
		},
		{
			name:       "profile_err",
			tripID:     "63460c3237970cf89b68d544",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			tripID:       "63460c3237970cf89b68d545",
			carVerifyERR: fmt.Errorf("car_verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			tripID:       "63460c3237970cf89b68d546",
			carUnlockErr: fmt.Errorf("car_unlock"),
			want:         golden,
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			mgo.NewObjIDWithValue(id.TripID(cc.tripID))
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyERR
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
