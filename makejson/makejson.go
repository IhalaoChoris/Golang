package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please, insert your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Now, insert your adress: ")
	scanner.Scan()
	adress := scanner.Text()

	commits := map[string]string{
		"name":   name,
		"adress": adress,
	}
	makejson, err := json.Marshal(commits)

	if err != nil {
		fmt.Println("Error creating json.")
	}

	fmt.Print(string(makejson))
}
