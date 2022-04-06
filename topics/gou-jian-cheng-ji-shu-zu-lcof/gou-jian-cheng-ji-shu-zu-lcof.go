// 剑指 Offer 66. 构建乘积数组
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//
//提示：
//
//所有元素乘积之和不会溢出 32 位整数
//a.length <= 100000
//通过次数100,167提交次数166,521
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/4/6 下午12:37
package gou_jian_cheng_ji_shu_zu_lcof

// 使用前缀乘积和后缀乘积
func constructArr(a []int) []int {
	n := len(a)
	prefix := make([]int, n+1)
	prefix[0] = 1
	for i, num := range a {
		prefix[i+1] = prefix[i] * num
	}
	suffix := 1
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = suffix * prefix[i]
		suffix *= a[i]
	}
	return ans
}

// 通过复用数组来降低空间复杂度
func constructArr2(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= a[i]
	}
	return ans
}
