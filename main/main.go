// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/6 上午11:35
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Inf(1))
	transformArray("[[1,1,2],[1,2,1],[2,1,1]]")

}

func transformArray(arr string) {
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(arr, "[", "{"), "]", "}"))
}
