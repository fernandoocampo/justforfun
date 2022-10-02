package concurrency

import "sync"

func ProcessDataOne(data []Data) []Result {
	responseStream := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(len(data))

	for index := range data {
		go processOne(data[index], responseStream, &wg)
	}

	result := processResultOne(responseStream, len(data))

	wg.Wait()

	return result
}

func processResultOne(responseStream chan Result, size int) []Result {
	result := make([]Result, 0)
	count := 0
	for response := range responseStream {
		count++
		result = append(result, response)
		if count == size {
			close(responseStream)
		}
	}
	return result
}

func processOne(value Data, responseStream chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	result := value.Value%2 == 0

	responseStream <- Result{
		Value: value.Value,
		Ok:    result,
	}
}
