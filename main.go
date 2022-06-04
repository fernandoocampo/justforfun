package mainds

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type Level int

const (
	// Debug log level
	Debug Level = 1 + iota
	// Info log level
	Info
	// Warn log level
	Warn
	// Error log level
	Error
)

func putOne() {
	var lock sync.Mutex
	data := make(map[int]string)

	put := func(key int, value string) error {
		lock.Lock()
		defer lock.Unlock()
		if _, ok := data[key]; ok {
			return errors.New("key exists")
		}
		data[key] = value
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			put(i, strconv.Itoa(i))
		}
	}()
	go func() {
		defer wg.Done()
		for i := 101; i < 200; i++ {
			put(i, strconv.Itoa(i))
		}
	}()
	wg.Wait()
	fmt.Println("the map has", len(data), "elements")
}

// func putTwo() {
// 	type item struct {
// 		key   int
// 		value string
// 	}
// 	inputStream := make(chan item)
// 	data := make(map[int]string)

// 	put := func(key int, value string) error {
// 		if _, ok := data[key]; ok {
// 			return errors.New("key exists")
// 		}
// 		data[key] = value
// 		return nil
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i < 100; i++ {
// 			put(i, strconv.Itoa(i))
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 101; i < 200; i++ {
// 			put(i, strconv.Itoa(i))
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println("the map has", len(data), "elements")
// }

func main() {
	// fmt.Println(Debug, Info, Warn, Error)
	// thes3.HeadObject()
	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.MinInt32)
	// fmt.Println("n << x is 'n times 2, x times', so, 1<<5 is 1 times 2, 5 times", 1<<5)
	// fmt.Println("2 << 5", 2<<5)
	// fmt.Println("-2 << 1", -2<<1)
	// fmt.Println("-2 >> 1", -2>>1)

	val := []int{1, 2}
	fmt.Println(val[len(val):], "len", len(val[len(val):]))

	w := "pepito"
	fmt.Println("w", w[:1])

	vals := make([]int, 0)
	fmt.Println("vals", vals, "cap", cap(vals), "len", len(vals))
	vals = append(vals, 1)
	fmt.Println("vals", vals, "cap", cap(vals), "len", len(vals))
	vals = append(vals, 2)
	fmt.Println("vals", vals, "cap", cap(vals), "len", len(vals))

	var vals2 [2]int
	fmt.Println("vals2", vals2, "cap", cap(vals2), "len", len(vals2))
}
