package jobs

import (
	"bufio"
	"fmt"
	"github.com/panjf2000/ants"
	"go-crack/types"
	"go-crack/util"
	"log"
	"os"
	"time"
)

func RockYou() {
	println(fmt.Sprintf("%v | Starting rockyou strategy...", time.Now().Format(time.DateTime)))

	p, err := ants.NewPoolWithFunc(util.SelectedOptions.TotalThreads, util.Operator, ants.WithPreAlloc(true))
	if err != nil {
		panic(err)
	}

	file, err := os.Open("./lists/rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total, err := util.LineCounter(file)

	util.Store.ListLineCount = total

	file, err = os.Open("./lists/rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}

	if util.SelectedOptions.IsLinear {
		err := util.Process(file)
		if err != nil {
			println(err)
			return
		}
		return
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if util.SelectedOptions.IsPooled {
			err := p.Invoke(types.Job{Ran: scanner.Text(), FileName: util.SelectedOptions.FileLocation})
			if err != nil {
				fmt.Println("Failed to invoke data")
			}
		} else {
			util.Operator(types.Job{Ran: scanner.Text(), FileName: util.SelectedOptions.FileLocation})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
