package go101

import "sync"

func LongTimeRequestQuestion(done <-chan interface{}, call func() int32) int32 {
	var result int32
	go func() {
		result = call()
	}()
	return result
}

func LongTimeRequest1(done <-chan interface{}, call func() int32) int32 { // doesn't work
	var result int32
	wait := make(chan interface{})
	go func() {
		result = call()
		close(wait)
	}()
	<-wait
	return result
}

func LongTimeRequest2(done <-chan interface{}, call func() int32) int32 {
	var result int32
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = call()
	}()
	wg.Wait()
	return result
}

func LongTimeRequest(done <-chan interface{}, call func() int32) int32 {
	result := make(chan int32)
	go func() {
		result <- call()
		close(result)
	}()
	return <-result
}

func LongTimeRequest4(done <-chan interface{}, call func() int32) int32 {
	var result int32
	resultChan := make(chan int32)
	go func() {
		resultChan <- call()
		close(resultChan)
	}()
	select {
	case <-done:
		return 0
	case value := <-resultChan:
		result = value
	}
	return result
}
