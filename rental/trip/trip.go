package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rentalpb "server/rental/api/gen/v1"
)

type Service struct {
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	s.Logger.Info("create trip", zap.String("start", req.Start))
	return nil, status.Error(codes.Unimplemented, "")
}
