// 面试题50. 第一个只出现一次的字符
//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
//示例 1:
//
//输入：s = "abaccdeff"
//输出：'b'
//示例 2:
//
//输入：s = ""
//输出：' '
//
//
//限制：
//
//0 <= s 的长度 <= 50000
//
//通过次数204,887提交次数329,653
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/9 上午9:51
package di_yi_ge_zhi_chu_xian_yi_ci_de_zi_fu_lcof

func firstUniqChar(s string) byte {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	for i := range s {
		if count[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
