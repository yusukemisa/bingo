package bingo

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type bingoMatrix [5][5]int

type bingo interface {
	run()
}

type miniBingo struct{ w io.Writer }
type regularBingo struct{ w io.Writer }

func newBingo(regular bool, writer io.Writer) bingo {
	if regular {
		return &regularBingo{w: writer}
	}
	return &miniBingo{w: writer}
}

func (b *miniBingo) run() {
	print(b.w, b.createBingoMatrix())
}

func (b *miniBingo) createBingoMatrix() [5][5]int {
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

func print(w io.Writer, b bingoMatrix) {
	fmt.Fprintf(w, " ------------------------ \n")
	fmt.Fprintf(w, "|  B |  I |  N |  G |  O |\n")
	fmt.Fprintf(w, "| -- | -- | -- | -- | -- |\n")

	for i, v := range b {
		for j, x := range v {
			if i == 2 && j == 2 {
				fmt.Fprintf(w, "|    ")
				continue
			}
			fmt.Fprintf(w, "| %2d ", x)
		}
		fmt.Fprintf(w, "|\n")
	}

	fmt.Fprintf(w, " ------------------------ \n")
}
func (b *regularBingo) run() {
	print(b.w, b.createBingoMatrix())
}

func (r *regularBingo) createBingoMatrix() [5][5]int {

	// B,I,N,G,O 各列に配置する数字の取り出し元となる数列を作成する
	var bsrc, isrc, nsrc, gsrc, osrc [15]int
	var count = 1
	for _, src := range []*[15]int{&bsrc, &isrc, &nsrc, &gsrc, &osrc} {
		for i, _ := range src {
			src[i] = count

			count++
		}
	}

	// 各取り出し元配列からランダムに5つ数字を取り出して各列に格納する
	var B, I, N, G, O [5]int
	for _, colum := range []struct {
		src *[15]int
		dst *[5]int
	}{
		{&bsrc, &B},
		{&isrc, &I},
		{&nsrc, &N},
		{&gsrc, &G},
		{&osrc, &O},
	} {
		rand.Seed(time.Now().UnixNano())
		random := rand.Perm(15)
		for i, _ := range colum.dst {
			colum.dst[i] = colum.src[random[i]]
		}
	}
	matrix := [5][5]int{B, I, N, G, O}

	// B, I, N, G, Oは行なので転置する
	var trance [5][5]int
	for i, array := range matrix {
		for j, _ := range array {
			trance[j][i] = matrix[i][j]
		}
	}
	return trance
}
