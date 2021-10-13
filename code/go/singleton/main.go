package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

func (DataBase) CreateConnection() {
	fmt.Println("Creating Singleton for Database...")
	time.Sleep(time.Second * 2)
	fmt.Println("Done!")
}

var db *DataBase
var lock sync.Mutex

func GetDataBase() *DataBase {
	lock.Lock()

	if db == nil {
		fmt.Println("Creating database...")
		db = &DataBase{}
		db.CreateConnection()
	} else {
		fmt.Println("DB Already Created!")
	}

	lock.Unlock()

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			GetDataBase()
			wg.Done()
		}()
	}

	wg.Wait()
}
