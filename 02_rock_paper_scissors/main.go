package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// scoring tables
var moves map[string]map[string]int
var scores map[string]int

func init() {
	moves = make(map[string]map[string]int)

	moves["A"] = map[string]int{
		"A": 3,
		"B": 6,
		"C": 0,
	}

	moves["B"] = map[string]int{
		"A": 0,
		"B": 3,
		"C": 6,
	}

	moves["C"] = map[string]int{
		"A": 6,
		"B": 0,
		"C": 3,
	}

	scores = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
}

type strategyFn func(string, string) string

func fixed(oponentMove, playerMove string) string {
	switch playerMove {
	case "X":
		return "A"
	case "Y":
		return "B"
	case "Z":
		return "C"
	}

	return ""
}

func clever(opponentMove, playerMove string) string {
	strat := map[string]map[string]string{
		"A": { // rock
			"X": "C", // lose
			"Y": "A", // draw
			"Z": "B", // win
		},
		"B": { // paper
			"X": "A",
			"Y": "B",
			"Z": "C",
		},
		"C": { //scissor
			"X": "B",
			"Y": "C",
			"Z": "A",
		},
	}

	return strat[opponentMove][playerMove]
}

func RockPaperScissors(input io.Reader, strategy strategyFn) int {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	result := 0
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		playerMove := strategy(parts[0], parts[1])

		result += moves[parts[0]][playerMove] + scores[playerMove]
	}

	return result
}

func main() {
	file := flag.String("file", "", "name of the input text file")
	cleverStrat := flag.Bool("clever", false, "uses the clever elf strategy")
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

	strategy := fixed
	if *cleverStrat {
		strategy = clever
	}

	fmt.Printf("total score: %d\n", RockPaperScissors(f, strategy))
}
