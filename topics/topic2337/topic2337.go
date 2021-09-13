// 面试题 17.21. 直方图的水量
//给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。
//
//
//
//上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 感谢 Marcos 贡献此图。
//
//示例:
//
//输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
//通过次数43,965提交次数69,136
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/13/21 9:02 PM
package topic2337

// 单调栈
func trap(height []int) int {
	var stack []int
	sum := 0
	for i, h := range height {
		for len(stack) > 0 && h >= height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			curH := min(h, height[left]) - height[top]
			sum += width * curH
		}
		stack = append(stack, i)
	}
	return sum
}

// 动态规划
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += min(leftMax[i], rightMax[i]) - height[i]
	}
	return sum
}

// 双指针法
func trap3(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, lMax, r, rMax := 0, 0, n-1, 0
	sum := 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		if height[l] < height[r] {
			sum += lMax - height[l]
			l++
		} else {
			sum += rMax - height[r]
			r--
		}
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
