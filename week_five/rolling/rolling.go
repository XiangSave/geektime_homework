package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type RollNumber struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

func NewRollNumber() *RollNumber {
	r := &RollNumber{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

func (r *RollNumber) getNowBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	return bucket
}

func (r *RollNumber) removeOldBucket() {
	old := time.Now().Unix() - 10
	for timestamp := range r.Buckets {
		if timestamp <= old {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *RollNumber) Increment(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getNowBucket()
	b.Value += i
	r.removeOldBucket()
}

func (r *RollNumber) GetRollValues() string {
	var str string
	for key, value := range r.Buckets {
		str = str + string(key) + fmt.Sprintf("%v", value)

	}
	return str

}

func main() {
	roll := NewRollNumber()
	for _, i := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19} {
		roll.Increment(i)
		time.Sleep(1 * time.Second)
		log.Println(roll.GetRollValues())
	}
}
