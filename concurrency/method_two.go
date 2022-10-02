package concurrency

func ProcessDataTwo(data []Data) []Result {
	responseStream := make(chan Result)

	for index := range data {
		go processTwo(data[index], responseStream)
	}

	result := processResultTwo(responseStream, len(data))

	return result
}

func processResultTwo(responseStream chan Result, size int) []Result {
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

func processTwo(value Data, responseStream chan Result) {
	result := value.Value%2 == 0

	responseStream <- Result{
		Value: value.Value,
		Ok:    result,
	}
}
