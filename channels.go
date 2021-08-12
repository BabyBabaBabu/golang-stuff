package main

import (
	"fmt"
	"sync"
)

// MUTEX EXAMPLE
var (
	myWg     sync.WaitGroup
	myMut    sync.Mutex
	customer string = "Kitaa"
	myBal    uint32 = 1000
)

func deposit(amount uint32, wg *sync.WaitGroup) {
	fmt.Println("Depositing...")
	myMut.Lock()
	myBal += amount
	myMut.Unlock()
	wg.Done()
	fmt.Println("Balance:", myBal)
}

func withdraw(amount uint32, wg *sync.WaitGroup) {
	fmt.Println("Withdrawing...")
	myMut.Lock()
	myBal -= amount
	myMut.Unlock()
	wg.Done()
	fmt.Println("Balance:", myBal)
}

func main() {
	fmt.Println("Hello", customer, ", your balance is: ", myBal)

	myWg.Add(2)

	go deposit(500, &myWg)

	go withdraw(700, &myWg)

	myWg.Wait()
	fmt.Println("...Done!")
}

// CHANNEL EXAMPLE
// // returns a receiver channel
// func write2Chan(arr []uint16) <-chan uint16 {

// 	myChan := make(chan uint16)
// 	go func() {
// 		for i := 0; i < len(arr); i++ {
// 			myChan <- arr[i]
// 		}
// 		close(myChan)
// 	}()
// 	return myChan
// }

// // receives a channel
// func sumChan(ch <-chan uint16) {

// 	var total uint16 = 0
// 	for val := range ch {
// 		total += val
// 	}
// 	fmt.Println("Total: ", total)
// }
// func main() {

// 	myArr := []uint16{78, 60, 89, 59, 70}
// 	myChan := write2Chan(myArr)
// 	sumChan(myChan)
// }
