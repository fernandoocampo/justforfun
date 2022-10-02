package concurrency

import (
	"context"
	"log"
	"runtime"
)

func ProcessDataFive(ctx context.Context, data []Data) []Result {
	responseStream := make(chan Result)
	inputStream := createDataStreamFive(ctx, data) // buffered channel

	// It’s more important for efficiency in resource usage.
	g := runtime.GOMAXPROCS(0) // give me the number of available system threads.
	for i := 0; i < g; i++ {
		go processFive(ctx, inputStream, responseStream)
	}

	result := processResultFive(ctx, responseStream, len(data))

	return result
}

func createDataStreamFive(ctx context.Context, data []Data) <-chan Data {
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

func processResultFive(ctx context.Context, responseStream chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("processResultFive", ctx.Err())
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

func processFive(ctx context.Context, inputStream <-chan Data, responseStream chan<- Result) {
	log.Println("starting processFive goroutine")

	for {
		select {
		case <-ctx.Done():
			log.Println("processFive", ctx.Err())
			return
		case newData, ok := <-inputStream:
			if !ok {
				log.Println("processFive", newData.Value, "input stream was closed")
				return
			}

			responseStream <- Result{
				Value: newData.Value,
				Ok:    newData.Value%2 == 0,
			}
		}
	}
}
