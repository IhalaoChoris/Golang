package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	names := make([]Name, 0)
	fmt.Print("Insert the file name: ")
	var filename string
	// reader := bufio.NewReader(os.Stdin)
	// filename, err := reader.ReadString('\n')
	// filename = strings.Trim(filename, "\n")

	// if err != nil {
	// 	fmt.Println("Error reading input.")
	// }
	fmt.Scan(&filename)

	opennedfile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error reading file. The error is: ", err)
		return
	}
	scanner := bufio.NewScanner(opennedfile)

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		var Person Name
		Person.fname, Person.lname = line[0], line[1]

		names = append(names, Person)
	}

	opennedfile.Close()

	for _, i := range names {
		fmt.Println(i.fname, i.lname)
	}
}
