// 面试题 16.10. 生存人数
//给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。
//
//你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。如果一个人在某一年的任意时期处于生存状态，那么他应该被纳入那一年的统计中。例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。
//
//如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。
//
//
//
//示例：
//
//输入：
//birth = {1900, 1901, 1950}
//death = {1948, 1951, 2000}
//输出： 1901
//
//
//提示：
//
//0 < birth.length == death.length <= 10000
//birth[i] <= death[i]
//通过次数10,251提交次数15,023
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/5/21 11:31 AM
package living_people_lcci

func maxAliveYear(birth []int, death []int) int {
	counts := make([]int, 101)
	for i := 0; i < len(birth); i++ {
		for j := birth[i]; j <= death[i]; j++ {
			counts[j-1900]++
		}
	}
	max := 0
	for i := 1; i < 101; i++ {
		if counts[i] > counts[max] {
			max = i
		}
	}
	return max + 1900
}

// 只记录每年的增量
func maxAliveYear2(birth []int, death []int) int {
	delta := make([]int, 102)
	for i := 0; i < len(birth); i++ {
		delta[birth[i]-1900]++
		delta[death[i]-1900+1]--
	}
	sum := delta[0]
	max := delta[0]
	maxYear := 1900
	for i := 1; i < 101; i++ {
		sum += delta[i]
		if sum > max {
			max = sum
			maxYear = i + 1900
		}
	}
	return maxYear
}
