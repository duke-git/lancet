package datastructure

import (
	"fmt"
	"math"

	"github.com/duke-git/lancet/datastructure"
	"github.com/duke-git/lancet/lancetconstraints"
)

func insertTreeNode[T any](rootNode, newNode *datastructure.TreeNode[T], comparator lancetconstraints.Comparator) {
	if comparator.Compare(newNode.Data, rootNode.Data) == -1 {
		if rootNode.Left == nil {
			rootNode.Left = newNode
		} else {
			insertTreeNode(rootNode.Left, newNode, comparator)
		}
	} else {
		if rootNode.Right == nil {
			rootNode.Right = newNode
		} else {
			insertTreeNode(rootNode.Right, newNode, comparator)
		}
	}
}

func printTreeNodes[T any](nodes []*datastructure.TreeNode[T], level, maxLevel int) {
	if len(nodes) == 0 || isAllNil(nodes) {
		return
	}

	floor := maxLevel - level
	endgeLines := int(math.Pow(float64(2), (math.Max(float64(floor)-1, 0))))
	firstSpaces := int(math.Pow(float64(2), float64(floor))) - 1
	betweenSpaces := int(math.Pow(float64(2), float64(floor)+1)) - 1

	printSpaces(firstSpaces)

	newNodes := []*datastructure.TreeNode[T]{}
	for _, node := range nodes {
		if node != nil {
			fmt.Printf("%v", node.Data)
			newNodes = append(newNodes, node.Left)
			newNodes = append(newNodes, node.Right)
		} else {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			printSpaces(1)
		}

		printSpaces(betweenSpaces)
	}

	fmt.Println("")

	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printSpaces(firstSpaces - i)
			if nodes[j] == nil {
				printSpaces(endgeLines + endgeLines + i + 1)
				continue
			}

			if nodes[j].Left != nil {
				fmt.Print("/")
			} else {
				printSpaces(1)
			}

			printSpaces(i + i - 1)

			if nodes[j].Right != nil {
				fmt.Print("\\")
			} else {
				printSpaces(1)
			}
			printSpaces(endgeLines + endgeLines - 1)
		}
		fmt.Println("")
	}

	printTreeNodes(newNodes, level+1, maxLevel)
}

// printSpaces
func printSpaces(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func isAllNil[T any](nodes []*datastructure.TreeNode[T]) bool {
	for _, v := range nodes {
		if v != nil {
			return false
		}
	}
	return true
}
