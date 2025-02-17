package jobs

import (
	"fmt"
	"github.com/panjf2000/ants"
	"go-crack/types"
	"go-crack/util"
)

func BruteForce() {
	p, err := ants.NewPoolWithFunc(util.SelectedOptions.TotalThreads, util.Operator, ants.WithPreAlloc(true))
	if err != nil {
		panic(err)
	}

	options := len(util.Letters)

	for l := util.SelectedOptions.MinimumCharacters; l <= util.SelectedOptions.MaximumCharacters; l++ {
		totalOutcomes := options * l
		println(totalOutcomes)
		for i := 0; i < totalOutcomes; i++ {
			ran := util.RandSeq(l, int64(i))

			if util.SelectedOptions.IsPooled {
				err := p.Invoke(types.Job{Ran: ran, FileName: util.SelectedOptions.FileLocation})
				if err != nil {
					fmt.Println("Failed to invoke data")
				}
			} else {
				util.Operator(types.Job{Ran: ran, FileName: util.SelectedOptions.FileLocation})
			}
		}
	}

	if util.SelectedOptions.IsPooled {

	} else {

	}
}
