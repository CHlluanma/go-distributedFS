package main

import (
	"bytes"
	"log"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbespicture"
	// expectedOriginKey := ""
	// expectedPathNameKey := ""
	pathKey := CASPathTransformFunc(key)

	// if pathKey.PathName != expectedOriginKey {
	// 	t.Errorf("have %s want %s", pathKey.PathName, expectedPathNameKey)
	// }
	// if pathKey.Original != expectedOriginKey {
	// 	t.Errorf("have %s want %s", pathKey.Original, expectedOriginKey)
	// }

	log.Printf("pathkey : %+v", &pathKey)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransform: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
