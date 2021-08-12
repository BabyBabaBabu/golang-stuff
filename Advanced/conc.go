package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	myFood := []string{"Chapati", "Beef Stew", "Steamed Cabbage", "Banana Cake", "Grilled Chicken"}
	// introducing concurrency
	// using waitgroups
	var myWg sync.WaitGroup
	myWg.Add(len(myFood))
	for _, f := range myFood {
		go func(f string) {
			kitchen(f)
			myWg.Done()
		}(f)
	}
	myWg.Wait()
	fmt.Println("...Done after: ", time.Since(startTime))
}

func kitchen(f string) {
	if f == "Banana Cake" {
		fmt.Println("Baking...", f)
	} else {
		fmt.Println("Cooking...", f)
	}
	time.Sleep(10 * time.Second)
	fmt.Println(f, "is ready!")
}
