package main

import (
	"container/list"
	"sort"
)

/**
 * lfu design
 * @param operators int整型二维数组 ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

type LFUCache struct {
	Cache    map[int]*Node
	Track    map[int]*list.List
	Capacity int
}

type Node struct {
	Key   int
	Value int
	Freq  int
}

func (lfu *LFUCache) set(key, value int) {
	node, ok := lfu.Cache[key]
	if ok {
		node.Value = value
	} else {
		node = &Node{key, value, 0}
	}
	lfu.Cache[key] = node
	lfu.updateFreq(node)
	if len(lfu.Cache) > lfu.Capacity {
		keys := []int{}
		for k := range lfu.Track {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		if list, ok := lfu.Track[keys[0]]; ok {
			first := list.Remove(list.Front())
			if list.Len() == 0 {
				delete(lfu.Track, keys[0])
			} else {
				lfu.Track[keys[0]] = list
			}
			delete(lfu.Cache, first.(Node).Key)
		}
	}
}

func (lfu *LFUCache) get(key int) int {
	node, ok := lfu.Cache[key]
	if !ok {
		return -1
	}
	lfu.updateFreq(node)
	return node.Value
}

func (lfu *LFUCache) updateFreq(node *Node) {
	if l, ok := lfu.Track[node.Freq]; ok {
		var e *list.Element
		for e = l.Front(); e != nil; e = e.Next() {
			if e.Value == node {
				l.Remove(e)
				break
			}
		}
		if l.Len() == 0 {
			delete(lfu.Track, node.Freq)
		} else {
			lfu.Track[node.Freq] = l
		}
	}
	node.Freq++
	l, ok := lfu.Track[node.Freq]
	if !ok {
		l = list.New()
	}
	l.PushBack(node)
	lfu.Track[node.Freq] = l
}

func LFU(operators [][]int, k int) []int {
	// write code here
	ans := []int{}
	lfu := &LFUCache{make(map[int]*Node), make(map[int]*list.List), k}
	for _, v := range operators {
		if v[0] == 1 {
			lfu.set(v[1], v[2])
		} else if v[0] == 2 {
			ans = append(ans, lfu.get(v[1]))
		}
	}
	return ans
}
