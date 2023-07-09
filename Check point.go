package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * time.Second) // simulate work

	fmt.Printf("Worker %d waiting for checkpoint\n", id)
	checkpoint <- true // signal that checkpoint has been reached
	<-resume           // wait for permission signal
	fmt.Printf("Worker %d resuming\n", id)
	// continue with the remaining work
}
func main() {
	numWorkers := 5
	checkpoint := make(chan bool)
	resume := make(chan bool)
	var wg sync.WaitGroup

	// launch the workers goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}

	// wait for all workers to reach the checkpoint
	for i := 0; i < numWorkers; i++ {
		<-checkpoint
	}

	fmt.Println("All workers reached the checkpoint")

	fmt.Println("Resuming all workers")

	// signal all workers to resume
	for i := 0; i < numWorkers; i++ {
		resume <- true
	}
	wg.Wait()
	fmt.Println("All workers completed their work")
}
