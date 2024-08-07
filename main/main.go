// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/6 上午11:35
package main

import (
	"fmt"
	"strings"
)

func main() {
	//transformArray("[[1,1,2],[1,2,1],[2,1,1]]")
	fmt.Println(byte(99))
	fmt.Println(uint(1))
	fmt.Println(uint8(1))
}

func transformArray(arr string) {
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(arr, "[", "{"), "]", "}"))
}
