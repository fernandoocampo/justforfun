package reversals

type Country struct {
	Name string
	Code string
}

type Handler struct {
	semaphore chan any
}

type Data struct {
	index int
	word  string
}

func NewHandler(workers int) *Handler {
	newHandler := Handler{
		semaphore: make(chan any, workers),
	}

	return &newHandler
}

func (h *Handler) Reverse(done <-chan any, data []Country) []string {
	resultStream := h.process(done, dataGenerator(done, data))

	result := make([]string, len(data))

	for {
		select {
		case <-done:
			return result
		case valueStream, ok := <-resultStream:
			if !ok {
				return result
			}
			select {
			case <-done:
				continue
			case value, ok := <-valueStream:
				if !ok {
					continue
				}
				result[value.index] = value.word
			}
		}
	}
}

func (h *Handler) process(done <-chan any, dataStream <-chan Data) <-chan (<-chan Data) {
	reversedStream := make(chan (<-chan Data))
	go func() {
		defer close(reversedStream)
		for {
			select {
			case <-done:
				return
			case data, ok := <-dataStream:
				if !ok {
					return
				}

				reversedStream <- h.newReverseWorker(done, data)
			}
		}
	}()
	return reversedStream
}

func dataGenerator(done <-chan any, data []Country) <-chan Data {
	dataStream := make(chan Data)
	go func() {
		defer close(dataStream)
		for index, country := range data {
			select {
			case <-done:
				return
			case dataStream <- country.toData(index):
			}
		}
	}()
	return dataStream
}

func (c Country) toData(index int) Data {
	return Data{
		index: index,
		word:  c.Name,
	}
}

func (d Data) reverseWord() Data {
	chars := []rune(d.word)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	d.word = string(chars)
	return d
}

func (h *Handler) newReverseWorker(done <-chan any, data Data) <-chan Data {
	dataStream := make(chan Data)
	h.semaphore <- 1
	go func() {
		defer close(dataStream)
		select {
		case <-done:
			return
		case dataStream <- data.reverseWord():
		}
		<-h.semaphore
	}()
	return dataStream
}
