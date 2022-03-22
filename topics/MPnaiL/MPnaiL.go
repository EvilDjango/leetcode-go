// 剑指 Offer II 014. 字符串中的变位词
//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
//
//
//
//示例 1：
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//示例 2：
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//
//提示：
//
//1 <= s1.length, s2.length <= 104
//s1 和 s2 仅包含小写字母
//
//
//注意：本题与主站 567 题相同： https://leetcode-cn.com/problems/permutation-in-string/
//
//通过次数14,779提交次数28,780
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/22 上午11:32
package MPnaiL

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	nums := [26]int{}
	for _, r := range s1 {
		nums[r-'a']--
	}
	left := 0
	for right, r := range s2 {
		nums[(r-'a')]++
		for nums[(r-'a')] > 0 {
			nums[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}
	return false
}
