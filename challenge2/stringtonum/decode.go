package stringtonum

import (
	"strconv"
	"strings"
)

func Stringtonum(input string) string {
	decrvalue := make([]int, len(input)+1)

	for i, char := range input {
		switch char {
		case '=':
			decrvalue[i+1] = decrvalue[i]
		case 'L':
			if decrvalue[i] != 0 {
				decrvalue[i+1] = 0
			} else {
				decrvalue[i+1] = -1
			}
		case 'R':
			decrvalue[i+1] = decrvalue[i] + 1
		}

		if decrvalue[i+1] < 0 {
			decrvalue[i+1]++
			for j := i; j >= 0; j-- {
				switch input[j] {
				case 'L':
					if decrvalue[j] > decrvalue[j+1] {
						break
					} else {
						decrvalue[j]++
					}
				case 'R':
					if decrvalue[j] < decrvalue[j+1] {
						break
					}
				case '=':
					if decrvalue[j] == decrvalue[j+1] {
						break
					} else {
						decrvalue[j] = decrvalue[j+1]
					}
				}
			}
		}
	}

	var result strings.Builder
	for _, value := range decrvalue {
		result.WriteString(strconv.Itoa(value))
	}

	return result.String()
}
