package integer

import "math"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1{
		return heights[0]
	}

	shortest, idx := math.MaxInt32, 0
	for i, v := range heights {
		if v < shortest {
			shortest = v
			idx = i
		}
	}

	return max(max(shortest*n, largestRectangleArea(heights[:idx])), largestRectangleArea(heights[idx+1:]))

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}