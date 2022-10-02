package concurrency

import (
	"context"
	"log"
)

func ProcessDataSeven(ctx context.Context, data []Data) []Result {

	inputStream := createDataStreamSeven(ctx, data) // buffered channel

	resultStream := processSeven(ctx, inputStream)

	result := processResultSeven(ctx, resultStream, len(data))

	return result
}

func createDataStreamSeven(ctx context.Context, data []Data) <-chan Data {
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

func processResultSeven(ctx context.Context, responseStream chan chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for {
		var stream <-chan Result
		select {
		case <-ctx.Done():
			log.Println("processResultSeven", ctx.Err())
			return nil
		case maybeStream, ok := <-responseStream:
			if !ok {
				return result
			}
			log.Println("receiving new channel")
			stream = maybeStream
		}

		for response := range stream {
			select {
			case <-ctx.Done(): // ignore
			default:
				count++
				result = append(result, response)
				if count == size {
					close(responseStream)
					return result
				}
			}
		}
		log.Println("finishing iterating stream")
	}
}

func processSeven(ctx context.Context, inputStream <-chan Data) chan chan Result {
	log.Println("starting processSeven goroutine")
	responseStream := make(chan (chan Result))
	for data := range inputStream {
		newResult := Result{
			Value: data.Value,
			Ok:    data.Value%2 == 0,
		}
		go func() {
			resultStream := make(chan Result)
			defer close(resultStream)

			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
					log.Println("processSeven", ctx.Err())
					return
				case responseStream <- resultStream:
				case resultStream <- newResult:
					return
				}
			}

			log.Println("closing goroutine")
		}()
	}

	return responseStream
}
