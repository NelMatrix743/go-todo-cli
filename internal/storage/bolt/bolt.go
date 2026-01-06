package bolt

import "go.etcd.io/bbolt"

type BoltDBStorage struct {
	db *bbolt.DB
}

func NewBoltDBStorage(dbPath string) (*BoltDBStorage, error) {
	return nil, nil
}
