package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open the file, defer close
	file, err := os.Open("./2024/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scan the lines of the file
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		var ints []int

		for _, num := range nums {
			int, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, int)
		}

		direction := getDirection(ints)

		fmt.Println("Checking line: ", ints)
		fmt.Printf("Direction: %s\n", direction)

		isSafe := checkSafety(ints, direction)
		if isSafe {
			count++
		}

		fmt.Println(line)
		fmt.Printf("Result: %v\n", isSafe)
		fmt.Printf("Running count: %d\n\n", count)
	}

	fmt.Printf("Number of safe reports: %d", count)
}

func getDirection(line []int) string {
	var direction string

	if line[len(line)-1] > line[0] {
		direction = "increasing"
	} else if line[len(line)-1] < line[0] {
		direction = "decreasing"
	}

	return direction
}

func checkDirection(currentNum, previousNum int, direction string) bool {
	var result bool
	var currentDirection string

	if currentNum > previousNum {
		currentDirection = "increasing"
	} else if currentNum < previousNum {
		currentDirection = "decreasing"
	} else {
		currentDirection = "neutral"
	}

	result = direction == currentDirection

	return result
}

func checkDiff(currentNum, previousNum int) bool {
	var result bool

	diff := int(math.Abs(float64(currentNum) - float64(previousNum)))

	if diff < 1 || diff > 3 {
		result = false
	} else {
		result = true
	}

	return result
}

func checkSafety(nums []int, direction string) bool {
	isSafe := false
	fmt.Printf("Checking safety on %v\n", nums)

	// loop through slice, creating a subset that omits current index to test for safety
	for i := range nums {
		fmt.Printf("Creating a new slice, removing int at index %d\n", i)

		newSlice := make([]int, 0, len(nums)-1)
		newSlice = append(newSlice, nums[:i]...)
		newSlice = append(newSlice, nums[i+1:]...)

		fmt.Printf("The new slice: %v\n", newSlice)

		faultCount := 0

		for j := range newSlice {
			if j == 0 {
				continue
			}

			currentNum := newSlice[j]
			previousNum := newSlice[j-1]

			validDir := checkDirection(currentNum, previousNum, direction)
			validDiff := checkDiff(currentNum, previousNum)

			if !validDir || !validDiff {
				faultCount += 1
			}
		}

		if faultCount == 0 {
			isSafe = true
		}
	}

	return isSafe
}
