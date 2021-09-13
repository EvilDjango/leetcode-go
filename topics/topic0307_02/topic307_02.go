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
package topic0307_02

// 线段树
type NumArray struct {
	segmentTree []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	segmentTree := make([]int, 4*n)
	buildTree(segmentTree, nums, 0, 0, n-1)
	return NumArray{segmentTree}
}

func buildTree(tree []int, nums []int, index, l, r int) int {
	var sum int
	if l == r {
		sum = nums[l]
	} else {
		mid := (r-l)/2 + l
		sum = buildTree(tree, nums, 2*index+1, l, mid) + buildTree(tree, nums, 2*index+2, mid+1, r)
	}
	tree[index] = sum
	return sum
}

func (this *NumArray) Update(index int, val int) {
	updateTree(this.segmentTree, 0, index, val, 0, len(this.segmentTree)/4-1)
}

func updateTree(tree []int, cur, target, val, l, r int) int {
	var increment int
	if l == r {
		increment = val - tree[cur]
	} else {
		mid := (r-l)/2 + l
		if target <= mid {
			increment = updateTree(tree, 2*cur+1, target, val, l, mid)
		} else {
			increment = updateTree(tree, 2*cur+2, target, val, mid+1, r)
		}
	}
	tree[cur] += increment
	return increment
}

func (this *NumArray) SumRange(left int, right int) int {
	return sumTree(this.segmentTree, left, right, 0, 0, len(this.segmentTree)/4-1)
}

func sumTree(tree []int, left, right, cur, l, r int) int {
	if right < l || left > r {
		return 0
	}
	if (left == l && right == r) || l == r {
		return tree[cur]
	}
	mid := (r-l)/2 + l
	return sumTree(tree, left, right, 2*cur+1, l, mid) + sumTree(tree, left, right, 2*cur+2, mid+1, r)
}
