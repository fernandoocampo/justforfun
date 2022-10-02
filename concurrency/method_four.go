package concurrency

import (
	"context"
	"log"
)

func ProcessDataFour(ctx context.Context, data []Data) []Result {
	responseStream := make(chan Result, len(data))
	inputStream := createDataStreamFour(ctx, data)

	for i := 0; i < len(data); i++ {
		go processFour(ctx, inputStream, responseStream)
	}

	result := processResultFour(ctx, responseStream, len(data))

	return result
}

func createDataStreamFour(ctx context.Context, data []Data) <-chan Data {
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

func processResultFour(ctx context.Context, responseStream chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("processResultFour", ctx.Err())
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

func processFour(ctx context.Context, inputStream <-chan Data, responseStream chan<- Result) {
	log.Println("starting processFour goroutine")

	select {
	case <-ctx.Done():
		log.Println("processFour", ctx.Err())
		return
	case newData, ok := <-inputStream:
		if !ok {
			log.Println("processFour", newData.Value, "input stream was closed")
			return
		}

		responseStream <- Result{
			Value: newData.Value,
			Ok:    newData.Value%2 == 0,
		}
	}
}
