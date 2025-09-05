package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// checksum calculates GTIN check digit according to GS1 standard
func checksum(digits []int) int {
	sum := 0
	for i, d := range digits {
		if (len(digits)-i)%2 == 1 {
			sum += d * 3
		} else {
			sum += d
		}
	}
	return (10 - (sum % 10)) % 10
}

// validate checks complete GTIN
func validate(digits []int) bool {
	return len(digits) > 1 && checksum(digits[:len(digits)-1]) == digits[len(digits)-1]
}

// toDigits converts string to digit slice
func toDigits(s string) []int {
	var digits []int
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		}
	}
	return digits
}

// intSliceToString converts []int to string
func intSliceToString(digits []int) string {
	var sb strings.Builder
	for _, d := range digits {
		sb.WriteString(strconv.Itoa(d))
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("1=Calculate 2=Validate q=Quit: ")
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
		if input == "q" || input == "quit" {
			break
		}
		
		if input != "1" && input != "2" {
			continue
		}
		
		fmt.Print("Digits: ")
		if !scanner.Scan() {
			break
		}
		
		digits := toDigits(scanner.Text())
		if len(digits) == 0 {
			continue
		}
		
		if input == "1" {
			if len(digits) < 1 {
				continue
			}
			c := checksum(digits)
			full := append(digits, c)
			fmt.Printf("Checksum: %d\n", c)
			fmt.Printf("Complete: %s\n", intSliceToString(full))
		} else {
			if len(digits) < 2 {
				continue
			}
			if validate(digits) {
				fmt.Println("Valid")
			} else {
				fmt.Println("Invalid")
			}
		}
	}
}
