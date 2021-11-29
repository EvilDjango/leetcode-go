// 面试题 05.04. 下一个数
//下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。
//
//示例1:
//
// 输入：num = 2（或者0b10）
// 输出：[4, 1] 或者（[0b100, 0b1]）
//示例2:
//
// 输入：num = 1
// 输出：[2, -1]
//提示:
//
//num的范围在[1, 2147483647]之间；
//如果找不到前一个或者后一个满足条件的正数，那么输出 -1。
//通过次数5,594提交次数14,638
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/29/21 9:23 PM
package closed_number_lcci

import "math/bits"

// 无他，就是酷炫的位运算而已
// 如何得到最接近且更大的数？从低到高（从右到左）找到第一个01，交换位置得到10，然后把右侧的1全部右移到最右端（这些1都是连续的，所以我们可以整体移位）
// 如何得到最接近且更小的数？从低到高（从右到左）找到第一个10，交换位置得到01，然后把右侧的1全部左移到紧邻01的位置（同样的，这些1都是连续的，所以我们可以整体移位）
func findClosedNumbers(num int) []int {
	more := num
	// 忽略尾部连续的0
	mask := (num & -num) << 1
	// 统计尾部连续0的个数，为后面移位做准备
	trailingZeroes := bits.OnesCount(uint(num&-num - 1))
	for mask < 1<<30 && mask&num > 0 {
		mask <<= 1
	}
	// 没有找到01
	if mask == 1<<30 {
		more = -1
	} else {
		// 找到了01
		// 将0改为1
		more |= mask
		mask >>= 1
		// 将1改为0
		more ^= mask
		// 取低位
		low := num & (mask - 1)
		// 将more的低位全部置为0
		more ^= low
		// 将低位右移，使低位的1全部移动到最右端，然后把低位加到more上
		more |= low >> trailingZeroes
	}

	mask = 1
	less := num
	// 在[0,num)范围内查找0
	for mask < num && mask&num > 0 {
		mask <<= 1
	}
	//没有找到0
	if mask >= num {
		less = -1
	} else {
		mask <<= 1
		zeroes := 1
		// 现在往左寻找1，同时统计0的数量，为后面的移位操作做准备
		for num&mask == 0 {
			mask <<= 1
			zeroes++
		}
		// 找到了10
		// 把1改为0
		less ^= mask
		mask >>= 1
		// 把0改为1
		less |= mask
		// 取10右侧的低位
		low := num & (mask - 1)
		// 在less中将低位全部置为0
		less ^= low
		// 将低位左移，是低位的1紧邻mask，然后把低位加到less上
		less |= low << (zeroes - 1)
	}

	return []int{more, less}
}
