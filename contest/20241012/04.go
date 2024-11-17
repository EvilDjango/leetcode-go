// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/10/12 23:19
package _0241012

func numberOfWays(n int, x int, y int) int {
	if n == 1 {
		return x * y
	}
	if x == 1 {
		return n * y
	}
	if y == 1 {
		return power(x, n)
	}

	ans := 0
	mod := int(1e9 + 7)
	maxPrograms := min(n, x)
	for i := 0; i <= maxPrograms; i++ {
		count := comb(n, i) * power(i, n-i) * power(y, i) % mod
		ans += count
		ans %= mod
	}

	return ans
}

func power(base, exponent int) int {
	ans := 1
	for ; exponent > 0; exponent >>= 1 {
		if exponent&1 == 1 {
			ans *= base
		}
		base *= base
	}
	return ans % (1e9 + 7)
}

func comb(n, k int) int {
	ans := 1
	for i := n - k + 1; i <= n; i++ {
		ans *= i
	}
	return ans % (1e9 + 7)
}
