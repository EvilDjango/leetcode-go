// 剑指 Offer II 067. 最大的异或
//给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
//
//
//
//示例 1：
//
//输入：nums = [3,10,5,25,2,8]
//输出：28
//解释：最大运算结果是 5 XOR 25 = 28.
//示例 2：
//
//输入：nums = [0]
//输出：0
//示例 3：
//
//输入：nums = [2,4]
//输出：6
//示例 4：
//
//输入：nums = [8,10,2]
//输出：10
//示例 5：
//
//输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//输出：127
//
//
//提示：
//
//1 <= nums.length <= 2 * 104
//0 <= nums[i] <= 231 - 1
//
//
//进阶：你可以在 O(n) 的时间解决这个问题吗？
//
//
//
//注意：本题与主站 421 题相同： https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
//
//通过次数2,274提交次数3,370
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/29/21 12:42 PM
package ms70jA

// 暴力算法，会超时
func findMaximumXOR(nums []int) int {
	ans := 0
	for _, num := range nums {
		for _, i := range nums {
			if num^i > ans {
				ans = num ^ i
			}
		}
	}
	return ans
}

// 从高到低依次确认每一位，利用哈希表
func findMaximumXOR2(nums []int) int {
	ans := 0
	for i := 30; i >= 0; i-- {
		seen := map[int]bool{}
		for _, num := range nums {
			seen[num>>i] = true
		}
		next := ans<<1 + 1
		found := false
		for _, num := range nums {
			if seen[num>>i^next] {
				found = true
				break
			}
		}
		if !found {
			next -= 1
		}
		ans = next
	}
	return ans
}

type trie struct {
	// left,right分表表示0，1
	left, right *trie
}

func (t *trie) add(num int) {
	curr := t
	mask := 1 << 30
	for i := 0; i < 31; i++ {
		if num&mask == 0 {
			if curr.left == nil {
				curr.left = &trie{}
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = &trie{}
			}
			curr = curr.right
		}
		mask >>= 1
	}
}

func (t *trie) check(num int) int {
	max := 0
	mask := 1 << 30
	curr := t
	for i := 0; i < 31; i++ {
		max <<= 1
		if mask&num == 0 {
			if curr.right != nil {
				max += 1
				curr = curr.right
			} else {
				curr = curr.left
			}
		} else {
			if curr.left != nil {
				max += 1
				curr = curr.left
			} else {
				curr = curr.right
			}
		}
		mask >>= 1
	}
	return max
}

// 前缀树
func findMaximumXOR3(nums []int) int {
	ans := 0
	root := &trie{}
	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		ans = max(ans, root.check(nums[i]))
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
