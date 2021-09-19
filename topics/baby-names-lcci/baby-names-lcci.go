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

func trulyMostPopular(names []string, synonyms []string) []string {
	frequencies := make(map[string]int, len(names))
	ancestors := make(map[string]string, len(names))
	for _, item := range names {
		index := strings.Index(item, "(")
		name := item[:index]
		frequencies[name], _ = strconv.Atoi(item[index+1 : len(item)-1])
		ancestors[name] = name
	}
	for _, item := range synonyms {
		index := strings.Index(item, ",")
		w1, w2 := item[1:index], item[index+1:len(item)-1]
		if ancestors[w1] == "" {
			ancestors[w1] = w1
		}
		if ancestors[w2] == "" {
			ancestors[w2] = w2
		}
		merge(ancestors, frequencies, w1, w2)
	}
	var ans []string
	for w, root := range ancestors {
		if w == root {
			word := fmt.Sprintf("%s(%d)", w, frequencies[w])
			ans = append(ans, word)
		}
	}
	return ans
}

func merge(ancestors map[string]string, frequencies map[string]int, w1, w2 string) {
	root1 := find(ancestors, w1)
	root2 := find(ancestors, w2)
	if root1 == root2 {
		return
	}
	if root1 < root2 {
		ancestors[root2] = root1
		frequencies[root1] += frequencies[root2]
	} else {
		ancestors[root1] = root2
		frequencies[root2] += frequencies[root1]
	}
}

func find(ancestors map[string]string, w string) string {
	for ancestors[w] != w {
		ancestors[w] = ancestors[ancestors[w]]
		w = ancestors[w]
	}
	return ancestors[w]
}
