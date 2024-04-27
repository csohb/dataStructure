package main

import (
	"log"
	"strconv"
	"sync"
)

type FizzBuzz struct {
	fizzArr []string
	mu      *sync.Mutex
}

func (f *FizzBuzz) print(i int) {
	log.Println("i:", i)
}

func (f *FizzBuzz) writingFizzBuzz(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	switch {
	case n%15 == 0:
		f.mu.Lock()
		defer f.mu.Unlock()
		f.fizzArr = append(f.fizzArr, "FizzBuzz")
	case n%3 == 0:
		f.mu.Lock()
		defer f.mu.Unlock()
		f.fizzArr = append(f.fizzArr, "Fizz")
	case n%5 == 0:
		f.mu.Lock()
		defer f.mu.Unlock()
		f.fizzArr = append(f.fizzArr, "Buzz")
	default:
		f.mu.Lock()
		defer f.mu.Unlock()
		f.fizzArr = append(f.fizzArr, strconv.Itoa(n))
	}
}

func main() {
	var mutex = &sync.Mutex{}
	fizzBuzz := &FizzBuzz{
		fizzArr: make([]string, 0),
		mu:      mutex,
	}
	var wg sync.WaitGroup // Create a WaitGroup to wait for all goroutines to finish
	for i := 1; i <= 100; i++ {
		wg.Add(1) // Increment the WaitGroup counter before starting a goroutine
		go fizzBuzz.writingFizzBuzz(i, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	for _, v := range fizzBuzz.fizzArr {
		log.Println(v)
	}
}
