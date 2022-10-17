package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	ch3 <- struct{}{}
	for i := 0; i < 100; i++ {
		var group sync.WaitGroup
		group.Add(3)
		go func() {
			Dog(ch3, ch1, i)
			defer group.Done()
		}()
		go func() {
			Cat(ch1, ch2, i)
			defer group.Done()
		}()
		go func() {
			Fish(ch2, ch3, i)
			defer group.Done()
		}()
		group.Wait()
	}

}

func Dog(in, out chan struct{}, i int) {
	<-in
	fmt.Println("dog----------" + strconv.Itoa(i+1))
	out <- struct{}{}
}

func Cat(in, out chan struct{}, i int) {
	<-in
	fmt.Println("cat----------" + strconv.Itoa(i+1))
	out <- struct{}{}
}

func Fish(in, out chan struct{}, i int) {
	<-in
	fmt.Println("fish----------" + strconv.Itoa(i+1))
	out <- struct{}{}
}
