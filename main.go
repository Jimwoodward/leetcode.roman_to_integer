package main

import (
	"fmt"
)

// Pseudo code:
// 1. Iterate through each byte of the input string.
// 2. If there is a byte past the ith position (if i is not the last byte in the string)...
// 3. If that preceding byte is greater than the ith byte (criteria #1 for a valid roman numeral combo)...
// 4. If the combination of the two bytes is a valid roman numeral combination i.e. "CD", "IX", etc., then add the
// corresponding value to the returned value and iterate i by 1 (since we have just evaluated the preceding byte we need
// to skip over it)
// 4. Else, then just add the corresponding value of the ith byte to the returned value
// 5. Else, then just add the corresponding value of the ith byte to the returned value
// 6. Else, then just add the corresponding value of the ith byte to the returned value and return
// 7. Return

func romanToInt(s string) int {
	var romanToInt = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	retVal := 0
	lengthOfInput := len(s)

	for i := 0; i < lengthOfInput; i++ {
		if i < lengthOfInput-1 {
			if romanToInt[string(s[i])] < romanToInt[string(s[i+1])] {
				if val, ok := romanToInt[string(s[i])+string(s[i+1])]; ok {
					retVal += val
					i += 1
				} else {
					retVal += romanToInt[string(s[i])]
				}
			} else {
				retVal += romanToInt[string(s[i])]
			}
		} else {
			retVal += romanToInt[string(s[i])]
			return retVal
		}
	}
	return retVal
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("CIX"))
}