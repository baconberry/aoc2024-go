package aoctypes

type IntRow []int

func (r IntRow) RemoveElement(i int) IntRow {
	if i == 0 {
		return r[1:]
	}
	if i == len(r)-1 {
		return r[:len(r)-1]
	}
	result := make(IntRow, len(r)-1)
	for idx, n := range r[0:i] {
		result[idx] = n
	}
	for idx, n := range r[i+1:] {
		result[i+idx] = n
	}
	return result
}
