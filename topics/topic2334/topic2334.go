// 面试题 17.18. 最短超串
//假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有的元素的最短子数组，其出现顺序无关紧要。
//
//返回最短子数组的左端点和右端点，如有多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。
//
//示例 1:
//
//输入:
//big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7]
//small = [1,5,9]
//输出: [7,10]
//示例 2:
//
//输入:
//big = [1,2,3]
//small = [4]
//输出: []
//提示：
//
//big.length <= 100000
//1 <= small.length <= 100000
//通过次数6,999提交次数15,692
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/14/21 2:33 PM
package topic2334

import "math"

// 深度搜索，会超时
func shortestSeq(big []int, small []int) []int {
	targets := make(map[int]bool)
	for _, num := range small {
		targets[num] = true
	}
	left, right := 0, math.MaxInt32
	var dfs func(j int, path []int)
	dfs = func(j int, path []int) {
		size := len(path)
		if size > 0 {
			l, r := path[0], path[size-1]
			if r-l >= right-left {
				return
			}
			if len(targets) == 0 {
				if right-left > r-l {
					left, right = l, r
				}
				return
			}
		}

		for k := j; k < len(big); k++ {
			num := big[k]
			if targets[num] {
				path = append(path, k)
				targets[num] = false
				dfs(k+1, path)
				path = path[:len(path)-1]
				targets[num] = true
			}
		}

	}
	dfs(0, []int{})
	if right != math.MaxInt32 {
		return []int{left, right}
	}
	return nil
}

// 滑动窗口
func shortestSeq2(big []int, small []int) []int {
	m, n := len(big), len(small)
	record := make(map[int]int, n)
	for _, num := range small {
		record[num] = 0
	}
	l, r := 0, 0
	left, right := 0, math.MaxInt32
	count := 0
	for r < m {
		for count < n && r < m {
			num := big[r]
			if cnt, ok := record[num]; ok {
				if cnt == 0 {
					count++
				}
				record[num]++
			}
			r++
		}
		for count == n && l < r {
			if r-l < right-left {
				left, right = l, r
			}
			num := big[l]
			if _, ok := record[num]; ok {
				record[num]--
				if record[num] == 0 {
					count--
				}
			}
			l++
		}
	}
	if right != math.MaxInt32 {
		return []int{left, right - 1}
	}
	return nil
}
