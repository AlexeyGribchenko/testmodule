package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlexeyGribchenko/testmodule/internal/add"
)

func processArgs(args []string) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("no data to sum")
	}

	numbers := make([]int, 0, len(args))
	for i, arg := range args {
		number, err := strconv.Atoi(arg)
		if err != nil {
			return 0, fmt.Errorf("argument №%d is not a number: %s", i+1, arg)
		}
		numbers = append(numbers, number)
	}
	return add.Add(numbers...), nil
}

func main() {
	args := os.Args[1:]

	result, err := processArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
