package pascal

func Triangle(n int) [][]int {

	if n == 1 {
		return [][]int{[]int{1}}
	}

	var (
		t    [][]int // Triangle result
		c    []int   // current line
		last = 1     // previous line index
	)

	t = make([][]int, n)
	t[0] = []int{1}
	t[1] = []int{1, 1}

	for i := 3; i <= n; i++ {
		c = make([]int, i)
		c[0] = 1
		c[i-1] = 1
		for j := 1; j < i-1; j++ {
			c[j] = t[last][j-1] + t[last][j]
		}
		t[i-1] = c
		last++
	}
	return t
}
