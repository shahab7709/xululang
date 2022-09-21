package xulu

import (
	"errors"
	"math"
)

const xuluSeprator = " "

var xuluChars = map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}

type xuluSubName struct {
	value  rune
	repeat int
}

type xuluName struct {
	subNames []xuluSubName
}

func NewXuluName(input string) (ixuluSentence, error) {
	result := xuluName{subNames: make([]xuluSubName, 0)}
	var oldItem = ' '
	var sub *xuluSubName = nil
	for _, item := range input {
		_, ok := xuluChars[item]
		if !ok {
			return nil, errors.New("invalid Name")
		}
		if item != oldItem {
			sub = &xuluSubName{value: item, repeat: 1}
			result.subNames = append(result.subNames, *sub)
		} else {
			result.subNames[len(result.subNames)-1].repeat++
		}
		oldItem = item
	}
	return &result, nil
}

func (xn *xuluName) compute() (int, error) {
	result := 0
	for _, item := range xn.subNames {
		itemvalue := xuluChars[item.value] * item.repeat
		itemvalue = int(math.Mod(float64(itemvalue), 5))
		itemvalue = int(math.Pow(float64(itemvalue), 2))
		result += itemvalue
	}
	return result, nil
}
