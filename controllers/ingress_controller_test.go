package controllers

import "testing"

func TestShortHash(t *testing.T) {
	if expect, got := "c09b5928e", shortHash("foo.com"+"/httpbin22222"); got != expect {
		t.Errorf("expected %v got %v", expect, got)
	}
}