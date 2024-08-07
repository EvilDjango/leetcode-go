// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/3 22:00
package minimum_factorization

import "math"

func smallestFactorization(num int) int {
	if num < 10 {
		return num
	}
	ans := 0
	base := 1
	for i := 9; i > 1 && num > 1; i-- {
		for num > 1 && num%i == 0 {
			if i > (math.MaxInt32-ans)/base {
				return 0
			}
			ans = base*i + ans
			num /= i
			base *= 10
		}
	}
	if num > 1 {
		return 0
	}
	return ans
}
