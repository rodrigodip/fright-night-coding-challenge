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
	//fmt.Println(coords)
	var answer int
	for i := 0; i < len(coords); i++ {
		if len(coords[i+1]) != 4 {
			break
		}
		sum, err := EuclideanDistance4D(coords[i], coords[i+1])
		if err != nil {
			fmt.Println("euclidian error")
		}
		answer = answer + int(sum)
	}
	fmt.Println(answer)
}
func parseDataSet(file []byte) [][]float64 {
	output := make([][]float64, 0)
	for coordSlice := range strings.SplitSeq(string(file), "\n") {
		seq := strings.Split(coordSlice, ",")
		parsed := stringToFloat(seq)
		output = append(output, parsed)
	}
	return output
}
func stringToFloat(s []string) []float64 {
	result := make([]float64, len(s))
	for i, str := range s {
		if len(s) != 4 {
			break
		}
		flt, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
		}
		//fmt.Printf("string: %v ,Type: %T\n", str, str)
		//fmt.Printf("float: %v ,Type: %T\n", flt, flt)
		result[i] = flt
	}
	return result
}
func EuclideanDistance4D(p1, p2 []float64) (float64, error) {
	if len(p1) != 4 || len(p2) != 4 {
		return 0, fmt.Errorf("ambos os pontos devem ter exatamente 4 dimensões")
	}
	dx := p2[0] - p1[0]
	dy := p2[1] - p1[1]
	dz := p2[2] - p1[2]
	dw := p2[3] - p1[3]
	return math.Ceil(math.Sqrt(dx*dx + dy*dy + dz*dz + dw*dw)), nil
}
