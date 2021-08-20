package heap

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/19/21 7:08 PM
*/

type Heap struct {
	data  []int
	isMax bool
}

func New(isMax bool) *Heap {
	return &Heap{isMax: isMax}
}

func Init(data []int, isMax bool) *Heap {
	h := New(isMax)
	h.data = data
	for i := len(data)/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *Heap) down(i int) {
	data := h.data
	n := len(data)
	for i < n {
		l, r := 2*i+1, 2*i+2
		extremum := i
		if l < n && data[l] > data[extremum] == h.isMax {
			extremum = l
		}
		if r < n && data[r] > data[extremum] == h.isMax {
			extremum = r
		}
		if extremum == i {
			break
		}
		data[extremum], data[i] = data[i], data[extremum]
		i = extremum
	}
}

func (h *Heap) up(i int) {
	data := h.data
	for i > 0 {
		p := (i - 1) / 2
		if data[p] >= data[i] == h.isMax {
			break
		}
		data[p], data[i] = data[i], data[p]
		i = p
	}
}

func (h *Heap) Push(num int) {
	h.data = append(h.data, num)
	h.up(len(h.data) - 1)
}

func (h *Heap) Pop() int {
	data := h.data
	top := data[0]
	data[0] = data[len(data)-1]
	h.data = data[:len(data)-1]
	h.down(0)
	return top
}

func (h *Heap) Len() int {
	return len(h.data)
}
