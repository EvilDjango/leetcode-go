// 剑指 Offer II 115. 重建序列
//请判断原始的序列 org 是否可以从序列集 seqs 中唯一地 重建 。
//
//序列 org 是 1 到 n 整数的排列，其中 1 ≤ n ≤ 104。重建 是指在序列集 seqs 中构建最短的公共超序列，即  seqs 中的任意序列都是该最短序列的子序列。
//
//
//
//示例 1：
//
//输入: org = [1,2,3], seqs = [[1,2],[1,3]]
//输出: false
//解释：[1,2,3] 不是可以被重建的唯一的序列，因为 [1,3,2] 也是一个合法的序列。
//示例 2：
//
//输入: org = [1,2,3], seqs = [[1,2]]
//输出: false
//解释：可以重建的序列只有 [1,2]。
//示例 3：
//
//输入: org = [1,2,3], seqs = [[1,2],[1,3],[2,3]]
//输出: true
//解释：序列 [1,2], [1,3] 和 [2,3] 可以被唯一地重建为原始的序列 [1,2,3]。
//示例 4：
//
//输入: org = [4,1,5,2,6,3], seqs = [[5,2,6,3],[4,1,5,2]]
//输出: true
//
//
//提示：
//
//1 <= n <= 104
//org 是数字 1 到 n 的一个排列
//1 <= segs[i].length <= 105
//seqs[i][j] 是 32 位有符号整数
//
//
//注意：本题与主站 444 题相同：https://leetcode-cn.com/problems/sequence-reconstruction/
//
//通过次数1,277提交次数4,149
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/10/21 10:30 PM
package ur2n8P

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	edges := map[int][]int{}
	inDegree := make([]int, n+1)
	seen := make([]bool, n)
	nodes := 0
	for _, seq := range seqs {
		for i, v := range seq {
			if v < 1 || v > n {
				return false
			}
			if !seen[v] {
				nodes++
				seen[v] = true
			}
			if i > 0 {
				u := seq[i-1]
				edges[u] = append(edges[u], v)
				inDegree[v]++
			}
		}
	}
	if nodes != n {
		return false
	}
	zero := -1
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			if zero != -1 {
				return false
			}
			zero = i
		}
	}
	index := 0
	for zero != -1 {
		if org[index] != zero {
			return false
		}
		nextZero := -1
		for _, to := range edges[zero] {
			inDegree[to]--
			if inDegree[to] == 0 {
				if nextZero != -1 {
					return false
				}
				nextZero = to
			}
		}
		zero = nextZero
		index++
	}
	return index == n
}
