package main

import (
	"crypto/md5"
)

func get_hash(testStr []byte) []byte {
	h := md5.New()
	h.Write(testStr)
	return h.Sum(nil)
}