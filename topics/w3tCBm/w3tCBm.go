// 剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
//给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
//
//
//
//示例 1:
//
//输入: n = 2
//输出: [0,1,1]
//解释:
//0 --> 0
//1 --> 1
//2 --> 10
//示例 2:
//
//输入: n = 5
//输出: [0,1,1,2,1,2]
//解释:
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
//
//
//说明 :
//
//0 <= n <= 105
//
//
//进阶:
//
//给出时间复杂度为 O(n*sizeof(integer)) 的解答非常容易。但你可以在线性时间 O(n) 内用一趟扫描做到吗？
//要求算法的空间复杂度为 O(n) 。
//你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount ）来执行此操作。
//
//
//注意：本题与主站 338 题相同：https://leetcode-cn.com/problems/counting-bits/
//
//通过次数31,809提交次数40,504
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/30 下午12:08
package w3tCBm

func countBits(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		decrease, mask := 0, 1
		for (i-1)&mask > 0 {
			decrease++
			mask <<= 1
		}
		ans[i] = ans[i-1] + 1 - decrease
	}
	return ans
}

//Brian Kernighan 算法
func countBits2(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = countOnes(i)
	}
	return ans
}

func countOnes(i int) int {
	cnt := 0
	for i != 0 {
		cnt++
		i = i & (i - 1)
	}
	return cnt
}

// 动态规划，最低设置位，定义正整数 x 的「最低设置位」为 xx 的二进制表示中的最低的 11 所在位
func countBits3(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
