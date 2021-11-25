// 面试题 08.10. 颜色填充
//编写函数，实现许多图片编辑软件都支持的「颜色填充」功能。
//
//待填充的图像用二维数组 image 表示，元素为初始颜色值。初始坐标点的行坐标为 sr 列坐标为 sc。需要填充的新颜色为 newColor 。
//
//「周围区域」是指颜色相同且在上、下、左、右四个方向上存在相连情况的若干元素。
//
//请用新颜色填充初始坐标点的周围区域，并返回填充后的图像。
//
//
//
//示例：
//
//输入：
//image = [[1,1,1],[1,1,0],[1,0,1]]
//sr = 1, sc = 1, newColor = 2
//输出：[[2,2,2],[2,2,0],[2,0,1]]
//解释:
//初始坐标点位于图像的正中间，坐标 (sr,sc)=(1,1) 。
//初始坐标点周围区域上所有符合条件的像素点的颜色都被更改成 2 。
//注意，右下角的像素没有更改为 2 ，因为它不属于初始坐标点的周围区域。
//
//
//提示：
//
//image 和 image[0] 的长度均在范围 [1, 50] 内。
//初始坐标点 (sr,sc) 满足 0 <= sr < image.length 和 0 <= sc < image[0].length 。
//image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535] 内。
//通过次数14,382提交次数25,748
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/24/21 7:39 PM
package color_fill_lcci

// 使用广度遍历
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	directions := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	originColor := image[sr][sc]
	if originColor == newColor {
		return image
	}
	q := [][2]int{{sr, sc}}
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		image[x][y] = newColor
		for _, direct := range directions {
			x1, y1 := x+direct[0], y+direct[1]
			if x1 >= 0 && x1 < len(image) && y1 >= 0 && y1 < len(image[0]) && image[x1][y1] == originColor {
				q = append(q, [2]int{x1, y1})
			}
		}
	}
	return image
}
