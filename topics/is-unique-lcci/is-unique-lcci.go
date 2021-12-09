// 面试题 01.01. 判定字符是否唯一
//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
//
//示例 1：
//
//输入: s = "leetcode"
//输出: false
//示例 2：
//
//输入: s = "abc"
//输出: true
//限制：
//
//0 <= len(s) <= 100
//如果你不使用额外的数据结构，会很加分。
//通过次数97,468提交次数135,930
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 11:31 PM
package is_unique_lcci

func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		for j := 0; j < i; j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}
