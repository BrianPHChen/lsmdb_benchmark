package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/tecbot/gorocksdb"
)

const num = 10000000

func main() {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "./rocksdb")

	defer db.Close()
	ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()
	wb := gorocksdb.NewWriteBatch()
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
	fmt.Println("Sequential write time cost: ", elapsed)

	// Delete data in ordered sequence
	start = time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		err = db.Delete(wo, []byte(s))
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
		err = db.Put(wo, []byte(s), []byte("test"))
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
		err = db.Delete(wo, []byte(s))
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
		err = db.Put(wo, []byte(s), []byte("test"))
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
		data, err := db.Get(ro, []byte(s))
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
		data, err := db.Get(ro, []byte(s))
		_ = data
		if err != nil {
			fmt.Println("err")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Random read time cost: ", elapsed)

	// Batch write in ordered sequence
	wb.Clear()
	start = time.Now()
	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		wb.Put([]byte(s), []byte("test"))
	}
	err = db.Write(wo, wb)
	if err != nil {
		fmt.Println("err")
	}
	elapsed = time.Since(start)
	wb.Clear()
	fmt.Println("Batch write in ordered sequence time cost: ", elapsed)

	// Batch Write in random sequence
	wb.Clear()
	start = time.Now()
	for _, value := range rand.Perm(num) {
		s := strconv.Itoa(value)
		wb.Put([]byte(s), []byte("test"))
	}
	err = db.Write(wo, wb)
	if err != nil {
		fmt.Println("err")
	}
	elapsed = time.Since(start)
	wb.Clear()
	fmt.Println("Batch write in random sequence time cost: ", elapsed)
}