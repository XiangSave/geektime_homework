package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
)

var letters = []rune("abcdefghjkmnpqrstuvwxyz123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func setKeys(n int) {
	log.Printf("test key len = %d start", n)
	c, err := redis.Dial("tcp", "172.16.1.185:6378")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect success")
	defer c.Close()

	for i := 0; i < 500000; i++ {
		intnStr := RandStringRunes(n)
		_, err = c.Do("Set", string(i), intnStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("test key len = %d end", n)
	time.Sleep(3 * time.Minute)
	_, err = c.Do("FLUSHALL")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Minute)

}

func main() {
	var kLen = []int{10, 20, 50, 100, 200, 1000, 5000}
	for _, klen := range kLen {
		setKeys(klen)
	}

}
