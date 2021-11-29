// 面试题 05.08. 绘制直线
//绘制直线。有个单色屏幕存储在一个一维数组中，使得32个连续像素可以存放在一个 int 里。屏幕宽度为w，且w可被32整除（即一个 int 不会分布在两行上），屏幕高度可由数组长度及屏幕宽度推算得出。请实现一个函数，绘制从点(x1, y)到点(x2, y)的水平线。
//
//给出数组的长度 length，宽度 w（以比特为单位）、直线开始位置 x1（比特为单位）、直线结束位置 x2（比特为单位）、直线所在行数 y。返回绘制过后的数组。
//
//示例1:
//
// 输入：length = 1, w = 32, x1 = 30, x2 = 31, y = 0
// 输出：[3]
// 说明：在第0行的第30位到第31为画一条直线，屏幕表示为[0b000000000000000000000000000000011]
//示例2:
//
// 输入：length = 3, w = 96, x1 = 0, x2 = 95, y = 0
// 输出：[-1, -1, -1]
//通过次数3,782提交次数6,837
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/29/21 3:04 PM
package draw_line_lcci

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	matrix := make([]int, length)
	// 列数
	cols := w / 32
	// 第y行的起点
	start := y * cols
	for i := (x1 - 1) / 32; i <= (x2-1)/32; i++ {
		pixel := int32(0)
		for col := 0; col < 32; col++ {
			x := i*32 + col
			if x >= x1 && x <= x2 {
				pixel |= 1 << (31 - col)
			}
		}
		matrix[start+i] = int(pixel)
	}
	return matrix
}

// 一种巧妙的解法，合并多次位移操作
func drawLine2(length int, w int, x1 int, x2 int, y int) []int {
	matrix := make([]int, length)
	low := (w*y + x1) / 32
	high := (w*y + x2) / 32
	for i := low + 1; i < high; i++ {
		matrix[i] = -1
	}
	ones := uint32(1<<32 - 1)
	lowValue := ones >> (x1 % 32)
	highValue := ones << (31 - x2%32)
	if low == high {
		highValue = lowValue & highValue
	}
	matrix[low] = int(int32(lowValue))
	matrix[high] = int(int32(highValue))
	return matrix
}
