package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BubbleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type an integer and press enter ten times to fullfil the array to be sorted: ")

	for i := 0; i < 10; i++ {
		scanner.Scan()
		integer, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Error converting type! Do you inputted an integer? ")
		}

		slice = append(slice, integer)
	}
	BubbleSort(slice)
	fmt.Println(slice)
}
