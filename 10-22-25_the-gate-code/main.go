package main

import (
	"fmt"
	"sort"
)

// The code to the gates is four digits.
// It is the first, the second, the third and the fourth most common digit in today's sequence, in that particular order.
// If there are equal number of appearances, higher digit should come first.
func main() {
	countNumbers("9987290779537181861718263722926737826136503025253740464459006479641844337167780423045527805486487485796564404005834514758433893744573734142679573351009826286774886155650743623777093969385167938866452748931324645805729522598782390277530128824744572871985484656664844546724961785418017359261504443345038054795671382901358766583122187229186258992426474876995530509125371887825623157371602041358423407110428760223423316002921873856765809321097798603691227129867807153310545626184837999882786515535577506857341506414325435790197046136410828479464619014073966738000698885495159472350690077648538124929621606279851847745324112486540337881161610966294844598652739599751876541605967612667816695705007844393487387377655757909615797393156816045360689228457735681985548987628074864714312098546888772681114992128801573263588531291947693766094451680041027421535064567997323371398336917985993690182826721488285396139374990168352575488232852198135973795171327799532997656758793714072998820768329347615989599167004651")
}
func countNumbers(num string) {
	runes := []rune(num)
	score := make(map[rune]int)
	for _, n := range runes {
		_, ok := score[n]
		if ok {
			score[n] += 1
		} else {
			score[n] = 1
		}
	}
	// Create a slice of digit-frequency pairs
	type DigitFreq struct {
		digit rune
		freq  int
	}

	var pairs []DigitFreq
	for digit, count := range score {
		pairs = append(pairs, DigitFreq{digit, count})
	}

	// Sort by frequency (descending), then by digit (descending for equal frequencies)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].freq == pairs[j].freq {
			return pairs[i].digit > pairs[j].digit
		}
		return pairs[i].freq > pairs[j].freq
	})
	// Get the first four digits as the code
	code := make([]rune, 4)
	for i := 0; i < 4 && i < len(pairs); i++ {
		code[i] = pairs[i].digit
	}

	fmt.Printf("The gate code is: %c%c%c%c\n", code[0], code[1], code[2], code[3])

	// Print all frequencies for verification
	fmt.Println("\nDigit frequencies:")
	for _, pair := range pairs {
		fmt.Printf("Digit %c: %d occurrences\n", pair.digit, pair.freq)
	}
}
