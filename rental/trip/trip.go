package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/trip/dao"
	"server/shared/auth"
	"server/shared/id"
)

type Service struct {
	ProfileManage ProfileManage
	CarManager    CarManager
	PoiManager    PoiManager
	Mongo         *dao.Mongo
	Logger        *zap.Logger
}

// ProfileManage ACL层（Anti Corruption Layer)防入侵
type ProfileManage interface {
	Verify(context.Context, id.AccountID) (id.IdentityID, error)
}

type CarManager interface {
	Verify(context.Context, id.CarID, *rentalpb.Location) error
	Unlock(context.Context, id.CarID) error
}

type PoiManager interface {
	Resolve(context.Context, *rentalpb.Location) (string, error)
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	//验证驾驶者身份
	iID, err := s.ProfileManage.Verify(c, aid)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	//检查车辆状态
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(c, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	//Tips：创建行程须在车辆开锁之前，如车辆开锁但创建行程失败没有补救措施，而创建行程成功但开锁失败还是有补救措施的
	//创建行程
	poi, err := s.PoiManager.Resolve(c, req.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", req.Start), zap.Error(err))
	}

	ls := &rentalpb.LocationStatus{
		Location: req.Start,
		PoiName:  poi,
	}
	trip, err := s.Mongo.CreateTrip(c, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      carID.String(),
		IdentityId: iID.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})
	if err != nil {
		s.Logger.Warn("cannot create trip", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}
	//车辆开锁
	go func() {
		err = s.CarManager.Unlock(context.Background(), carID)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   trip.ID.Hex(),
		Trip: trip.Trip,
	}, nil
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if req.Current != nil {
		tr.Trip.Current = s.calcCurrentStatus(tr.Trip, req.Current)
	}
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	s.Mongo.UpdateTrip(c, tid, aid, tr.UpdatedAt, tr.Trip)
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
