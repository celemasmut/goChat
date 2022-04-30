package main

import (
	"fmt"
	"sync"
	"time"
)

// Patron de diseño creacional que se asegura que solo exista una instancia de una clase

type Database struct {}

func (Database) CreatingSingletonConnection(){
	fmt.Println("creating singleton for data base")
	time.Sleep(2 * time.Second)
	fmt.Println("creation Done")
}

var db *Database
var lock sync.Mutex  // Mutex para evitar que se cree más de una instancia de la base de datos

func getDatabaseInstance() *Database  {
	lock.Lock()
	defer lock.Unlock()
	if db == nil{
		fmt.Println("creating DB connection ")
		db=&Database{}
		db.CreatingSingletonConnection()
	}else{
		fmt.Println("DB already created")
	}
	return db
}

func main(){
	var wg sync.WaitGroup

	// Lanzamos 10 gorutinas para pedir la instancia de la base de datos
	for i := 0 ; i < 10 ; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}