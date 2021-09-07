// Oops, forgot to write comments. Good luck, bro.
//累加数是一个字符串，组成它的数字可以形成累加序列。
//
//一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
//
//给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。
//
//说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
//
//示例 1:
//
//输入: "112358"
//输出: true
//解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
//示例 2:
//
//输入: "199100199"
//输出: true
//解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
//进阶:
//你如何处理一个溢出的过大的整数输入?
//
//通过次数16,207提交次数48,059
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/additive-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/1/21 1:08 PM
package topic306

import "strconv"

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	return find(num, 0, 0, 0, 0)
}

func find(num string, index, count, i, j int) bool {
	n := len(num)
	if index == n {
		return count >= 3
	}
	if count < 2 {
		for k := index + 1; k <= n; k++ {
			digit, _ := strconv.Atoi(num[index:k])
			if count == 0 && find(num, k, 1, digit, 0) || count == 1 && find(num, k, 2, i, digit) {
				return true
			}
			if num[index] == '0' {
				break
			}
		}
		return false
	}
	next := strconv.Itoa(i + j)
	length := len(next)
	if length > n-index || num[index:index+length] != next {
		return false
	}
	return find(num, index+length, count+1, j, i+j)
}
