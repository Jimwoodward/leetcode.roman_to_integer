package main

import (
	"fmt"
)

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

	// Iterate through each byte of the input string
	for i := 0; i < lengthOfInput; i++ {
		fmt.Println("We should enter here 1")
		// If there is a byte past the i-th position...
		if i < lengthOfInput-1 {
			fmt.Println("We should enter here 2 times")
			// If the next byte in our string is greater than the indexed byte AND the combination of the two is in
			// our map, then look up the combined string representation of the bytes to see if they're one of the
			// special cases in roman numerals
			if romanToInt[string(s[i])] < romanToInt[string(s[i+1])] {
				fmt.Println("We should not enter here 1")
				// If that combo of bytes (as a string) is in our map, add it to retVal
				if val, ok := romanToInt[string(s[i])+string(s[i])]; ok {
					retVal += val
					// If the combo of bytes (as a string) is not in our map, then just add the current indexed byte...
				} else {
					retVal += romanToInt[string(s[i])]
				}
			} else {
				fmt.Println("We should enter here 2")
				retVal += romanToInt[string(s[i])]
				fmt.Println(retVal)
			}
			//retVal += romanToInt[string(s[i])]
			fmt.Printf("retVal = %v\n", retVal)

			// If i is that last byte in our input string, don't bother to look for a preceding rune
		} else {
			fmt.Println("We should enter here 3")
			retVal += romanToInt[string(s[i])]
			return retVal
		}
	}
	return retVal
}

func main() {
	fmt.Println(romanToInt("III"))
}