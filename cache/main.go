package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func fibonacci(n int) int{
	if n <= 1{
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Memory holds a function and a map of results
type Memory struct {
	f Function // Function to be used
	cache map [int]FunctionResult  // Map of results for a given key
	lock sync.Mutex
}

// A function has to recive a value and return a value and an error
type Function func( key int)(interface{}, error)

// The result of a function
type FunctionResult struct {
	value interface{}
	err error
}

// NewCache creates a new cache
func NewCache(f Function)*Memory{
	return &Memory{
		f: f,
		cache: make(map[int]FunctionResult),
	}
}

// Get returns the value for a given key
func(m *Memory)Get(key int)(interface{},error){
	m.lock.Lock() 	// Lock the cache
	result,exist := m.cache[key] 	// Check if the value is in the cache
	m.lock.Unlock() 	// Unlock the cache

	if !exist { // If the value is not in the cache, calculate it
		m.lock.Lock()
		result.value, result.err = m.f(key) // Calculate the value
		m.cache[key]=result  // Store the value in the cache
		m.lock.Unlock()
	}

	return result.value, result.err
}

// Function to be used in the cache
func GETFibonacci(n int)(interface{}, error)  {
	return fibonacci(n), nil
}


func main()  {
	start := time.Now()

	// Create a cache and some values
	cache := NewCache(GETFibonacci)
	fibo := []int{42, 40, 41, 42, 38, 41, 42}

	var wg sync.WaitGroup

	channel := make(chan int, 2)


	// For each value to calculate, get the value and print the time it took to calculate
	for _, n := range fibo{
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			channel <- 1
			start := time.Now()
			value, err := cache.Get(n)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d \n", n, time.Since(start),value)
			<-channel
		}(n)
		wg.Add(1)
	}
	fmt.Printf("Total time: %v\n", time.Since(start))
}