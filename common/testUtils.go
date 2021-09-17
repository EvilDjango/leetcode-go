// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 12:27 AM
package common

import "math/rand"

func RandomInts() []int {
	size := rand.Intn(1000)
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = rand.Int()
	}
	return ints
}
