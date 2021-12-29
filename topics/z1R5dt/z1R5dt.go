// 剑指 Offer II 066. 单词之和
//实现一个 MapSum 类，支持两个方法，insert 和 sum：
//
//MapSum() 初始化 MapSum 对象
//void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
//int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
//
//
//示例：
//
//输入：
//inputs = ["MapSum", "insert", "sum", "insert", "sum"]
//inputs = [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
//输出：
//[null, null, 3, null, 5]
//
//解释：
//MapSum mapSum = new MapSum();
//mapSum.insert("apple", 3);
//mapSum.sum("ap");           // return 3 (apple = 3)
//mapSum.insert("app", 2);
//mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
//
//
//提示：
//
//1 <= key.length, prefix.length <= 50
//key 和 prefix 仅由小写英文字母组成
//1 <= val <= 1000
//最多调用 50 次 insert 和 sum
//
//
//注意：本题与主站 677 题相同： https://leetcode-cn.com/problems/map-sum-pairs/
//
//通过次数2,463提交次数3,793
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/29/21 1:49 PM
package z1R5dt

type trie struct {
	prefix, val int
	children    [26]*trie
}

func (t *trie) add(s string, val int) {
	t1 := t.find(s)
	diff := val
	if t1 != nil {
		diff = val - t1.val
	}
	curr := t
	for _, b := range s {
		i := b - 'a'
		if curr.children[i] == nil {
			curr.children[i] = &trie{}
		}
		curr = curr.children[i]
		curr.prefix += diff
	}
	curr.val = val
}

func (t *trie) find(s string) *trie {
	curr := t
	for _, b := range s {
		i := b - 'a'
		curr = curr.children[i]
		if curr == nil {
			break
		}
	}
	return curr
}

func (t *trie) sum(prefix string) int {
	curr := t
	for _, b := range prefix {
		curr = curr.children[b-'a']
		if curr == nil {
			return 0
		}
	}
	return curr.prefix
}

type MapSum struct {
	t *trie
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{&trie{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.t.add(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	return this.t.sum(prefix)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
