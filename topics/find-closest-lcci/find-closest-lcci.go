// 面试题 17.11. 单词距离
//有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
//
//示例：
//
//输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
//输出：1
//提示：
//
//words.length <= 100000
//通过次数14,207提交次数20,692
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 3:36 PM
package find_closest_lcci

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	indexes := make(map[string][]int)
	for i, w := range words {
		indexes[w] = append(indexes[w], i)
	}
	return minDis(indexes[word1], indexes[word2])
}

func minDis(ints, ints2 []int) int {
	minDis := math.MaxInt32
	m, n := len(ints), len(ints2)
	i, j := 0, 0
	for i < m && j < n {
		dis := abs(ints[i] - ints2[j])
		if dis < minDis {
			minDis = dis
		}
		if ints[i] < ints2[j] {
			i++
		} else if ints[i] > ints2[j] {
			j++
		} else {
			return 0
		}
	}
	return minDis
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findClosest2(words []string, word1 string, word2 string) int {
	minDis := math.MaxInt32
	i, j := -1, -1
	for k, w := range words {
		if w == word1 {
			i = k
		}
		if w == word2 {
			j = k
		}
		if i != -1 && j != -1 && abs(i-j) < minDis {
			minDis = abs(i - j)
		}
	}
	return minDis
}
