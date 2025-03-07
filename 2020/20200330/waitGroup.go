package main

import (
	"fmt"
	"runtime"
	"sync"
)

func PrintChars(name string, group *sync.WaitGroup) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s: %c\n", name, ch)
		runtime.Gosched()
	}
	group.Done()
}

func main() {
	group := &sync.WaitGroup{}

	group.Add(3)
	go PrintChars("1", group)
	go PrintChars("2", group)

	PrintChars("3", group)
	group.Wait()
}
