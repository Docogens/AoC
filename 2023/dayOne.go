package main

import (
	"log"
	"os"
)

func main() {
	localFile, err := os.ReadFile("AoC//input.txt")

	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(localFile)
}

/*
func Solution() {

}
*/
