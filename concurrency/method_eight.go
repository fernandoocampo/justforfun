package concurrency

import (
	"context"
	"log"
)

func ProcessDataEight(ctx context.Context, data []Data) []Result {
	inputStream := createDataStreamEight(ctx, data) // buffered channel
	responseStream := processEight(ctx, inputStream)
	result := processResultEight(ctx, responseStream)

	return result
}

func createDataStreamEight(ctx context.Context, data []Data) <-chan Data {
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

func processResultEight(ctx context.Context, responseStream <-chan Result) []Result {
	result := make([]Result, 0)
	for {
		select {
		case <-ctx.Done():
			log.Println("processResultEight", ctx.Err())
			return nil
		case response, ok := <-responseStream:
			if !ok {
				log.Println("processResultEight: closing response stream")
				return result
			}

			result = append(result, response)
		}
	}
}

func processEight(ctx context.Context, inputStream <-chan Data) <-chan Result {
	log.Println("starting processEight goroutine")

	resultStream := make(chan Result)

	go func() {
		defer close(resultStream)
		for {
			select {
			case <-ctx.Done():
				log.Println("processEight", ctx.Err())
				return
			case newData, ok := <-inputStream:
				if !ok {
					log.Println("processEight", newData.Value, "input stream was closed")
					return
				}

				resultStream <- Result{
					Value: newData.Value,
					Ok:    newData.Value%2 == 0,
				}
			}
		}
	}()

	return resultStream
}
