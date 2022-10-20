package words

import (
	"errors"
	"strings"
)

func ToNumber(word string) (int, error) {
	word = strings.TrimSpace(word)
	if word == "" {
		return 0, errors.New("err: invalid word")
	}

	word = strings.ReplaceAll(word, "-", " ")

	// i have map with value
	numbWord := map[string]int{
		"zero":      0,
		"one":       1,
		"two":       2,
		"three":     3,
		"four":      4,
		"five":      5,
		"six":       6,
		"seven":     7,
		"eight":     8,
		"nine":      9,
		"ten":       10,
		"eleven":    11,
		"twelve":    12,
		"thirteen":  13,
		"fourteen":  14,
		"fifteen":   15,
		"sixteen":   16,
		"seventeen": 17,
		"eighteen":  18,
		"nineteen":  19,
		"twenty":    20,
		"thirty":    30,
		"forty":     40,
		"fifty":     50,
		"sixty":     60,
		"seventy":   70,
		"eighty":    80,
		"ninety":    90,
	}

	scale := map[string]int{
		"hundred":  100,
		"thousand": 1_000,
		"million":  1_000_000,
		"billion":  1_000_000_000,
	}

	words := strings.Split(word, " ")
	res := 0
	lastScale := 1
	lastScaleLabel := ""
	isMinus := false
	for i := len(words) - 1; i >= 0; i-- {
		wr := words[i]
		// find number
		val, ok := numbWord[wr]
		if ok {
			res += lastScale * val

			if lastScaleLabel == "hundred" {
				lastScale = 1
			}
			continue
		}

		// find scale
		val, ok = scale[wr]
		if ok {
			lastScaleLabel = wr
			lastScale *= val
			continue
		}

		if wr == "minus" {
			isMinus = true
			continue
		}

		// find and
		if wr != "and" {
			return 0, errors.New("err: Invalid word")
		}
	}

	if isMinus {
		res *= -1
	}

	return res, nil
}
