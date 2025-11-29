package main

import (
	"container/heap"
	"os"
	"path/filepath"
)

type Node struct {
	Name     string
	Size     int64
	Children []*Node
}

func buildTree(path string) *Node {
	root := &Node{Name: filepath.Base(path)}
	populate(root, path)
	return root
}

// paths to skip
var blacklist = map[string]bool{
	"/proc": true,
	"/sys":  true,
	"/dev":  true,
	"/run":  true,
}

func populate(node *Node, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return // skip unaccessable dirs
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())

		if blacklist[childPath] {
			continue
		}

		if entry.IsDir() {
			child := &Node{Name: entry.Name()}
			populate(child, childPath)
			node.Children = append(node.Children, child)
			node.Size += child.Size
		} else {
			info, err := entry.Info()
			if err != nil {
				continue // skip unaccessable files
			}
			child := &Node{
				Name: entry.Name(),
				Size: info.Size(),
			}
			node.Children = append(node.Children, child)
			node.Size += child.Size
		}
	}
}

// Since folders are the size of the sum of their children this can be greedy
func prune(root *Node, size int) *Node {
	keep := make(map[*Node]bool)
	pq := &NodeHeap{root}
	heap.Init(pq)

	for pq.Len() > 0 && len(keep) < size {
		node := heap.Pop(pq).(*Node)
		keep[node] = true

		for _, child := range node.Children {
			heap.Push(pq, child)
		}
	}

	var reduced func(node *Node) *Node
	reduced = func(node *Node) *Node {
		if node == nil || !keep[node] {
			return nil
		}

		newChildren := []*Node{}
		for _, child := range node.Children {
			rc := reduced(child)
			if rc == nil {
				continue
			}
			newChildren = append(newChildren, rc)
		}
		node.Children = newChildren
		return node
	}

	return reduced(root)
}

// Max Heap (this reminds me of old school JS classes ;.;)
type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Size > h[j].Size }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x any)        { *h = append(*h, x.(*Node)) }
func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
