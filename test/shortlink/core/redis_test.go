package core_test

import (
	"gustz.com/rocket/v1.0/src/shortlink/core"
	"testing"
)

var (
	r  *core.RedisCli
	sl string
)

const (
	url = "http://wwww.baidu.com"
)

func TestNewRedisCli(t *testing.T) {
	addr := "127.0.0.1:6381"
	pass := "5zktXpVO2MIwCZE5"
	db := 0

	r = core.NewRedisCli(addr, pass, db)
	if err := r.Ping(); err != nil {
		t.Errorf("No server running at localhost:6379")
	}
}

func TestShorten(t *testing.T) {
	var err error
	exp := 60
	sl, err = r.Shorten(url, int64(exp))
	if err != nil {
		t.Error("Shorten failed")
	}
}

func TestUnshorten(t *testing.T) {
	u, err := r.Unshorten(sl)
	if err != nil {
		t.Error("Unshorten return error")
	}

	if u != url {
		t.Errorf("Expected receive url %s. Got url %s", url, u)
	}
}
