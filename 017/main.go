package main

import "fmt"

func main() {
	letters := buildMap()
	count := 0
	for i := 1; i < 1001; i++ {
		word := ""
		j := i
		// thousands
		thou := int(j / 1000)
		if thou > 0 {
			j = j % 1000
			word += letters[thou] + letters[1000]
		}
		// hundreds
		hund := int(j / 100)
		if hund > 0 {
			j = j % 100
			word += letters[hund] + letters[100]
		}
		// tens
		if i >= 100 && j > 0 {
			word += "and"
		}
		if j <= 20 {
			word += letters[j]
		} else {
			tens := int(j / 10)
			word += letters[tens*10]
			j = j % 10
			if j > 0 {
				word += letters[j]
			}
		}
		count += len(word)
	}
	fmt.Printf("length = %d\n", count)
}

func buildMap() map[int]string {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
	m[6] = "six"
	m[7] = "seven"
	m[8] = "eight"
	m[9] = "nine"
	m[10] = "ten"
	m[11] = "eleven"
	m[12] = "twelve"
	m[13] = "thirteen"
	m[14] = "fourteen"
	m[15] = "fifteen"
	m[16] = "sixteen"
	m[17] = "seventeen"
	m[18] = "eighteen"
	m[19] = "nineteen"
	m[20] = "twenty"
	m[30] = "thirty"
	m[40] = "forty"
	m[50] = "fifty"
	m[60] = "sixty"
	m[70] = "seventy"
	m[80] = "eighty"
	m[90] = "ninety"
	m[100] = "hundred"
	m[1000] = "thousand"
	return m
}
