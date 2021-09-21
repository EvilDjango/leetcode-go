// 面试题 17.05.  字母与数字
//给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
//
//返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。
//
//示例 1:
//
//输入: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
//
//输出: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
//示例 2:
//
//输入: ["A","A"]
//
//输出: []
//提示：
//
//array.length <= 100000
//通过次数5,341提交次数13,807
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/21/21 12:24 PM
package find_longest_subarray_lcci

import "unicode"

// 暴力算法, tle
func findLongestSubarray(array []string) []string {
	n := len(array)
	l, r := 0, -1
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += cal(array[j])
			if sum == 0 && j-i > r-l {
				l, r = i, j
			}
		}
	}
	return array[l : r+1]
}

func findLongestSubarray2(array []string) []string {
	n := len(array)
	prefix := make([]int, n+1)
	sum := 0
	for i := 0; i < n; i++ {
		sum += cal(array[i])
		prefix[i+1] = sum
	}
	l, r := 0, -1
	for i := 0; i < n; i++ {
		for j := r - l + 1 + i; j < n; j++ {
			if prefix[j+1]-prefix[i] == 0 && j-i > r-l {
				l, r = i, j
			}
		}
	}
	return array[l : r+1]
}

// 我们找到两个相同的前缀和，下标分别为i,j，那么我们可以知道 (i,j]区间内的总和为0，是符合要求的区间。
// 那么问题转化为找到距离最远的相同前缀和
func findLongestSubarray3(array []string) []string {
	n := len(array)
	if n == 0 {
		return nil
	}
	sum, l, r := 0, 1, 0
	// indexes[i]表示前缀和为i的第一个下标
	indexes := make([]int, n<<1+1)
	for i := 0; i < len(indexes); i++ {
		indexes[i] = -2
	}
	// 需要把0第一次出现的位置设置为-1，才不会漏掉从第一个元素开始的最长序列
	indexes[n] = -1
	for i := 0; i < n; i++ {
		sum += cal(array[i])
		if indexes[sum+n] == -2 {
			indexes[sum+n] = i
		} else if i-indexes[sum+n] > r-l+1 {
			l = indexes[sum+n] + 1
			r = i
		}
	}
	return array[l : r+1]
}

func cal(s string) int {
	if unicode.IsLetter(rune(s[0])) {
		return 1
	}
	return -1
}
