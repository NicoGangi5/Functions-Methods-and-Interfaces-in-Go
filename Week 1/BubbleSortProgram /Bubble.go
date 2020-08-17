package main

import (
	"fmt"
	"strconv"
)

func main()  {

	sli := make([]int, 0)
	var input string

	for i := 0; i < 10; i++ {
		fmt.Print("Enter a number (", i + 1, "of 10): ")
		fmt.Scan(&input)

		number, _ := strconv.Atoi(input)
		sli = append(sli, number)
	}

	fmt.Println("Slice no sort: ", sli)
	BubbleSort(sli)
	fmt.Println("Slice sort: ", sli)

}

func BubbleSort(sliToSort []int) {
	for i := 0; i < len(sliToSort); i++ {
		for j := 0; j < len(sliToSort) - 1 - i; j++ {
			if sliToSort[j] > sliToSort[j+1] {
				Swap(sliToSort, j)
			}
		}
	}
}

func Swap(sliToSwap []int, position int) {
	var aux int

	aux = sliToSwap[position]
	sliToSwap[position] = sliToSwap[position + 1]
	sliToSwap[position + 1] = aux
}