// 面试题 05.07. 配对交换
//配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。
//
//示例1:
//
// 输入：num = 2（或者0b10）
// 输出 1 (或者 0b01)
//示例2:
//
// 输入：num = 3
// 输出：3
//提示:
//
//num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。
//通过次数12,037提交次数17,059
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/29/21 4:54 PM
package exchange_lcci

func exchangeBits(num int) int {
	evenMask := 1
	oddMask := 1 << 1
	ans := 0
	for i := 0; i < 15; i++ {
		ans |= (num & evenMask) << 1
		ans |= (num & oddMask) >> 1
		evenMask <<= 2
		oddMask <<= 2
	}
	return ans
}

func exchangeBits2(num int) int {
	even := num & 0x55555555
	odd := num & 0xaaaaaaaa
	return (even << 1) | (odd >> 1)
}
