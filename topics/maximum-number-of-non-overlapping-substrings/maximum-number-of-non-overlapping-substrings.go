// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/7 22:48
package maximum_number_of_non_overlapping_substrings

func maxNumOfSubstrings(s string) []string {
	firstIndices := make(map[byte]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		firstIndices[s[i]] = i
	}
	lastIndices := make(map[byte]int, 26)
	for i, r := range s {
		lastIndices[byte(r)] = i
	}

	var subs [][2]int
outer:
	for i := 0; i < len(s); i++ {
		// 每个字符的最小子串第一次遇见的时候就计算过了
		if firstIndices[s[i]] < i {
			continue
		}

		right := lastIndices[s[i]]
		for j := i + 1; j < right; j++ {
			if s[j] == s[i] {
				continue
			}
			last := lastIndices[s[j]]
			first := firstIndices[s[j]]
			// 这个字符的子串和前面的另一个字符是交叉耦合在一起的，所以它俩的最小子串是同一个，前面已经计算过了
			if first < i {
				continue outer
			}

			// 中间字符延伸到后边去了，我们不得不扩大范围
			if last > right {
				right = last
			}
		}

		// 前面可能计算出了一个大子串，后面如果遇到了包在它里面的子串，显然小子串更好，
		// 可以获得更多的子串数量，即使子串数量相同也可以得到更小的总长度
		if len(subs) > 0 {
			prev := subs[len(subs)-1]
			if prev[1] > right {
				subs = subs[:len(subs)-1]
			}
		}
		subs = append(subs, [2]int{i, right})
	}

	// 排序是为了方便后面的查找
	var ans []string
	for _, sub := range subs {
		ans = append(ans, s[sub[0]:sub[1]+1])
	}

	return ans
}
