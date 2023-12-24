package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
	a file, which I need to parse through to determine
	a certain value by adding up two digits from within the string

	example:
	xx2xxxx5 = 25

	therefore, the current objectives of this problem:
	- Create a function capable of examining a string to form an integer and return said integer
	- Design a storage for list of two digits
	- Function to iterate through said storage and sum up elements for answer

*/

func main() {
	localFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// make sure to close file afterwards
	defer localFile.Close()

	scanner := bufio.NewScanner(localFile)

	var total int

	for scanner.Scan() {
		value := findDigit(scanner.Text())
		i, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		total += i
	}

	fmt.Println(total)

}

func findDigit(str string) string {
	var result string
	var count int

	for i := range str {
		var currentUnicode int = int(str[i])

		if currentUnicode >= 48 && currentUnicode <= 57 {
			strs := string(currentUnicode)
			result += strs
			count++
		}
	}

	if count < 2 {
		result = result + result
	} else if count > 2 {
		start := result[0]
		end := result[len(result)-1]

		strs := string(start) + string(end)

		result = strs
	}

	return result
}
