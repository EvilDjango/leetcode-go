// 剑指 Offer II 096. 字符串交织
//给定三个字符串 s1、s2、s3，请判断 s3 能不能由 s1 和 s2 交织（交错） 组成。
//
//两个字符串 s 和 t 交织 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
//s = s1 + s2 + ... + sn
//t = t1 + t2 + ... + tm
//|n - m| <= 1
//交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//提示：a + b 意味着字符串 a 和 b 连接。
//
//
//
//示例 1：
//
//
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
//示例 2：
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false
//示例 3：
//
//输入：s1 = "", s2 = "", s3 = ""
//输出：true
//
//
//提示：
//
//0 <= s1.length, s2.length <= 100
//0 <= s3.length <= 200
//s1、s2、和 s3 都由小写英文字母组成
//
//
//注意：本题与主站 97 题相同： https://leetcode-cn.com/problems/interleaving-string/
//
//通过次数2,283提交次数4,644
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/20/21 7:39 PM
package IY6buf

// 记忆化递归
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	cache := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		cache[i] = make([]int, l2+1)
	}
	cache[l1][l2] = 1
	// s1[i:]，s2[j:]能否交织组成s3[i+j:]
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if cache[i][j] > 0 {
			return cache[i][j] == 1
		}
		cache[i][j] = 2
		k := i + j
		if i < l1 && s1[i] == s3[k] && dfs(i+1, j) {
			cache[i][j] = 1
		} else if j < l2 && s2[j] == s3[k] && dfs(i, j+1) {
			cache[i][j] = 1
		}
		return cache[i][j] == 1
	}
	return dfs(0, 0)
}

// 动态规划
func isInterleave2(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	// dp[i][j]表示s1[:i],s2[:j]能否交织组成s3[:i+j]
	// 这里为了节省空间使用滚动数组，省去了第一层数组
	dp := make([]bool, l2+1)

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			dp[j] = i+j == 0 || (i > 0 && dp[j] && s1[i-1] == s3[i+j-1]) || (j > 0 && dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l2]
}
