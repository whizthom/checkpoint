package main

import (
	"fmt"
	"os"
	"strconv"
)



// Write a program that finds all pairs of elements in an integer array that sum up to a given target value. T
// he program should output a list of pairs, each representing the indices of the elements that form the pair.

// In this exercise you must take in consideration the following:

// Ensure it's possible to have positive or negative integers in the array.
// Ensure each element is used only once in a pair, although the element can be repeated in different pairs.
// Allow for multiple pairs to sum up to the target value.
// The output messages should follow the one given in the examples bellow.
// Return the message "No pairs found." when no pair is present.
// Return the message "Invalid target sum." if the target is invalid.
// Return the message "Invalid number: " if the number in the array is invalid.
// For any input format that deviates from the specified format "[1, 2, 3, 4, 5]" "6", the program will return an "Invalid input." error message.

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	// validate array format: must start with [ and end with ]
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// validate target: must be a single integer
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// strip brackets
	inner := arrStr[1 : len(arrStr)-1]

	// handle empty array
	if len(trimSpace(inner)) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	// split by comma manually
	parts := splitByComma(inner)

	// parse each number
	nums := []int{}
	for _, part := range parts {
		part = trimSpace(part)
		n, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Invalid number: " + part)
			return
		}
		nums = append(nums, n)
	}

	// find pairs
	pairs := [][2]int{}
	used := make([]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				pairs = append(pairs, [2]int{i, j})
				used[i] = true
				used[j] = true
				break
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	// format output
	result := fmt.Sprintf("Pairs with sum %d: [", target)
	for i, pair := range pairs {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("[%d %d]", pair[0], pair[1])
	}
	result += "]"
	fmt.Println(result)
}

// trimSpace removes leading and trailing spaces manually
func trimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

// splitByComma splits a string by comma manually
func splitByComma(s string) []string {
	parts := []string{}
	current := ""
	for _, ch := range s {
		if ch == ',' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if len(current) > 0 {
		parts = append(parts, current)
	}
	return parts
}