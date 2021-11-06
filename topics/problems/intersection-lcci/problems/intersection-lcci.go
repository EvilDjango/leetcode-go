// 面试题 16.03. 交点
//给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。
//
//要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。
//
//
//
//示例 1：
//
//输入：
//line1 = {0, 0}, {1, 0}
//line2 = {1, 1}, {0, -1}
//输出： {0.5, 0}
//示例 2：
//
//输入：
//line1 = {0, 0}, {3, 3}
//line2 = {1, 1}, {2, 2}
//输出： {1, 1}
//示例 3：
//
//输入：
//line1 = {0, 0}, {1, 1}
//line2 = {1, 0}, {2, 1}
//输出： {}，两条线段没有交点
//
//
//提示：
//
//坐标绝对值不会超过 2^7
//输入的坐标均是有效的二维坐标
//通过次数9,654提交次数21,839
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/6/21 4:20 PM
package problems

// 使用直线的一般式Ax+By+C=0来求解。 已经两点 (x1,y1),(x2,y2) 可以得到A=y2-y1,B=x1-x2,c=x2y1-x1y2
// 已知两条直线A1x+B1y+C1=0 和 A2x+B2y+C2=0, 记D=A1B2-A2B1，若D==0则两直线平行，此时若A1C2-A2C1==0则两直线重合，否则没有交点
// 若D!=0则两直线相交，交点记为(x,y),x=(B1C2-B2C1)/D, y=(A2C1-A1C2)/D
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	if start1[0] > end1[0] {
		start1, end1 = end1, start1
	}
	if start2[0] > end2[0] {
		start2, end2 = end2, start2
	}
	if start1[0] > start2[0] {
		start1, end1, start2, end2 = start2, end2, start1, end1
	}
	A1 := end1[1] - start1[1]
	B1 := start1[0] - end1[0]
	C1 := end1[0]*start1[1] - start1[0]*end1[1]
	A2 := end2[1] - start2[1]
	B2 := start2[0] - end2[0]
	C2 := end2[0]*start2[1] - start2[0]*end2[1]
	D1 := A1*B2 - A2*B1
	D2 := A1*C2 - A2*C1

	// 平行
	if D1 == 0 && D2 != 0 {
		return nil
	}

	// 两直线重合
	if D1 == 0 && D2 == 0 {
		//不垂直于x轴的情形
		if start1[0] != end1[0] {
			// 线段重合
			if end1[0] >= start2[0] {
				return []float64{float64(start2[0]), float64(start2[1])}
			}
			// 线段不重合
			return nil
		}

		// 垂直于x轴的情形
		if start1[1] > end1[1] {
			start1, end1 = end1, start1
		}
		if start2[1] > end2[1] {
			start2, end2 = end2, start2
		}
		if start1[1] > start2[1] {
			start1, end1, start2, end2 = start2, end2, start1, end1
		}

		// 线段重合
		if end1[1] >= start2[1] {
			return []float64{float64(start2[0]), float64(start2[1])}
		}
		// 线段不重合
		return nil
	}

	// 余下的是直线相交的情况
	x := float64(B1*C2-B2*C1) / float64(D1)
	y := float64(A2*C1-A1*C2) / float64(D1)
	if x >= float64(start1[0]) && x <= float64(end1[0]) && x >= float64(start2[0]) && x <= float64(end2[0]) &&
		y >= float64(min(start1[1], end1[1])) &&
		y <= float64(max(start1[1], end1[1])) &&
		y >= float64(min(start2[1], end2[1])) &&
		y <= float64(max(start2[1], end2[1])) {
		return []float64{x, y}
	}
	return nil
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
