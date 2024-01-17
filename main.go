package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(`Parallel Client for URL.
Run this program like you would use curl.
Use "--parallel <number>" to run multiple requests in parallel.`)
		return
	}

	start := time.Now()
	parallel := 1
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		if args[i] == "--parallel" && i+1 < len(args) {
			p, err := strconv.Atoi(args[i+1])
			if err != nil || p < 1 {
				fmt.Println("input valid number for --parallel")
				os.Exit(1)
			}

			parallel = p
			args = append(args[:i], args[i+2:]...)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, args []string) {
			defer wg.Done()
			cmd := exec.Command("curl", args...)

			// cmd.Stdin = os.Stdin
			// cmd.Stdout = os.Stdout
			// cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("error:", err)
			}
		}(&wg, args)
	}

	wg.Wait()
	fmt.Println("took:", time.Since(start))
}
