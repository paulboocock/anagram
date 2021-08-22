package sorting

type SortAlphabetically []rune

func (a SortAlphabetically) Len() int           { return len(a) }
func (a SortAlphabetically) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortAlphabetically) Less(i, j int) bool { return a[i] < a[j] }
