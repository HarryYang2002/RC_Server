package car

import (
	"context"
	"fmt"
	carpb "server/car/api/gen/v1"
	rentalpb "server/rental/api/gen/v1"
	"server/shared/id"
)

type Manager struct {
	CarService carpb.CarServiceClient
}

func (m *Manager) Verify(c context.Context, cid id.CarID, lco *rentalpb.Location) error {
	car, err := m.CarService.GetCar(c, &carpb.GetCarRequest{
		Id: cid.String(),
	})
	if err != nil {
		return fmt.Errorf("cannot get car: %v", err)
	}
	if car.Status != carpb.CarStatus_LOCKED {
		return fmt.Errorf("cannot unlock; car status is: %v", car.Status)
	}
	return nil
}

func (m *Manager) Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error {
	_, err := m.CarService.UnlockCar(c, &carpb.UnlockCarRequest{
		Id: cid.String(),
		Driver: &carpb.Driver{
			Id:        aid.String(),
			AvatarUrl: avatarURL,
		},
		TripId: tid.String(),
	})
	if err != nil {
		return fmt.Errorf("cannot unlock car: %v", err)
	}
	return nil
}

func (m *Manager) Lock(c context.Context, cid id.CarID) error {
	_, err := m.CarService.LockCar(c, &carpb.LockCarRequest{
		Id: cid.String(),
	})
	if err != nil {
		return fmt.Errorf("cannot lock car: %v", err)
	}
	return nil
}
