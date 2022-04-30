package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int)int {
	fmt.Printf("Calculate expensive Fibonacci for %d \n",n)
	time.Sleep(5*time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPEnding map[int] [] chan int
	Lock sync.RWMutex
}

func (s *Service)Work(job int){
	s.Lock.RLock()//lock lectura
	exist := s.InProgress[job]
	if exist{
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPEnding[job] = append(s.IsPEnding[job],response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for Response job : %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n",resp)
		return
	}
	s.Lock.RUnlock()
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calulate Fibonacci for %d\n",job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendignWorkers, exist := s.IsPEnding[job]
	s.Lock.RUnlock()


	if exist{
		for _, pendignWorker := range pendignWorkers{
			pendignWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job : %d\n",job)
	}
	s.Lock.Lock()
	s.InProgress[job]=false
	s.IsPEnding[job] = make([] chan int,0)
	s.Lock.Unlock()
}

func NewService()*Service{
	return &Service{
		InProgress: make(map[int]bool),
		IsPEnding: make(map[int][]chan int),
	}
}


func main(){
	service := NewService()
	jobs := []int{3,4,5,5,4,8,8,8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, n := range jobs{
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}