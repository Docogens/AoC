package main

import (
	"log"
	"os"
)

func main() {
	localFile, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	//os.Stdout.Write(localFile)
	readByte := make([]byte, 5)

}

/*
func Solution() {

}
*/
