package main

import (
	"fmt"
	"sync"
	"time"
)

var symbols = []string{"+", "-", "*", "/"}

func produce(in chan string) {
	defer close(in)
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d", i)
		in <- fmt.Sprintf("%d", i)
	}
}

func consume(in, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range in {
		if s == "500" {
			out <- s
		}
	}
}

func stop(out chan string, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// func main() {
// 	in, out := make(chan string), make(chan string)
// 	wg := &sync.WaitGroup{}
// 	go produce(in)
// 	for i := 0; i < runtime.NumCPU(); i++ {
// 		wg.Add(1)
// 		go consume(in, out, wg)
// 	}
// 	go stop(out, wg)
// 	fmt.Println(<-out)
// }

func main() {
	mutex := &sync.Mutex{}
	m := make(map[interface{}]interface{})
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
			time.Sleep(1)
		}(i)
	}
	wg.Wait()
	fmt.Printf("%v\n", m)
}
