// 面试题 10.03. 搜索旋转数组
//搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
//
//示例1:
//
// 输入: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
// 输出: 8（元素5在该数组中的索引）
//示例2:
//
// 输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
// 输出：-1 （没有找到）
//提示:
//
//arr 长度范围在[1, 1000000]之间
//通过次数13,780提交次数35,003
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/13/21 4:15 PM
package search_rotate_array_lcci

func search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[0] == target {
		return 0
	}

	l, r := 1, len(arr)-1
	// 跳过开头的N个相同元素，N可能为1。这是为了避免首尾元素相同导致无法判断l,r是否在分割点同侧
	for l <= r && arr[l] == arr[l-1] {
		l++
	}
	for l <= r {
		mid := (r-l)>>1 + l
		// l,r在分割点同侧
		if arr[l] < arr[r] {
			if arr[mid] >= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
			continue
		}
		// l,r 在分割点异侧
		// mid在分割点左侧
		if arr[mid] >= arr[l] {
			if arr[mid] < target {
				l = mid + 1
			} else if arr[l] > target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// mid在分割点右侧
			if arr[mid] >= target {
				r = mid - 1
			} else if arr[r] < target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	if l < len(arr) && arr[l] == target {
		return l
	}
	return -1
}
