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

// 20220723重做此题，题目描述变了，条件得到了增强，seqs中的全部序列都是nums的子序列。所以只要能得到一个唯一的拓扑排序，那么这个拓扑排序必然与
//nums完全相等，就不用再一一比对元素了。
//新的题目描述如下：
//给定一个长度为 n 的整数数组 nums ，其中 nums 是范围为 [1，n] 的整数的排列。还提供了一个 2D 整数数组 sequences ，其中 sequences[i] 是 nums 的子序列。
//检查 nums 是否是唯一的最短 超序列 。最短 超序列 是 长度最短 的序列，并且所有序列 sequences[i] 都是它的子序列。对于给定的数组 sequences ，可能存在多个有效的 超序列 。
//
//例如，对于 sequences = [[1,2],[1,3]] ，有两个最短的 超序列 ，[1,2,3] 和 [1,3,2] 。
//而对于 sequences = [[1,2],[1,3],[1,2,3]] ，唯一可能的最短 超序列 是 [1,2,3] 。[1,2,3,4] 是可能的超序列，但不是最短的。
//如果 nums 是序列的唯一最短 超序列 ，则返回 true ，否则返回 false 。
//子序列 是一个可以通过从另一个序列中删除一些元素或不删除任何元素，而不改变其余元素的顺序的序列。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3], sequences = [[1,2],[1,3]]
//输出：false
//解释：有两种可能的超序列：[1,2,3]和[1,3,2]。
//序列 [1,2] 是[1,2,3]和[1,3,2]的子序列。
//序列 [1,3] 是[1,2,3]和[1,3,2]的子序列。
//因为 nums 不是唯一最短的超序列，所以返回false。
//示例 2：
//
//输入：nums = [1,2,3], sequences = [[1,2]]
//输出：false
//解释：最短可能的超序列为 [1,2]。
//序列 [1,2] 是它的子序列：[1,2]。
//因为 nums 不是最短的超序列，所以返回false。
//示例 3：
//
//输入：nums = [1,2,3], sequences = [[1,2],[1,3],[2,3]]
//输出：true
//解释：最短可能的超序列为[1,2,3]。
//序列 [1,2] 是它的一个子序列：[1,2,3]。
//序列 [1,3] 是它的一个子序列：[1,2,3]。
//序列 [2,3] 是它的一个子序列：[1,2,3]。
//因为 nums 是唯一最短的超序列，所以返回true。
//
//
//提示：
//
//n == nums.length
//1 <= n <= 104
//nums 是 [1, n] 范围内所有整数的排列
//1 <= sequences.length <= 104
//1 <= sequences[i].length <= 104
//1 <= sum(sequences[i].length) <= 105
//1 <= sequences[i][j] <= n
//sequences 的所有数组都是 唯一 的
//sequences[i] 是 nums 的一个子序列
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ur2n8P
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sequenceReconstruction2(org []int, seqs [][]int) bool {
	n := len(org)
	edges := make([][]int, n+1)
	inDegrees := make([]int, n+1)
	for _, seq := range seqs {
		for i := 1; i < len(seq); i++ {
			u, v := seq[i-1], seq[i]
			edges[u] = append(edges[u], v)
			inDegrees[v]++
		}
	}
	var q []int
	for i := 1; i <= n; i++ {
		if inDegrees[i] == 0 {
			q = append(q, i)
			if len(q) > 1 {
				return false
			}
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range edges[cur] {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				q = append(q, v)
				if len(q) > 1 {
					return false
				}
			}
		}
	}
	return true
}
