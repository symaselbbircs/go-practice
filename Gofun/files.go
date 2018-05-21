package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test_files.go.txt") //os.Create creates a file

	if err != nil {
		log.Fatal(err) //logging the errors
	}
	file.WriteString("This is a test file created by golang")
	defer file.Close()

	chad, err := ioutil.ReadFile("test_files.go.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(chad)

	updatedstring := strings.Replace(readString, "golang", "go", 1) //replace text

	file.WriteString(updatedstring)

	fmt.Println(updatedstring)
}
