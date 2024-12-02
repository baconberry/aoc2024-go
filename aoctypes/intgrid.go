package aoctypes

type IntGrid [][]int

func (g IntGrid) Row(i int) []int {
	return g[i]
}
