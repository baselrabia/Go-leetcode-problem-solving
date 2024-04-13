func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix) , len(matrix[0])
	cache := make(map[string]int)

	var helper func(r int, c int) int

	helper = func(r int, c int) int {
		if r >= rows || c >= cols {
			return 0
		}
        key := fmt.Sprintf("%d,%d", r, c)

		if _, exist := cache[key]; !exist {
			down := helper(r+1, c)
			right := helper(r, c+1)
			diag := helper(r+1, c+1)

			cache[key] = 0
            val, _ := strconv.Atoi(string(matrix[r][c]))
			if val == 1 {
				cache[key] = 1 + min(down, right, diag)
			}
		}

        return cache[key]
	}

    helper(0,0)


    max := 0
    for _, v := range cache {
        if v > max {
            max = v
        }
    }
    return max * max
}
 
func min(a, b, c int) int {
    min := a // Assume 'a' is the minimum initially
    if b < min {
        min = b // If 'b' is smaller than the current minimum, update 'min'
    }
    if c < min {
        min = c // If 'c' is smaller than the current minimum, update 'min'
    }
    return min
}

 