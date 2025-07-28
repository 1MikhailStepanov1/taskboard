package storage

type Roles struct {
	*BaseStorage
}

func NewRolesStorage(storage *BaseStorage) *Roles {
	return &Roles{storage}
}
