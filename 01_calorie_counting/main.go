package main

import (
	"bufio"
	"flag"
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

	if sum > max {
		max = sum
	}

	return max
}

func Top3CalorieCount(input io.Reader) int {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	sum := 0
	top3 := [3]int{}

	// simple sort to ensure the first element is the smallest
	sort3 := func() {
		if top3[0] > top3[1] {
			top3[0], top3[1] = top3[1], top3[0]
		}

		if top3[1] > top3[2] {
			top3[1], top3[2] = top3[2], top3[1]
		}
	}

	for s.Scan() {
		line := s.Text()
		if line == "" {
			if sum > top3[0] {
				top3[0] = sum
				sort3()
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

	// ensure we don't miss the last entry
	if sum > top3[0] {
		top3[0] = sum
		sort3()
	}

	return top3[0] + top3[1] + top3[2]
}

func main() {
	top3 := flag.Bool("top3", false, "sums the top 3 highest calory count")
	file := flag.String("file", "", "name of the input text file")
	flag.Parse()

	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if *top3 {
		fmt.Printf("sum of top 3 calories: %d\n", Top3CalorieCount(f))
	} else {
		fmt.Printf("max calories: %d\n", MaxCalorieCount(f))
	}
}
