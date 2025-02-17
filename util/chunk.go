package util

import (
	"go-crack/types"
	"math"
	"strings"
	"sync"
)

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool) {

	var wg sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	chunkSize := 300
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < (noOfThread); i++ {

		wg.Add(1)
		go func(s int, e int) {
			defer wg.Done() //to avoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}

				Operator(types.Job{Ran: text, FileName: SelectedOptions.FileLocation})
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg.Wait()
	logsSlice = nil
}
