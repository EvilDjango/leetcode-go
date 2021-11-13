// 面试题 10.10. 数字流的秩
//假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。请实现数据结构和算法来支持这些操作，也就是说：
//
//实现 track(int x) 方法，每读入一个数字都会调用该方法；
//
//实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。
//
//注意：本题相对原题稍作改动
//
//示例:
//
//输入:
//["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
//[[], [1], [0], [0]]
//输出:
//[null,0,null,1]
//提示：
//
//x <= 50000
//track 和 getRankOfNumber 方法的调用次数均不超过 2000 次
//通过次数4,882提交次数7,828
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/10/21 12:42 PM
package re_space_lcci

// 树状数组树状数组必须从下标1开始，所以输入的数字（可能为0）都加1
type StreamRank struct {
	counts [50002]int
}

func Constructor() StreamRank {
	return StreamRank{}
}

func (this *StreamRank) Track(x int) {
	for i := x + 1; i < 50001; i += i & -i {
		this.counts[i]++
	}
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	rank := 0
	for i := x + 1; i > 0; i -= i & -i {
		rank += this.counts[i]
	}
	return rank
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
