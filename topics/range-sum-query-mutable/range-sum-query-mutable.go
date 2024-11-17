// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/21 22:26
package range_sum_query_mutable

type NumArray []int

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)<<2)
	var build func(int, int, int)
	build = func(left, right, index int) {
		if left == right {
			arr[index] = nums[left]
			return
		}
		mid := left + ((right - left) >> 1)
		build(left, mid, index<<1)
		build(mid+1, right, index<<1+1)
		arr[index] = arr[index<<1] + arr[index<<1+1]
	}
	build(0, len(nums)-1, 1)
	return arr
}

func (this *NumArray) Update(index int, val int) {
	this.update(index, val, 1, 0, len(*this)>>2-1)
}

func (this *NumArray) update(target, val, index, left, right int) {
	arr := *this
	if left == right {
		arr[index] = val
		return
	}
	mid := left + ((right - left) >> 1)
	if target <= mid {
		this.update(target, val, index<<1, left, mid)
	} else {
		this.update(target, val, index<<1+1, mid+1, right)
	}
	arr[index] = arr[index<<1] + arr[index<<1+1]
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(left, right, 1, 0, len(*this)>>2-1)
}

func (this *NumArray) query(targetLeft, targetRight, index, left, right int) int {
	arr := *this
	if targetLeft == left && targetRight == right {
		return arr[index]
	}
	mid := left + (right-left)>>1
	lSum, rSum := 0, 0
	if targetLeft <= mid {
		lSum = this.query(targetLeft, targetRight, index<<1, left, mid)
	}
	if targetRight > mid {
		rSum = this.query(targetLeft, targetRight, index<<1+1, mid+1, right)
	}
	return lSum + rSum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
