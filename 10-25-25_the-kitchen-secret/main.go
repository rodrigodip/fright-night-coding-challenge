package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	path := "./10-25-25_the-kitchen-secret/hyperskill-dataset-117389797.txt"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	instructions := parseDataSet(file)
	var letters []string
	for _, instruction := range instructions {
		//instruction := []string{"DOWN", "LEFT", "RIGHT", "UP", "UP", "RIGHT"}
		letters = append(letters, getLetter(instruction))
	}
	fmt.Println(letters)
}
func parseDataSet(file []byte) [][]string {
	output := make([][]string, 0)
	for instructions := range strings.SplitSeq(string(file), "\n") {
		lvl2 := strings.Split(instructions, " ")
		output = append(output, strings.Split(lvl2[0], ","))
	}
	return output
}
func getLetter(in []string) string {
	var table = [4][4]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
	}
	type cooredenada struct {
		linha  int
		coluna int
	}
	var coord cooredenada
	for _, move := range in {
		switch move {
		case "RIGHT":
			if (coord.coluna + 1) <= 3 {
				coord.coluna++
			}
			continue
		case "LEFT":
			if (coord.coluna - 1) < 0 {
				continue
			}
			coord.coluna--
		case "UP":
			if (coord.linha - 1) < 0 {
				continue
			}
			coord.linha--
		case "DOWN":
			if (coord.linha + 1) <= 3 {
				coord.linha++
			}
		}
	}
	return table[coord.linha][coord.coluna]
}
