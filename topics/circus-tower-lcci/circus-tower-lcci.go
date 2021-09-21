// 面试题 17.08. 马戏团人塔
//有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。
//
//示例：
//
//输入：height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
//输出：6
//解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
//提示：
//
//height.length == weight.length <= 10000
//通过次数8,162提交次数29,995
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/19/21 3:59 PM
package circus_tower_lcci

import (
	"sort"
)

// 动态规划,会超时
func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	t := &TwoInts{height, weight}
	sort.Sort(t)
	dp := make([]int, len(weight))
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if height[i] > height[j] && weight[i] > weight[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 贪心+二分查找,身高相同时对体重降序排列
func bestSeqAtIndex2(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	t := &TwoInts{height, weight}
	sort.Sort(t)
	// minEnd[i]表示长度为i的子序列末尾元素的最小值
	minEnd := make([]int, n+1)
	maxLen := 0
	for _, w := range weight {
		l, r := 1, maxLen
		for l <= r {
			mid := (r-l)/2 + l
			if minEnd[mid] >= w {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		minEnd[l] = w
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

// 贪心+二分查找，身高相同时体重无需排序，在二分查找时，要考虑身高
// todo 这里情况很复杂，目前还无法AC
func bestSeqAtIndex3(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	t := &TwoInts{height, weight}
	sort.Sort(t)
	// minEnd[i]表示长度为i的子序列末尾元素的最小值
	minEnd := make([][]int, n+1)
	maxLen := 0
	for i := 0; i < n; i++ {
		l, r := 1, maxLen
		for l <= r {
			mid := (r-l)/2 + l
			if minEnd[mid][1] >= weight[i] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if l > 1 && height[i] == minEnd[l-1][0] {
			continue
		}
		minEnd[l] = []int{height[i], weight[i]}
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

type TwoInts struct {
	one, two []int
}

func (t *TwoInts) Len() int {
	return len(t.one)
}

func (t *TwoInts) Less(i, j int) bool {
	//if t.one[i] == t.one[j] {
	//	return t.two[i] > t.two[j]
	//}
	return t.one[i] < t.one[j]
}

func (t *TwoInts) Swap(i, j int) {
	t.one[i], t.one[j] = t.one[j], t.one[i]
	t.two[i], t.two[j] = t.two[j], t.two[i]
}
