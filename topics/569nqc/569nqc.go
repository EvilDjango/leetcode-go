// 剑指 Offer II 035. 最小时间差
//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
//示例 1：
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//示例 2：
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//提示：
//
//2 <= timePoints <= 2 * 104
//timePoints[i] 格式为 "HH:MM"
//
//
//注意：本题与主站 539 题相同： https://leetcode-cn.com/problems/minimum-time-difference/
//
//通过次数5,278提交次数7,956
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/6/22 2:56 PM
package _69nqc

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 24*60 {
		return 0
	}
	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})
	ans := math.MaxInt64
	for i := 0; i < len(timePoints); i++ {
		diff := getDiff(timePoints[i], timePoints[(i+1)%len(timePoints)])
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

func getDiff(t1, t2 string) int {
	h1, _ := strconv.Atoi(t1[:2])
	h2, _ := strconv.Atoi(t2[:2])
	m1, _ := strconv.Atoi(t1[3:])
	m2, _ := strconv.Atoi(t2[3:])
	diff := (h1-h2)*60 + m1 - m2
	if diff < 0 {
		diff = -diff
	}
	complement := 60*24 - diff
	if complement < diff {
		diff = complement
	}
	return diff
}
