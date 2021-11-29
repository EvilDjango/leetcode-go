// 面试题 05.03. 翻转数位
//给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
//
//示例 1：
//
//输入: num = 1775(110111011112)
//输出: 8
//示例 2：
//
//输入: num = 7(01112)
//输出: 4
//通过次数11,586提交次数30,187
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/29/21 11:45 PM
package reverse_bits_lcci

func reverseBits(num int) int {
	// prefix[i]表示位置i前缀（左侧的）连续1的数量
	prefix := make([]int, 32)
	// suffix[i]表示位置i后缀（右侧的）连续1的数量
	suffix := make([]int, 32)
	for i := 1; i < 32; i++ {
		prev := num & (1 << (32 - i))
		if prev > 0 {
			prefix[i] = prefix[i-1] + 1
		}
	}
	for i := 30; i >= 0; i-- {
		next := num & (1 << (30 - i))
		if next > 0 {
			suffix[i] = suffix[i+1] + 1
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		if prefix[i]+suffix[i] > ans {
			ans = prefix[i] + suffix[i]
		}
	}
	return ans
}

// 一次遍历，每次遇到0时将这个0转换为1（这个转换只在脑子里进行），并减去上一个0之前的1的计数，那么我们通过一次遍历就能完全统计把任意位置的0转变为1之后能得到的最长连续1长度
func reverseBits2(num int) int {
	max := 0
	// 当前连续1的数量
	count := 0
	// 被我们转变为1的0后面的连续1数量，为-1表示当前还没有遇到0
	afterZero := -1
	for i := 0; i < 32; i++ {
		// 当前位为1
		if num&(1<<i) > 0 {
			count++
			if afterZero != -1 {
				afterZero++
			}
		} else {
			// 第一次遇到0，将其转换为1，1计数增加
			if afterZero == -1 {
				count++
			} else {
				// 后续遇到0，首先将前面的0恢复，然后将当前的0转变为1，当前长度改变
				count = afterZero + 1
			}
			afterZero = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
