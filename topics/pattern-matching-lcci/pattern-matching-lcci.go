// 面试题 16.18. 模式匹配
//你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。
//
//示例 1：
//
//输入： pattern = "abba", value = "dogcatcatdog"
//输出： true
//示例 2：
//
//输入： pattern = "abba", value = "dogcatcatfish"
//输出： false
//示例 3：
//
//输入： pattern = "aaaa", value = "dogcatcatdog"
//输出： false
//示例 4：
//
//输入： pattern = "abba", value = "dogdogdogdog"
//输出： true
//解释： "a"="dogdog",b=""，反之也符合规则
//提示：
//
//1 <= len(pattern) <= 1000
//0 <= len(value) <= 1000
//你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。
//通过次数15,113提交次数43,911
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/24/21 8:38 AM
package pattern_matching_lcci

// 回溯
func patternMatching(pattern string, value string) bool {
	if len(pattern) == 1 {
		return true
	}
	return dfs(pattern, 0, value, 0, map[byte]string{}, map[string]bool{})
}

func dfs(pattern string, i int, value string, j int, patterns map[byte]string, used map[string]bool) bool {
	if i == len(pattern) {
		return i == len(pattern) && j == len(value)
	}
	p := pattern[i]
	if v, ok := patterns[p]; ok {
		if j+len(v) <= len(value) && value[j:j+len(v)] == v && dfs(pattern, i+1, value, j+len(v), patterns, used) {
			return true
		}
		return false
	}
	for k := j; k <= len(value); k++ {
		v := value[j:k]
		if used[v] {
			continue
		}
		patterns[p] = v
		used[v] = true
		if dfs(pattern, i+1, value, k, patterns, used) {
			return true
		}
		delete(patterns, p)
		delete(used, v)
	}
	return false
}

// 由于模式串中只包含两个字符，所以可以做简化处理，枚举两个匹配字符的可能长度即可。
func patternMatching2(pattern string, value string) bool {
	nA, nB := 0, 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == pattern[0] {
			nA++
		} else {
			nB++
		}
	}
	for lenA := 0; lenA <= len(value)/nA; lenA++ {
		lenB := 0
		if nB != 0 {
			lenB = (len(value) - lenA*nA) / nB
		}
		if lenA*nA+lenB*nB != len(value) {
			continue
		}
		vA, vB := value[0:lenA], ""
		i, j := 1, lenA
		for ; i < len(pattern); i++ {
			if pattern[i] == pattern[0] {
				if value[j:j+lenA] != vA {
					break
				}
				j += lenA
			} else {
				if len(vB) != lenB {
					vB = value[j : j+lenB]
				} else if vA == vB || value[j:j+lenB] != vB {
					break
				}
				j += lenB
			}
		}
		if i == len(pattern) {
			return true
		}
	}
	return false
}
