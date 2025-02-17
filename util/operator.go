package util

import (
	"fmt"
	"go-crack/types"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Operator(i interface{}) {
	job := i.(types.Job)

	password := fmt.Sprintf("-p%s", job.Ran)

	cmd := exec.Command("unrar", "x", job.FileName, password)
	stdout, err := cmd.Output()

	if err != nil {
		//fmt.Println(err.Error())
	}

	// Print the output
	output := string(stdout)

	Store.CompletedOps += 1

	if strings.Contains(output, "exit status 11") {
		println(Printer.Sprintf("Found password '%s' after %d tries!", password, Store.CompletedOps))

		log.Fatal("password found")
	}

	go print(Printer.Sprintf("\r%v | Completed Operators: %d out of %d", time.Now().Format(time.DateTime), Store.CompletedOps, Store.ListLineCount))
}
