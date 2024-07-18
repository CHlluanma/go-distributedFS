package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Original: hashStr,
	}
}

type PathTransformFunc func(string) PathKey

var DefaultPathTransform PathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Original: key,
	}
}

type PathKey struct {
	PathName string
	Original string
}

func (p PathKey) Filename() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}

type StoreOpts struct {
	PathTransform PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransform(key)
	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	pathAndFilename := pathKey.Filename()

	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}
	// defer f.Close()

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("wrote %d bytes to  disk :%s", n, pathKey.PathName)

	return nil
}
