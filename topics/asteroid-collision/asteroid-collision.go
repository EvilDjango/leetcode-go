// 735. 行星碰撞
//给定一个整数数组 asteroids，表示在同一行的行星。
//
//对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
//
//找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
//
//
//
//示例 1：
//
//输入：asteroids = [5,10,-5]
//输出：[5,10]
//解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
//示例 2：
//
//输入：asteroids = [8,-8]
//输出：[]
//解释：8 和 -8 碰撞后，两者都发生爆炸。
//示例 3：
//
//输入：asteroids = [10,2,-5]
//输出：[10]
//解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
//
//
//提示：
//
//2 <= asteroids.length <= 104
//-1000 <= asteroids[i] <= 1000
//asteroids[i] != 0
//通过次数54,229提交次数125,844
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/7/13 下午10:51
package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	n := len(asteroids)
	if n < 2 {
		return asteroids
	}
	var left, right []int
	for i := 0; i < n; i++ {
		asteroid := asteroids[i]
		if asteroid < 0 {
			m := len(right)
			collapse := false
			for j := m - 1; j >= 0; j-- {
				if right[j] > -asteroid {
					break
				}
				if right[j] == -asteroid {
					collapse = true
					right = right[:j]
					break
				}
				right = right[:j]
			}
			if len(right) == 0 && !collapse {
				left = append(left, asteroid)
			}
		} else {
			right = append(right, asteroid)
		}
	}
	return append(left, right...)
}

// 参考官方题解，使用栈
func asteroidCollision2(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		alive := true
		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			last := len(stack) - 1
			alive = stack[last] < -asteroid
			if stack[last] <= -asteroid {
				stack = stack[:last]
			}
		}
		if alive {
			stack = append(stack, asteroid)
		}
	}
	return stack
}
