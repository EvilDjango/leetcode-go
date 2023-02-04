// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2023/2/4 下午1:45
package longest_palindromic_substring

import "strings"

// 动态规划
func longestPalindrome2(s string) string {
	size := len(s)
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}
	left, right := 0, 1
	for length := 1; length <= size; length++ {
		for i := 0; i+length-1 < size; i++ {
			dp[i][i+length-1] = length == 1 || (s[i] == s[i+length-1] && (length == 2 || dp[i+1][i+length-2]))
			if dp[i][i+length-1] && length > right-left {
				left, right = i, i+length
			}
		}
	}
	return s[left:right]
}

func longestPalindrome(s string) string {
	builder := strings.Builder{}
	builder.WriteString("$#")
	for _, r := range s {
		builder.WriteRune(r)
		builder.WriteRune('#')
	}
	builder.WriteRune('%')
	start, end := longest(builder.String())
	return s[(start-1)/2 : (end-1)/2]
}

func longest(s string) (int, int) {
	radius := make([]int, len(s))
	pivot, mx := 2, 3
	for i := 2; i < len(s)-2; i++ {
		length := 0
		if i < mx {
			length = min(mx-i, radius[2*pivot-i])
		}
		radius[i] = longestPalindromeFrom(s, i, length)
		if i+radius[i] > mx {
			pivot = i
			mx = i + radius[i]
		}
	}
	start, end := 1, 3
	for i, r := range radius {
		if 2*r+1 > end-start+1 {
			start = i - r + 1
			end = i + r - 1
		}
	}
	return start, end
}

// 计算以下标pivot为中心的最大回文串臂长，若回文串长度为2n+1，那么臂长为n
func longestPalindromeFrom(s string, pivot, length int) int {
	radius := length
	low, high := pivot-length, pivot+length
	for s[low] == s[high] {
		low--
		high++
		radius++
	}
	return radius
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
