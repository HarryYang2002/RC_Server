package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "server/auth/api/gen/v1"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	Logger         *zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	OpenID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}
	return &authpb.LoginResponse{
		AccessToken: "token for open id" + OpenID,
		ExpiresIn:   7200,
	}, nil
}
