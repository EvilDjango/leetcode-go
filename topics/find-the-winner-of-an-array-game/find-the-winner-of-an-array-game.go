// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/30 23:01
package find_the_winner_of_an_array_game

func getWinner(arr []int, k int) int {
	n := len(arr)
	if k >= n-1 {
		maxVal := arr[0]
		for _, v := range arr {
			maxVal = max(maxVal, v)
		}
		return maxVal
	}
	for i := 0; i < n; i++ {
		if i > 0 && arr[i] <= arr[i-1] {
			continue
		}
		cnt := 1
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				cnt++
				if cnt == k {
					return arr[i]
				}
			} else {
				break
			}
		}
	}
	panic("won't reach")
}
