package store

import (
	"example.com/arena/internal/model"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = []string{"records", "audits", "workflows", "attachments"}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error { s.mu.Lock(); defer s.mu.Unlock(); return s.db.Close() }
func (s *Store) put(bucket, key string, v any) error {
	data, e := model.Encode(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		data := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if data == nil {
			return bbolt.ErrBucketNotFound
		}
		return model.Decode(append([]byte(nil), data...), v)
	})
}
func (s *Store) SaveRecord(r model.Record) error { return s.put("records", r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.get("records", id, &r)
	return r, e
}
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := model.Decode(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveAudit(a model.AuditEvent) error      { return s.put("audits", a.ID, a) }
func (s *Store) SaveWorkflow(w model.Workflow) error     { return s.put("workflows", w.ID, w) }
func (s *Store) SaveAttachment(a model.Attachment) error { return s.put("attachments", a.ID, a) }
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
