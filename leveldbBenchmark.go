package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/syndtr/goleveldb/leveldb"
)

const num = 10000000

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
	fmt.Println("Sequential write time cost: ", elapsed)

	// Delete data in ordered sequence
	start = time.Now()

	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Delete([]byte(s), nil)
		if err != nil {
			fmt.Println("err")
		}
	}

	elapsed = time.Since(start)
	fmt.Println("Delete data in ordered sequence time cost: ", elapsed)

	// Sequential Write
	start = time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Put([]byte(s), []byte("test"), nil)
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Sequential write time cost: ", elapsed)

	// Delete data in random sequence
	start = time.Now()

	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		err = db.Delete([]byte(s), nil)
		if err != nil {
			fmt.Println("err")
		}
	}

	elapsed = time.Since(start)
	fmt.Println("Delete data in random sequence time cost: ", elapsed)

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
	fmt.Println("Random write time cost: ", elapsed)

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
	fmt.Println("Sequential read time cost: ", elapsed)

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
	fmt.Println("Random read time cost: ", elapsed)

	// Batch write in ordered sequence
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
	fmt.Println("Batch write in ordered sequence time cost: ", elapsed)

	// Batch Write in random sequence
	start = time.Now()
	batchR := new(leveldb.Batch)
	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		batchR.Put([]byte(s), []byte("test"))
	}
	err = db.Write(batchR, nil)
	if err != nil {
		fmt.Println("err")
	}
	elapsed = time.Since(start)
	fmt.Println("Batch write in random sequence time cost: ", elapsed)
}