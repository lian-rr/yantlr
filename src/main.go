package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/lian-rr/yantlr/src/grammar"
	"github.com/lian-rr/yantlr/src/utils"
)

var (
	logger *utils.Logger

	inputPath string
)

func init() {
	logger = utils.NewLogger(os.Stderr)

	flag.StringVar(&inputPath, "input", "", ".yant file with grammar")

	flag.Parse()
}

func main() {
	lines := readFile(inputPath)

	lexer := grammar.NewLexer()

	lexer.LoadTokens(lines)

	for _, t := range lexer.Tokens {
		logger.Info(t.String())
	}

}

func readFile(path string) []string {
	var lines = make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		logger.Fatal("error opening input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		logger.Fatal("error reading the input", err)
	}

	return lines
}
