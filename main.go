package main

import (
	"fmt"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb"
)

const num = 100000

func main() {
	db, err := leveldb.OpenFile("./db", nil)
	defer db.Close()

	if err != nil {
		fmt.Println("err")
	}

	for i:=0; i < num; i++ {
		s := strconv.Itoa(i)
		fmt.Println(s)
		err = db.Put([]byte(s), []byte("hello"), nil)
	}

	fmt.Println("success")
}