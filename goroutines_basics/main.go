package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 goroutine execution")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}

	fmt.Println("f1 goroutine finished")
	wg.Done()
	// (*wg).Done()
}

func f2() {
	fmt.Println("f2 goroutine execution")

	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}

	fmt.Println("f2 goroutine finished")
}

func main() {
	fmt.Println("numbers of cpus:", runtime.NumCPU())
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	fmt.Println("Go max procs", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("number of goroutines after f1:", runtime.NumGoroutine())

	f2()
	wg.Wait()
	// fmt.Println("number of goroutines after f2:", runtime.NumGoroutine())
	// time.Sleep(time.Second * 2)
	fmt.Println("final main")
}
