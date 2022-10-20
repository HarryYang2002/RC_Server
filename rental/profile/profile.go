package profile

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rentalpb "server/rental/api/gen/v1"
	"server/rental/profile/dao"
	"server/shared/auth"
	"time"
)

// Service defines a profile service.
type Service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

// GetProfile gets profile for the current account.
func (s *Service) GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		fmt.Println(111)
		if err == mongo.ErrNoDocuments {
			fmt.Println(222)
			return &rentalpb.Profile{}, nil
		}
		s.Logger.Error("cannot get profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return pr, nil
}

// SubmitProfile submits a profile.
func (s *Service) SubmitProfile(c context.Context, i *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       i,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_UNSUBMITTED, p)
	if err != nil {
		s.Logger.Error("cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	go func() {
		time.Sleep(3 * time.Second)
		err := s.Mongo.UpdateProfile(context.Background(), aid,
			rentalpb.IdentityStatus_PENDING, &rentalpb.Profile{
				Identity:       i,
				IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
			})
		if err != nil {
			s.Logger.Error("cannot verify identity", zap.Error(err))
		}
	}()
	return p, nil
}

// ClearProfile clears profile for an account.
func (s *Service) ClearProfile(c context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		s.Logger.Error("cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return p, nil
}

func (s *Service) logAndConvertProfileErr(err error) codes.Code {
	if err == mongo.ErrNoDocuments {
		return codes.NotFound
	}
	s.Logger.Error("cannot get profile", zap.Error(err))
	return codes.Internal
}
