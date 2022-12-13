package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func MaxCalorieCount(input io.Reader) int {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	sum := 0
	max := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}

		cur, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("failed to convert (%s) with error (%s)\n", line, err)
			continue
		}
		sum += cur
	}

	return max
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("you must specify a txt file as input")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Printf("max calories: %d\n", MaxCalorieCount(f))
}
