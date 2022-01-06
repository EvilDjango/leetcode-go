// 剑指 Offer II 038. 每日温度
//请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
//
//示例 1:
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//示例 2:
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//示例 3:
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//
//提示：
//
//1 <= temperatures.length <= 105
//30 <= temperatures[i] <= 100
//
//
//注意：本题与主站 739 题相同： https://leetcode-cn.com/problems/daily-temperatures/
//
//通过次数7,162提交次数9,209
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/5/22 7:15 PM
package iIQa4I

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	ans := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			index := stack[len(stack)-1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
