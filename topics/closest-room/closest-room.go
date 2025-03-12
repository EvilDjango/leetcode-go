// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/12/21 14:41
package closest_room

import (
	"slices"
)

func closestRoom(rooms [][]int, queries [][]int) (ans []int) {
	slices.SortFunc(rooms, func(a, b []int) int {
		return b[1] - a[1]
	})
	queryIndices := make([]int, len(queries))
	for i := range queries {
		queryIndices[i] = i
	}
	slices.SortFunc(queryIndices, func(a, b int) int {
		return queries[b][1] - queries[a][1]
	})
	j := 0
	ans := make([]int, len(queries))
	roomSet := redBlackTree.New[int, bool]()
	for _, queryIndex := range queryIndices {
		prefferred, minSize := queries[queryIndex][0], queries[queryIndex][1]
		for ; j < len(rooms) && rooms[j][1] >= minSize; j++ {
			roomSet.Put(rooms[j][0], true)
		}
		roomId:=
		ans[queryIndex] = slices.Min(ans[queryIndex], len(rooms))
	}
}
