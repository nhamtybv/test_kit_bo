package database

import (
	"fmt"
	"os"
	"time"

	"go.etcd.io/bbolt"
)

var (
	dbname              = ".config.db"
	dbperms os.FileMode = 0644
	options             = &bbolt.Options{Timeout: 1 * time.Second}
)

func NewDBDB() (*bbolt.DB, error) {
	db, err := bbolt.Open(dbname, dbperms, options)
	if err != nil {
		return nil, fmt.Errorf("bolt database initial error %w", err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("settings"))
		if err != nil {
			return fmt.Errorf("create bucket error %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
