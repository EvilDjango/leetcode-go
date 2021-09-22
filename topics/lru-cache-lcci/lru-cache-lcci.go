// 面试题 16.25. LRU 缓存
//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
//
//它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//通过次数31,133提交次数57,105
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 1:33 PM
package lru_cache_lcci

type DLinkNode struct {
	k, v       int
	prev, next *DLinkNode
}

type DLinkList struct {
	head, tail *DLinkNode
}

func NewList() *DLinkList {
	head, tail := &DLinkNode{}, &DLinkNode{}
	head.next = tail
	tail.prev = head
	return &DLinkList{head, tail}
}

func (this *DLinkList) PushBack(node *DLinkNode) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *DLinkList) MoveToFront(node *DLinkNode) {
	if this.head.next == node {
		return
	}
	this.Remove(node)
	this.PushFront(node)
}

func (this *DLinkList) PushFront(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *DLinkList) Remove(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *DLinkList) RemoveLast() *DLinkNode {
	last := this.tail.prev
	this.Remove(last)
	return last
}

type LRUCache struct {
	capacity int
	list     *DLinkList
	dict     map[int]*DLinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, NewList(), make(map[int]*DLinkNode, capacity+1)}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.dict[key]; ok {
		this.list.MoveToFront(ele)
		return ele.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.dict[key]; ok {
		ele.v = value
		this.list.MoveToFront(ele)
		return
	}
	node := &DLinkNode{k: key, v: value}
	this.list.PushFront(node)
	this.dict[key] = node
	if len(this.dict) > this.capacity {
		last := this.list.RemoveLast()
		delete(this.dict, last.k)
	}
}
