package core

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Env is wrapper of all sorts db clients
type Env struct {
	S Storage
}

func GetEnv() *Env {
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6380,localhost:6381,localhost:6382,localhost:6383"
	}
	passwd := os.Getenv("APP_REDIS_PASSWD")
	if passwd == "" {
		passwd = "5zktXpVO2MIwCZE5"
	}
	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = "0"
	}
	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to redis (addr: %s ,db: %d)", addr, db)

	st := NewRedisCli(strings.Split(addr, ","), passwd, db)
	return &Env{S: st}
}
