// 面试题 17.17. 多次搜索
//给定一个较长字符串big和一个包含较短字符串的数组smalls，设计一个方法，根据smalls中的每一个较短字符串，对big进行搜索。输出smalls中的字符串在big里出现的所有位置positions，其中positions[i]为smalls[i]出现的所有位置。
//
//示例：
//
//输入：
//big = "mississippi"
//smalls = ["is","ppi","hi","sis","i","ssippi"]
//输出： [[1,4],[8],[],[3],[1,4,7,10],[5]]
//提示：
//
//0 <= len(big) <= 1000
//0 <= len(smalls[i]) <= 1000
//smalls的总字符数不会超过 100000。
//你可以认为smalls中没有重复字符串。
//所有出现的字符均为英文小写字母。
//通过次数7,654提交次数17,306
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/15/21 10:53 AM
package topic2344

import "strings"

type Trie struct {
	IsLeaf   bool
	Children [26]*Trie
}

func (t *Trie) Insert(s string) {
	curr := t
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if curr.Children[index] == nil {
			curr.Children[index] = &Trie{}
		}
		curr = curr.Children[index]
	}
	curr.IsLeaf = true
}

func multiSearch(big string, smalls []string) [][]int {
	root := &Trie{}
	indexes := make(map[string]int)
	for i, s := range smalls {
		root.Insert(s)
		indexes[s] = i
	}
	ans := make([][]int, len(smalls))
	for i := 0; i < len(big); i++ {
		curr := root
		for j := i; j < len(big); j++ {
			curr = curr.Children[big[j]-'a']
			if curr == nil {
				break
			}
			if curr.IsLeaf {
				w := big[i : j+1]
				index := indexes[w]
				ans[index] = append(ans[index], i)
			}
		}
	}
	return ans
}

// 朴素算法
func multiSearch2(big string, smalls []string) [][]int {
	ans := make([][]int, len(smalls))
	for i, s := range smalls {
		if s == "" {
			continue
		}
		for j := 0; j < len(big); {
			index := strings.Index(big[j:], s)
			if index < 0 {
				break
			}
			ans[i] = append(ans[i], index)
			j += index + 1
		}
	}
	return ans
}

// 马拉车算法
func multiSearch3(big string, smalls []string) [][]int {
	ans := make([][]int, len(smalls))
	for i, s := range smalls {
		if s == "" {
			continue
		}
		for j := 0; j < len(big); {
			index := strings.Index(big[j:], s)
			if index < 0 {
				break
			}
			ans[i] = append(ans[i], index)
			j += index + 1
		}
	}
	return ans
}
