// 剑指 Offer II 111. 计算除法
//给定一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
//另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
//
//返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
//
//注意：输入总是有效的。可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
//
//
//
//示例 1：
//
//输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
//输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
//解释：
//条件：a / b = 2.0, b / c = 3.0
//问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
//结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
//示例 2：
//
//输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
//输出：[3.75000,0.40000,5.00000,0.20000]
//示例 3：
//
//输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
//输出：[0.50000,2.00000,-1.00000,-1.00000]
//
//
//提示：
//
//1 <= equations.length <= 20
//equations[i].length == 2
//1 <= Ai.length, Bi.length <= 5
//values.length == equations.length
//0.0 < values[i] <= 20.0
//1 <= queries.length <= 20
//queries[i].length == 2
//1 <= Cj.length, Dj.length <= 5
//Ai, Bi, Cj, Dj 由小写英文字母与数字组成
//
//
//注意：本题与主站 399 题相同： https://leetcode-cn.com/problems/evaluate-division/
//
//通过次数1,178提交次数1,754
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/13/21 2:47 PM
package vlzXQL

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	parent := map[string]string{}
	multiple := map[string]float64{}
	var find func(string) string
	find = func(s string) string {
		if v, ok := parent[s]; !ok {
			parent[s] = s
			multiple[s] = 1
		} else if v != s {
			p := parent[s]
			parent[s] = find(parent[s])
			multiple[s] *= multiple[p]
		}
		return parent[s]
	}
	union := func(a, b string, ratio float64) {
		x, y := find(a), find(b)
		// 条件矛盾
		if x == y {
			if multiple[a]/multiple[b] != ratio {
				panic("contradiction")
			}
			return
		}
		parent[x] = y
		multiple[x] = ratio * multiple[b] / multiple[a]
	}

	calculate := func(a, b string) float64 {
		if _, ok := parent[a]; !ok {
			return -1
		}
		if _, ok := parent[b]; !ok {
			return -1
		}
		x, y := find(a), find(b)
		if x != y {
			return -1
		}
		return multiple[a] / multiple[b]
	}
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		union(e[0], e[1], values[i])
	}
	ans := make([]float64, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, calculate(query[0], query[1]))
	}
	return ans
}

// 使用Floyd-Warshall算法，计算出任意两点间的距离
func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, e := range equations {
		if _, ok := id[e[0]]; !ok {
			id[e[0]] = len(id)
		}
		if _, ok := id[e[1]]; !ok {
			id[e[1]] = len(id)
		}
	}
	n := len(id)
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}
	for i, e := range equations {
		u, v := id[e[0]], id[e[1]]
		matrix[u][v] = values[i]
		matrix[v][u] = 1 / values[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if matrix[j][k] != -1 {
					continue
				}
				if matrix[j][i] != -1 && matrix[i][k] != -1 {
					matrix[j][k] = matrix[j][i] * matrix[i][k]
				}
			}
		}
	}
	ans := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		u, ok1 := id[queries[i][0]]
		v, ok2 := id[queries[i][1]]
		if ok1 && ok2 {
			ans[i] = matrix[u][v]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
