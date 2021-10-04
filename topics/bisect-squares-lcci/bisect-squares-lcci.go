// 面试题 16.13. 平分正方形
//给定两个正方形及一个二维平面。请找出将这两个正方形分割成两半的一条直线。假设正方形顶边和底边与 x 轴平行。
//
//每个正方形的数据square包含3个数值，正方形的左下顶点坐标[X,Y] = [square[0],square[1]]，以及正方形的边长square[2]。所求直线穿过两个正方形会形成4个交点，请返回4个交点形成线段的两端点坐标（两个端点即为4个交点中距离最远的2个点，这2个点所连成的线段一定会穿过另外2个交点）。2个端点坐标[X1,Y1]和[X2,Y2]的返回格式为{X1,Y1,X2,Y2}，要求若X1 != X2，需保证X1 < X2，否则需保证Y1 <= Y2。
//
//若同时有多条直线满足要求，则选择斜率最大的一条计算并返回（与Y轴平行的直线视为斜率无穷大）。
//
//示例：
//
//输入：
//square1 = {-1, -1, 2}
//square2 = {0, -1, 2}
//输出： {-1,0,2,0}
//解释： 直线 y = 0 能将两个正方形同时分为等面积的两部分，返回的两线段端点为[-1,0]和[2,0]
//提示：
//
//square.length == 3
//square[2] > 0
//通过次数1,948提交次数4,514
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/4/21 10:11 PM
package bisect_squares_lcci

func cutSquares(square1 []int, square2 []int) []float64 {
	// p1,p2分别是两正方形的中心，很显然同时平分两正方形的直线一定经过这两点
	p1 := []float64{float64(square1[0]) + float64(square1[2])/2, float64(square1[1]) + float64(square1[2])/2}
	p2 := []float64{float64(square2[0]) + float64(square2[2])/2, float64(square2[1]) + float64(square2[2])/2}

	ans := intersect(square1, p1, p2)
	points := intersect(square2, p1, p2)
	if points[0] < ans[0] || (points[0] == ans[0] && points[1] < ans[1]) {
		ans[0] = points[0]
		ans[1] = points[1]
	}
	if points[2] > ans[2] || (points[2] == ans[2] && points[3] > ans[3]) {
		ans[2] = points[2]
		ans[3] = points[3]
	}
	return ans
}

// 求正方形和直线的交点，由于直线经过正方形的中心，所以必然有两个交点
func intersect(square []int, p1, p2 []float64) []float64 {
	left, right, up, down := float64(square[0]), float64(square[0])+float64(square[2]), float64(square[1])+float64(square[2]), float64(square[1])
	if p1[0] == p2[0] {
		return []float64{p1[0], down, p1[0], up}
	} else if p1[1] == p2[1] {
		return []float64{left, p1[1], right, p1[1]}
	}
	points := make([]float64, 4)

	//直线和左边交点的纵坐标
	y := p1[1] + (p2[1]-p1[1])*(left-p1[0])/(p2[0]-p1[0])
	// 直线和左右两边相交的情况
	if y >= down && y <= up {
		points[0] = left
		points[1] = y
		points[2] = right
		points[3] = p1[1] + (p2[1]-p1[1])*(right-p1[0])/(p2[0]-p1[0])
	} else {
		// 直线和上下两边相交的情况
		points[0] = p1[0] + (p2[0]-p1[0])*(up-p1[1])/(p2[1]-p1[1])
		points[1] = up
		points[2] = p1[0] + (p2[0]-p1[0])*(down-p1[1])/(p2[1]-p1[1])
		points[3] = down
		if points[0] > points[2] {
			points[0], points[1], points[2], points[3] = points[2], points[3], points[0], points[1]
		}
	}
	return points
}
