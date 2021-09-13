// 面试题 17.25. 单词矩阵
//给定一份单词的清单，设计一个算法，创建由字母组成的面积最大的矩形，其中每一行组成一个单词(自左向右)，每一列也组成一个单词(自上而下)。不要求这些单词在清单里连续出现，但要求所有行等长，所有列等高。
//
//如果有多个面积最大的矩形，输出任意一个均可。一个单词可以重复使用。
//
//示例 1:
//
//输入: ["this", "real", "hard", "trh", "hea", "iar", "sld"]
//输出:
//[
//   "this",
//   "real",
//   "hard"
//]
//示例 2:
//
//输入: ["aa"]
//输出: ["aa","aa"]
//说明：
//
//words.length <= 1000
//words[i].length <= 100
//数据保证单词足够随机
//通过次数1,382提交次数2,803
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/9/21 2:46 PM
package topic2341

import (
	"math"
	"strings"
)

// 暴力算法
func maxRectangle(words []string) []string {
	dict := make(map[string]bool)
	wordsByLength := make(map[int][]string)
	maxLen := math.MinInt32
	for _, word := range words {
		dict[word] = true
		l := len(word)
		wordsByLength[l] = append(wordsByLength[len(word)], word)
		if l > maxLen {
			maxLen = l
		}
	}
	for l := maxLen; l > 0; l-- {
		if _, ok := wordsByLength[l]; !ok {
			continue
		}
		var rectangles [][]string
		getRectangles(dict, wordsByLength[l], l, 0, []string{}, &rectangles)
		if len(rectangles) > 0 {
			return rectangles[len(rectangles)-1]
		}
	}
	return nil
}

func isValid(dict map[string]bool, rectangle []string) bool {
	n := len(rectangle[0])
	cols := make([]strings.Builder, n)
	for _, word := range rectangle {
		for i := 0; i < n; i++ {
			cols[i].WriteByte(word[i])
		}
	}
	for _, col := range cols {
		if !dict[col.String()] {
			return false
		}
	}
	return true
}

func getRectangles(dict map[string]bool, words []string, length, i int, rectangle []string, rectangles *[][]string) {
	if i > 0 && isValid(dict, rectangle) {
		arr := make([]string, len(rectangle))
		copy(arr, rectangle)
		*rectangles = append(*rectangles, arr)
	}
	if i == length {
		return
	}
	for _, word := range words {
		rectangle = append(rectangle, word)
		getRectangles(dict, words, length, i+1, rectangle, rectangles)
		rectangle = rectangle[:len(rectangle)-1]
	}
}
