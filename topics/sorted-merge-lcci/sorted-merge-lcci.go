// 面试题 10.01. 合并排序的数组
//给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
//
//初始化 A 和 B 的元素数量分别为 m 和 n。
//
//示例:
//
//输入:
//A = [1,2,3,0,0,0], m = 3
//B = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
//说明:
//
//A.length == n + m
//通过次数57,785提交次数104,099
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/15/21 11:35 AM
package sorted_merge_lcci

func merge(A []int, m int, B []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	//依次将最大的元素放到尾部
	for j >= 0 && i >= 0 {
		if A[i] >= B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
		k--
	}
	// 将B数组剩余元素加入到A中。注意，如果此时j=-1而i>=0,是无需操作的，因为A[0]...A[i]本身就在A数组中且已经正确排序
	for j >= 0 {
		A[k] = B[j]
		k--
		j--
	}
}
