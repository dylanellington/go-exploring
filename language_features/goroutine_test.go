package language_features

import (
	//"fmt"
	"sync"
	"testing"
)

// Test_Goroutine_ChannelProcessing uses goroutines to make and receive items.
// A channel is passed as a sender to the make function, and as a receiver to the receiving function.
// A wait group is used to ensure the make and receive processes finish before validating the test.
func Test_Goroutine_ChannelProcessing(t *testing.T) {
	itemsToMake := 5
	channel := make(chan int)
	processedItems := make([]int, 0)

	makeItems := func(itemsToMake int, channel chan<- int) {
		slice := make([]int, itemsToMake)

		for index := 0; index < itemsToMake; index++ {
			slice[index] = index + 1
		}

		for _, item := range slice {
			channel <- item
		}

		close(channel)
	}

	receiveItems := func(channel <-chan int) {
		for item := range channel {
			//not thread safe but okay for this test
			processedItems = append(processedItems, item)
		}
	}

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		makeItems(itemsToMake, channel)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		receiveItems(channel)
	}()

	waitGroup.Wait()

	for index, item := range processedItems {
		//fmt.Printf("\nItem Processed: %v", item)
		if processedItems[index] != index + 1 {
			t.Errorf("Processed item %v does not match expected item %v.", item, index + 1)
		}
	}
}
