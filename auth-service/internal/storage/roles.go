package storage

type Roles struct {
	*Base
}

func NewRoles(storage *Base) *Roles {
	return &Roles{storage}
}
