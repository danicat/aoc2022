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

	// A = Rock, X = Rock
	// B = Paper, Y = Paper
	// C = Scissors, Z = Scissors
	moves["A"] = map[string]int{
		"X": 3,
		"Y": 6,
		"Z": 0,
	}

	moves["B"] = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	moves["C"] = map[string]int{
		"X": 6,
		"Y": 0,
		"Z": 3,
	}

	scores = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
}

func RockPaperScissors(input io.Reader) int {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	result := 0
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		result += moves[parts[0]][parts[1]] + scores[parts[1]]
	}

	return result
}

func main() {
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

	fmt.Printf("total score: %d\n", RockPaperScissors(f))
}
