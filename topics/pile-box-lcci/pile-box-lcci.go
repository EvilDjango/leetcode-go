// 面试题 08.13. 堆箱子
//堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。
//
//输入使用数组[wi, di, hi]表示每个箱子。
//
//示例1:
//
// 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
// 输出：6
//示例2:
//
// 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
// 输出：10
//提示:
//
//箱子的数目不大于3000个。
//通过次数7,160提交次数14,266
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 8:54 AM
package pile_box_lcci

import "sort"

// 这是一个拓扑排序的问题。若全部箱子有拓扑排序则结果是全部箱子的高度值和。如果没有拓扑排序，则要考虑要舍弃哪些箱子，情况相对复杂。
func pileBox(box [][]int) int {
	n := len(box)
	sortByWidth := make([]int, n)
	sortByDepth := make([]int, n)
	sortByHeight := make([]int, n)
	for i := 0; i < n; i++ {
		sortByWidth[i] = i
		sortByDepth[i] = i
		sortByHeight[i] = i
	}
	sort.Slice(sortByWidth, func(i, j int) bool {
		return box[i][0] < box[i][0]
	})
	sort.Slice(sortByDepth, func(i, j int) bool {
		return box[i][1] < box[i][1]
	})
	sort.Slice(sortByHeight, func(i, j int) bool {
		return box[i][2] < box[i][2]
	})
	adjacentList := map[int][]int{}
	for i := 1; i < n; i++ {
		adjacentList[sortByWidth[i-1]] = append(adjacentList[sortByWidth[i-1]], sortByWidth[i])
	}
	for i := 1; i < n; i++ {
		adjacentList[sortByDepth[i-1]] = append(adjacentList[sortByDepth[i-1]], sortByDepth[i])
	}
	for i := 1; i < n; i++ {
		adjacentList[sortByHeight[i-1]] = append(adjacentList[sortByHeight[i-1]], sortByHeight[i])
	}
	// 每个箱子已经达到的最大高度
	maxHeights := make([]int, n)
	for i := 0; i < n; i++ {
		// -1表示这个箱子还没有访问过
		maxHeights[i] = -1
	}
	maxHeight := 0
	dfs(adjacentList, -1, sortByWidth[0], maxHeights, &maxHeight)
	return maxHeight
}

func dfs(list map[int][]int, prev, curr int, maxHeights []int, maxHeight *int) {
	// 当前的节点已经访问过，成环了,要继续向后拓展，必须移除当前节点，然后从前一节点继续遍历
	if maxHeights[curr] > 0 {
		maxHeights[prev] -= maxHeights[curr]
		dfs(list, curr, prev, maxHeights, maxHeight)
	}
}
