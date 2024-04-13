func maximalRectangle(matrix [][]byte) int {
      if len(matrix) == 0 {
        return 0
    }
    maxArea := 0
    // Initialize a slice to store the heights of rectangles
    heights := make([]int, len(matrix[0]))
    // Iterate over each row
    for _, row := range matrix {
        // Update the heights slice based on the current row
        for j, val := range row {
            if val == '0' {
                heights[j] = 0
            } else {
                heights[j]++
            }
        }
        // Calculate the maximal rectangle area using the updated heights
        maxArea = max(maxArea, largestRectangleArea(heights))
    }
    return maxArea
}

// Function to calculate the largest rectangle area in a histogram
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    stack := make([]int, 0)
    maxArea := 0
    for i, h := range heights {
        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            // Pop the stack and calculate the area for each popped element
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            height := heights[idx]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            area := height * width
            maxArea = max(maxArea, area)
        }
        stack = append(stack, i)
    }
    // Calculate the remaining areas for the elements left in the stack
    for len(stack) > 0 {
        idx := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        height := heights[idx]
        width := len(heights)
        if len(stack) > 0 {
            width = len(heights) - stack[len(stack)-1] - 1
        }
        area := height * width
        maxArea = max(maxArea, area)
    }
    return maxArea
}

// Function to find the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

 