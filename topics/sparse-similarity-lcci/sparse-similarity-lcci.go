// 面试题 17.26. 稀疏相似度
//两个(具有不同单词的)文档的交集(intersection)中元素的个数除以并集(union)中元素的个数，就是这两个文档的相似度。例如，{1, 5, 3} 和 {1, 7, 2, 3} 的相似度是 0.4，其中，交集的元素有 2 个，并集的元素有 5 个。给定一系列的长篇文档，每个文档元素各不相同，并与一个 ID 相关联。它们的相似度非常“稀疏”，也就是说任选 2 个文档，相似度都很接近 0。请设计一个算法返回每对文档的 ID 及其相似度。只需输出相似度大于 0 的组合。请忽略空文档。为简单起见，可以假定每个文档由一个含有不同整数的数组表示。
//
//输入为一个二维数组 docs，docs[i] 表示 id 为 i 的文档。返回一个数组，其中每个元素是一个字符串，代表每对相似度大于 0 的文档，其格式为 {id1},{id2}: {similarity}，其中 id1 为两个文档中较小的 id，similarity 为相似度，精确到小数点后 4 位。以任意顺序返回数组均可。
//
//示例:
//
//输入:
//[
//  [14, 15, 100, 9, 3],
//  [32, 1, 9, 3, 5],
//  [15, 29, 2, 6, 8, 7],
//  [7, 10]
//]
//输出:
//[
//  "0,1: 0.2500",
//  "0,2: 0.1000",
//  "2,3: 0.1429"
//]
//提示：
//
//docs.length <= 500
//docs[i].length <= 500
//通过次数3,309提交次数9,700
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/9/21 12:11 PM
package sparse_similarity_lcci

import (
	"fmt"
	"sort"
)

// 先排序，然后计算相似度
// 不知道为什么，只能通过一半的测试用例
func computeSimilarities(docs [][]int) []string {
	for _, doc := range docs {
		sort.Ints(doc)
	}
	n := len(docs)
	var ans []string
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			similarity := getSimilarity(docs[i], docs[j])
			if similarity > 0 {
				ans = append(ans, fmt.Sprintf("%d,%d: %.4f", i, j, similarity))
			}
		}
	}
	return ans
}

func getSimilarity(doc1, doc2 []int) float64 {
	m, n := len(doc1), len(doc2)
	if m*n == 0 {
		return 0
	}
	if doc1[0] > doc2[0] {
		doc1, doc2 = doc2, doc1
		m, n = n, m
	}
	if doc1[m-1] < doc2[0] {
		return 0
	}
	i := m - 1
	for ; i > 0 && doc1[i] > doc2[0]; i-- {
	}
	j := 0
	common := 0
	for i < m && j < n {
		if doc1[i] > doc2[j] {
			j++
		} else if doc1[i] < doc2[j] {
			i++
		} else {
			common++
			i++
			j++
		}
	}
	return float64(common) / float64(m+n-common)
}

// 利用哈希表
func computeSimilarities2(docs [][]int) []string {
	n := len(docs)
	idsByWord := make(map[int][]int)
	commons := make([][]int, n)
	for i := 0; i < n; i++ {
		commons[i] = make([]int, n)
	}
	for id, doc := range docs {
		for _, word := range doc {
			ids := idsByWord[word]
			for _, sharer := range ids {
				commons[id][sharer]++
				commons[sharer][id]++
			}
			idsByWord[word] = append(ids, id)
		}
	}
	var ans []string
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if commons[i][j] > 0 {
				similarity := float64(commons[i][j]) / float64(len(docs[i])+len(docs[j])-commons[i][j])
				ans = append(ans, fmt.Sprintf("%d,%d: %.4f", i, j, similarity+1e-9))
			}
		}
	}
	return ans
}
