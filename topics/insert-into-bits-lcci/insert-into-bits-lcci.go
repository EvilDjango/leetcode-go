// insert-into-bits-lcci/
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 1:04 PM
package insert_into_bits_lcci

func insertBits(N int, M int, i int, j int) int {
	mask := (1<<(j-i+1) - 1) << i
	N ^= N & mask
	N |= M << i
	return N
}
