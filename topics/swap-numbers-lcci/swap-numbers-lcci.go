// 面试题 16.01. 交换数字
//编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
//
//示例：
//
//输入: numbers = [1,2]
//输出: [2,1]
//提示：
//
//numbers.length == 2
//-2147483647 <= numbers[i] <= 2147483647
//通过次数26,080提交次数31,437
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/6/21 5:54 PM
package swap_numbers_lcci

// 利用加减法
func swapNumbers(numbers []int) []int {
	numbers[0] += numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] -= numbers[1]
	return numbers
}

// 利用异或
func swapNumbers2(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}
