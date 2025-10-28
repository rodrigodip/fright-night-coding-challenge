package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "./10-27-25_runtime-warp-again/hyperskill-dataset-117426216.txt"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	coords := parseDataSet(file)
	fmt.Println("Instructions:", coords)
}
func parseDataSet(file []byte) [][]float64 {
	output := make([][]float64, 0)
	var number []float64
	for instructions := range strings.SplitSeq(string(file), "\n") {
		lvl2 := strings.Split(instructions, " ")
		for _, n := range lvl2 {
			parsed := strings.Fields(n)
			for _, m := range parsed {
				f, err := strconv.ParseFloat(m, 64)
				if err != nil {
					panic(err)
				}
				fmt.Print(f)
			}
			output = append(output, number)
		}
	}
	return output
}
func EuclideanDistance4D(p1, p2 []float64) (float64, error) {
	if len(p1) != 4 || len(p2) != 4 {
		return 0, fmt.Errorf("ambos os pontos devem ter exatamente 4 dimensões")
	}
	dx := p2[0] - p1[0]
	dy := p2[1] - p1[1]
	dz := p2[2] - p1[2]
	dw := p2[3] - p1[3]
	return math.Sqrt(dx*dx + dy*dy + dz*dz + dw*dw), nil
}
