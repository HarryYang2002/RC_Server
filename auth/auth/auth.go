package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "server/auth/api/gen/v1"
	"server/auth/dao"
	"time"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         *zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	OpenID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(c, OpenID)
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	tkn, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
