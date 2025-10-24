package main

import (
	"fmt"
)

func main() {
	findHim("bfkvrvsmdybdyxlxmlgwgapfalqwuowgpflplxoyugyxzapdajgjqpvowrammfzfdyjujsmobwpyvfmqbamumrbabrlpmspzfmdlgoyafusujqvqxzgxzmyrdfmifwwqlqyrdfzsfjfqgbpqoqpzalsuupbyrabdmlzpofqwvqjfxwrpmoxfrlyrsmjdxxurpyosbbdwvuvzmzggoblfvdfwbuzqvgqwgxaouvwyxjvfylfdsdlrbxgplbxmlbvwbfzysrbdjallubgduvdyyawjrjuuuygfgwvmavofzwwltxbwybbbvaqxsgsdvflwpopjjsvbzvmrgqbwsdxpvvowjgarpjqfxflwcblfxgpujyrvolzmqwfgbmqfbdgldggmbxdqdzjorwvwlogxyvfxybwrsardprsqgapgyrlpypmspmympslpjlrzvlaoxsjfbguzyyafadbpabsglzuwyqdvpbbfjppoyqxryuvyuugysbfaohassxrgvjswrasaqrdsolfugddxjupxqbmbxzdrvyouyuxlqapxyovsgvoqajxdpjsgjxxflqqwmmaoruggsbjzoymggqwwbsbxjyqyjwvsdegqpmrlqfavqrxmudfdusbqbmasxvwqwogysrazljjflfoyrzjuujrzgpypodoudgvmfqbrswwbfszppvjprgvduazyoujofzgfoarorwzwfllffjjaufwvblxvprafzbmpmbpuzpjsbouwbqwgrjwqzlqwdgxlqmfoluqxwvgsdzdbzlrpwsywlwpvsffzawfxzusjswmdywfbmbqjzpfgryjplqsgudylzyrzrdzxwdrguburrgylbmggbbflzdmrgvxygjgjrrvrlwafbozfllzbadvdsxdguqylsmdodumvrujdzbmzxdqaogxjlzgqfrdumzbdarxxormaslulllbgfzrymsjzxvoasxdpmggwnwuqmmwsusxxapzrloovpplm")

}
func findHim(s string) {
	runes := []rune(s)
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
	//type DigitFreq struct {
	//	digit string
	//	freq  int
	//}

	var uniques []string
	for digit, count := range score {
		if count == 1 {
			uniques = append(uniques, string(digit))
		}
	}
	//sort.Slice(pairs, func(i, j int) bool {
	//	return pairs[i].freq < pairs[j].freq
	//})
	fmt.Println(uniques)
}
