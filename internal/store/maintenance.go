package store

import (
	"example.com/arena/internal/model"
	"go.etcd.io/bbolt"
)

func (s *Store) SaveBatch(records []model.Record) error {
	for _, r := range records {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
func (s *Store) Export() ([]byte, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	return model.Encode(all)
}
