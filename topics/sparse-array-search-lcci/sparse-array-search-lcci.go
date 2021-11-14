// 面试题 10.05. 稀疏数组搜索
//稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
//
//示例1:
//
// 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
// 输出：-1
// 说明: 不存在返回-1。
//示例2:
//
// 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
// 输出：4
//提示:
//
//words的长度在[1, 1000000]之间
//通过次数20,074提交次数36,223
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/13/21 2:53 PM
package sparse_array_search_lcci

func findString(words []string, s string) int {
	l, r := 0, len(words)
	for l < r {
		mid := (r-l)/2 + l
		for mid > l && words[mid] == "" {
			mid--
		}
		if words[mid] > s {
			r = mid
		} else if words[mid] < s {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
