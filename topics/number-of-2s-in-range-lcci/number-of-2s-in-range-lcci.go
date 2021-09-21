// 面试题 17.06. 2出现的次数
//编写一个方法，计算从 0 到 n (含 n) 中数字 2 出现的次数。
//
//示例:
//
//输入: 25
//输出: 9
//解释: (2, 12, 20, 21, 22, 23, 24, 25)(注意 22 应该算作两次)
//提示：
//
//n <= 10^9
//通过次数4,580提交次数10,197
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/20/21 8:08 PM
package number_of_2s_in_range_lcci

func numberOf2sInRange(n int) int {
	count := 0
	base := 1
	for i := 0; n > base; i++ {
		num := n / base
		if num%10 == 2 {
			count += n%(base*10) - 2*base + 1
		} else if num%10 > 2 {
			count += base
		}
		count += num / 10 * base
		base *= 10
	}
	return count
}
