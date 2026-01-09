package bolt

import "go.etcd.io/bbolt"

var tasksBucket = []byte("tasks")

type BoltDBStorage struct {
	db *bbolt.DB
}

func NewBoltDBStorage(dbPath string) (*BoltDBStorage, error) {
	db, err := bbolt.Open(dbPath, 0o600, nil)

	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})

	if err != nil {
		db.Close()
		return nil, err
	}

	return &BoltDBStorage{db: db}, nil
}
