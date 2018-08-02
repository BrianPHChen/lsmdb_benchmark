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
	fmt.Println("Start to sequential write")
	start := time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Put([]byte(s), []byte("test"), nil)
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Sequential write time cost: ", elapsed, "\n")

	// Delete data in ordered sequence
	fmt.Println("Start to delete data in ordered sequence")
	start = time.Now()

	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Delete([]byte(s), nil)
		if err != nil {
			fmt.Println("err")
		}
	}

	elapsed = time.Since(start)
	fmt.Println("Delete data in ordered sequence time cost: ", elapsed, "\n")

	// Sequential Write
	fmt.Println("Start to sequential write")
	start = time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Put([]byte(s), []byte("test"), nil)
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Sequential write time cost: ", elapsed, "\n")

	// Delete data in random sequence
	fmt.Println("Start to delete data in random sequence")
	start = time.Now()

	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		err = db.Delete([]byte(s), nil)
		if err != nil {
			fmt.Println("err")
		}
	}

	elapsed = time.Since(start)
	fmt.Println("Delete data in random sequence time cost: ", elapsed, "\n")

	// Random Write
	fmt.Println("Start to random write")
	start = time.Now()
	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		err = db.Put([]byte(s), []byte("test"), nil)
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Random write time cost: ", elapsed, "\n")

	// Sequential Read
	fmt.Println("Start to sequential read")
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
	fmt.Println("Sequential read time cost: ", elapsed, "\n")

	// Random Read
	fmt.Println("Start to random read")
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
	fmt.Println("Random read time cost: ", elapsed, "\n")

	// Batch write
	fmt.Println("Start to batch write in ordered sequence")
	start = time.Now()
	batch := new(leveldb.Batch)
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		batch.Put([]byte(s), []byte("test"))
	}
	err = db.Write(batch, nil)
	if err != nil {
		fmt.Println("err")
	}
	elapsed = time.Since(start)
	fmt.Println("Batch write in ordered sequence time cost: ", elapsed, "\n")
}