package main

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	SetGetRedis()
}

func SetGetRedis() {
	// bikin pool untuk connection ke redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:4455")
	}, 10)

	pool.MaxActive = 10

	// ambil 1 koneksi
	conn := pool.Get()
	defer conn.Close()

	// coba untuk cache data dari data map
	user := map[string]string{
		"fullName": "emanuel dina prasetyawan",
		"lastName": "prasetyawan",
		"kota":     "jogja",
	}

	// untuk ambil cache di redis
	reply, _ := redis.Bytes(conn.Do("GET", "mahasiswa"))
	if reply != nil { // ketika ada datanya di redis
		log.Println("data dari redis ada, ini datanya")
		log.Println(string(reply))
	} else { // ketika ngga ada datanya di redis
		log.Println("data dari redis ngga ada, coba ambil dari db")
		output, _ := json.Marshal(user)

		// untuk set cache di redis
		conn.Do("SETEX", "mahasiswa", 20, string(output))
	}
}
