package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func populate(node *Node, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return // skip unaccessable dirs
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())

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

func printTree(root *Node, prefix string) {
	fmt.Println(prefix + root.Name + " : " + strconv.FormatInt(root.Size, 10))
	for _, entry := range root.Children {
		if entry.Children != nil {
			printTree(entry, prefix+"  ")
		} else {
			fmt.Println(prefix + entry.Name + " : " + strconv.FormatInt(entry.Size, 10))
		}
	}
}
