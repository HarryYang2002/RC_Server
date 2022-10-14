package profile

import (
	"context"
	"server/shared/id"
)

type Manager struct {
}

func (p *Manager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("a1"), nil
}
