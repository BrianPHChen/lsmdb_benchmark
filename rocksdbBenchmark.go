package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/tecbot/gorocksdb"
)

const num = 1000000

func main() {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "./rocksdb")

	defer db.Close()
	ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()
	fmt.Println("Data Number: ", num)

	// Sequential Write
	start := time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Put(wo, []byte(s), []byte("test"))
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Sequential Write: ", elapsed)

	// Random Write
	start = time.Now()
	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		err = db.Put(wo, []byte(s), []byte("test"))
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Random Write: ", elapsed)

	// Sequential Read
	start = time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		data, err := db.Get(ro, []byte(s))
		_ = data
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Sequential Read: ", elapsed)

	// Random Read
	start = time.Now()
	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		data, err := db.Get(ro, []byte(s))
		_ = data
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Random Read: ", elapsed)
}