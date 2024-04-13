func trap(height []int) int {
	len := len(height)
	if len == 0 {
		return 0
	}

	l, r := 0, len-1
	maxL, maxR := height[l], height[r]
	res := 0

	for l < r {
		if maxL < maxR {
			l++
			maxL = max(maxL, height[l])
			res += maxL - height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			res += maxR - height[r]
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}