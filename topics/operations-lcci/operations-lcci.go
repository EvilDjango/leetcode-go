// 面试题 16.09. 运算
//请实现整数数字的乘法、减法和除法运算，运算结果均为整数数字，程序中只允许使用加法运算符和逻辑运算符，允许程序中出现正负常数，不允许使用位运算。
//
//你的实现应该支持如下操作：
//
//Operations() 构造函数
//minus(a, b) 减法，返回a - b
//multiply(a, b) 乘法，返回a * b
//divide(a, b) 除法，返回a / b
//示例：
//
//Operations operations = new Operations();
//operations.minus(1, 2); //返回-1
//operations.multiply(3, 4); //返回12
//operations.divide(5, -2); //返回-2
//提示：
//
//你可以假设函数输入一定是有效的，例如不会出现除法分母为0的情况
//单个用例的函数调用次数不会超过1000次
//通过次数2,173提交次数3,965
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/14/21 7:36 AM
package operations_lcci

import "math"

type Operations struct {
	negatives, positives [31]int
}

func Constructor() Operations {
	negatives, positives := [31]int{}, [31]int{}
	neg, pos := -1, 1
	for i := 0; i < 31; i++ {
		negatives[i] = neg
		neg += neg
		positives[i] = pos
		pos += pos
	}
	return Operations{negatives, positives}
}

// 取相反数
func (this *Operations) opposite(a int) int {
	if a == 0 {
		return 0
	}
	result := 0
	if a > 0 {
		for i := 30; i >= 0; i-- {
			if a+this.negatives[i] < 0 {
				continue
			}
			a += this.negatives[i]
			result += this.negatives[i]
		}
	} else {
		for i := 30; i >= 0; i-- {
			if a+this.positives[i] > 0 {
				continue
			}
			a += this.positives[i]
			result += this.positives[i]
		}
	}
	return result
}

func (this *Operations) Minus(a int, b int) int {
	return a + this.opposite(b)
}

func (this *Operations) Multiply(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		a, b = this.opposite(a), this.opposite(b)
	}
	result, times := b, 1
	// 防止溢出
	for times+times > 0 && times+times < a {
		result += result
		times += times
	}
	return result + this.Multiply(this.Minus(a, times), b)
}

func (this *Operations) Divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if a == b {
		return 1
	}
	if b == 1 {
		return a
	}
	if b == -1 {
		return this.opposite(a)
	}
	// 只处理同号的情况，异号时转化为同号处理
	if a > 0 {
		if b == math.MinInt32 {
			return 0
		}
		if b < 0 {
			return this.opposite(this.Divide(a, this.opposite(b)))
		}
		if a < b {
			return 0
		}
		result, sum := 1, b
		for sum+sum > 0 && sum+sum < a {
			result += result
			sum += sum
		}
		return result + this.Divide(this.Minus(a, sum), b)
	} else {
		if b > 0 {
			return this.opposite(this.Divide(a, this.opposite(b)))
		}
		if b < a {
			return 0
		}
		result, sum := 1, b
		for sum+sum < 0 && sum+sum > a {
			result += result
			sum += sum
		}
		return result + this.Divide(this.Minus(a, sum), b)
	}
}
