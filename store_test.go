package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransform: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestPathTransformFunc(t *testing.T) {
	key := "momsbespicture"
	expectedOriginKey := "71c208b7f5e7e40b01a6e329f6286d9a6dd56de8"
	expectedPathNameKey := "71c20/8b7f5/e7e40/b01a6/e329f/6286d/9a6dd/56de8"
	pathKey := CASPathTransformFunc(key)

	if pathKey.PathName != expectedPathNameKey {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathNameKey)
	}
	if pathKey.Filename != expectedOriginKey {
		t.Errorf("have %s want %s", pathKey.Filename, expectedOriginKey)
	}

	log.Printf("pathkey : %+v", &pathKey)
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 500; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		// read test
		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}
		b, _ := io.ReadAll(r)

		log.Printf("%s\n", b)

		if string(b) != string(data) {
			t.Errorf("want %s hava %s", data, b)
		}
		s.Delete(key)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	s := newStore()
	key := "momsspecials"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}
