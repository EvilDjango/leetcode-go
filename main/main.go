// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/6 上午11:35
package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func main() {
	for {
		foo()
	}
}
func foo() {
	bound := 1000
	size := rand.Intn(bound)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(bound)
	}
	indices := make([]int, size)
	for i := 0; i < size; i++ {
		indices[i] = i
	}
	//sort.Slice(indices, func(i, j int) bool {
	//	return arr[indices[i]] < arr[indices[j]]
	//})

	slices.SortFunc(indices, func(i, j int) int {
		return arr[i] - arr[j]
	})

	for i := 1; i < size; i++ {
		if arr[indices[i]] < arr[indices[i-1]] {
			panic("gotcha")
		}
	}
	fmt.Println("finished")
}

func transformArray(arr string) {
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(arr, "[", "{"), "]", "}"))
}
