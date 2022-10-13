package id

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type TripID string

func (a TripID) String() string {
	return string(a)
}

type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

type CarID string

func (c CarID) String() string {
	return string(c)
}
