// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2023/2/19 上午11:17
package maximum_average_pass_ratio

import (
	"container/heap"
	"sort"
)

type data [][]int

func (this data) Len() int {
	return len(this)
}

func (this data) Swap(i, j int) {
	this[j], this[j] = this[j], this[i]
}
func (this data) Less(i, j int) bool {
	d1 := (this[i][0]+1)/(this[i][1]+1) - (this[i][0] / this[i][1])
	d1 := (this[i][0]+1)/(this[i][1]+1) - (this[i][0] / this[i][1])
}

func (this *data) Push(v interface{}) {
	*this = append(*this, v.([]int))
}

func (this *data) Pop() interface{} {
	a := *this
	v := a[len(a)-1]
	*this = a[:len(a)-1]
	return v
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	heap.Init()

}
