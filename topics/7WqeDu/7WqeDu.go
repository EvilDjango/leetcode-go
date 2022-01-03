// 剑指 Offer II 057. 值和下标之差都在给定的范围内
//给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
//
//如果存在则返回 true，不存在返回 false。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,1], k = 3, t = 0
//输出：true
//示例 2：
//
//输入：nums = [1,0,1,1], k = 1, t = 2
//输出：true
//示例 3：
//
//输入：nums = [1,5,9,1,5,9], k = 2, t = 3
//输出：false
//
//
//提示：
//
//0 <= nums.length <= 2 * 104
//-231 <= nums[i] <= 231 - 1
//0 <= k <= 104
//0 <= t <= 231 - 1
//
//
//注意：本题与主站 220 题相同： https://leetcode-cn.com/problems/contains-duplicate-iii/
//
//通过次数2,927提交次数8,228
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 7:37 PM
package _WqeDu

// 笨办法
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var intervals [][2]int
	for i, num := range nums {
		interval := [2]int{num, num + t}
		// 移除最左侧元素
		if len(intervals) == k+1 {
			start := [2]int{nums[i-k-1], nums[i-k-1] + t}
			l, r := 0, k+1
			for l < r {
				mid := (r-l)>>1 + l
				if intervals[mid][0] < start[0] {
					l = mid + 1
				} else if intervals[mid][0] > start[0] {
					r = mid
				} else {
					l, r = mid, mid
				}
			}
			for j := l; j < k; j++ {
				intervals[j] = intervals[j+1]
			}
			intervals = intervals[:k]
		}
		// 加入新元素
		l, r := 0, len(intervals)
		for l < r {
			mid := (r-l)>>1 + l
			if intervals[mid][0] <= interval[1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		// 有重合区间
		if l-1 >= 0 && intervals[l-1][1] >= interval[0] {
			return true
		}
		intervals = append(intervals, [2]int{0, 0})
		for j := len(intervals) - 1; j > l; j-- {
			intervals[j] = intervals[j-1]
		}
		intervals[l] = interval
	}
	return false
}

// 笨办法2
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0 && j >= i-k; j-- {
			if nums[j] >= nums[i]-t && nums[j] <= nums[i]+t {
				return true
			}
		}
	}
	return false
}

// 桶
func containsNearbyAlmostDuplicate3(nums []int, k int, t int) bool {
	buckets := map[int]int{}
	for i, num := range nums {
		if len(buckets) > k {
			delete(buckets, getId(nums[i-k-1], t+1))
		}
		id := getId(num, t+1)
		if _, ok := buckets[id]; ok {
			return true
		}
		if i, ok := buckets[id-1]; ok && abs(i-num) <= t {
			return true
		}
		if i, ok := buckets[id+1]; ok && abs(i-num) <= t {
			return true
		}
		buckets[id] = num
	}
	return false
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
func getId(x, w int) int {
	if x > 0 {
		return x / w
	}
	return x/w - 1
}
