// 面试题 08.06. 汉诺塔问题
//在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
//(1) 每次只能移动一个盘子;
//(2) 盘子只能从柱子顶端滑出移到下一根柱子;
//(3) 盘子只能叠在比它大的盘子上。
//
//请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
//
//你需要原地修改栈。
//
//示例1:
//
// 输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
//示例2:
//
// 输入：A = [1, 0], B = [], C = []
// 输出：C = [1, 0]
//提示:
//
//A中盘子的数目不大于14个。
//通过次数25,094提交次数38,264
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/25/21 4:47 PM
package hanota_lcci

// 分治
func hanota(A []int, B []int, C []int) []int {
	dfs(&A, len(A), &B, &C)
	return C
}

func dfs(src *[]int, n int, middle *[]int, dest *[]int) {
	if n == 1 {
		*dest = append(*dest, (*src)[len(*src)-1])
		*src = (*src)[:len(*src)-1]
		return
	}
	dfs(src, n-1, dest, middle)
	dfs(src, 1, middle, dest)
	dfs(middle, n-1, src, dest)
}
