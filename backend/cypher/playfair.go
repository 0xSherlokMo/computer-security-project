package cypher

import (
	"sort"
	"strings"

	"cs-be/utilities"
)

type PlayFair struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

func (p *PlayFair) Encrypt() string {
	var encryption string
	graph := p.setGraph()
	pairs := p.setPairs()
	for _, pair := range pairs {
		c1, c2 := p.coordination(pair[0], graph), p.coordination(pair[1], graph)
		if c1[1] == c2[1] {
			// Same row
			c1[0] = (c1[0] + 1) % 5
			c2[0] = (c2[0] + 1) % 5
		} else if c1[0] == c2[0] {
			// Same column
			c1[1] = (c1[1] + 1) % 5
			c2[1] = (c2[1] + 1) % 5
		} else {
			// Get the other corners of the square
			c1[0], c2[0] = c2[0], c1[0]
		}
		encryption += p.letterFromCoordination(c1[0], c1[1], graph)
		encryption += p.letterFromCoordination(c2[0], c2[1], graph)
	}

	return encryption
}

func (p *PlayFair) Decrypt() string {
	var encryption string
	graph := p.setGraph()
	pairs := p.setPairs()
	for _, pair := range pairs {
		c1, c2 := p.coordination(pair[0], graph), p.coordination(pair[1], graph)
		if c1[1] == c2[1] {
			// Same row
			c1[0] = c1[0] - 1
			if c1[0] == -1 {
				c1[0] = 4
			}

			c2[0] = c2[0] - 1
			if c2[0] == -1 {
				c2[0] = 4
			}
		} else if c1[0] == c2[0] {
			// Same column
			c1[1] = c1[1] - 1
			if c1[1] == -1 {
				c1[1] = 4
			}

			c2[1] = c2[1] - 1
			if c2[1] == -1 {
				c2[1] = 4
			}
		} else {
			// Get the other corners of the square
			c1[0], c2[0] = c2[0], c1[0]
		}
		encryption += p.letterFromCoordination(c1[0], c1[1], graph)
		encryption += p.letterFromCoordination(c2[0], c2[1], graph)
	}

	return encryption
}

func (p *PlayFair) setGraph() []string {
	str := strings.ToUpper(p.Key)
	alphabet := map[string]int{
		"A": 65,
		"B": 66,
		"C": 67,
		"D": 68,
		"E": 69,
		"F": 70,
		"G": 71,
		"H": 72,
		"I": 73,
		"K": 75,
		"L": 76,
		"M": 77,
		"N": 78,
		"O": 79,
		"P": 80,
		"Q": 81,
		"R": 82,
		"S": 83,
		"T": 84,
		"U": 85,
		"V": 86,
		"W": 87,
		"X": 88,
		"Y": 89,
		"Z": 90,
	}
	var strArr []string
	for _, character := range str {
		char := string(character)
		if char == "J" {
			strArr = append(strArr, "I")
		} else {
			strArr = append(strArr, char)
		}
		delete(alphabet, char)
	}
	var alphaArr []int
	for _, character := range alphabet {
		alphaArr = append(alphaArr, character)
	}
	sort.Ints(alphaArr)
	for _, character := range alphaArr {
		char := string(rune(character))
		strArr = append(strArr, char)
	}
	return strArr
}

func (p *PlayFair) setPairs() [][2]string {
	s := strings.ToUpper(p.Message)
	strArr := utilities.Split(s)

	var lp [][2]string
	for {
		// If there are no letters, then we're done
		if len(strArr) == 0 {
			break
		}

		var innerLp [2]string

		// If there is only one letter left
		if len(strArr) == 1 {
			innerLp[0] = strArr[0]
			innerLp[1] = "Z"
			lp = append(lp, innerLp)
			break
		}

		// Add the first letter to the list
		innerLp[0] = strArr[0]

		// Add an x if the next letter is the same, and loop with [1:]
		if strArr[0] == strArr[1] {
			innerLp[1] = "X"
			lp = append(lp, innerLp)
			strArr = strArr[1:]
		} else {
			innerLp[1] = strArr[1]
			lp = append(lp, innerLp)
			strArr = strArr[2:]
		}
	}
	return lp
}

func (p *PlayFair) coordination(s string, l []string) [2]int {
	s = strings.ToUpper(s)
	var index int
	for i, letter := range l {
		if letter == s {
			index = i
			break
		}
	}
	y := index / 5
	x := index - (5 * y)
	return [2]int{x, y}
}

func (p *PlayFair) letterFromCoordination(x int, y int, l []string) string {
	letter := l[y*5+x]
	return letter
}
