// 剑指 Offer II 004. 只出现一次的数字
//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
//示例 1：
//
//输入：nums = [2,2,3,2]
//输出：3
//示例 2：
//
//输入：nums = [0,1,0,1,0,1,100]
//输出：100
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//-231 <= nums[i] <= 231 - 1
//nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//
//
//注意：本题与主站 137 题相同：https://leetcode-cn.com/problems/single-number-ii/
//
//通过次数26,793提交次数37,810
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/30 上午11:11
package WGki4K

// 依次确定每一个二进制位。对于第i位，由于其他数字都出现了三次，那么全部数字第i位的和除以3的余数就是答案的第i位
func singleNumber(nums []int) int {
	ans := int32(0)
	for bit := 0; bit < 32; bit++ {
		sum := int32(0)
		for _, num := range nums {
			sum += (int32(num) >> bit) & 1
		}
		ans |= (sum % 3) << bit
	}
	return int(ans)
}
