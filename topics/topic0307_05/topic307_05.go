// 给你一个数组 nums ，请你完成两类查询，其中一类查询要求更新数组下标对应的值，另一类查询要求返回数组中某个范围内元素的总和。
//
//实现 NumArray 类：
//
//NumArray(int[] nums) 用整数数组 nums 初始化对象
//void update(int index, int val) 将 nums[index] 的值更新为 val
//int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和（即，nums[left] + nums[left + 1], ..., nums[right]）
//
//
//示例：
//
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 9 ，sum([1,3,5]) = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 8 ，sum([1,2,5]) = 8
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//-100 <= nums[i] <= 100
//0 <= index < nums.length
//-100 <= val <= 100
//0 <= left <= right < nums.length
//最多调用 3 * 104 次 update 和 sumRange 方法
//通过次数23,000提交次数42,351
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/range-sum-query-mutable
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/1/21 2:32 PM
package topic0307_05

//
//// 线段树
//type NumArray struct {
//	container []int
//}
//
//func Constructor(nums []int) NumArray {
//	container := make([]int, len(nums)<<1)
//	buildTree(container, nums)
//	return NumArray{container}
//}
//
//func buildTree(container, nums []int) {
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		container[i+n] = nums[i]
//	}
//	for i := n - 1; i > 0; i-- {
//		container[i] = container[i<<1] + container[(i<<1)+1]
//	}
//}
//
//func (this *NumArray) Update(index int, val int) {
//	n := len(this.container) >> 1
//	index += n
//	increment := val - this.container[index]
//	for index > 0 {
//		this.container[index] += increment
//		index >>= 1
//	}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	n := len(this.container) >> 1
//	left += n
//	right += n
//	sum := 0
//	for left <= right {
//		if left&1 == 1 {
//			sum += this.container[left]
//			left++
//		}
//		if right&1 == 0 {
//			sum += this.container[right]
//			right--
//		}
//		left >>= 1
//		right >>= 1
//	}
//	return sum
//}

// 树状数组
type NumArray struct {
	tree, nums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	buildTree(tree, nums)
	return NumArray{tree, nums}
}
func buildTree(tree, nums []int) {
	n := len(nums)
	for i := 1; i <= n; i++ {
		tree[i] += nums[i-1]
		j := i + lowBit(i)
		if j <= n {
			tree[j] += tree[i]
		}
	}
}

func (this *NumArray) Update(index int, val int) {
	increment := val - this.nums[index]
	for i := index + 1; i < len(this.tree); i += lowBit(i) {
		this.tree[i] += increment
	}
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right) - this.query(left-1)
}

// nums前i个元素的和
func (this *NumArray) query(i int) int {
	sum := 0
	for j := i + 1; j > 0; j -= lowBit(j) {
		sum += this.tree[j]
	}
	return sum
}

func lowBit(x int) int {
	return x & -x
}
