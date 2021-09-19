// 面试题 17.07. 婴儿名字
//每年，政府都会公布一万个最常见的婴儿名字和它们出现的频率，也就是同名婴儿的数量。有些名字有多种拼法，例如，John 和 Jon 本质上是相同的名字，但被当成了两个名字公布出来。给定两个列表，一个是名字及对应的频率，另一个是本质相同的名字对。设计一个算法打印出每个真实名字的实际频率。注意，如果 John 和 Jon 是相同的，并且 Jon 和 Johnny 相同，则 John 与 Johnny 也相同，即它们有传递和对称性。
//
//在结果列表中，选择 字典序最小 的名字作为真实名字。
//
//
//
//示例：
//
//输入：names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"], synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
//输出：["John(27)","Chris(36)"]
//
//
//提示：
//
//names.length <= 100000
//通过次数8,581提交次数20,848
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/19/21 9:33 PM
package baby_names_lcci_02

import (
	"fmt"
	"strconv"
	"strings"
)

// 试图使用名字下标代替字符串来加速，实际上加速效果优先，而且代码看起来很臃肿。没有什么必要。
func trulyMostPopular(names []string, synonyms []string) []string {
	n := len(names)
	realNames := make([]string, n)
	nameIndexes := make(map[string]int)
	frequencies := make([]int, n)
	ancestors := make([]int, n)
	for i, item := range names {
		index := strings.Index(item, "(")
		name := item[:index]
		realNames[i] = name
		nameIndexes[name] = i
		frequencies[i], _ = strconv.Atoi(item[index+1 : len(item)-1])
		ancestors[i] = i
	}
	for _, item := range synonyms {
		index := strings.Index(item, ",")
		word1, word2 := item[1:index], item[index+1:len(item)-1]
		w1, w2 := nameIndexes[word1], nameIndexes[word2]
		if _, ok := nameIndexes[word1]; !ok {
			ancestors[w1] = w1
		}
		if _, ok := nameIndexes[word2]; !ok {
			ancestors[w2] = w2
		}
		merge(ancestors, frequencies, realNames, w1, w2)
	}
	var ans []string
	for w, root := range ancestors {
		if w == root {
			word := fmt.Sprintf("%s(%d)", realNames[w], frequencies[w])
			ans = append(ans, word)
		}
	}
	return ans
}

func merge(ancestors, frequencies []int, names []string, w1, w2 int) {
	root1 := find(ancestors, w1)
	root2 := find(ancestors, w2)
	if root1 == root2 {
		return
	}
	if names[root1] < names[root2] {
		ancestors[root2] = root1
		frequencies[root1] += frequencies[root2]
	} else {
		ancestors[root1] = root2
		frequencies[root2] += frequencies[root1]
	}
}

func find(ancestors []int, w int) int {
	for ancestors[w] != w {
		ancestors[w] = ancestors[ancestors[w]]
		w = ancestors[w]
	}
	return ancestors[w]
}
