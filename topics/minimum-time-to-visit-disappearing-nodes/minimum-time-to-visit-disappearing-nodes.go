// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/18 下午10:41
package minimum_time_to_visit_disappearing_nodes

func minimumTime(n int, edges [][]int, disappear []int) []int {
	minimumTime := func(time int) map[int]int {
		adjList := make([]map[int]int, n)
		for i := 0; i < n; i++ {
			adjList[i] = make(map[int]int)
		}

		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			if disappear[u] <= time || disappear[v] <= time {
				continue
			}
			adjList[u][v] = w
			adjList[v][u] = w
		}

		answer := make([]int, n)
		determined := make([]bool, n)
		determined[0] = true
		for i := 1; i < n; i++ {
			determined[i] = false
			if w, ok := adjList[0][i]; ok {
				answer[i] = w
			} else {
				answer[i] = 1e10
			}
		}
		for i := 1; i < n; i++ {
			idx, shortest := -1, int(1e10)
			for j, v := range answer[1:] {
				if determined[j] {
					continue
				}
				if v < shortest {
					shortest, idx = v, j
				}
			}
			if idx == -1 {
				break
			}
			determined[idx] = true
			if shortest >= time {
				break
			}
			for v, w := range adjList[idx] {
				answer[v] = min(answer[v], answer[idx]+w)
			}
		}

		ret := make(map[int]int)
		for i, v := range answer {
			if v < time {
				ret[i] = v
			}
		}
		return ret
	}

	maxTime := 1
	for _, time := range disappear {
		maxTime = max(maxTime, time)
	}

	res := make([]int, n)
	for time := 1; time <= maxTime; time++ {
		ret := minimumTime(time)
		if len(ret) == 0 {
			break
		}
		for i, v := range ret {
			if res[i] != 0 {
				res[i] = v
			}
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
