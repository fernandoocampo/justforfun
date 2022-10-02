package concurrency

import (
	"context"
	"log"
	"runtime"
)

func ProcessDataSix(ctx context.Context, data []Data) []Result {
	responseStream := make(chan Result, len(data))
	inputStream := createDataStreamSix(ctx, data) // buffered channel

	// It’s more important for efficiency in resource usage.
	g := runtime.GOMAXPROCS(0) // give me the number of available system threads.
	semaphore := make(chan bool, g)
	for i := 0; i < len(data); i++ {
		go processSix(ctx, inputStream, responseStream, semaphore)
	}

	result := processResultSix(ctx, responseStream, len(data))

	return result
}

func createDataStreamSix(ctx context.Context, data []Data) <-chan Data {
	inputStream := make(chan Data)

	go func() {
		defer close(inputStream)

		for index := range data {
			newData := data[index]
			select {
			case <-ctx.Done():
				log.Println("createDataStream", ctx.Err())
				return
			case inputStream <- newData:
			}
		}

		log.Println("closing stream")
	}()

	return inputStream
}

func processResultSix(ctx context.Context, responseStream chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("processResultSix", ctx.Err())
			return nil
		case response := <-responseStream:
			count++
			result = append(result, response)
			if count == size {
				close(responseStream)
				return result
			}
		}
	}
}

func processSix(ctx context.Context, inputStream <-chan Data, responseStream chan<- Result, semaphore chan bool) {
	log.Println("starting processSix goroutine")

	semaphore <- true
	defer func() {
		<-semaphore
	}()

	select {
	case <-ctx.Done():
		log.Println("processSix", ctx.Err())
		return
	case newData, ok := <-inputStream:
		if !ok {
			log.Println("processSix", newData.Value, "input stream was closed")
			return
		}

		responseStream <- Result{
			Value: newData.Value,
			Ok:    newData.Value%2 == 0,
		}
	}
}
