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

		isSafe := true

		var direction string
		if ints[1] > ints[0] {
			direction = "increasing"
		} else if ints[1] < ints[0] {
			direction = "decreasing"
		}

		// check unilateral direction
		for i := range ints {
			if i == 0 {
				continue
			}

			if direction == "increasing" {
				if ints[i] < ints[i-1] {
					isSafe = false
					break
				}
			} else if direction == "decreasing" {
				if ints[i] > ints[i-1] {
					isSafe = false
					break
				}
			}
		}

		if !isSafe {
			continue
		}

		// check diffs
		var diffs []int

		for i := range ints {
			if i == 0 {
				continue
			}

			diff := int(math.Abs(float64(ints[i]) - float64(ints[i-1])))
			diffs = append(diffs, diff)
		}

		for _, v := range diffs {
			if v < 1 || v > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			count += 1
		}

		fmt.Println(line)
		fmt.Println(diffs)
		fmt.Printf("Result: %v\n", isSafe)
	}

	fmt.Printf("Number of safe reports: %d", count)
}
