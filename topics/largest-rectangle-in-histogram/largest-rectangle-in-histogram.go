// 84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//示例 1:
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//示例 2：
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//提示：
//
//1 <= heights.length <=105
//0 <= heights[i] <= 104
//通过次数202,329提交次数465,277
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/4/22 2:55 PM
package largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	var stack []int
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	for i, h := range heights {
		area := (right[i] - left[i] - 1) * h
		if area > ans {
			ans = area
		}
	}
	return ans
}

// 一种更简洁的解法
func largestRectangleArea2(heights []int) int {
	ans := 0
	stack := []int{-1}
	for i, h := range heights {
		j := len(stack) - 1
		for ; j > 0 && heights[stack[j]] >= h; j-- {
			area := heights[stack[j]] * (i - stack[j-1] - 1)
			ans = max(ans, area)
		}
		stack = stack[:j+1]
		stack = append(stack, i)
	}
	// 对于还在栈里的柱子，说明他们右边没有更低的柱子，那么其右边界就是len(heights)
	for i := len(stack) - 1; i > 0; i-- {
		area := heights[stack[i]] * (len(heights) - stack[i-1] - 1)
		ans = max(ans, area)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
