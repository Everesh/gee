package main

import (
	"fmt"
	"math"
	"strings"
)

func printTree(root *Node, width int) {
	lines := []string{}
	populateLines(root, width, int(root.Size), "", &lines)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func populateLines(root *Node, width, max int, prefix string, lines *[]string) {
	widthTree := int(math.Min(45, float64(width/3)))

	name := prefix + root.Name
	if len(name) > widthTree {
		name = name[:widthTree]
	}

	*lines = append(*lines, " "+
		humanSize(root.Size)+" "+
		name+strings.Repeat(" ", widthTree-len(name))+" │"+
		makeBar(max, int(root.Size), width-widthTree-20)+"│ "+
		fmt.Sprintf("%3d", int(100*float64(root.Size)/float64(max)))+"% ")

	for _, child := range root.Children {
		populateLines(child, width, max, prefix+"  ", lines)
	}
}

func humanSize(size int64) string {
	units := []string{"B ", "KB", "MB", "GB", "TB", "PB"}

	if size < 1024 {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	fsize := float64(size)
	unitIndex := 0

	for fsize >= 1024 && unitIndex < len(units)-1 {
		fsize /= 1024
		unitIndex++
	}

	if fsize >= 100 {
		return fmt.Sprintf("%4.0f %s ", fsize, units[unitIndex])
	} else if fsize >= 10 {
		return fmt.Sprintf("%4.1f %s ", fsize, units[unitIndex])
	} else {
		return fmt.Sprintf("%4.2f %s ", fsize, units[unitIndex])
	}
}

func makeBar(max, size, width int) string {
	percent := float64(size) / float64(max)
	filled := int(percent * float64(width))

	filled = int(math.Min(float64(filled), float64(width)))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return bar
}
