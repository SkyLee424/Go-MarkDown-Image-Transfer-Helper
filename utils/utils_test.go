package utils

import (
	"testing"
	"time"
)

func TestGenHashVal(t *testing.T) {
	s := "2024-03-25-08-43-000.jpg"
	s1 := GenHashValue(s)
	time.Sleep(time.Second)
	s2 := GenHashValue(s)
	if s1 != s2 {
		t.Fatalf("hash value is not equal")
	}
}
