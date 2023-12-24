package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(ParseFile("input.txt"))
}

func ParseFile(path string) int {
	// assigning the file and handling error if no file exists
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	// can't forget to close the file after its done being used
	defer file.Close()

	// behaves similar to the scanner class in java
	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		value := findDigit(scanner.Text())

		i, err := strconv.Atoi(value)

		if err != nil {
			//log.Fatal(err)
			result += 0
		}

		result += i
	}

	return result
}

func findDigit(str string) string {
	var result string
	var count int

	for i := range str {
		var currentUnicode int = int(str[i])

		if currentUnicode >= 48 && currentUnicode <= 57 {
			strs := string(rune(currentUnicode))
			//strs := strconv.Itoa(currentUnicode)

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
	} else if count == 0 {
		return "0"
	}

	return result
}
