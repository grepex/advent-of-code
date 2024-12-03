package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// open the file, defer close
	file, err := os.Open("./2024/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scan the lines of the file
	scanner := bufio.NewScanner(file)

	mulFunc := regexp.MustCompile(`mul\(\d+,\d+\)`)
	// mulFunc := regexp.MustCompile(`where`)

	var matches []string

	for scanner.Scan() {
		line := scanner.Text()

		lineMatches := mulFunc.FindAllString(line, -1)

		// for _, match := range lineMatches {
		// 	fmt.Println("Found match: ", match)
		// }
		matches = append(matches, lineMatches...)
	}

	digits := regexp.MustCompile(`\d+,\d+`)
	sum := 0

	for _, match := range matches {
		digitPair := digits.FindString(match)
		splitDigits := strings.Split(digitPair, ",")
		firstDigit, err := strconv.Atoi(splitDigits[0])
		if err != nil {
			fmt.Println(err)
		}
		secondDigit, err := strconv.Atoi(splitDigits[1])
		if err != nil {
			fmt.Println(err)
		}
		result := firstDigit * secondDigit
		sum += result
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Your sum is: ", sum)

}
