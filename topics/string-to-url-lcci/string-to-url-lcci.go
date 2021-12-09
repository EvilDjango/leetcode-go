// 面试题 01.03. URL化
//URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）
//
//
//
//示例 1：
//
//输入："Mr John Smith    ", 13
//输出："Mr%20John%20Smith"
//示例 2：
//
//输入："               ", 5
//输出："%20%20%20%20%20"
//
//
//提示：
//
//字符串长度在 [0, 500000] 范围内。
//通过次数48,022提交次数83,070
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 10:55 PM
package string_to_url_lcci

import "strings"

func replaceSpaces(S string, length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		if S[i] != ' ' {
			builder.WriteByte(S[i])
		} else {
			builder.WriteString("%20")
		}
	}
	return builder.String()
}

// 直接操作底层的字符数组
func replaceSpaces2(S string, length int) string {
	bytes := []byte(S)
	i, j := len(S)-1, length-1
	for j >= 0 {
		if bytes[j] == ' ' {
			bytes[i] = '0'
			bytes[i-1] = '2'
			bytes[i-2] = '%'
			i -= 3
		} else {
			bytes[i] = bytes[j]
			i--
		}
		j--
	}
	return string(bytes[i+1:])
}
