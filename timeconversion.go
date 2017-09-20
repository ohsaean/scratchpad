package main

import (
	"fmt"
	"strconv"
)

/*
Given a time in -hour AM/PM format, convert it to military (-hour) time.

Note: Midnight is  on a -hour clock, and  on a -hour clock. Noon is  on a -hour clock, and  on a -hour clock.

Input Format

A single string containing a time in -hour clock format (i.e.:  or ), where  and .

Output Format

Convert and print the given time in -hour format, where .

Sample Input

07:05:45PM
Sample Output

19:05:45
https://www.hackerrank.com/challenges/time-conversion
https://www.mathsisfun.com/time.html
*/

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input string

	fmt.Scanf("%s", &input)

	var hh, mm, ss, apm string

	var pos int = 0
	for _, r := range input {

		char := string(r)
		if char == ":" {
			continue
		}

		if pos < 2 {
			hh += char
		} else if pos < 4 {
			mm += char
		} else if pos < 6 {
			ss += char
		} else {
			apm += char
		}
		pos++
	}

	hhInt, _ := strconv.Atoi(hh)

	if apm == "AM" {
		if hhInt == 12 {
			// For the first hour of the day (12 Midnight to 12:59 AM), subtract 12 Hours
			hhInt -= 12
			hh = strconv.Itoa(hhInt)
			if hhInt == 0 {
				hh = "00"
			}
		}
	} else if apm == "PM" {
		if hhInt >= 1 && hhInt <= 11 {
			hhInt += 12
			hh = strconv.Itoa(hhInt)
		}
	}

	fmt.Printf("%s:%s:%s", hh, mm, ss)
}
