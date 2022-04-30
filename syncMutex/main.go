package main

import (
	"fmt"
	"sync"
)

var(
	balance int = 100
)

// Problem: Race condition
// build: go build --race main.go

//emulamos un deposito
func Deposit(amount int, w *sync.WaitGroup, lock *sync.RWMutex){
	defer  w.Done()
	lock.Lock() // Lock the mutex//bloquea el programa en esta linea . indica que alguien esta escribiendo en la variable b y que se debe esperar , como un semaforo
	b := balance
	balance = b + amount
	lock.Unlock()  // Unlock is a method of the Mutex struct
}

func Balance(lock *sync.RWMutex)int{
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}//1 solo deposita
// pero puede haber n consulta al balance

func main()  {
	var wg sync.WaitGroup// patra bloquear el porgrama

	// Mutex is a struct that implements the Lock and Unlock methods
	var lock sync.RWMutex // Mutex is a struct that implements the Lock and Unlock methods
	//evita que deposite uitilice la misma variables en distintas go-utinas

	for i := 1 ; i <= 5; i++{
		wg.Add(1)
		go Deposit(i *100,&wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}

/*
sync.Mutex.Lock() nos ayudara a bloquear el acceso a valores compartidos en diferentes go routines
sync.Mutex.Unlock()  desbloqueara nuevamente el valor al que necesitamos acceder para que otro go routine lo utilice.
*/

/*
Lock bloquea lecturas (con RLock) y escrituras (con Lock) de otras goroutines
Unlock permite nuevas lecturas (con Rlock) y/o otra escritura (con Lock)
RLock bloquea escrituras (Lock) pero no bloquea lecturas (RLock)
RUnlock permite nuevas escrituras (y también lecturas, pero por la naturaleza de RLock, estas no se vieron bloqueadas nunca)
En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por nuevos escritores, a diferencia de no usar nada.

Sacado de aquí mero
*/