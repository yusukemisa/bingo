package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(" ------------------------ ")
	fmt.Println("|  B |  I |  N |  G |  O |")
	fmt.Println("| -- | -- | -- | -- | -- |")

	for i, v := range createBingoMatrix() {
		for j, x := range v {
			if i == 2 && j == 2 {
				fmt.Print("|    ")
				continue
			}
			fmt.Printf("| %2d ", x)
		}
		fmt.Println("|")
	}
	fmt.Println(" ------------------------ ")
}

func createBingoMatrix() [5][5]int {
	var numSeq []int
	rand.Seed(time.Now().Unix())
	perm := rand.Perm(26)
	for _, v := range perm {
		if v != 0 {
			numSeq = append(numSeq, v)
		}
	}

	var bingoMatrix [5][5]int
	var k int
	for i, row := range bingoMatrix {
		for j := range row {
			bingoMatrix[i][j] = numSeq[k]
			k++
		}
	}
	return bingoMatrix
}
