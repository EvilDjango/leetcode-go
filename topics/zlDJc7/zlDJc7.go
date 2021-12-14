// 剑指 Offer II 109. 开密码锁
//一个密码锁由 4 个环形拨轮组成，每个拨轮都有 10 个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
//锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
//列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
//字符串 target 代表可以解锁的数字，请给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
//
//
//
//示例 1:
//
//输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//输出：6
//解释：
//可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
//注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，因为当拨动到 "0102" 时这个锁就会被锁定。
//示例 2:
//
//输入: deadends = ["8888"], target = "0009"
//输出：1
//解释：
//把最后一位反向旋转一次即可 "0000" -> "0009"。
//示例 3:
//
//输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
//输出：-1
//解释：
//无法旋转到目标数字且不被锁定。
//示例 4:
//
//输入: deadends = ["0000"], target = "8888"
//输出：-1
//
//
//提示：
//
//1 <= deadends.length <= 500
//deadends[i].length == 4
//target.length == 4
//target 不在 deadends 之中
//target 和 deadends[i] 仅由若干位数字组成
//
//
//注意：本题与主站 752 题相同： https://leetcode-cn.com/problems/open-the-lock/
//
//通过次数2,052提交次数3,448
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/14/21 6:38 PM
package zlDJc7

// 广度优先遍历
func openLock(deadends []string, target string) int {
	dead := map[string]bool{}
	for _, v := range deadends {
		dead[v] = true
	}
	start := "0000"
	if start == target {
		return 0
	}
	if dead[start] || dead[target] {
		return -1
	}

	type Pair struct {
		status string
		steps  int
	}
	q := []Pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		pair := q[0]
		q = q[1:]
		for _, status := range next(pair.status) {
			if status == target {
				return pair.steps + 1
			}
			if !seen[status] && !dead[status] {
				seen[status] = true
				q = append(q, Pair{status, pair.steps + 1})
			}
		}
	}
	return -1
}

// 深度优先遍历
//todo 这种做法虽然效率不高，但理论上是可行的。可是目前还无法ac，后面有时间改改。
func openLock2(deadends []string, target string) int {
	cache := map[string]int{target: 0}
	for _, v := range deadends {
		cache[v] = -1
	}
	start := "0000"
	if start == target {
		return 0
	}
	if cache[start] == -1 {
		return -1
	}
	seen := map[string]bool{}
	var dfs func(s string) int
	dfs = func(s string) int {
		if i, ok := cache[s]; ok {
			return i
		}
		seen[s] = true
		minSteps := -1
		for _, status := range next(s) {
			if seen[status] {
				continue
			}
			nextSteps := dfs(status)
			if nextSteps != -1 && (minSteps == -1 || nextSteps+1 < minSteps) {
				minSteps = nextSteps + 1
			}
		}
		delete(seen, s)
		cache[s] = minSteps
		return minSteps
	}
	return dfs(start)
}

// 双向广度遍历
func openLock3(deadends []string, target string) int {
	start := "0000"
	if start == target {
		return 0
	}
	visited := map[string]bool{}
	for _, s := range deadends {
		visited[s] = true
	}
	if visited[start] {
		return -1
	}
	visited[start] = true
	visited[target] = true
	forward := map[string]bool{start: true}
	backward := map[string]bool{target: true}
	length := 0
	for len(forward) > 0 && len(backward) > 0 {
		if len(forward) > len(backward) {
			forward, backward = backward, forward
		}
		length++
		nextForward := map[string]bool{}
		for s := range forward {
			for _, status := range next(s) {
				if backward[status] {
					return length
				}
				if visited[status] {
					continue
				}
				nextForward[status] = true
				visited[status] = true
			}
		}
		forward = nextForward
	}
	return -1
}

func next(s string) []string {
	ret := make([]string, 0, 8)
	for i := 0; i < 4; i++ {
		up := []byte(s)
		up[i]++
		if up[i] > '9' {
			up[i] = '0'
		}
		ret = append(ret, string(up))

		down := []byte(s)
		down[i]--
		if down[i] < '0' {
			down[i] = '9'
		}
		ret = append(ret, string(down))
	}
	return ret
}
