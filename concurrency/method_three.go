package concurrency

import (
	"context"
	"log"
)

func ProcessDataThree(ctx context.Context, data []Data) []Result {
	responseStream := make(chan Result)

	for index := range data {
		go processThree(ctx, data[index], responseStream)
	}

	result := processResultThree(ctx, responseStream, len(data))

	return result
}

func processResultThree(ctx context.Context, responseStream chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return nil
		case response, ok := <-responseStream:
			if !ok { // channel was closed
				return result
			}
			count++
			result = append(result, response)
			if count == size {
				close(responseStream)
			}
		}
	}
}

func processThree(ctx context.Context, value Data, responseStream chan Result) {
	result := Result{
		Value: value.Value,
		Ok:    value.Value%2 == 0,
	}

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case responseStream <- result:
		return
	}

}
