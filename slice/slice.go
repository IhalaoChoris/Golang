package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)

	for {
		fmt.Print("Insert a integer of your choice, finish the program typing X:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		stringnumber := scanner.Text()

		fmt.Println(stringnumber)

		if stringnumber != "X" {
			integernumber, err := strconv.Atoi(stringnumber)

			if err != nil {
				fmt.Println("Error during conversion")
				return
			}

			fmt.Println(integernumber)
			sli = append(sli, integernumber)
			sort.Ints(sli)
			fmt.Println(sli)

		} else {
			fmt.Println("The program has finished.")
			return
		}
	}
}
