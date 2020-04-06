package bauch

import (
	"math/rand"
	"time"
	"unicode"
)

const maxLen = 4

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

// Say speaks the given input in Bauch format (lIKe ThiS :o)
func Say(input string) string {
	output := make([]rune, len(input))
	nUpper := 0
	i := -1
	for _, c := range input {
		i++
		if !unicode.IsLetter(c) {
			output[i] = c
			continue
		}
		if rand.Intn(2*maxLen) > maxLen+nUpper {
			output[i] = unicode.ToUpper(c)
			if nUpper >= 0 {
				nUpper++
			} else {
				nUpper = 0
			}
		} else {
			output[i] = unicode.ToLower(c)
			if nUpper <= 0 {
				nUpper--
			} else {
				nUpper = 0
			}
		}
	}
	output = output[:i+1]

	return string(output)
}

// Smile is a Bauch smile
const Smile = ":o)"
