package problems

import (
	"fmt"
	"os"

	"github.com/shodgesio/euler-solutions/problems/problem1"
	"github.com/shodgesio/euler-solutions/problems/problem2"
)

func Run(problem int) {
	problemmap := map[int]func(){
		1: problem1.RunSolution,
		2: problem2.RunSolution,
	}

	if p, ok := problemmap[problem]; ok {
		p()
	} else {
		fmt.Println("No solution found for the provided problem number.")
		os.Exit(1)
	}
}
