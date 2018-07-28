package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/syndtr/goleveldb/leveldb"
)

const num = 1000000

func main() {
	db, err := leveldb.OpenFile("./leveldb", nil)
	defer db.Close()

	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Data Number: ", num)

	// Sequential Write
	start := time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Put([]byte(s), []byte("test"), nil)
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
		err = db.Put([]byte(s), []byte("test"), nil)
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
		data, err := db.Get([]byte(s), nil)
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
		data, err := db.Get([]byte(s), nil)
		_ = data
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Random Read: ", elapsed)
}