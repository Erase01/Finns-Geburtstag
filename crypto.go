package main

import (
	"crypto/md5"
	"encoding/hex"
)

func get_hash(testStr string) string {
	h := md5.New()
	h.Write([]byte(testStr))
	return hex.EncodeToString(h.Sum(nil))
}